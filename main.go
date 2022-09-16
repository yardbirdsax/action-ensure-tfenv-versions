/*
Copyright © 2022 Joshua Feierman <josh@sqljosh.com>
*/
package main

import (
	"os"

	"github.com/yardbirdsax/ensure-tfenv-versions/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
