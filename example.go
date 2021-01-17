package main

import (
	"bytes"
	"fmt"
	"github.com/justintime50/mockcmd/mockcmd"
	"os/exec"
)

func main() {
	out, _ := mockedCmd(exec.Command)
	fmt.Println(out.String())
}

func mockedCmd(cmdContext mockcmd.ExecContext) (*bytes.Buffer, error) {
	cmd := cmdContext("echo", "Hello World")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to run command: %s", err)
		return nil, err
	}

	return &outb, nil
}
