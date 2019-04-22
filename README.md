# SDKMAN! Golang CLI

[![Build Status](https://travis-ci.org/sdkman/sdk.svg?branch=master)](https://travis-ci.org/sdkman/sdk)

**This project has been downgraded to a mere spike, the learnings of which will lead to the next phase of exploration of the CLI. We are also looking at Rust as an alternative implementation solution.**

The current bash implementation is still completely fit for purpose and remains in full production use.

### Living Documentation

SDKMAN's development is always driven by tests. More than that, we use Cucumber to describe the behaviour of the CLI in plain English. We do so using Cucumber features, all of which can be found under the [_features/_](features) folder of this repo. These Features form a body of Living Documentation that evolves with the software implementation.

The Cucumber Features are backed by Step Definitions, snippets of matched code that are invoked in order as the Cucumber Feature are 


### Development

#### Prerequisites

Ensure that [Go is installed](https://golang.org/doc/install) on your system.

Optionally, install Godog for running the Cucumber specifications directly:

    $ go get github.com/DATA-DOG/godog/cmd/godog

#### Running the tests

To run all tests using Go's builtin test support (unit and cukes):

    $ cd path/to/the/repo/sdk
    $ go test --godog.format=pretty

If Godog was installed earlier, run the Cucumber specs directly in isolation with the following command:

    $ godog

### Run

To kick the tyres before building:

    $ go run sdk.go version

### Build

To build and run the binary executable:

    $ go build
    $ ./sdk
