package yaml

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	file string
}

func NewConfig(configPath string) *Config {
	return &Config{
		file: configPath,
	}
}

func NewDefaultConfigGet(path string) *Result  {
	c := NewConfig("./etc/config.yaml")
	return c.Get(path)
}

func (c *Config) ConfigParse() (map[interface{}]interface{}, error) {
	var m = make(map[interface{}]interface{})
	f, err := os.Open(c.file)
	if err != nil {
		return m, err
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return m, err
	}
	err = yaml.Unmarshal(content, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}


type Result struct {
	res interface{}
	error error
}

func (c *Config) Get(path string) *Result {
	var r interface{}
	paths := strings.Split(path, ".")
	configs, err := c.ConfigParse()
	if err != nil {
		return &Result{
			res: "",
			error:err,
		}
	}
	var p = make(map[string]interface{})
	for k, v := range configs {
		p[k.(string)] = v
	}
	for i := range paths {
		// 判断类型, 是map继续循环 不是map终止循环
		switch p[paths[i]].(type) {
		case map[string]interface{}:
			p = p[paths[i]].(map[string]interface{})
			r = p
		default:
			r = p[paths[i]]
			break
		}
	}
	return &Result{
		res:r,
		error: nil,
	}
}

func (r *Result)Result() interface{}  {
	return r.res
}

func (r *Result)Int() int  {
	return r.res.(int)
}

func (r *Result)Uint8() uint8  {
	return r.res.(uint8)
}

func (r *Result)String() string {
	return r.res.(string)
}

func (r *Result)Bool() bool  {
	return r.res.(bool)
}

func (r *Result)Error() error {
	return r.error
}