package main

import (
	"flag"
	"github.com/sdkman/sdk/steps"
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/fatih/color"
)

var opt = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

// integrate godog with `go test`
func TestMain(m *testing.M) {

	color.NoColor = true

	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		steps.CliContext(s)
		steps.EnvContext(s)
		steps.VersionFeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
