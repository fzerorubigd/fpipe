package main

import (
	"fmt"
	"log"
)

func errorCheck(err error) {
	if (err != nil) {
		log.Fatal(err)
	}
}

func main() {
	gc, err := parseCmdLine()
	errorCheck(err)

	if gc.Mode == "add" {
		// Its time to add an user entry
		var fc FriendConfiguration
		if !gc.ForceAdd && hasFriendConfiguration(gc.Friend) {
			//There is already a friend config file
			err := fmt.Errorf("There is a config for this friend, for overwriting this use -force flag")
			errorCheck(err)
		}
		fc.Port = gc.Port
		fc.Ip = gc.Ip

		saveFriendConfiguration(gc.Friend, fc)

		log.Printf("Added new friend '%s'", gc.Friend)

		return
	} else if gc.Mode == "server" {
		fc, err := loadFriendConfiguration(gc.Friend)
		errorCheck(err)

		errorCheck(initializeServer(*fc, *gc))
	} else if gc.Mode == "client" {
		fc, err := loadFriendConfiguration(gc.Friend)
		errorCheck(err)
		errorCheck(initializeClient(*fc, *gc))
	}
}
