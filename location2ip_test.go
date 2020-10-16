package location2ip

import (
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		geo string
		err error
	}{
		{
			geo: "156002000000",
		},
		{
			geo: "中国",
		},
		{
			geo: "山东",
		},
		{
			geo: "北京",
		},
		{
			geo: "成都",
		},
		{
			geo: "日本",
		},
		{
			geo: "invalid",
			err: ErrInvalidGeo,
		},
	}

	for _, test := range tests {
		_, err := Generate(test.geo)
		if err != test.err {
			t.Errorf("error mismatch:%s", test.geo)
		}
	}
}
