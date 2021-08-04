package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

// 作业：封装一个consul操作的代码
type Server struct {
	ID      string
	Name    string
	Address string
	Port    int
}

func init() {
	server := &Server{
		ID:      "go-tcp-server-1",
		Name:    "go-tcp-server",
		Address: "172.100.99.10",
		Port:    3333,
	}
	registeConsul(server)
}

// 服务启动的时候自动祖册到consul
func registeConsul(server *Server) {
	s, _ := json.Marshal(server) // bytes
	// 172.17.0.3  在windows上直接go run 测试通过，未在docker执行测试
	// 可以向任意consul节点进行注册，server或者client节点都可以
	var url string = "http://172.100.99.21:8521/v1/agent/service/register"
	req, _ := http.NewRequest("PUT", url, bytes.NewReader(s)) // bytes.NewReader 将bytes转换成io.reader类型
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	fmt.Println(res)
}

// docker network inspect [netid]
func main() {
	fmt.Println("启动服务端 ： tcp://127.0.0.1:3333")
	// 1. 监听端口 tcp://0.0.0.0:3333  监听的网络主要以本机可用ip为主
	listen, err := net.Listen("tcp", ":3333")
	if err != nil {
		fmt.Println("err : ", err)
		return // return 表示程序结束
	}
	for {
		// 2. 接收客户向服务端建立的连接
		conn, err := listen.Accept() // 可以与客户端建立连接 ， 如果没有连接挂起阻塞状态
		if err != nil {
			fmt.Println("err : ", err)
			return // return 表示程序结束
		}
		// 3. 处理用户的连接信息
		go handler(conn)
	}
}

// 处理用户的连接信息
func handler(c net.Conn) {
	defer c.Close() // 一定要写 ，关闭连接
	for {
		var data [1024]byte // 数组 - 》定义每一次数据读取的量
		//  Read(p []byte) 需要采用切片接收
		n, err := bufio.NewReader(c).Read(data[:])
		if err != nil {
			fmt.Println("err : ", err)
			break
		}
		fmt.Println("n", string(data[:n]))
		// Write(b []byte) (n int, err error)
		c.Write([]byte("hello world i'm is server"))
	}
}
