# zlxGo
> 自己的golang依赖库

``stringfile  github.com/tianyazc/zlxGo/zlxmodule/stringfile``

``color github.com/tianyazc/zlxGo/zlxmodule/color``

``yaml github.com/tianyazc/zlxGo/zlxmodule/yaml``

## Ex:
#### config file
```yaml
redis:
  server: 127.0.0.1:6379
  db: 0
  pass:
elsticsearch: http://10.6.201.133:49200
logLevel: 1  # 1-->debug 2-->Info 2-->Warn 4-->Error
```

```go
package main

import (
	"fmt"
	"github.com/tianyazc/zlxGo/zlxmodule/yaml"
)

func main() {
	// 指定配置文件路径
	c := yaml.NewConfig("./etc/config.yaml")
	r :=c.Get("redis.server").Result()
	// NewDefaultConfigGet 会读取默认配置文件路径为程序所在目录下的./etc/config.yaml
	e := yaml.NewDefaultConfigGet("elsticsearch").Result()
	fmt.Printf("redis:%s\nes:%s\n",r,e)
}
// 输出
$ ./config          
redis:127.0.0.1:6379
es:http://10.6.201.133:49200
```