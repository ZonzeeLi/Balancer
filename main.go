package main

import (
	"balance/lb"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	insts := make([]*lb.Instance, 0)
	// 自定义添加服务器的地址、数量、实例信息等...
	for i := 0; i < 4; i++ {
		host := fmt.Sprintf("192.168.%d.%d\n", rand.Intn(255), rand.Intn(255))
		insts = append(insts, &lb.Instance{
			Host: host, Port: "8080",
		})
	}
	var balanceName = "random" // 默认随机
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}
	for {
		inst, err := lb.Select(balanceName, insts)
		if err != nil {
			fmt.Println("do lb err:", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
