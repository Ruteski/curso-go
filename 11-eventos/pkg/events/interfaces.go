package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

// EventHandlerInterface operações que serão executadas quando um evento é chamado
type EventHandlerInterface interface {
	Handle(event EventInterface)
}

// EventDispatcherInterface gerenciador dos nossos eventos/operações
// registra os eventos e suas operações
// despachar/fire no evento para que suas operações sejam executadas
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}
