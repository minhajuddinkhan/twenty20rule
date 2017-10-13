package sys

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"
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

func Command(os string) func() {

	fmt.Println(os)

	if os == "darwin" {
		return func() {
			fmt.Println("darwin!")
			notification := fmt.Sprintf("display notification \"%s\" with title \"%s\" ", messages[getRandomNumber()], TITLE)
			exec.Command("osascript", "-e", notification)
		}

	}
	return func() {

		fmt.Println("running on linux")
		cmd := "notify-send " + "\"" + TITLE + "\"" + " " + "\"" + messages[getRandomNumber()] + "\""
		exec.Command("sh", "-c", cmd).Output()

	}
}

func getRandomNumber() int {

	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(RANGE)
}
