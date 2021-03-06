package publishers

import (
	"testing"

	"github.com/westcottian/vwap-calculation-engine/internal/aggregates"
	"github.com/westcottian/vwap-calculation-engine/test/mocks"
)

const (
	SendFunc = "Send"
)

func TestLocalPublisher(t *testing.T) {
	t.Run("publish aggregate successfully", func(t *testing.T) {
		aggregate := aggregates.MatchAggregate{
			ProductID:         "BTC-USD",
			SlidingWindowSize: 200,
		}
		clientMock := new(mocks.VWAPMessageBrokerClient)
		underTest := NewLocalPublisher(clientMock)
		clientMock.On(SendFunc, &aggregate).Return(nil)

		underTest.Publish(&aggregate)

		clientMock.AssertCalled(t, SendFunc, &aggregate)
	})
}
