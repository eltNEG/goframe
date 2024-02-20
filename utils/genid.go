package utils

import (
	"strings"

	"github.com/google/uuid"
)

type genIDparamFunc func(string) string

type generateIDparam struct{}

var GenIDArg = generateIDparam{}

func (generateIDparam) WithPrefix(prefix string) genIDparamFunc {
	return func(id string) string {
		if prefix == "" {
			return id
		}
		return prefix + "_" + id

	}
}

func (generateIDparam) WithSurfix(surfix string) genIDparamFunc {
	return func(id string) string {
		if surfix == "" {
			return id
		}
		return id + "_" + surfix
	}
}

func GenID(fix ...genIDparamFunc) string {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")
	for _, f := range fix {
		id = f(id)
	}
	return id
}
