package atom

import (
	"github.com/k0kubun/pp"
	"testing"
)

func TestRuntimeFunc(t *testing.T) {
	type args struct {
		skip int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ok", args{2}, "TestRuntimeFunc"},
	}
	pp.Println(RuntimeFunc(1))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RuntimeFunc(tt.args.skip); got != tt.want {
				t.Errorf("RuntimeFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
