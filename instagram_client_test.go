package gosaic

import (
	"testing"
)

func TestUrl(t *testing.T) {
	i := InstagramClient{"client123"}

	cases := []struct {
		searchTerm, expectedUrl string
	}{
		{"dogs", "https://api.instagram.com/v1/tags/dogs/media/recent?client_id=client123&count=50"},
		{"star wars", "https://api.instagram.com/v1/tags/starwars/media/recent?client_id=client123&count=50"},
		{"spock/kirk", "https://api.instagram.com/v1/tags/spockkirk/media/recent?client_id=client123&count=50"},
		{"/etc/passwd", "https://api.instagram.com/v1/tags/etcpasswd/media/recent?client_id=client123&count=50"},
		{"../foo/bar", "https://api.instagram.com/v1/tags/foobar/media/recent?client_id=client123&count=50"},
	}

	for _, c := range cases {
		actual, err := i.instagramAPIUrl(c.searchTerm)
		if err != nil {
			t.Fatal(err)
		}

		if actual != c.expectedUrl {
			t.Fatalf("instagramAPIUrl(%q) => expected: %s, actual: %s", c.searchTerm, c.expectedUrl, actual)
		}
	}
}
