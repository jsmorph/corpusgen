package gen

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

var punctuation = regexp.MustCompile(`[^-a-zA-Z0-9]+`)

type Dims struct {
	Atoms     int
	Bytes     int
	AtomBytes int
	Depth     int
	Width     int
}

const (
	DimsCSVHeader = "atoms,bytes,atombytes,depth,width"
	DataCSVHeader = "op,payload"
)

func (d *Dims) CSV() string {
	return fmt.Sprintf("%d,%d,%d,%d,%d",
		d.Atoms,
		d.Bytes,
		d.AtomBytes,
		d.Depth,
		d.Width,
	)
}

func JSON(x interface{}) string {
	js, _ := json.Marshal(&x)
	return string(js)
}

func Depth(x interface{}) int {
	switch vv := x.(type) {
	case map[string]interface{}:
		max := 0
		for _, v := range vv {
			if d := Depth(v); max < d {
				max = d
			}
		}
		return max + 1
	case []interface{}:
		max := 0
		for _, v := range vv {
			if d := Depth(v); max < d {
				max = d
			}
		}
		return max + 1
	default:
		return 0
	}

}

func Width(x interface{}) int {
	switch vv := x.(type) {
	case map[string]interface{}:
		max := len(vv)
		for _, v := range vv {
			if d := Depth(v); max < d {
				max = d
			}
		}
		return max
	case []interface{}:
		max := len(vv)
		for _, v := range vv {
			if d := Depth(v); max < d {
				max = d
			}
		}
		return max
	default:
		return 0
	}

}

func ComputeDims(s string) *Dims {
	var x interface{}
	if err := json.Unmarshal([]byte(s), &x); err != nil {
		panic(err)
	}
	d := &Dims{
		Bytes: len(s),
		Depth: Depth(x),
		Width: Width(x),
	}
	s = punctuation.ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)
	for _, atom := range strings.Split(s, " ") {
		d.Atoms++
		d.AtomBytes += len(atom)
	}

	return d
}
