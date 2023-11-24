package messagingQueue

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Connection struct {
	Connection *amqp.Connection
}

type Exchange struct {
	Name       string
	Type       string
	Passive    bool
	Durable    bool
	AutoDelete bool
	Arguments  amqp.Table
}

type Queue struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	RoutingKey string
}

func NewConnection(url string) (*Connection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to create RabbitMQ connection: %v", err)
	}
	return &Connection{
		Connection: conn,
	}, err
}

func (c *Connection) Close() {
	c.Connection.Close()
}

func (c *Connection) BindQueueToExchangeAndConsume(exchange Exchange, queue Queue) (<-chan amqp.Delivery, error) {
	ch, err := c.Connection.Channel()
	if err != nil {
		return nil, fmt.Errorf("Failed to create RabbitMQ channel: %v", err)
	}

	err = declareExchange(ch, exchange)
	if err != nil {
		return nil, fmt.Errorf("Failed to create RabbitMQ exchange: %v", err)
	}

	_, err = declareQueue(ch, queue)
	if err != nil {
		return nil, fmt.Errorf("Failed to declare a queue: %v", err)
	}

	err = bindQueue(ch, exchange, queue)
	if err != nil {
		return nil, fmt.Errorf("Failed to bind a queue: %v", err)
	}

	messages, err := consumeQueue(ch, queue)
	if err != nil {
		return nil, fmt.Errorf("Failed to register a consumer: %v", err)
	}

	return messages, err
}

func (c *Connection) PublishToExchange(exchange Exchange, routingKey string, payload string, delay int) error {
	ch, err := c.Connection.Channel()
	if err != nil {
		return fmt.Errorf("Failed to create RabbitMQ channel: %v", err)
	}

	err = declareExchange(ch, exchange)
	if err != nil {
		return fmt.Errorf("Failed to create RabbitMQ exchange: %v", err)
	}

	headers := make(amqp.Table)
	if delay != 0 {
		headers["x-delay"] = delay
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
		ch.Close()
	}()
	err = ch.PublishWithContext(ctx, exchange.Name, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(payload),
		Headers:     headers,
	})
	if err != nil {
		return fmt.Errorf("Failed to publish a message: %v", err)
	}
	return nil
}

func declareExchange(channel *amqp.Channel, exchange Exchange) error {
	err := channel.ExchangeDeclare(
		exchange.Name,       // name
		exchange.Type,       // type
		exchange.Durable,    // durable
		exchange.AutoDelete, // auto-deleted
		false,               // internal
		false,               // no-wait
		exchange.Arguments,  // arguments
	)
	return err
}

func declareQueue(channel *amqp.Channel, queue Queue) (amqp.Queue, error) {
	q, err := channel.QueueDeclare(
		queue.Name,       // name
		queue.Durable,    // durable
		queue.AutoDelete, // delete when unused
		queue.Exclusive,  // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	return q, err
}

func bindQueue(channel *amqp.Channel, exchange Exchange, queue Queue) error {
	err := channel.QueueBind(
		queue.Name,       // queue name
		queue.RoutingKey, // routing key
		exchange.Name,    // exchange
		false,
		nil)
	return err
}

func consumeQueue(channel *amqp.Channel, queue Queue) (<-chan amqp.Delivery, error) {
	messages, err := channel.Consume(
		queue.Name,      // queue
		"",              // consumer
		false,           // auto ack
		queue.Exclusive, // exclusive
		false,           // no local
		false,           // no wait
		nil,             // args
	)
	return messages, err
}
