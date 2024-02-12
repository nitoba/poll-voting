package core_test

import (
	"testing"
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CustomAggregateCreatedEvent struct {
	ocurredAt time.Time
	aggregate CustomAggregate
}

func (e *CustomAggregateCreatedEvent) Name() string {
	return "event.custom.aggregate.created"
}

func (e *CustomAggregateCreatedEvent) GetAggregateId() core.UniqueEntityId {
	return core.NewUniqueEntityId()
}

func (e *CustomAggregateCreatedEvent) OcurredAt() time.Time {
	return e.ocurredAt
}

func NewCustomAggregateCreatedEvent(aggregate CustomAggregate) *CustomAggregateCreatedEvent {
	return &CustomAggregateCreatedEvent{
		aggregate: aggregate,
		ocurredAt: time.Now(),
	}
}

type CustomAggregate struct {
	core.AggregateRoot
}

func NewCustomAggregate() *CustomAggregate {
	customAggregate := &CustomAggregate{}

	// We need to add a domain event here to make sure that the aggregate is marked for dispatch
	customAggregate.AddDomainEvent(NewCustomAggregateCreatedEvent(*customAggregate))

	return customAggregate
}

type MockCallBackFunction struct {
	mock.Mock
}

func (m *MockCallBackFunction) callbackFunction(event interface{}) {
	m.Called(event)
}

func TestDomainEvents(t *testing.T) {
	t.Run("it should dispatch events for aggregate", func(t *testing.T) {
		mockCallBackFunction := &MockCallBackFunction{}

		mockCallBackFunction.On("callbackFunction", mock.Anything).Return(nil)

		core.DomainEvents().Register(mockCallBackFunction.callbackFunction, "event.custom.aggregate.created")
		core.DomainEvents().Register(mockCallBackFunction.callbackFunction, "event.custom.aggregate")

		customAggregate := NewCustomAggregate()
		assert.Len(t, customAggregate.GetEvents(), 1)

		core.DomainEvents().DispatchEventsForAggregate(customAggregate.Id)

		mockCallBackFunction.AssertNumberOfCalls(t, "callbackFunction", 1)
		assert.Len(t, customAggregate.GetEvents(), 0)
		mockCallBackFunction.AssertExpectations(t)

	})
}
