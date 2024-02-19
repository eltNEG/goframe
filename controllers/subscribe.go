package controllers

import (
	"context"
	"goframe/enum"
	"goframe/sse"
	"goframe/utils"
	"slices"
	"time"
)

type SupportedEvent string

var SupportedEvents = enum.MakeEnum[SupportedEvent](struct {
	timer SupportedEvent
	name  SupportedEvent
}{
	timer: "event.time",
	name:  "event.name",
})

type Subscribe struct {
	ID     string           `json:"id,omitempty" validate:"required"`
	Events []SupportedEvent `json:"events,omitempty" validate:"not oneof=time name"`
}

type SubscribeRes struct {
	Message string `json:"message,omitempty"`
}

var subscriptions = make(map[string][]SupportedEvent)

var eventResponders = make(map[string]*sse.EventResponder[EventData])

type EventData struct {
	Event string      `json:"event,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func (s *Subscribe) Controller(ctx context.Context) (int, string, *SubscribeRes, error) {
	res := &SubscribeRes{
		Message: "success",
	}

	for _, e := range s.Events {
		if !SupportedEvents.IsValid(e) {
			return 400, "error", nil, nil
		}
	}

	subscriptions[s.ID] = s.Events

	r := utils.GetOriginalReq(ctx)
	w := utils.GetOriginalRes(ctx)

	responder, err := sse.Http2EventHandler[EventData](w, r)

	if err != nil {
		return 400, "error", nil, nil
	}
	eventResponders[s.ID] = responder

	if slices.Contains(s.Events, SupportedEvents.V.timer) {
		ticker := time.NewTicker(30 * time.Second)
		for {
			select {
			case <-ticker.C:
				t := time.Now().Format("2006-01-02 15:04:05")
				eventData := &EventData{
					Event: "server.ping",
					Data:  t,
				}
				responder.Respond(eventData)
			case <-r.Context().Done():
				delete(eventResponders, s.ID)
				return 200, "success", res, nil
			}
		}
	}

	return 200, "success", res, nil
}
