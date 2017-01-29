package crawlerclient

import ("testing"
        "strings"
)

func TestCrawlerclient(t *testing.T) {
	cases := []struct {
		in , want string
	}{
		{"example.com", "urls found:"},
    {"", "Failed to parse"},
	}

	for _, c := range cases {
		got:= Crawlerclient(c.in)
		if !strings.Contains(got,c.want) {
			t.Errorf("Crawlerclient(%q) contains %q, wanted  %q inside it", c.in, got, c.want)
    }
	}
}
