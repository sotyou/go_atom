package atom

import (
	"github.com/k0kubun/pp"
	"reflect"
	"testing"
	"time"
)

func TestMicroseconds(t *testing.T) {
	type args struct {
		start time.Time
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"ok", args{time.Now()}, 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Microseconds(tt.args.start); got != tt.want {
				t.Errorf("Microseconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZeroTime(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
		{"ok", time.Unix(0,0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZeroTime() = %v, want %v", got, tt.want)
			} else {
				pp.Println(got.String())
			}
		})
	}
}

func TestWaitForPeriod(t *testing.T) {
	type args struct {
		period time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{"second", args{period: time.Second}},
		{"second", args{period: 10*time.Second}},
		{"minute", args{period: time.Minute}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WaitForPeriod(tt.args.period)
			pp.Println(time.Now().String())
		})
	}
}

func TestPeriod(t *testing.T) {
	type args struct {
		unit  string
		multi int64
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"ok", args{ unit:  "second", multi: 10 }, 10*time.Second},
		{"ok", args{ unit:  "minute", multi: 10 }, 10*time.Minute},
		{"ok", args{ unit:  "hour", multi: 10 }, 10*time.Hour},
		{"ok", args{ unit:  "day", multi: 10 }, 240*time.Hour},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Period(tt.args.unit, tt.args.multi); got != tt.want {
				t.Errorf("Period() = %v, want %v", got, tt.want)
			}
		})
	}
}