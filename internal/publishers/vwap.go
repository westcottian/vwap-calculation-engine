package publishers

import (
	"github.com/westcottian/vwap-calculation-engine/internal/aggregates"
	"github.com/westcottian/vwap-calculation-engine/pkg/broker"
)

type VWAPPublisher interface {
	Publish(aggregate aggregates.VWAPAggregate)
}

type LocalPublisher struct {
	cli broker.VWAPMessageBrokerClient
}

func NewLocalPublisher(cli broker.VWAPMessageBrokerClient) LocalPublisher {
	return LocalPublisher{cli: cli}
}

func (l LocalPublisher) Publish(aggregate aggregates.VWAPAggregate) {
	l.cli.Send(aggregate)
}
