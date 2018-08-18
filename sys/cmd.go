package sys

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

// ExceShell 执行
func ExecShell(s string) error {
	cmd := exec.Command("/bin/bash", "-c", s)

	var (
		out    bytes.Buffer
		stderr bytes.Buffer
	)

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err, stderr.String())
		return err
	}
	fmt.Println(out.String())
	return nil
}

// ExecRowShell 按行执行
func ExecRowShell(s string, n int) error {
	cmd := exec.Command("/bin/bash", "-c", s)
	// fmt.Println(cmd.Args)
	// var (
	// 	out    bytes.Buffer
	// 	stderr bytes.Buffer
	// )

	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer out.Close()

	cmd.Start()

	nr := bufio.NewReader(out)

	for i := 0; i < n; i++ {
		l, err := nr.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(l)
	}

	cmd.Wait()

	return nil
}
