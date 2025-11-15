package product

type ProductConfig struct {
	Exchange   string `envconfig:"EXCHANGE" default:"product"`
	Queue      string `envconfig:"QUEUE" default:"product"`
	RoutingKey string `envconfig:"ROUTING_KEY" default:"product"`
}
