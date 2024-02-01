package tests

import (
	"errors"
	"goframe/controllers"
	"goframe/utils"
	"testing"
)

func TestControllers(t *testing.T) {
	tests := []*struct {
		name string
		req  func(t *testing.T)
	}{
		{
			name: "ping success",
			req: MakeReqRes(
				&controllers.Ping{
					Name:  "a",
					Name2: "f",
					Name3: "c",
				},
				&controllers.PingRes{
					Name:    "a",
					Version: 1,
				},
				200,
				"success",
				errors.New(""),
				nil,
			),
		},
		{
			name: "ping success with no data",
			req: MakeReqRes(
				&controllers.Ping{
					Name:  "e",
					Name2: "f",
					Name3: "c",
				},
				nil,
				200,
				"success",
				errors.New(""),
				nil,
			),
		},
		{
			name: "ping fail with name required",
			req: MakeReqRes(
				&controllers.Ping{
					// Name:  "e",
					Name2: "f",
					Name3: "c",
				},
				nil,
				400,
				"validation error",
				errors.New(""),
				&utils.Response{
					Message: "validation error",
					Data:    []string{utils.MakeValidationErr("Name", "required", "", "")},
				},
			),
		},
		{
			name: "ping fail with name2 oneof and name3 required",
			req: MakeReqRes(
				&controllers.Ping{
					Name:  "e",
					Name2: "g",
					// Name3: "c",
				},
				nil,
				400,
				"validation error",
				errors.New(""),
				&utils.Response{
					Message: "validation error",
					Data: []string{
						utils.MakeValidationErr("Name2", "oneof", "e f", "g"),
						utils.MakeValidationErr("Name3", "required", "", ""),
					},
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.req(t)
		})
	}
}
