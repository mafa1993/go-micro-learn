package main

import (
	"fmt"
	"net"
)

// tcp客户端
func main() {
	// 1. 创建建立连接
	conn, _ := net.Dial("tcp", "172.100.100.100:3333")
	// conn, _ := net.Dial("tcp", "192.168.169.204:3000")
	fmt.Println("与tcp://172.100.100.100:3333建立连接")
	defer conn.Close()
	// 2. 进行数据的发送&接收数据
	conn.Write([]byte("你好 server"))
	var data [1024]byte
	n, _ := conn.Read(data[:])
	fmt.Println("n : ", string(data[:n])) // 切片获取信息
	// 3. 关闭 // 不关闭不会造成太大 ，如果服务端没有心跳会存在问题
}
