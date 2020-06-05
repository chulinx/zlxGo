# zlxGo
> 自己的golang依赖库

``stringfile github.com/chulinx/zlxGo/stringfile``

``color github.com/chulinx/zlxGo/color``

``yaml github.com/chulinx/zlxGo/yaml``

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
	"github.com/chulinx/zlxGo/yaml"
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
#### Terminal 
```go
package main
import (
    "fmt"
    zssh "github.com/chulinx/zlxGo/ssh"
    "golang.org/x/crypto/ssh"
)
func main() {
	sshConfig := &ssh.ClientConfig{
		User: "vagrant",
		Auth: []ssh.AuthMethod{
			ssh.Password("vagrant"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "127.0.0.1:2222", sshConfig)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	err = zssh.NewTerminal(client)
	if err != nil {
		fmt.Println(err)
	}
}
```