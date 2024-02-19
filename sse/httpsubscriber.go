package sse

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	log "log/slog"
	"net/http"
)

type EventResponder[T any] struct {
	events chan *T
	cancel chan struct{}
}

func (c *EventResponder[T]) Respond(data *T) {
	c.events <- data
}

func Http2EventHandler[Event any](w http.ResponseWriter, r *http.Request) (c *EventResponder[Event], err error) {
	origin := r.Header.Get("Origin")
	log.With("origin", origin).Debug("https event origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	f, ok := w.(http.Flusher)
	if !ok {
		log.Error("no flusher")
		return c, errors.New("no flusher: invalid http2 request")
	}

	c = &EventResponder[Event]{
		events: make(chan *Event),
	}

	go func() {
		for ev := range c.events {
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			err := enc.Encode(ev)
			if err != nil {
				log.With("err", err).Error("error encoding event data")
			}
			fmt.Fprintf(w, "data:%v\n\n", buf.String())
			f.Flush()
			if c.cancel != nil {
				close(c.cancel)
				return
			}
		}
	}()
	return c, nil
}
