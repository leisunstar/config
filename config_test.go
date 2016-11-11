package config

import (
	"testing"
)

type testC struct {
	Id         int               `json:"id"`
	IString    string            `json:"i_string"`
	IntList    []int             `json:"int_list"`
	StringList []string          `json:"string_list"`
	StrMap     map[string]string `json:"string_map"`
}

func TestConfig(t *testing.T) {
	c := &testC{}
	err := Unmarshal("test.conf", c)
	if err != nil {
		t.Error("err %v", err)
	}
	t.Logf("c (%v)", c)
}
