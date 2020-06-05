package yaml

import (
	"github.com/tianyazc/zlxGo/zlxmodule/stringfile"
	"testing"
)

var confContent = `
name: Tony
age: 22
have:
  computer:
    - Python
    - Go
    - Bash
  son: midou
`

func TestYaml(t *testing.T)  {
	if err :=stringfile.WriteStringToFile(confContent,"./test.yaml","rw");err !=nil{
		panic(err)
	}
	c := NewConfig("./test.yaml")
	if c.Get("name").String() != "Tony" && c.Get("age").Int() == 22 {
		t.Error("simple test result error")
	}else {
		t.Log("simple test result ok")
	}
	if c.Get("have.son").String()=="midou" {
		t.Log("yaml mutil level test result ok")
	}else {
		t.Fatal("yaml mutil level result error")
	}
	if c.Get("have.computer.0").String()=="Python" {
		t.Log("yaml list test result ok")
	}else {
		t.Fatal("yaml list test result error")
	}
}