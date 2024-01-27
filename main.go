package main

import (
	"fmt"
)
import "syscall"
import "errors"

func InitServer(addr syscall.Sockaddr) (int, error) {
	fd, SockErr := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if SockErr != nil {
		return 0, errors.New("error when creating socket")
	}
	BindErr := syscall.Bind(fd, addr)
	if BindErr != nil {
		return 0, errors.New("error when binding a socket")
	}

	return fd, nil
}

func main() {
	fmt.Println("Hello world")
	server, SockErr := InitServer(&syscall.SockaddrInet4{Port: 8000, Addr: [4]byte{127, 0, 0, 1}})
	if SockErr != nil {
		return
	}
	listenErr := syscall.Listen(server, 3)
	if listenErr != nil {
		fmt.Println(listenErr)
		return
	}
	for {
		req, sa, acceptErr := syscall.Accept(server)
		if acceptErr != nil {
			fmt.Println(acceptErr)
			return
		}
		data := make([]byte, 0, 1000)
		_, readArr := syscall.Read(req, data)
		if readArr != nil {
			fmt.Println(readArr)
			return
		}
		fmt.Println(string(data[:]))
		response := []byte("HTTP/1.1 200 OK\n\nHello World")

		sendMsgErr := syscall.Sendto(req, response, syscall.MSG_SEND, sa)
		err := syscall.Close(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		if sendMsgErr != nil {
			fmt.Println(sendMsgErr)
			return
		}

	}

}
