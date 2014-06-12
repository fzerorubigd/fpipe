package main

import (
	"flag"
	"fmt"
)

type GlobalConfig struct {
	Mode       string
	Friend     string
	Ip         string
	Port       int
	ForceAdd   bool
}

func parseCmdLine() (*GlobalConfig, error) {
	var gc GlobalConfig

	flag.StringVar(&gc.Mode, "mode", "client", "Recive data(client), Send data (server), add friend (add)")
	flag.StringVar(&gc.Friend, "friend", "", "The target friend")
	flag.StringVar(&gc.Ip, "address", "", "The friend address")
	flag.IntVar(&gc.Port, "port", 0, "The port to use with this friend, must set this port on the other side too")
	flag.BoolVar(&gc.ForceAdd, "force", false, "Force overwrite the friend config file in add?")
	flag.Parse()

	if gc.Friend == "" { //TODO : validate friend name
		err := fmt.Errorf("The friend name is mandatory")
		return nil, err
	}

	if gc.Mode != "client" && gc.Mode != "server" && gc.Mode != "add" {
		err := fmt.Errorf("The mode is invalid must be one of server, client, add")
		return nil, err
	}

	if gc.Mode == "add" {
		if gc.Port <= 0 || gc.Ip == "" {
			err := fmt.Errorf("For adding a friend you should set the port and address")
			return nil, err
		}
	}
	return &gc, nil
}
