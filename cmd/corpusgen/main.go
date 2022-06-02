package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/jsmorph/corpusgen"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	cfg := struct {
		Value   *corpusgen.Value
		Trimmer *corpusgen.Trimmer
	}{
		Value:   corpusgen.DefaultValue,
		Trimmer: corpusgen.DefaultTrimmer,
	}

	var (
		cfgFilename = flag.String("cfgfile", "", "filename for cfg JOSN")
		cfgJSON     = flag.String("cfgjson", "{}", "JSON cfg overlay")
		printCfg    = flag.Bool("cfg", false, "print cfg and exit")
		nonmaps     = flag.Bool("nonmaps", false, "events and patterns don't have to be maps")

		dims             = flag.Bool("dims", false, "emit dimensions")
		numEvents        = flag.Int("events", 10, "number of events")
		patternsPerEvent = flag.Int("pats-per-event", 2, "number of patterns per event")
		numPatterns      = flag.Int("patterns", 5, "number of independent patterns")
		seed             = flag.Int64("seed", time.Now().UnixNano(), "random number generator seed")
	)

	flag.Parse()

	rand.Seed(*seed)

	if *cfgFilename != "" {
		bs, err := ioutil.ReadFile(*cfgFilename)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(bs, &cfg); err != nil {
			return err
		}
	}

	if err := json.Unmarshal([]byte(*cfgJSON), &cfg); err != nil {
		return err
	}

	if *printCfg {
		bs, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", bs)
		return nil
	}

	printDims := func(tag string, x interface{}) {
		if !*dims {
			return
		}
		switch vv := x.(type) {
		case []byte:
			x = string(vv)
		case string:
		default:
		}
		d := corpusgen.ComputeDims(x.(string))
		js, err := json.Marshal(&d)
		if err != nil {

			fmt.Printf("%s dims %#v\n", tag, d)
			return
		}
		fmt.Printf("%s dims %s\n", tag, js)
	}

	for i := 0; i < *numEvents; i++ {
		event, err := cfg.Value.GenerateEvent(!*nonmaps)
		if err != nil {
			return err
		}
		fmt.Printf("event %s\n", event)
		printDims("event", event)
		for j := 0; j < *patternsPerEvent; j++ {
			pattern, err := cfg.Trimmer.DerivePattern(event)
			if err != nil {
				return err
			}
			fmt.Printf("pattern %s\n", pattern)
			printDims("pattern", event)
		}
	}

	for i := 0; i < *numPatterns; i++ {
		event, err := cfg.Value.GenerateEvent(!*nonmaps)
		if err != nil {
			return err
		}
		pattern, err := cfg.Trimmer.DerivePattern(event)
		if err != nil {
			return err
		}
		fmt.Printf("pattern %s\n", pattern)
		printDims("pattern", event)
	}

	return nil
}
