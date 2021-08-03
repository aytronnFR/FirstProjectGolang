package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func searchCommand(args []string) bool {
	search(args)
	return true
}

var userList []string

func search(args []string) {
	if len(args) != 2 {
		info("Invalid argument. search <pseudo>")
		return
	}
	first := args[1]
	if contains(userList, first) {
		fmt.Printf(NoticeColor, "The player \"" + first + "\" is a beta player !")
	} else {
		fmt.Printf(ErrorColor, "The player \"" + first + "\" is a not beta player !")
	}
}

func getBetaUser() {
	defer wg.Done()
	resp, err := http.Get(config.ApiUrl)
	failOnError(err, "Cannot send request !")
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	failOnError(err, "Cannot read the body !")

	//Convert the body to type string
	sb := string(body)

	replace := []string{"[", "]", "\"", "\n", " "}

	for _, entry := range replace {
		sb = strings.ReplaceAll(sb, entry, "")
	}

	userList = strings.Split(sb, ",")
}


func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
