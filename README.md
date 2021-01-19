<div align="center">

# mockcmd

Mocks the `exec.Command` interface in Golang.

[![Build](https://github.com/Justintime50/mockcmd/workflows/build/badge.svg)](https://github.com/Justintime50/mockcmd/actions)
[![Coverage Status](https://coveralls.io/repos/github/Justintime50/mockcmd/badge.svg?branch=main)](https://coveralls.io/github/Justintime50/mockcmd?branch=main)
[![Licence](https://img.shields.io/github/license/justintime50/mockcmd)](LICENSE)

<img src="assets/showcase.png" alt="Showcase">

</div>

Mocking the `exec.Command` interface in Golang is an absolute pain. You run the risk of running system commands in a unit test context or are required to build out a custom solution to run your commands in an isolated and safe manner which still isn't ideal. Enter `mockcmd`, the easy and safe way to mock your `exec.Command` functions.

`mockcmd` accomplishes a safe and easy mocking of the `exec.Command` interface by creating a fake `cmd` function and passing your command and args through to it. We assert that the `Stdout` gets passed through and returned correctly for success and that an `err` gets returned in the case of failure.

## Install

```bash
go get github.com/justintime50/mockcmd
```

Once `mockcmd` is installed in your project, simply import it:

```go
package yours

import (
	"github.com/justintime50/mockcmd/mockcmd"
	...
)

...
```

## Usage

**Your Command**

Below is an example on how to setup your CMD command in a way it can be easily mocked:

```go
// Pass in `exec.Command` as the context for your real command
func main() {
	out, _ := myCommandFunction(exec.Command)
	fmt.Println(out.String())
}

func myCommandFunction(cmdContext mockcmd.ExecContext) (*bytes.Buffer, error) {
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
```

**Your Test**

Below is an example on how to setup your test to mock your command from above:

```go
// Mock a success and failure response from exec.Command
func TestMyCommandFunctionSuccess(t *testing.T) {
	stdout, err := mockedCmd(mockcmd.MockExecSuccess)
	mockcmd.Success(t, stdout, err)
}

func TestMyCommandFunctionFailure(t *testing.T) {
	_, err := mockedCmd(mockcmd.MockExecFailure)
	mockcmd.Fail(t, err)
}

// The following functions are required in each package that will mock an exec.Command
// Do not alter the following when placed into your package_test.go file
func TestMockProcessSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Stdout.WriteString("mocked Stdout")
	os.Exit(0)
}

func TestMockProcessFailure(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Exit(1)
}
```

## Development

```bash
# Build the project
make build

# Install the project globally from source
make install

# Clean the executables
make clean

# Test the project
make test

## Get test coverage
make coverage

# Lint the project (requires golangci-lint be installed)
make lint
```

## Attribution

* Inspired by the articles by [Jamie Thompson](https://jamiethompson.me/posts/Unit-Testing-Exec-Command-In-Golang) and [Nate Finch](https://npf.io/2015/06/testing-exec-command/)
* Icons made by <a href="https://www.flaticon.com/authors/pixel-perfect" title="Pixel perfect">Pixel perfect</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a>
