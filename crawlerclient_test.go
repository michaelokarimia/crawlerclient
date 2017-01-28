package crawlerclient

import ("testing"
        "strings"
      //  "fmt"
)

func TestCrawlerclient(t *testing.T) {
	cases := []struct {
		in , want string
	}{
		{"example.com", "example"},
    {"", "Failed to parse"},

    //{"http://example.com/", "<!doctype html>"},

	}

	for _, c := range cases {
		got:= Crawlerclient(c.in)
		if !strings.ContainsAny(got,c.want) {
			t.Errorf("Crawlerclient(%q) contains %q, wanted  %q inside it", c.in, got, c.want)
    }
    //fmt.Printf(got)


	}





}
