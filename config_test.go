package config

import(
    "testing"
    "encoding/json"
    "fmt"
)

type testC struct {
    Id  int `json:"id"`
    IString  string `json:"i_string"`
    IntList []int `json:"int_list"`
    StringList []string `json:"string_list"`
}

func TestConfig(t *testing.T) {
    a := &testC{}
    a.Id = 2
    a.IntList = []int{1,2,3,4}
    a.IString = "test"
     testJosn,_ := json.Marshal(a)
    fmt.Println(string(testJosn))

    jsonStr,err := Unmarshal("/Users/heguangxiu/go/src/config/test.conf")
    if err != nil{
        t.Error("err %v", err)
    }
    C := &testC{}
    err = json.Unmarshal([]byte(jsonStr), C)
    fmt.Printf("testC %v \n", C)
}