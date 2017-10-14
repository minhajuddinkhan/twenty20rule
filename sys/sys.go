package sys

import (
	"fmt"
	"os/exec"
	"runtime"
)

//OpSys interface for executing notifications from cmd
type OpSys interface {
	Execute() error
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

func (d *darwin) Execute() error {

	notification := fmt.Sprintf("display notification \"%s\" with title \"%s\"", d.messages[d.index], d.title)
	_, err := exec.Command("osascript", "-e", notification).Output()
	return err

}
func (l *linux) Execute() error {

	notification := fmt.Sprintf("notify-send \"%s\" \"%s\"", l.title, l.messages[l.index])
	_, err := exec.Command("sh", "-c", notification).Output()
	return err

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
