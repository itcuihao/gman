package main

import (
	"flag"
	"fmt"
	"os"

	chc "github.com/itcuihao/gman/commands"
)

func main() {
	os.Exit(realMain())
}

var com string

func realMain() int {
	flags := flag.NewFlagSet(gman, flag.ExitOnError)
	flags.Usage = func() { printUsage() }

	flags.StringVar(&com, "command", "", "linux command")

	if err := flags.Parse(os.Args[1:]); err != nil {
		flags.Usage()
		return 1
	}

	if len(flags.Args()) > 1 {
		flags.Usage()
		return 1
	}

	v := chc.GetCommand(flags.Arg(0))

	if len(v) > 0 {
		fmt.Fprintf(os.Stderr, v)
	} else {
		printUsage()
	}

	return 0
}

func printUsage() {
	fmt.Fprintf(os.Stderr, helpText)
}

const helpText = `
Usage: gman [options]

Linux Command Assistant

eg:
	gman ls 

source:
	http://www.cnblogs.com/peida/tag/每日一linux命令/
`

const gman = "gman"
