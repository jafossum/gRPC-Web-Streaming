package nats

import "github.com/nats-io/nats.go"

type nc struct {
	conn *nats.Conn
	subs []*nats.Subscription
}

func NewNats() *nc {
	// Connect to a server
	conn, _ := nats.Connect(nats.DefaultURL)
	return &nc{
		conn: conn,
	}
}

func (c *nc) Subscribe(topic string) (<-chan *nats.Msg, error) {
	// Channel Subscriber
	ch := make(chan *nats.Msg)
	sub, err := c.conn.ChanSubscribe(topic, ch)
	if err != nil {
		return nil, err
	}

	c.subs = append(c.subs, sub)
	return ch, nil
}

func (c *nc) Close() {
	for _, sub := range c.subs {
		sub.Unsubscribe()
		sub.Drain()
	}
	c.conn.Close()
}
