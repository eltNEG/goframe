package controllers

import (
	"context"
	"goframe/sse"
	log "log/slog"
	"time"
)

type Subscribe2 struct {
	ID    string `json:"id,omitempty" validate:"required"`
	Event string `json:"event,omitempty" validate:"omitempty,oneof=event.time event.name"`
}

func (s *Subscribe2) Controller(ctx context.Context) (int, string, *SubscribeRes, error) {

	events := []string{s.Event}

	if s.Event == "" {
		events = []string{string(SupportedEvents.V.name), string(SupportedEvents.V.timer)}
	}

	done, err := sse.Http2EventHandlerWithSubManager(ctx, s.ID, events)

	if err != nil {
		return 400, "error", nil, nil
	}

	if done == nil {
		return 200, "subscribed", nil, nil
	}

	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			t := time.Now().Format("2006-01-02 15:04:05")
			eventData := &sse.Event{
				Event: "event.time",
				Data:  t,
			}
			err := sse.Respond(s.ID, eventData)
			if err != nil {
				log.With("id", s.ID).With("err", err).Error("error responding")
			}
		case <-done:
			log.With("id", s.ID).Debug("context done in sub2")
			delete(eventResponders, s.ID)
			ticker.Stop()
			return 0, "success", nil, nil
		}
	}
}
