package battery

import "testing"

func TestLowBatteryGateNotifiesOnceUntilChargeRecovers(t *testing.T) {
	var gate LowBatteryGate
	for _, level := range []int{11, 10, 8, 3} {
		want := level == 10
		if got := gate.ShouldNotify(level); got != want {
			t.Errorf("ShouldNotify(%d) = %t, want %t", level, got, want)
		}
	}
	if got := gate.ShouldNotify(20); got {
		t.Fatal("ShouldNotify(20) = true; recovery must only rearm")
	}
	if got := gate.ShouldNotify(10); !got {
		t.Fatal("ShouldNotify(10) = false after recovery; want a new warning")
	}
}

func TestLowBatteryGateWarnsWhenStartedAtCriticalLevel(t *testing.T) {
	var gate LowBatteryGate
	if got := gate.ShouldNotify(7); !got {
		t.Fatal("ShouldNotify(7) = false; want startup warning")
	}
	if got := gate.ShouldNotify(7); got {
		t.Fatal("repeated critical report should not repeat the warning")
	}
}
