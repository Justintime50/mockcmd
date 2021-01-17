package main

import (
	"github.com/justintime50/mockcmd/mockcmd"
	"os"
	"testing"
)

func TestExampleSuccess(t *testing.T) {
	stdout, err := mockedCmd(mockcmd.MockExecSuccess)
	mockcmd.Success(t, stdout, err)
}

func TestMockProcessSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Stdout.WriteString("mocked Stdout")
	os.Exit(0)
}

func TestExampleFailure(t *testing.T) {
	_, err := mockedCmd(mockcmd.MockExecFailure)
	mockcmd.Fail(t, err)
}

func TestMockProcessFailure(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Exit(1)
}
