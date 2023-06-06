package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:12", true},
		{"12:05:57", true},
		{"bad", false},
		{"1:-3:44", false},
		{"", false},
		{"0:59:59", true},
		{"aa:bb:cc", false},
		{"11:22", false},
		{"5:23", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)

		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}
}
