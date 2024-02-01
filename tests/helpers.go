package tests

import (
	"goframe/utils"
	"reflect"
	"testing"
)

func MakeReqRes[M utils.Model[R], R any](req M, res R, code int, msg string, err error, verr *utils.Response) func(t *testing.T) {
	return func(t *testing.T) {
		vr, err := utils.ValidateData(req)
		if err != nil {
			if vr.Message != msg {
				t.Errorf("res = %+v want %v", vr.Message, msg)
			}
			if !reflect.DeepEqual(vr.Data, verr.Data) {
				t.Errorf("res = %+v want %v", vr.Data, verr.Data)
			}
			if code != 400 {
				t.Errorf("code = %v, want %v", code, 400)
			}
			return
		}

		c, m, r, e := req.Controller()
		if c != code {
			t.Errorf("code = %v, want %v", c, code)
		}
		if m != msg {
			t.Errorf("msg = %v, want %v", m, msg)
		}
		if !reflect.DeepEqual(r, res) {
			t.Errorf("res = %v, want %v", r, res)
		}

		if (err == nil && e != nil) || (err != nil && e == nil) {
			t.Errorf("err = %v, want %v", e, err)
		}
		if err != nil && e != nil && e.Error() != err.Error() {
			t.Errorf("err = %v, want %v", e, err)
		}
	}
}
