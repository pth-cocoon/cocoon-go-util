package stringutil

import (
	"errors"
	"testing"
)

func TestToString(t *testing.T) {
	kv := make(map[interface{}]string)
	kv[true] = "true"
	kv[false] = "false"
	kv[0.123] = "0.123"
	kv[errors.New("error")] = "error"
	var i1 int = 1
	var i8 int8 = 8
	var i16 int16 = 16
	var i64 int64 = 64
	var f32 float32 = 0.32
	var f64 float64 = 0.64
	kv[i1] = "1"
	kv[i8] = "8"
	kv[i16] = "16"
	kv[i64] = "64"
	kv[f32] = "0.32"
	kv[f64] = "0.64"
	for k, v := range kv {
		a := ToString(k)
		println(a)
		if a == v {
			continue
		}
		t.Errorf("except: %s,but got %s", v, a)
	}
}

func TestReverseString(t *testing.T) {
	kv := make(map[string]string)
	kv["123"] = "321"
	kv["HelloWorld"] = "dlroWolleH"
	kv["  123"] = "321  "
	for k, v := range kv {
		a := ReverseString(k)
		if a == v {
			continue
		}
		t.Errorf("except: %s,but got %s", v, a)
	}
}
