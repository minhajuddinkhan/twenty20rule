package main

import (
	"runtime"

	"github.com/jasonlvhit/gocron"
	"github.com/minhajuddinkhan/twenty20rule/sys"
)

//RANGE range for rand num

func main() {

	gocron.Every(1).Second().Do(TwentyRule)
	<-gocron.Start()
}

//TwentyRule In the case of the eyes, the rule is to take 20 seconds to look at something 20 feet away (instead of your computer), and repeat this every 20 minutes.
func TwentyRule() {
	sys.Command(runtime.GOOS)()

}
