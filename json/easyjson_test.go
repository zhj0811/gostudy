package json

import (
	"encoding/json"
	"testing"

	"github.com/zhj0811/gostudy/define"
)

// Unmarshal 前需判断[]byte 是否为 nil

func TestEasyJSON(t *testing.T) {
	var b []byte
	var s define.Data
	if b == nil {
		t.Log("b is nil")
	}
	if err := s.UnmarshalJSON(b); err != nil {
		t.Error(err.Error())
	}
}

func TestJSON(t *testing.T) {
	var b []byte
	var s define.Data
	if b != nil {
		t.Log("b is not nil")
	}
	if err := json.Unmarshal(b, &s); err != nil {
		t.Error(err.Error())
	}
}
