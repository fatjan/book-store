package main

import (
	"book-store/cmd"
)

func main() {
	go cron.Execute()
	cmd.Execute()
}
