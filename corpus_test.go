package corpusgen

import (
	"fmt"
	"testing"

	quamina "github.com/timbray/quamina/core"
	"github.com/timbray/quamina/pruner"
)

func TestCorpus(t *testing.T) {
	s, e := Defaults()

	e.Matcher = quamina.NewCoreMatcher()

	f := func() {
		d, _, err := s.Exec(e)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%T elapsed %s\n", e.Matcher, d)
	}

	f()

	e.Matcher = pruner.NewMatcher(nil)

	f()
}
