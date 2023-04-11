package jsoniter_test

import (
	"math"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestStructNanError(t *testing.T) {
	type t0 struct {
		C float64 `json:"c,omitempty"`
	}

	type t1 struct {
		A float64 `json:"a,omitempty"`
		B t0      `json:"b,omitempty"`
		D float64 `json:"d,omitempty"`
	}

	v := t1{B: t0{C: math.NaN()}, D: 3}

	d, err := jsoniter.Marshal(v)
	if err == nil {
		t.Fatalf("expect an error - encoded %s", string(d))
	}

	if s := err.Error(); s != "jsoniter_test.t1.B: jsoniter_test.t0.C: unsupported value: NaN" {
		t.Fatalf("error was %q", s)
	}
}
