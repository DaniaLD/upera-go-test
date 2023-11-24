package productRevisionConsumerHandler

import (
	"encoding/json"
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"github.com/rabbitmq/amqp091-go"
)

func (h *ProductRevisionConsumerHandler) ProductRevisionCreationConsumerHandler(msgs <-chan amqp091.Delivery) {
	var forever chan struct{}
	go func() {
		for d := range msgs {
			var message globalModel.ProductUpdateEvent
			_ = json.Unmarshal(d.Body, &message)
			h.productRevisionAppService.Create(message.ProductId, message.PreviousData, message.NewData)
			d.Ack(false)
		}
	}()
	<-forever
}
