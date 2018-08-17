package sys

import (
	"bytes"
	"log"
	"os/exec"
)

// ExceShell 执行
func ExceShell(s string) error {
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
	log.Println(out.String())
	return nil
}
