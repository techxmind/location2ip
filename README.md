# 命令行
```
go get github.com/techxmind/location2ip/cmd/loc2ip

loc2ip -h
loc2ip -n 5 北京
```

# 示例
```
import (
	"fmt"

	"github.com/techxmind/location2ip"
)


var locs = []string{
    "成都",
    "四川",
    "156007000004", // 位置编码：www.ipcaa.org
}

for loc := range locs {
    if ip, err = location2ip.Generate(loc); err == nil {
        fmt.Println("ip: %s", ip)
    }
}
```

