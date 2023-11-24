package consumer

import (
	productRevisionConsumerHandler "github.com/DaniaLD/upera-go-test/internal/app/consumer-handler/product-revision"
	messagingQueue "github.com/DaniaLD/upera-go-test/internal/infrastructure/messaging-queue"
)

type Consumer struct {
	rabbitMqConnection             *messagingQueue.Connection
	productRevisionConsumerHandler *productRevisionConsumerHandler.ProductRevisionConsumerHandler
}

func NewConsumer(
	rabbitMqConnection *messagingQueue.Connection,
	productRevisionConsumerHandler *productRevisionConsumerHandler.ProductRevisionConsumerHandler,
) *Consumer {
	return &Consumer{
		rabbitMqConnection:             rabbitMqConnection,
		productRevisionConsumerHandler: productRevisionConsumerHandler,
	}
}

func (c *Consumer) InitConsumer() {
	c.productRevisionConsumers()
}
