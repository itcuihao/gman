package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"

	gc "gman/commands"
	gs "gman/sys"
)

var errCommand = errors.New("command not found")

func main() {
	if err := realMain(); err != nil {
		printUsage()
		os.Exit(1)
	}
}

func realMain() error {

	var (
		v    string
		line int
		err  error
	)

	flags := flag.NewFlagSet(gman, flag.ExitOnError)
	flags.Usage = func() { printRemind(err) }

	err = flags.Parse(os.Args[1:])
	if err != nil {
		flags.Usage()
		return err
	}

	lf := len(flags.Args())
	switch lf {
	case 0:
		return errCommand
	case 1:
		v = gc.GetCommand(flags.Arg(0))
		if len(v) == 0 {
			err = errCommand
			flags.Usage()
			return errCommand
		}
		if err := gs.ExecShell(fmt.Sprintf("echo '%s' | less", v)); err != nil {
			flags.Usage()
			return err
		}
	case 2:
		v = gc.GetCommand(flags.Arg(0))
		if len(v) == 0 {
			err = errCommand
			flags.Usage()
			return errCommand
		}

		line, err = strconv.Atoi(flags.Arg(1))
		if err != nil {
			flags.Usage()
			return err
		}

		if err := gs.ExecRowShell(fmt.Sprintf("echo '%s' | less", v), line); err != nil {
			flags.Usage()
			return err
		}
	default:
		err = errCommand
		flags.Usage()
	}

	return nil
}

func printUsage() {
	fmt.Fprintf(os.Stderr, helpText)
}

func printRemind(err error) {
	fmt.Fprintf(os.Stderr, remindText, err)
}

const helpText = `
Linux Command Assistant

Usage: gman [options] [line]

eg:
	gman ch

	灵感：由于Linux命令繁杂，利用gman能够快速查找对应的命令的中文解释。
	数据来源于网络。
										--我的名字叫浩仔丶
	
eg:
	gman ch 1
	灵感：由于Linux命令繁杂，利用gman能够快速查找对应的命令的中文解释。

source:
	http://www.cnblogs.com/peida/tag/每日一linux命令/
`
const remindText = "Please confirm your input: %v\n"

const gman = "gman"
