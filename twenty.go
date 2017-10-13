package main

import (
	"math/rand"
	"os/exec"
	"time"

	"github.com/jasonlvhit/gocron"
)

//RANGE range for rand num
const RANGE int = 4

var messages = map[int]string{
	0: "Get your eyes off the screen, Nerd.",
	1: "You've been eyeing the screen too much, take a break",
	2: "You need to see some green color",
	3: "All work and no play makes Jack dull.",
}

func main() {

	gocron.Every(20).Minutes().Do(TwentyRule)
	<-gocron.Start()
}

//TwentyRule In the case of the eyes, the rule is to take 20 seconds to look at something 20 feet away (instead of your computer), and repeat this every 20 minutes.
func TwentyRule() {

	cmd := "notify-send TwentyTwentyRule " + messages[getRandomNumber()]
	exec.Command("sh", "-c", cmd).Output()
}

func getRandomNumber() int {

	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(RANGE)
}
