package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/mackerelio/checkers"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	cases := map[string]struct {
		r                  *Runner
		st                 checkers.Status
		msg                string
		rebootRequired     *string
		rebootRequiredPerm fs.FileMode
	}{
		"OK if reboot-required file does not exist": {
			r: &Runner{
				Dir: t.TempDir(), Warning: true,
			},
			st:  checkers.OK,
			msg: "",
		},
		"Warning if reboot-required file exists": {
			r: &Runner{
				Dir: t.TempDir(),
			},
			st:                 checkers.WARNING,
			msg:                "System restart required",
			rebootRequired:     toPtr("a"),
			rebootRequiredPerm: 0666,
		},
		"Warning if reboot-required file exists and warning flag is true": {
			r: &Runner{
				Dir: t.TempDir(), Warning: true,
			},
			st:                 checkers.WARNING,
			msg:                "System restart required",
			rebootRequired:     toPtr("a"),
			rebootRequiredPerm: 0666,
		},
		"Critical if reboot-required file exists and critical flag is true": {
			r: &Runner{
				Dir: t.TempDir(), Critical: true,
			},
			st:                 checkers.CRITICAL,
			msg:                "System restart required",
			rebootRequired:     toPtr("a"),
			rebootRequiredPerm: 0666,
		},
		"Unknown if both warning & critical flag is true": {
			r: &Runner{
				Dir: "", Warning: true, Critical: true,
			},
			st:  checkers.UNKNOWN,
			msg: "cannot specify both -critical and -warning",
		},
	}
	for name, tt := range cases {
		if tt.rebootRequired != nil {
			if err := os.WriteFile(
				filepath.Join(tt.r.Dir, "reboot-required"),
				[]byte(*tt.rebootRequired),
				tt.rebootRequiredPerm,
			); err != nil {
				t.Error(err)
			}
		}
		t.Run(name, func(t *testing.T) {
			st, msg := tt.r.Run()
			assert.Equal(t, tt.st, st)
			assert.Equal(t, tt.msg, msg)
		})
	}
}

func toPtr[T any](v T) *T {
	return &v
}
