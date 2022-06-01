package corpusgen

import (
	"context"

	"github.com/timbray/quamina/core"
)

type Check struct {
	PatternError string `json:",omitempty"`
	Matches      bool
	MatchError   string `json:",omitempty"`
}

type LegalPattern struct {
	Quamina     *Check
	EventBridge *Check `json:",omitempty"`
}

func CheckPattern(ctx context.Context, p, e string, eb bool) *LegalPattern {
	lp := &LegalPattern{}

	{
		m := core.NewCoreMatcher()
		err := m.AddPattern(1, p)
		xs, merr := m.MatchesForJSONEvent([]byte(e))
		serr := ""
		if err != nil {
			serr = err.Error()
		}
		smerr := ""
		if merr != nil {
			smerr = merr.Error()
		}
		lp.Quamina = &Check{
			PatternError: serr,
			Matches:      0 < len(xs),
			MatchError:   smerr,
		}
	}

	if eb {
		matches, err := EventBridgeMatches(ctx, nil, p, e)
		serr := ""
		if err != nil {
			serr = err.Error()
		}
		lp.EventBridge = &Check{
			PatternError: serr,
			Matches:      matches,
			MatchError:   serr,
		}

	}

	return lp

}
