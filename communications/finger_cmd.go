package main

import (
	"flag"
	"net"
	"bufio"
	"os/user"
	"io/ioutil"
	"errors"
)

/*
Finger daemon

Write a finger daemon that works with the finger(1) command.

From the Debian package description:

Fingerd is a simple daemon based on RFC 1196 that provides an interface to the "finger" program at most network sites.
The program is supposed to return a friendly, human-oriented status report on either the system at the moment or a
particular person in depth.

Stick to the basics and only support a username argument. If the user has a .plan file show the contents of that file.
So your program needs to be able to figure out:

Does the user exist?
If the user exists, show the contents of the .plan file.
*/

func main()  {
	flag.Parse()
	ln, err := net.Listen("tcp", ":79")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	usr, _, _ := reader.ReadLine()

	if info, err := getUserInfo(string(usr)); err != nil {
		conn.Write([]byte(err.Error()))
	} else {
		conn.Write(info)
	}
}

func getUserInfo(usr string) ([]byte, error) {
	u, e := user.Lookup(usr)

	if e != nil {
		return nil, e
	}

	data, err := ioutil.ReadFile(u.HomeDir + ".plan")
	if err != nil {
		return data, errors.New("User doesn't have a .plan file!\n")
	}

	return data, nil
}
