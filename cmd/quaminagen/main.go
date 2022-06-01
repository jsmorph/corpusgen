package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	quamina "github.com/timbray/quamina/core"
	"github.com/timbray/quamina/pruner"

	gen "github.com/jsmorph/corpusgen"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

type Cfg struct {
	Spec *gen.CorpusSpec
	Exec *gen.Exec
}

func run() error {
	// Warning: flag action below!

	cfg := Cfg{
		Spec: gen.DefaultCorpusSpec,
		Exec: gen.DefaultExec,
	}

	flag.IntVar(&cfg.Exec.Ops, "ops",
		cfg.Exec.Ops, "ops per goroutine")
	flag.IntVar(&cfg.Exec.Goroutines, "threads",
		cfg.Exec.Goroutines, "number of goroutines")

	var (
		core        = flag.Bool("core", false, "use CoreMatcher instead of Pruner")
		cfgFilename = flag.String("cfgfile", "", "filename for cfg JOSN")
		cfgJSON     = flag.String("cfgjson", "{}", "JSON cfg overlay")
		printCfg    = flag.Bool("cfg", false, "print cfg and exit")
	)

	flag.Parse()

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

	if *core {
		cfg.Exec.Matcher = quamina.NewCoreMatcher()
	} else {
		cfg.Exec.Matcher = pruner.NewMatcher(nil)
	}

	d, _, err := cfg.Spec.Exec(cfg.Exec)
	if err != nil {
		return err
	}
	fmt.Printf("elapsed %s\n", d)

	return nil
}
