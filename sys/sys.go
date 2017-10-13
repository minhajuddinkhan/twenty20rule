package sys

import (
	"fmt"
	"os/exec"
	"runtime"
)

//OpSys interface for executing notifications from cmd
type OpSys interface {
	Execute()
}

type linux struct {
	messages map[int]string
	title    string
	index    int
}
type darwin struct {
	messages map[int]string
	title    string
	index    int
}

func (d *darwin) Execute() {

	notification := fmt.Sprintf("'display notification \"%s\" with title \"%s\"'", d.messages[d.index], d.title)
	exec.Command("osascript", "-e", notification).Output()

}
func (l *linux) Execute() {

	notification := fmt.Sprintf("notify-send \"%s\" \"%s\"", l.title, l.messages[l.index])
	exec.Command("sh", "-c", notification).Output()

}

//Command returns a function for notification execution.
func Command(messages map[int]string, num int, title string) (OpSys, error) {

	if runtime.GOOS == "darwin" {
		return &darwin{messages, title, num}, nil
	}
	if runtime.GOOS == "linux" {
		return &linux{messages, title, num}, nil
	}
	return nil, fmt.Errorf("os not configured")
}
