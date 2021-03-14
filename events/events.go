package events



type Event interface {
	FormatMessage() string
	GetBaseEvent() *EventBase
}

type EventType int
const (
	EventTypeCluster EventType = iota
	EventTypeHost
)

type EventSeverity int
const (
	EventSeverityInfo EventSeverity = iota
	EventSeverityWarning
)

type EventBase struct {
	Format string
	Id     string
	Type EventType
	Severity EventSeverity
}


type EventParams map[string]interface{}

type  Sender interface {
	SendEvent(event Event)
}


