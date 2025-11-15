package consumer

import (
	"github.com/afiflampard/boilerplate-consumer/cmd/config"
	"github.com/afiflampard/boilerplate-consumer/cmd/service"
	"github.com/afiflampard/boilerplate-domain/infra/logger"
	"github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	service service.ProductService
}

func NewConsumer(service service.ProductService) Consumer {
	return Consumer{service: service}
}

func (c *Consumer) ListenAllQueues(conn *amqp091.Connection, cfg config.Config) error {
	if err := c.listenQueue(
		cfg,
		conn,
		cfg.Product.Exchange,
		cfg.Product.RoutingKey,
		cfg.Product.Queue,
		c.CreateProductConsumer,
	); err != nil {
		return err
	}
	return nil
}

func (c *Consumer) listenQueue(
	cfg config.Config,
	conn *amqp091.Connection,
	exchange, routingKey, queue string,
	handler func([]byte) error,
) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	if err := ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil); err != nil {
		return err
	}

	q, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	if err := ch.QueueBind(q.Name, routingKey, exchange, false, nil); err != nil {
		return err
	}

	// Prefetch only 1 unacknowledged message per consumer
	if err := ch.Qos(cfg.Rabbit.PrefetchCount, 0, false); err != nil {
		return err
	}

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			err := handler(d.Body)
			if err != nil {
				logger.Debug("❌ Handler error (ACK anyway): %v", err)
			}

			// ✅ Always ACK to prevent retry
			if err := d.Ack(false); err != nil {
				logger.Debug("⚠️ Failed to ACK: %v", err)
			}
		}
	}()

	return nil
}
