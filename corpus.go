package corpusgen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"sync"
	"time"

	quamina "github.com/timbray/quamina/core"
	"github.com/timbray/quamina/pruner"
)

type nothing struct{}

var na = nothing{}

type Corpus struct {
	MatchingEvents   []string
	MatchingPatterns []string
	OtherEvents      []string
	OtherPatterns    []string
	Spec             *CorpusSpec
}

func (c *Corpus) Copy() *Corpus {
	acc := *c
	return &acc
}

type CorpusSpec struct {
	V                *Value
	Trimmer          *Trimmer
	MatchingEvents   int
	MatchingPatterns int
	OtherEvents      int
	OtherPatterns    int
	PatternIds       int
}

func (s *CorpusSpec) Copy() *CorpusSpec {
	acc := *s
	return &acc
}

func isMap(x interface{}) bool {
	_, is := x.(map[string]interface{})
	return is
}

func (s *CorpusSpec) Gen() (*Corpus, error) {
	corpus := &Corpus{
		MatchingEvents:   make([]string, 0, s.MatchingEvents),
		MatchingPatterns: make([]string, 0, s.MatchingPatterns),
		OtherEvents:      make([]string, 0, s.OtherEvents),
		OtherPatterns:    make([]string, 0, s.OtherPatterns),
		Spec:             s,
	}

	for len(corpus.MatchingEvents) < s.MatchingEvents {

		event := s.V.Sample(nil)
		if !isMap(event) {
			continue
		}
		eventjs, err := json.Marshal(&event)
		if err != nil {
			return nil, err
		}

		corpus.MatchingEvents = append(corpus.MatchingEvents, string(eventjs))
	}
	log.Printf("made %d matching events", len(corpus.MatchingEvents))

MATCHING_PATTERNS:
	for len(corpus.MatchingPatterns) < s.MatchingPatterns {
		for _, eventjs := range corpus.MatchingEvents {
			if len(corpus.MatchingPatterns) >= s.MatchingPatterns {
				break MATCHING_PATTERNS
			}
			var event interface{}
			if err := json.Unmarshal([]byte(eventjs), &event); err != nil {
				return nil, err
			}

			trimmed := s.Trimmer.Trim(event)

			pattern := Arrayify(trimmed)
			if !isMap(event) {
				// js, _ := json.Marshal(&trimmed)
				// log.Printf("not a map %s", js)
				continue
			}

			patternjs, err := json.Marshal(&pattern)
			if err != nil {
				return nil, err
			}
			corpus.MatchingPatterns = append(corpus.MatchingPatterns, string(patternjs))
		}
	}

	log.Printf("made %d matching patterns", len(corpus.MatchingPatterns))

	for len(corpus.OtherEvents) < s.OtherEvents {

		event := s.V.Sample(nil)
		if !isMap(event) {
			continue
		}
		eventjs, err := json.Marshal(&event)
		if err != nil {
			return nil, err
		}

		corpus.OtherEvents = append(corpus.OtherEvents, string(eventjs))
	}

	log.Printf("made %d other events", len(corpus.OtherEvents))

	for len(corpus.OtherPatterns) < s.OtherPatterns {
		event := s.V.Sample(nil)
		if !isMap(event) {
			continue
		}
		trimmed := s.Trimmer.Trim(event)

		pattern := Arrayify(trimmed)
		if !isMap(event) {
			continue
		}

		patternjs, err := json.Marshal(&pattern)
		if err != nil {
			return nil, err
		}
		corpus.OtherPatterns = append(corpus.OtherPatterns, string(patternjs))
	}

	log.Printf("made %d other patterns", len(corpus.OtherPatterns))

	return corpus, nil
}

type OpsMix struct {
	Adds, Matches, Deletes int
}

func (m *OpsMix) Copy() *OpsMix {
	acc := *m
	return &acc
}

