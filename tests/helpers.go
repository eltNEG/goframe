package tests

import (
	"context"
	"goframe/utils"
	"reflect"
	"testing"
)

type ResponseObject[R any] struct {
	Code    int
	Message string
	Data    R
	ErrMsg  []string
}

func MakeReqRes[M utils.Model[R], R any](req M, res *ResponseObject[R]) func(t *testing.T) {
	return func(t *testing.T) {
		vr, err := utils.ValidateData(req)
		if err != nil {
			if vr.Message != res.Message {
				t.Errorf("res = %+v want %v", vr.Message, res.Message)
			}
			if !reflect.DeepEqual(vr.Data, res.ErrMsg) {
				t.Errorf("res = %+v want %v", vr.Data, res.ErrMsg)
			}
			if res.Code != 400 {
				t.Errorf("code = %v, want %v", res.Code, 400)
			}
			return
		}

		c, m, r, e := req.Controller(context.Background())
		if c != res.Code {
			t.Errorf("code = %v, want %v", c, res.Code)
		}
		if m != res.Message {
			t.Errorf("msg = %v, want %v", m, res.Message)
		}
		if e != nil {
			if e.Error() != res.ErrMsg[0] {
				t.Errorf("err = %v, want %v", e, res)
			}

		}
		if !reflect.DeepEqual(r, res.Data) {
			t.Errorf("res = %v, want %v", r, res.Data)
		}
	}
}
