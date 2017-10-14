package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/minhajuddinkhan/twenty20rule/sys"
)

const (
	RANGE int    = 4
	TITLE string = "Twenty Twenty Rule"
)

var (
	messages = map[int]string{
		0: "Get your eyes off the screen, Nerd.",
		1: "You've been eyeing the screen too much, take a break",
		2: "You need to see some green color, like plants or something.",
		3: "All work and no play makes Jack dull.",
	}
)

func main() {

	gocron.Every(1).Second().Do(TwentyRule)
	<-gocron.Start()
}

//TwentyRule In the case of the eyes, the rule is to take 20 seconds to look at something 20 feet away (instead of your computer), and repeat this every 20 minutes.
func TwentyRule() {

	command, err := sys.Command(messages, getRandomNumber(), TITLE)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = command.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
func getRandomNumber() int {

	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(RANGE)
}
