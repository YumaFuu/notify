package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	for i := 1; i <= 12; i++ {
		for j := 0; j <= 60; j++ {
			for k := 0; k <= 60; k++ {
				var s string

				switch k % 4 {
				case 0:
					s = "|"
				case 1:
					s = "\\"
				case 2:
					s = "-"
				case 3:
					s = "/"
				}
				fmt.Printf("\r%v %v min            ", s, j)
				time.Sleep(1 * time.Second)
			}
		}

		title := "💪" + strconv.Itoa(i) + "時間経過"
		message := "10分間散歩しましょう"

		exec.
			Command("terminal-notifier",
				"-title",
				title,
				"-message",
				message,
				"-sound",
				"Ping").
			Run()
	}
}
