package mockcmd

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"testing"
)

// NOTE: We don't test the `if os.Getenv("MOCKCMD_INTERNAL_TEST") != "1"` blocks because then tests fail due to calling `t.Error()`
// As such, we won't be able to reach 100% coverage because we can't run the couple lines that are unreachable in a test setting

func TestMockExecSuccess(t *testing.T) {
	mockExecSuccess := MockExecSuccess("echo", "Hello World")

	expectedCmdType := "exec.Cmd"
	if reflect.TypeOf(mockExecSuccess) != reflect.TypeOf(&exec.Cmd{}) {
		t.Errorf("MockExecSuccess is not the expected type %s, got: %s", expectedCmdType, reflect.TypeOf(mockExecSuccess))
	}

	actualTestEnv := mockExecSuccess.Env
	expectedTestEnv := []string{"GO_TEST_PROCESS=1"}
	if actualTestEnv[0] != expectedTestEnv[0] {
		t.Errorf("GO_TEST_PROCESS is not the expected value, got: %s, want: %s", actualTestEnv[0], expectedTestEnv[0])
	}
}

func TestMockExecFailure(t *testing.T) {
	mockExecFailure := MockExecFailure("echo", "Hello World")

	expectedCmdType := "exec.Cmd"
	if reflect.TypeOf(mockExecFailure) != reflect.TypeOf(&exec.Cmd{}) {
		t.Errorf("MockExecSuccess is not the expected type %s, got: %s", expectedCmdType, reflect.TypeOf(mockExecFailure))
	}

	actualTestEnv := mockExecFailure.Env
	expectedTestEnv := []string{"GO_TEST_PROCESS=1"}
	if actualTestEnv[0] != expectedTestEnv[0] {
		t.Errorf("GO_TEST_PROCESS is not the expected value, got: %s, want: %s", actualTestEnv[0], expectedTestEnv[0])
	}
}

// nilErr returns a nil error which is a helper function for some of these tests
func nilErr() (err error) {
	return
}

// TestFailTrue checks that a failure state was successful
func TestFailTrue(t *testing.T) {
	err := errors.New("Some error message")
	fail := Fail(t, err)
	expectedReturn := true
	if fail != expectedReturn {
		t.Errorf("Fail should return %s because an error was returned.", strconv.FormatBool(expectedReturn))
	}
}

// TestFailFalse checks that a failure state was unsuccessful
func TestFailFalse(t *testing.T) {
	os.Setenv("MOCKCMD_INTERNAL_TEST", "1")
	nilErr := nilErr()
	fail := Fail(t, nilErr)
	expectedReturn := false
	if fail != expectedReturn {
		t.Errorf("Fail should return %s because no error was returned.", strconv.FormatBool(expectedReturn))
	}
}

// TestSuccessTrue checks that a success state was successful
func TestSuccessTrue(t *testing.T) {
	mockStdout := bytes.NewBufferString("mocked Stdout")
	nilErr := nilErr()
	success := Success(t, mockStdout, nilErr)
	expectedReturn := true
	if success != expectedReturn {
		t.Errorf("Success should return %s because no errors were returned.", strconv.FormatBool(expectedReturn))
	}
}

// TestSuccessFalseError checks that a success state was unsuccessful due to an error
func TestSuccessFalseError(t *testing.T) {
	mockStdout := bytes.NewBufferString("mocked Stdout")
	err := errors.New("Some error message")
	success := Success(t, mockStdout, err)
	expectedReturn := false
	if success != expectedReturn {
		t.Errorf("Success should return %s because an error was returned.", strconv.FormatBool(expectedReturn))
	}
}

// TestSuccessFalseMismatchedStdout checks that a success state was unsuccessful due to a mismatching Stdout
func TestSuccessFalseMismatchedStdout(t *testing.T) {
	mockStdout := bytes.NewBufferString("mocked Stdout that does not match")
	nilErr := nilErr()
	success := Success(t, mockStdout, nilErr)
	expectedReturn := false
	if success != expectedReturn {
		t.Errorf("Success should return %s because the Stdout mismatched.", strconv.FormatBool(expectedReturn))
	}
}
