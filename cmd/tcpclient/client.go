package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ilaziness/gokit/server/tcp"
)

// TCPClient 定义TCP客户端结构体
type TCPClient struct {
	conn net.Conn
}

// NewTCPClient 创建一个新的TCP客户端
func NewTCPClient(addr string) (*TCPClient, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &TCPClient{conn: conn}, nil
}

func NewTCPTLSClient(addr string) (*TCPClient, error) {
	conn, err := tls.Dial("tcp", addr, &tls.Config{
		InsecureSkipVerify: true,
	})
	fmt.Println("connect to tls server")
	if err != nil {
		return nil, err
	}
	return &TCPClient{conn: conn}, nil
}

// Start 启动客户端，处理用户输入和服务器响应
func (c *TCPClient) Start(wg *sync.WaitGroup) {
	defer wg.Done()
	// 用于等待读取服务器响应的goroutine完成
	var readWg sync.WaitGroup
	readWg.Add(2) // 一个用于读取服务器响应，另一个用于处理用户输入

	// 单独开一个协程读取服务器发来的消息
	go c.readServer(&readWg)
	go c.heartbeat(&readWg)

	// 处理用户输入
	c.handleUserInput(&readWg)

	// 等待所有操作完成
	readWg.Wait()
}

// readServer 读取服务器发送的消息
func (c *TCPClient) readServer(wg *sync.WaitGroup) {
	defer wg.Done()
	codec := tcp.NewPackCodec()
	for {
		pack, err := codec.Decode(c.conn)
		if err != nil {
			fmt.Println("Error reading from server:", err)
			os.Exit(1)
			return
		}
		if pack.Head.OpCode == uint16(tcp.OpCodePong) {
			fmt.Println("Pong received")
			continue
		}
		// 打印服务器返回的数据，使用实际读取的长度
		fmt.Printf("op code: %d, Message from server: %s\n", pack.Head.OpCode, pack.Payload)
	}
}

// heartbeat 发送心跳包
func (c *TCPClient) heartbeat(wg *sync.WaitGroup) {
	defer wg.Done()
	codec := tcp.NewPackCodec()
	tk := time.NewTicker(time.Second * 20)
	for {
		select {
		case <-tk.C:
			pack := &tcp.Pack{
				Head: tcp.PackHead{
					Len:    uint32(10),
					SQID:   0,
					OpCode: uint16(tcp.OpCodePing),
				},
			}
			err := codec.Encode(c.conn, pack)
			if err != nil {
				fmt.Println("Error sending data to server:", err)
			}
		}
	}
}

// handleUserInput 处理用户输入
func (c *TCPClient) handleUserInput(wg *sync.WaitGroup) {
	defer wg.Done()
	inputReader := bufio.NewReader(os.Stdin)
	log.Println("Enter text (type 'quit' to exit):")
	codec := tcp.NewPackCodec()
	for {
		// 读取用户输入
		input, _ := inputReader.ReadString('\n')
		// 去除首尾空格
		input = strings.TrimSpace(input)

		// 如果用户输入'quit'，则退出
		if input == "quit" {
			fmt.Println("Closing connection...")
			c.conn.Close()
			return
		}

		pack := &tcp.Pack{
			Head: tcp.PackHead{
				Len:    uint32(len(input) + 10),
				SQID:   0,
				OpCode: uint16(1000),
			},
			Payload: []byte(input),
		}

		// 发送数据到服务器
		err := codec.Encode(c.conn, pack)
		if err != nil {
			fmt.Println("Error sending data to server:", err)
			return
		}
	}
}

func main() {
	// 创建TCP客户端，连接本地8080端口
	// client, err := NewTCPClient("127.0.0.1:8080")
	client, err := NewTCPTLSClient("127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	// 启动客户端
	go client.Start(&wg)

	// 等待所有操作完成
	wg.Wait()
}
