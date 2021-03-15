package events

import (
	"fmt"
)

type ClusterEvent interface {
	FormatMessage() string
	GetBaseEvent() *ClusterBaseEvent
}

type HostEvent interface {
	FormatMessage() string
	GetBaseEvent() *HostBaseEvent
}

type EventSeverity int
const (
	EventSeverityInfo EventSeverity = iota
	EventSeverityWarning
)

type ClusterBaseEvent struct {
	Format string
	Id     string
	Severity EventSeverity
}

/*
func (e *ClusterBaseEvent) GetBaseEvent() *ClusterBaseEvent {
	return nil
}

func (e *ClusterBaseEvent) FormatMessage() string {
	return ""
}
 */

type HostBaseEvent struct {
	Format string
	Id     string
	Severity EventSeverity
}

/*
func (e *HostBaseEvent) GetBaseEvent() *HostBaseEvent {
	return nil
}

func (e *HostBaseEvent) FormatMessage() string {
	return ""
}

 */


type EventParams map[string]interface{}

type  Sender interface {
	SendHostEvent(clusterId string, hostId string, event HostEvent)
	SendClusterEvent(clusterId string, event ClusterEvent)
}


type PrintSender struct {}

func (p PrintSender) SendHostEvent(clusterId string, hostId string, event HostEvent) {
	fmt.Printf("Host Event(%s): %s\n", event.GetBaseEvent().Id, event.FormatMessage())
}

func (p PrintSender) SendClusterEvent(clusterId string, event ClusterEvent) {
	fmt.Printf("Cluster Event(%s): %s\n", event.GetBaseEvent().Id, event.FormatMessage())
}
