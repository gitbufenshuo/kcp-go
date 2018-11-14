package main

import (
	"time"

	kcp "github.com/gitbufenshuo/kcp-go"
)

func main() {
	kcpconn, err := kcp.DialWithOptions("192.168.0.1:10000", nil, 10, 3)
	if err != nil {
		panic(err)
	}
	for {
		kcpconn.Write([]byte("hello-world--1"))
		// kcpconn.Write([]byte("hello-world--1"))
		// kcpconn.Write([]byte("hello-world--1"))
		// kcpconn.Write([]byte("hello-world--1"))
		// kcpconn.Write([]byte("hello-world--1"))

		// time.Sleep(time.Second * 10)
		// kcpconn.Write([]byte("hello-world--2"))
		time.Sleep(time.Second * 100000)
	}
}
