package consumer

import (
	messagingQueue "github.com/DaniaLD/upera-go-test/internal/infrastructure/messaging-queue"
	"github.com/spf13/viper"
)

func (c *Consumer) productRevisionConsumers() {
	c.productUpdatedConsumer()
}

func (c *Consumer) productUpdatedConsumer() {
	msgs, err := c.rabbitMqConnection.BindQueueToExchangeAndConsume(messagingQueue.Exchange{
		Name:       viper.GetString("messagingQueue.rabbit.productUpdateEvent.exchange"),
		Type:       "topic",
		Passive:    false,
		Durable:    true,
		AutoDelete: false,
	}, messagingQueue.Queue{
		Name:       "product_updated",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		RoutingKey: viper.GetString("messagingQueue.rabbit.productUpdateEvent.routingKey"),
	})

	if err != nil {
		panic(err)
	}
	go c.productRevisionConsumerHandler.ProductRevisionCreationConsumerHandler(msgs)
}
