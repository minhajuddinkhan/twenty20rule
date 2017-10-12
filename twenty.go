package main

import (
	"os/exec"

	"github.com/jasonlvhit/gocron"
)

func main() {

	gocron.Every(20).Minutes().Do(TwentyRule)

	<-gocron.Start()
}

//TwentyRule In the case of the eyes, the rule is to take 20 seconds to look at something 20 feet away (instead of your computer), and repeat this every 20 minutes.
func TwentyRule() {
	cmd := "notify-send TwentyTwentyRules \"Get your eyes off the screen, Nerd.\""
	exec.Command("sh", "-c", cmd).Output()
}
