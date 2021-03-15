package main

import (
	events "github.com/ronniel1/eventstest/event"
	eventslib "github.com/ronniel1/eventstest/events"
)

func main() {
	var ev eventslib.ClusterEvent
	ev = events.NewClusterVipUpdateEvent("my-vip-ip", "my-ingress-ip")
	sender := eventslib.PrintSender{}
	sender.SendClusterEvent("cluster1", ev)
	sender.SendHostEvent("cluster2", "host-10", events.NewHostVipUpdatedEvent(66))
}

/*
// event

type ClusterVipUpdateEvent struct {
	eventslib.EventBase
	ApiVip     string
	IngressVip string
}

func NewClusterVipUpdateEvent(apiVip string, ingressVip string) *ClusterVipUpdateEvent {
	return &ClusterVipUpdateEvent{
		EventBase: eventslib.EventBase{
			Format:   "Cluster was updated with api-vip {api_vip}, ingress-vip {ingress_vip}",
			Id:       "AI000001",
			Type:     eventslib.EventTypeCluster,
			Severity: eventslib.EventSeverityInfo,
		},
		ApiVip:     apiVip,
		IngressVip: ingressVip,
	}
}

func (e ClusterVipUpdateEvent) GetBaseEvent() *eventslib.EventBase {
	return &e.EventBase
}
func (e ClusterVipUpdateEvent) FormatMessage() string {
	r := strings.NewReplacer("{api_vip}", e.ApiVip,
		"{ingress_vip}", e.IngressVip)
	return r.Replace(e.EventBase.Format)
}

- name: cluster_vip_update
id: "AI000001"
event_type: cluster
severity: "info"
properties:
api_vip: string
ingress_vip: string

}
func format_msg(format string, params EventParams) string {
	args2 := make([]string, len(args))
	for i, v := range args {
		if i%2 == 0 {
			args2[i] = fmt.Sprintf("{%v}", v)
		} else {
			args2[i] = fmt.Sprint(v)
		}
	}
	r := strings.NewReplacer(args2...)
	return r.Replace(format)
}

func (e Sender) SendEvent(ctx context.Context, eventTime time.Time) {
	//format_msg()

	//output event
	fmt.Printf("Event: %s\n", "asf")
}

*/
