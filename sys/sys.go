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
	messages string
	title    string
}
type darwin struct {
	messages string
	title    string
}

func (d *darwin) Execute() error {

	notification := fmt.Sprintf("display notification \"%s\" with title \"%s\"", d.messages, d.title)
	_, err := exec.Command("osascript", "-e", notification).Output()
	return err

}
func (l *linux) Execute() error {

	notification := fmt.Sprintf("notify-send \"%s\" \"%s\"", l.title, l.messages)
	_, err := exec.Command("sh", "-c", notification).Output()
	return err

}

//Command returns a function for notification execution.
func Command(messages string, title string) (OpSys, error) {

	if runtime.GOOS == "darwin" {
		return &darwin{messages, title}, nil
	}
	if runtime.GOOS == "linux" {
		return &linux{messages, title}, nil
	}
	return nil, fmt.Errorf("os not configured")
}