func (c *Corpus) Event() string {
	switch n := rand.Intn(len(c.MatchingEvents) + len(c.OtherEvents)); {
	case n < len(c.MatchingEvents):
		return c.MatchingEvents[n]
	default:
		return c.OtherEvents[n-len(c.MatchingEvents)]
	}
}

func (c *Corpus) Pattern() string {
	switch n := rand.Intn(len(c.MatchingPatterns) + len(c.OtherPatterns)); {
	case n < len(c.MatchingPatterns):
		return c.MatchingPatterns[n]
	default:
		return c.OtherPatterns[n-len(c.MatchingPatterns)]
	}
}

type Op struct {
	Add    string    `json:",omitempty"`
	Delete quamina.X `json:",omitempty"`
	Match  string    `json:",omitempty"`
}

func (o *Op) Exec(c *Corpus, m quamina.Matcher, printDims, printData bool) error {
	if o.Add != "" {
		if printDims {
			d := ComputeDims(o.Add)
			fmt.Printf("dims,add,%s\n", d.CSV())
		}
		p := rand.Intn(c.Spec.PatternIds)
		if printData {
			fmt.Printf("data,addPattern,%s\n", JSON(o.Add))
		}
		return m.AddPattern(p, o.Add)
	}
	if o.Delete != nil {
		return m.DeletePattern(o.Delete)
	}
	if printDims {
		d := ComputeDims(o.Match)
		fmt.Printf("dims,match,%s\n", d.CSV())
	}
	if printData {
		fmt.Printf("data,matchEvent,%s\n", JSON(o.Match))
	}
	_, err := m.MatchesForJSONEvent([]byte(o.Match))
	return err
}

func (o *OpsMix) Op(c *Corpus) *Op {
	switch n := rand.Intn(o.Adds + o.Matches + o.Deletes); {
	case n < o.Adds:
		return &Op{
			Add: c.Pattern(),
		}
	case n < o.Adds+o.Matches:
		return &Op{
			Match: c.Event(),
		}
	default:
		return &Op{
			Delete: rand.Intn(c.Spec.PatternIds),
		}
	}

	return nil
}

func (c *Corpus) Dump(filename string) error {
	js, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, js, 0644)
}

type Exec struct {
	Mix        *OpsMix
	Matcher    quamina.Matcher `json:"-"`
	Filename   string
	Goroutines int
	Ops        int
	PrintDims  bool
	PrintData  bool
}

func (e *Exec) Copy() *Exec {
	acc := *e
	return &acc
}

func (s *CorpusSpec) Exec(e *Exec) (time.Duration, *Corpus, error) {
	corpus, err := s.Gen()
	if err != nil {
		return 0, nil, err
	}

	if e.Filename != "" {
		if err = corpus.Dump(e.Filename); err != nil {
			return 0, nil, err
		}
	}

	ops := make([]*Op, e.Ops)

	for i, _ := range ops {
		ops[i] = e.Mix.Op(corpus)
	}

	if e.PrintDims {
		fmt.Printf("dims,op,%s\n", DimsCSVHeader)
	}
	if e.PrintDims {
		fmt.Printf("data,op,%s\n", DataCSVHeader)
	}
	wg := &sync.WaitGroup{}

	then := time.Now()
	for i := 0; i < e.Goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			executed := 0
			want := e.Ops
			for 0 < want {
				i := rand.Intn(len(ops))
				o := ops[i]
				if err := o.Exec(corpus, e.Matcher, e.PrintDims, e.PrintData); err != nil {
					continue
					// ToDo: limit
				}
				executed++
				want--
			}
			log.Printf("%d executed %d ops", i, executed)
			wg.Done()
		}(i)
	}

	wg.Wait()
	elapsed := time.Now().Sub(then)

	if p, is := e.Matcher.(*pruner.Matcher); is {
		s := p.Stats()
		js, _ := json.MarshalIndent(&s, "", "  ")
		fmt.Printf("%s\n", js)
	}

	return elapsed, corpus, nil
}
