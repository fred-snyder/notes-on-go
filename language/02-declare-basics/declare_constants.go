package main

import "fmt"

const (
	speedLimitLow  = 30
	speedLimitHigh = 130
	country        = "Netherlands"
)

// grouped constants // https://talks.golang.org/2012/10things.slide#2
var config struct {
	APIKey string
}

func declareConstants() {
	fmt.Println(speedLimitLow, country)

	config.APIKey = "BADC0C0A"
}
