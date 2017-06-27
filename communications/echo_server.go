package main

import (
	"net"
	"bufio"
	"fmt"
)

/*
Echo server

Write a simple echo server. Make it listen to TCP port number 8053 on localhost.
It should be able to read a line (up to the newline), echo back that line and then close the connection.

Make the server concurrent so that every request is taken care of in a separate goroutine.
*/

func main()  {
	l, err := net.Listen("tcp", "127.0.0.1:8053")

	if err != nil {
		fmt.Printf("Failure to listen: %s\n", err.Error())
	}

	for {
		if c, err := l.Accept(); err == nil {
			Echo(c)
		}
	}
}

func Echo(c net.Conn)  {
	defer c.Close()

	line, err := bufio.NewReader(c).ReadString('\n')

	if err != nil {
		fmt.Printf("Failure to read: %s\n", err.Error())
		return
	}

	_, err = c.Write([]byte(line))

	if err != nil {
		fmt.Printf("failure to write: %s\n", err.Error())
		return
	}
}
