package main

import "testing"

func TestHumanSize(t *testing.T) {
	cases := []struct {
		b    int64
		want string
	}{
		{0, "0 B"},
		{512, "512 B"},
		{1024, "1.0 KiB"},
		{1536, "1.5 KiB"},
		{1048576, "1.0 MiB"},
	}
	for _, c := range cases {
		if got := humanSize(c.b); got != c.want {
			t.Errorf("humanSize(%d) 期望 %q 得到 %q", c.b, c.want, got)
		}
	}
}
