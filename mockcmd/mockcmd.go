package mockcmd

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

const (
	mockStdout = "mocked Stdout"
)

// ExecContext tells a cmd function what context to run as (real vs mocked)
type ExecContext = func(name string, arg ...string) *exec.Cmd

// MockExecSuccess is a function that initialises a new exec.Cmd.
// This only serves to call TestMockProcessSuccess rather than the real CMD.
// It will also pass through the command and its arguments as an argument to TestMockProcessSuccess.
func MockExecSuccess(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestMockProcessSuccess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

// MockExecFailure is a function that initialises a new exec.Cmd.
// This only serves to call TestMockProcessFailure rather than the real CMD.
// It will also pass through the command and its arguments as an argument to TestMockProcessFailure.
func MockExecFailure(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestMockProcessFailure", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

// Success handles logging on a cmd test you expect to succeed
func Success(t *testing.T, stdout *bytes.Buffer, err error) bool {
	// Ensure the command was indeed successful
	if err != nil {
		// Don't throw errors on internal mockcmd tests
		if os.Getenv("MOCKCMD_INTERNAL_TEST") != "1" {
			t.Error(err)
		}
		return false
	}

	// Check to make sure the stdout is returned properly
	// Note: value matching is not checked since the command is not run
	stdoutStr := stdout.String()
	if stdoutStr != mockStdout {
		// Don't throw errors on internal mockcmd tests
		if os.Getenv("MOCKCMD_INTERNAL_TEST") != "1" {
			t.Errorf("stdout mismatch:\n%s\n vs \n%s", stdoutStr, mockStdout)
		}
		return false
	}
	return true
}

// Fail handles error logging on a cmd test you expect to fail
func Fail(t *testing.T, err error) bool {
	if err == nil {
		// Don't throw errors on internal mockcmd tests
		if os.Getenv("MOCKCMD_INTERNAL_TEST") != "1" {
			t.Errorf("Expected error due to shell command exiting with non-zero exit code")
		}
		return false
	}
	return true
}
