package enum

import (
	"encoding/json"
)

// type keyval map[K comparable, V bool] map[K]V

type Enum[T any, E comparable] struct {
	V      *T
	keyval map[E]bool
}

func MakeEnum[E comparable, T any](enum T) *Enum[T, E] {
	kvbyte, _ := json.Marshal(enum)
	kv := map[string]E{}
	keyval := map[E]bool{}
	json.Unmarshal(kvbyte, &kv)
	for _, v := range kv {
		keyval[v] = true
	}
	return &Enum[T, E]{&enum, keyval}
}

func (e *Enum[T, E]) IsValid(data E) bool {
	if _, ok := e.keyval[data]; !ok {
		return false
	}
	return true
}
