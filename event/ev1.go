// Code generated by from events definition in swagger; DO NOT EDIT.
package event

import (
    eventslib "github.com/ronniel1/eventstest/events"
    "strings"
)


//
// Event cluster_vip_update
//

type ClusterVipUpdateEvent struct {
    eventslib.EventBase
    ApiVip string
    IngressVip string
}

func NewClusterVipUpdateEvent(
    apiVip string,
    ingressVip string,
) *ClusterVipUpdateEvent {
    return &ClusterVipUpdateEvent {
		EventBase: eventslib.EventBase {
			Format:   "Cluster was updated with api-vip {api_vip}, ingress-vip {ingress_vip}",
			Id:       "AI000001",
			Type:     eventslib.EventTypeCluster,
			Severity: eventslib.EventSeverityInfo,
		},
        ApiVip: apiVip,
        IngressVip: ingressVip,
    }
}

func (e *ClusterVipUpdateEvent) GetEventBase() *eventslib.EventBase {
    return &e.EventBase
}

func (e *ClusterVipUpdateEvent) FormatMessage() string {
    r := strings.NewReplacer(
    "{{ api_vip }}", e.ApiVip,
    "{{ ingress_vip }}", e.IngressVip,
                  )
          return r.Replace(e.EventBase.Format)
}
//
// Event cluster_vip_removed
//

type ClusterVipRemovedEvent struct {
    eventslib.EventBase
    Vip string
}

func NewClusterVipRemovedEvent(
    vip string,
) *ClusterVipRemovedEvent {
    return &ClusterVipRemovedEvent {
		EventBase: eventslib.EventBase {
			Format:   "Vip {vip_id} was removed from cluster",
			Id:       "AI000001",
			Type:     eventslib.EventTypeCluster,
			Severity: eventslib.EventSeverityInfo,
		},
        Vip: vip,
    }
}

func (e *ClusterVipRemovedEvent) GetEventBase() *eventslib.EventBase {
    return &e.EventBase
}

func (e *ClusterVipRemovedEvent) FormatMessage() string {
    r := strings.NewReplacer(
    "{{ vip }}", e.Vip,
                  )
          return r.Replace(e.EventBase.Format)
}