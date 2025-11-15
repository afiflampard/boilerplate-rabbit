package rabbit

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func Connect(cfg RabbitMQConfig) (*RabbitMQ, error) {
	uri := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		cfg.Protocol,
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		normalizeVHost(cfg.VHost),
	)

	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Printf("❌ Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("❌ Failed to open channel: %v", err)
		return nil, err
	}
	log.Println("✅ Connected to RabbitMQ")
	return &RabbitMQ{conn: conn, channel: ch}, nil
}

func normalizeVHost(vhost string) string {
	if vhost == "" || vhost == "/" {
		return "%2f" // URL-encoded "/"
	}
	return vhost
}

func (r *RabbitMQ) Connection() *amqp.Connection {
	return r.conn
}

func (r *RabbitMQ) Close() {
	if r.conn != nil {
		if err := r.conn.Close(); err != nil {
			log.Printf("⚠️ Error closing RabbitMQ connection: %v", err)
		} else {
			log.Println("✅ RabbitMQ connection closed")
		}
	}
}

func (r *RabbitMQ) PublishQueue(queue string, body []byte) error {
	// declare queue
	_, err := r.channel.QueueDeclare(
		queue,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// publish message
	return r.channel.Publish(
		"",    // default exchange
		queue, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)
}

func (r *RabbitMQ) Publish(exchange, routingKey string, body []byte) error {
	// declare exchange
	err := r.channel.ExchangeDeclare(
		exchange,
		"topic",
		true,  // durable
		false, // autoDelete
		false, // internal
		false, // noWait
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// publish
	return r.channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)
}
