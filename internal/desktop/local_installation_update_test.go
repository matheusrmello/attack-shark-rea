package desktop

import (
	"context"
	"io/fs"
	"testing"

	packaged "github.com/blak0p/attack-shark-linux/packaging/udev"
)

func TestLocalInstallationUpdatesInspectWithoutUpdaterOrApply(t *testing.T) {
	for _, tc := range []struct {
		name   string
		source policyFixture
		want   string
	}{
		{"current", policyFixture{data: []byte(packaged.Policy())}, "current"},
		{"missing", policyFixture{missing: true}, "missing"},
		{"custom", policyFixture{data: []byte("custom rule")}, "custom"},
		{"unreadable", policyFixture{readError: fs.ErrPermission}, "error"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := &Service{}
			ConfigureLocalInstallationUpdates(s, &tc.source)
			info, err := s.CheckUnifiedUpdates(context.Background())
			if err != nil || info == nil {
				t.Fatalf("local check: info=%+v err=%v", info, err)
			}
			if !info.LocalInstallation || info.PolicyState != tc.want || info.AppAvailable || info.AppInstalled || info.PolicyInstalled {
				t.Fatalf("unexpected local result: %+v", info)
			}
			if (info.PolicyError != "") != (tc.want == "error") {
				t.Fatalf("inspection error not preserved: %+v", info)
			}
			if _, err := s.ApplyUnifiedUpdate(context.Background()); err == nil {
				t.Fatal("local installation must not authorize automatic updates")
			}
			if tc.source.calls != 0 || tc.source.checks != 0 {
				t.Fatal("local inspection invoked remote check or privileged apply")
			}
			if tc.want == "current" {
				tc.source.missing = true
				again, err := s.CheckUnifiedUpdates(context.Background())
				if err != nil || again.PolicyState != "missing" {
					t.Fatalf("recheck did not inspect current files: %+v, %v", again, err)
				}
			}
		})
	}
}
