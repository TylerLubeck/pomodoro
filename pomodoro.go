package main

import (
	"flag"
	"fmt"
	"github.com/deckarep/gosx-notifier"
	"time"
)

const (
	tomato       = "🍅 "
	clearScreen  = "\033[2J\033[0;0H\n"
	formatLine   = "\033[2J\033[0;0H%s: %s\n"
	finishedLine = "\033[2J\033[0;0HYou're done!\n"
)

var pomodoroLength = flag.Int64("length", 25, "The length of the pomodoro in minutes")

func durationToReadableMinutes(duration time.Duration) string {
	hours := duration / time.Hour
	duration -= hours * time.Hour

	minutes := duration / time.Minute
	duration -= minutes * time.Minute

	seconds := duration / time.Second

	return fmt.Sprintf("%d:%d:%d", hours, minutes, seconds)
}

func finishTimer() {
	fmt.Println(finishedLine)
	note := gosxnotifier.NewNotification("You're done!")
	note.Sound = gosxnotifier.Default
	note.Group = "com.tylerlubeck.pomodoro.finished"
	note.Title = "Pomodoro Complete"
	note.Push()
}

func main() {
	flag.Parse()

	duration := time.Duration(*pomodoroLength) * time.Minute

	fmt.Print(clearScreen)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			duration -= (1 * time.Second)

			if duration <= 0 {
				finishTimer()
				return
			}

			fmt.Printf(formatLine, tomato, durationToReadableMinutes(duration))
		}
	}
}
