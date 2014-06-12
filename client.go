package main

import (
	"net"
	"fmt"
	"io"
	"os"
)

type Client struct {
	conn net.Conn
}

func initializeClient(fc FriendConfiguration, gc GlobalConfig) error {
	var (
		client Client
		err error
	)
	if gc.Ip != "" {
		//Overwrite address
		fc.Ip = gc.Ip
	}
	if gc.Port > 0 {
		fc.Port = gc.Port
	}

	client.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", fc.Ip, fc.Port))
	if err != nil {
		return err
	}

	fmt.Fprintf(client.conn, "HI, FRIEND_PIPE")

	buf := make([]byte, 1024)
	for {
		len, err := client.conn.Read(buf)

		if err != nil && err != io.EOF {
			return err
		}

		if len > 0 {
			os.Stdout.Write(buf[:len])
		}

		if len < 1024 {
			break
		}
	}

	return nil
}
