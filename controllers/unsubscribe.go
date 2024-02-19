package controllers

import (
	"context"
	"fmt"
	"goframe/sse"
	"log/slog"
)

type Unsubscribe struct {
	ID    string `json:"id,omitempty" validate:"required"`
	Event string `json:"event,omitempty" validate:"omitempty,oneof=event.time event.name"`
}

func (s *Unsubscribe) Controller(ctx context.Context) (int, string, *SubscribeRes, error) {

	events := []string{}

	if s.Event == "" {
		events = append(events, string(s.Event))
	}

	if len(events) == 0 {
		events = []string{string(SupportedEvents.V.timer), string(SupportedEvents.V.name)}
	}
	err := sse.UnSubscribe(s.ID, events...)
	if err != nil {
		slog.With("id", s.ID).With("err", err).Error("error unsubscribing")
		return 400, fmt.Sprintf("error unsubscribing: %v", err), nil, nil
	}

	return 200, "ok", nil, nil
}
