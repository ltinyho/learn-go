package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	data := []byte("[这里才是一个完整的数据包]")
	conn, err := net.DialTimeout("tcp", "localhost:4044", time.Second*30)
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	for range time.Tick(time.Microsecond*1) {
		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}
