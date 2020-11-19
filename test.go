package main

import (
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
)

func main() {
	username := ""
	apikey := ""
	// 1. Create a session
	sess := session.New(username, apikey)
	// 2. Get a service
	accountService := services.GetAccountService(sess)
	services
	// 3. Invoke a method:
	account, err := accountService.GetObject()


}
