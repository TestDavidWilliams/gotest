package main

import (
"fmt"
"github.com/PagerDuty/go-pagerduty"
)

var	authtoken = "y_NbAkKc66ryYTWUXYEu" // Set your auth token here

func main() {
	var opts pagerduty.ListEscalationPoliciesOptions
	client := pagerduty.NewClient(authtoken)
	eps, err := client.ListEscalationPolicies(opts)
	if err != nil {
		panic(err)
	}
	for _, p := range eps.EscalationPolicies {
		fmt.Println(p.Name)
		fmt.Println(p.ID)
		names := p.EscalationRules[0].Targets[0].Summary
		fmt.Println(names)
		break
	}
	
	
}
