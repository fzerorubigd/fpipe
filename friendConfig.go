package main

import (
	"os"
	"os/user"
	"encoding/json"
)

type FriendConfiguration struct {
	Ip   string
	Port int
}

func hasFriendConfiguration(friend string) bool {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	configFile := usr.HomeDir + "/.config/fpipe/" + friend + ".json"

	if _, err := os.Stat(configFile); err == nil {
		return true
	}

	return false
}

func loadFriendConfiguration(friend string) (*FriendConfiguration, error) {
	var config FriendConfiguration
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	configFile := usr.HomeDir + "/.config/fpipe/" + friend + ".json"
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	decoder := json.NewDecoder(file)

	if err = decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func saveFriendConfiguration(friend string, config FriendConfiguration) error {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	configFile := usr.HomeDir + "/.config/fpipe/" + friend + ".json"

	file, err := os.Create(configFile)
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	encoder := json.NewEncoder(file)
	if err = encoder.Encode(config); err != nil {
		return err
	}

	return nil
}
