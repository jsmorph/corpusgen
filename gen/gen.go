package gen

// It's so tempting to get carried away with interfaces for general
// probability distributions.  I WILL resist this time.  I WILL.

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Int struct {
	Min, Max int
}

func (s *Int) Sample() int {
	return rand.Intn(s.Max-s.Min) + s.Min
}

type Float struct {
	Min, Max float64
}

func (s *Float) Sample() float64 {
	return 0
}

type Map struct {
	NumProperties Int
	Properties    String
}

type Decays struct {
	Map   float64
	Array float64
}

func (d *Decays) Decay(r *Decays) *Decays {
	return &Decays{
		Map:   d.Map * r.Map,
		Array: d.Array * r.Array,
	}
}

// var bab = (func() babble.Babbler {
// 	b := babble.NewBabbler()
// 	b.Count = 1
// 	return b
// })()

func (s *Map) Sample(v *Value, decays *Decays) map[string]interface{} {
	n := s.NumProperties.Sample()
	acc := make(map[string]interface{}, n)
	for i := 0; i < n; i++ {
		p := s.Properties.Sample()
		// p := bab.Babble()
		acc[p] = v.Sample(decays)
	}
	return acc
}

type Array struct {
	Length Int
}

func (s *Array) Sample(v *Value, decays *Decays) []interface{} {
	n := s.Length.Sample()
	acc := make([]interface{}, n)
	for i := 0; i < n; i++ {
		acc[i] = v.Sample(decays)
	}
	return acc
}

type Char struct{}

var Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (s *Char) Sample() byte {
	i := rand.Intn(len(Chars))
	return Chars[i]
}

var char = &Char{}

type String struct {
	Length      Int
	Cardinality int

	acc []string
}

func (s *String) Sample() string {
	if 0 < s.Cardinality && s.Cardinality <= len(s.acc) {
		return s.acc[rand.Intn(len(s.acc))]
	}
	n := s.Length.Sample()
	acc := make([]byte, n)
	for i := range acc {
		acc[i] = char.Sample()
	}
	x := string(acc)

	if 0 < s.Cardinality {
		if s.acc == nil {
			s.acc = make([]string, 0, s.Cardinality)
		}
		s.acc = append(s.acc, x)
	}
	return x
}

type Value struct {
	Map, Array, Int, Float, Bool, Null, String float64

	Maps    Map
	Strings String
	Ints    Int
	Arrays  Array
	Decays  Decays
}

func (v *Value) Copy() *Value {
	// Don't copy non-exported fields, which can be shamelessly
	// used in unholy ways.
	js, _ := json.Marshal(v)
	var acc Value
	json.Unmarshal(js, &acc)
	return &acc
}

func (s *Value) Sample(decays *Decays) interface{} {
	if decays == nil {
		decays = &Decays{1, 1}
	}
	x := rand.Float64()
	f := s.Map * decays.Map
	if x < f {
		return s.Maps.Sample(s, decays.Decay(&s.Decays))
	}
	f += s.Array * decays.Array
	if x < f {
		return s.Arrays.Sample(s, decays.Decay(&s.Decays))
	}
	f += s.Int
	if x < f {
		return s.Ints.Sample()
	}

	return s.Strings.Sample()
}

type Trimmer struct {
	Map   float64
	Array float64
}

func (p *Trimmer) Trim(x interface{}) interface{} {
	switch vv := x.(type) {
	case map[string]interface{}:
		acc := make(map[string]interface{})
		for k, v := range vv {
			if rand.Float64() < p.Map {
				continue
			}
			acc[k] = p.Trim(v)
		}
		if len(acc) == 0 {
			for k, v := range vv {
				acc[k] = p.Trim(v)
				break
			}
		}
		return acc
	case []interface{}:
		acc := make([]interface{}, 0, len(vv))
		for _, v := range vv {
			if rand.Float64() < p.Array {
				continue
			}
			acc = append(acc, p.Trim(v))
		}
		return acc
	default:
		return x
	}
}

type Leaf struct {
	Strings    String
	Ints       Int
	Additional Int
	Match      bool
}

func (l *Leaf) ArrayifyAtom(x interface{}) interface{} {
	more := l.Additional.Sample()

	acc := make([]interface{}, 0, more+1)
	if l.Match {
		acc = append(acc, x)
	}
	switch x.(type) {
	case string:
		for i := 0; i < more; i++ {
			acc = append(acc, l.Strings.Sample())
		}
	case int:
		for i := 0; i < more; i++ {
			acc = append(acc, l.Ints.Sample())
		}
	case float64:
		// ToDo: Floats.
		for i := 0; i < more; i++ {
			acc = append(acc, l.Ints.Sample())
		}
	}

	return acc
}

func (l *Leaf) Arrayify(x interface{}) interface{} {
	switch vv := x.(type) {
	case map[string]interface{}:
		acc := make(map[string]interface{})
		for k, v := range vv {
			acc[k] = l.Arrayify(v)
		}
		return acc
	case []interface{}:
		if len(vv) == 0 {
			return nil
		}
		// If there's a map in here, deal with it recursively
		// and ignore the rest.
		for _, v := range vv {
			if m, is := v.(map[string]interface{}); is {
				return l.Arrayify(m)
			}
		}
		// If there's an array in here, ignore it.  Return a
		// subset of the array of atoms.
		atomics := make([]interface{}, 0, len(vv))
		for _, v := range vv {
			if _, is := v.([]interface{}); is {
				continue
			}
			atomics = append(atomics, v)
		}
		if 0 == len(atomics) {
			return nil
		}
		want := rand.Intn(len(atomics)) + 1
		acc := make([]interface{}, 0, want)
		for _, i := range rand.Perm(len(vv)) {
			v := vv[i]
			acc = append(acc, v)
			if len(acc) == want {
				break
			}
		}

		return acc
	default:
		return l.ArrayifyAtom(vv)
	}
}

func (v *Value) GenerateEvent(requireMap bool) ([]byte, error) {
	for i := 0; i < 100; i++ {
		x := v.Sample(nil)
		if !requireMap || isMap(x) {
			js, err := json.Marshal(&x)
			if err != nil {
				return nil, err
			}

			return js, nil
		}
	}
	return nil, fmt.Errorf("GenerateEvent failed to produce a map")
}

func (t *Trimmer) DerivePattern(l *Leaf, event []byte) (string, error) {
	var e interface{}
	if err := json.Unmarshal(event, &e); err != nil {
		return "", err
	}
	p := l.Arrayify(t.Trim(e))
	js, err := json.Marshal(&p)
	if err != nil {
		return "", err
	}
	return string(js), nil

}

func isMap(x interface{}) bool {
	_, is := x.(map[string]interface{})
	return is
}
