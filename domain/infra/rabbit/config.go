package rabbit

type RabbitMQConfig struct {
	Host          string `envconfig:"HOST" default:"localhost"`
	Port          string `envconfig:"PORT" default:"5672"`
	Username      string `envconfig:"USERNAME" default:"guest"`
	Password      string `envconfig:"PASSWORD" default:"guest"`
	Protocol      string `envconfig:"PROTOCOL" default:"amqp"`
	VHost         string `envconfig:"VHOST" default:"/"`
	PrefetchCount int    `envconfig:"PREFETCH_COUNT" default:"1"`
}
