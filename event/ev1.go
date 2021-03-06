// Code generated by from events definition in swagger; DO NOT EDIT.
package event

import (
	"fmt"
	eventslib "github.com/ronniel1/eventstest/events"
	"strings"
)

//
// Event cluster_vip_update
//
type ClusterVipUpdateEvent struct {
	eventslib.ClusterBaseEvent
	ApiVip     string
	IngressVip string
}

func NewClusterVipUpdateEvent(
	apiVip string,
	ingressVip string,
) *ClusterVipUpdateEvent {
	return &ClusterVipUpdateEvent{
		ClusterBaseEvent: eventslib.ClusterBaseEvent{
			Format:   "Cluster was updated with api-vip {api_vip}, ingress-vip {ingress_vip}",
			Id:       "AI000001",
			Severity: eventslib.EventSeverityInfo,
		},
		ApiVip:     apiVip,
		IngressVip: ingressVip,
	}
}

func (e *ClusterVipUpdateEvent) GetBaseEvent() *eventslib.ClusterBaseEvent {
	return &e.ClusterBaseEvent
}

func (e *ClusterVipUpdateEvent) FormatMessage() string {
	r := strings.NewReplacer(
		"{api_vip}", fmt.Sprint(e.ApiVip),
		"{ingress_vip}", fmt.Sprint(e.IngressVip),
	)
	return r.Replace(e.ClusterBaseEvent.Format)
}

//
// Event host_vip_updated
//
type HostVipUpdatedEvent struct {
	eventslib.HostBaseEvent
	Num int
}

func NewHostVipUpdatedEvent(
	num int,
) *HostVipUpdatedEvent {
	return &HostVipUpdatedEvent{
		HostBaseEvent: eventslib.HostBaseEvent{
			Format:   "Vip {num} was removed from host",
			Id:       "AI000002",
			Severity: eventslib.EventSeverityInfo,
		},
		Num: num,
	}
}

func (e *HostVipUpdatedEvent) GetBaseEvent() *eventslib.HostBaseEvent {
	return &e.HostBaseEvent
}

func (e *HostVipUpdatedEvent) FormatMessage() string {
	r := strings.NewReplacer(
		"{num}", fmt.Sprint(e.Num),
	)
	return r.Replace(e.HostBaseEvent.Format)
}
