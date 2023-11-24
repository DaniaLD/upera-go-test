package productRepositoryImpl

import (
	messagingQueue "github.com/DaniaLD/upera-go-test/internal/infrastructure/messaging-queue"
	"github.com/spf13/viper"
)

func (r *productRepositoryImpl) DispatchUpdateEvent(revisionData []byte) {
	exchange := messagingQueue.Exchange{
		Name:       viper.GetString("messagingQueue.rabbit.productUpdateEvent.exchange"),
		Type:       "topic",
		Passive:    false,
		Durable:    true,
		AutoDelete: false,
	}
	routingKey := viper.GetString("messagingQueue.rabbit.productUpdateEvent.routingKey")
	_ = r.rabbitMQConn.PublishToExchange(exchange, routingKey, revisionData, 0)
}
