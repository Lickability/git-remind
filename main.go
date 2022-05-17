package main

import (
	"github.com/Lickability/git-remind/app/cli"
	"os"
)

func main() {
	cli.App.Run(os.Args)
}
