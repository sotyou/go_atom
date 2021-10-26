package atom

import "testing"

func TestHmac256(t *testing.T) {
	type args struct {
		base string
		key  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"OK", args{
			base: `GET
api.huobi.pro
/ws/v2
accessKey=yh4fhmvs5k-5f082ed4-76287c7e-0f372&signatureMethod=HmacSHA256&signatureVersion=2.1&timestamp=2020-11-19T00%3A06%3A47`, key: `1d0e2d29-2ff3b40f-10abccaf-13145`,
		}, "T2Lhwy+nHw2GHiSwo4ZJTd6UyXUGFaxtiuWG0N6DqsE="},
		{"OK", args{
			base: `GET
api.huobi.pro
/ws/v2
accessKey=yh4fhmvs5k-5f082ed4-76287c7e-0f372&signatureMethod=HmacSHA256&signatureVersion=2.1&timestamp=2020-11-19T00%3A03%3A01`, key: `1d0e2d29-2ff3b40f-10abccaf-13145`,
		}, "c7LE/8yXtQvVqU4rEPGbvUC7Kb7aRFdUi7mlBpg4X9Q="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hmac256(tt.args.base, tt.args.key); got != tt.want {
				t.Errorf("Hmac256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomString(t *testing.T) {
	type args struct {
		len int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2", args{ len: 2 }, 2},
		{"16", args{ len: 16 }, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomString(tt.args.len)
			if len(got) != tt.want {
				t.Errorf("RandomString() = %v, want %v, got %d", got, tt.want, len(got))
			}
		})
	}
}