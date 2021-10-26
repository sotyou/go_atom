package atom

import "testing"

func TestConvertCamel(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ok", args{"abc_def"}, "AbcDef"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.args.in); got != tt.want {
				t.Errorf("ConvertCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ok",args{"AbcDef"}, "abc_def" },
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCase(tt.args.in); got != tt.want {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHost(t *testing.T) {
	type args struct {
		rawUrl string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"http", args{"https://www.google.co.jp/health"}, "www.google.co.jp"},
		{"wss", args{"wss://google.co.jp/health"}, "google.co.jp"},
		{"ng", args{"wss://"}, ""},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHost(tt.args.rawUrl); got != tt.want {
				t.Errorf("GetHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatToRoundString(t *testing.T) {
	type args struct {
		f float64
		p int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"RoundUp", args{ f: 9.0089, p: 2, }, "9.01"},
		{"RoundDown", args{ f: 9.0101, p: 2 }, "9.01"},
		{"RoundDown", args{ f: 9.1, p: 2 }, "9.10"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatToRoundString(tt.args.f, tt.args.p); got != tt.want {
				t.Errorf("FloatToRoundString() = %v, want %v", got, tt.want)
			}
		})
	}
}