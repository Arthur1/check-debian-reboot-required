package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mackerelio/checkers"
)

func main() {
	runner := new(Runner)

	flag.StringVar(&runner.Dir, "dir", "/var/run/", "directory of reboot-required file [for debug]")
	flag.BoolVar(&runner.Critical, "critical", false, "create critical check report when reboot required")
	flag.BoolVar(&runner.Warning, "warning", true, "create warning check report when reboot required")
	flag.Parse()

	status, message := runner.Run()
	checker := checkers.NewChecker(status, message)
	checker.Name = "check-debian-reboot-required"
	checker.Exit()
}

type Runner struct {
	Dir      string
	Critical bool
	Warning  bool
}

func (r *Runner) Run() (status checkers.Status, message string) {
	if r.Critical && r.Warning {
		status = checkers.UNKNOWN
		message = "cannot specify both -critical and -warning"
		return
	}

	statusWhenRebootRequired := checkers.WARNING
	if r.Critical {
		statusWhenRebootRequired = checkers.CRITICAL
	}

	if _, err := os.Stat(filepath.Join(r.Dir, "reboot-required")); err == nil {
		if b, err := os.ReadFile(filepath.Join(r.Dir, "reboot-required.pkgs")); err == nil {
			status = statusWhenRebootRequired
			message = fmt.Sprintf("System restart required because of following packages:\n%s", string(b))
		} else if os.IsNotExist(err) {
			status = statusWhenRebootRequired
			message = "System restart required"
		} else {
			status = checkers.UNKNOWN
			message = err.Error()
		}
	} else if os.IsNotExist(err) {
		status = checkers.OK
	} else {
		status = checkers.UNKNOWN
		message = err.Error()
	}
	return
}
