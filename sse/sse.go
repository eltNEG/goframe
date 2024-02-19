package sse

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"goframe/utils"
	log "log/slog"
	"net/http"
	"slices"
)

type Event struct {
	Event string `json:"event"`
	Data  any    `json:"data"`
}

type EventString interface {
	~string
}

var idResponders = make(map[string]*EventResponder[Event])
var idEvents = make(map[string][]string)

func Respond(id string, data *Event) error {
	if evs, ok := idEvents[id]; !ok || !slices.Contains(evs, data.Event) {
		return errors.New("event not subscribed")
	}

	if c, ok := idResponders[id]; ok {
		c.Respond(data)
		return nil
	}
	return errors.New("no responder found")
}

func UnSubscribe[SE ~string](id string, events ...SE) error {
	log.With("id", id).With("events", events).Debug("unsubscribing")
	c, ok := idResponders[id]
	if !ok {
		return errors.New("no responder found")
	}
	newEvents := []string{}

	for _, oe := range idEvents[id] {
		if !slices.Contains(events, SE(oe)) {
			newEvents = append(newEvents, oe)
		}
	}

	log.With("id", id).With("events", newEvents).Debug("new events")

	if len(newEvents) == 0 {
		log.Debug("no more events to subscribe")
		c.cancel <- struct{}{}
		return nil
	}

	idEvents[id] = newEvents
	return nil
}

func Http2EventHandlerWithSubManager(ctx context.Context, subid string, events []string) (done <-chan struct{}, err error) {
	r := utils.GetOriginalReq(ctx)
	w := utils.GetOriginalRes(ctx)
	if r == nil {
		return nil, errors.New("invalid request or response in context")
	}
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
		return nil, errors.New("no flusher: invalid http2 request")
	}

	c := &EventResponder[Event]{
		events: make(chan *Event),
		cancel: make(chan struct{}),
	}

	if idResponders[subid] != nil {
		idEvents[subid] = append(idEvents[subid], events...)
		return nil, nil
	}

	idResponders[subid] = c
	idEvents[subid] = events

	ctx.Err()

	doneChan := make(chan struct{})

	go func() {
		for {
			select {
			case <-ctx.Done():
				delete(idResponders, subid)
				delete(idEvents, subid)
				log.With("err", ctx.Err()).With("id", subid).Debug("context done in sse")
				doneChan <- struct{}{}
				close(doneChan)
				return
			case ev := <-c.events:
				var buf bytes.Buffer
				enc := json.NewEncoder(&buf)
				err := enc.Encode(ev)
				if err != nil {
					fmt.Println("encode err: ", err)
				}
				fmt.Fprintf(w, "data:%v\n\n", buf.String())
				f.Flush()
			case <-c.cancel:
				delete(idResponders, subid)
				delete(idEvents, subid)
				log.With("id", subid).Debug("context done in sse by cancel")
				close(doneChan)
				close(c.events)
				close(c.cancel)
				return
			}
		}
	}()
	return doneChan, nil
}
