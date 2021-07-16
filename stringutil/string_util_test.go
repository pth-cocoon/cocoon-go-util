package stringutil

import (
	"errors"
	"testing"
)

func TestToString(t *testing.T) {
	kv := make(map[interface{}]string)
	kv[123] = "123"
	kv[true] = "true"
	kv[false] = "false"
	kv[0.123] = "0.123"
	kv[errors.New("error")] = "error"
	for k, v := range kv {
		a := ToString(k)
		if a == v {
			continue
		}
		t.Errorf("except: %s,but got %s", k, v)
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
		t.Errorf("except: %s,but got %s", k, v)
	}
}
