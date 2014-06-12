package main

import (
	"net"
	"fmt"
	"log"
	"os"
	"io"
)

type Server struct {
	ln       net.Listener
	accepted bool
}

func initializeServer(fc FriendConfiguration, gc GlobalConfig) error {
	var (
		sv Server
		err error
	)

	if gc.Ip != "" {
		//Overwrite address
		fc.Ip = gc.Ip
	}
	if gc.Port > 0 {
		fc.Port = gc.Port
	}
	sv.ln, err = net.Listen("tcp", fmt.Sprintf(":%d", fc.Port))
	sv.accepted = false

	if err != nil {
		return err
	}

	for {
		conn, err := sv.ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Its not a real concurrency, so this is not a go routine
		err = handleRequest(conn)
		if err == nil {
			sv.accepted = true // Ok I know its very cheap, but I have plans for this ;)
		} else {
			log.Print(err)
		}
		
		if sv.accepted {
			break
		}
	}

	return sv.ln.Close()
}

func handleRequest(conn net.Conn) error {
	buf := make([]byte, 1024)

	for {
		len, err := os.Stdin.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if (len > 0) {
			conn.Write(buf[:len])
		}

		if len < 1024 {
			break
		}
	}

	return conn.Close()
}
