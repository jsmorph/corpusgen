package corpusgen

// It's so tempting to get carried away with interfaces for general
// probability distributions.  I WILL resist this time.  I WILL.

import (
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

type Value struct {
	Map, Array, Int, Float, Bool, Null, String float64

	Maps    Map
	Strings String
	Ints    Int
	Arrays  Array
	Decays  Decays
}

func (v *Value) Copy() *Value {
	acc := *v
	return &acc
}

type String struct {
	Length Int
}

type Char struct{}

func (s *Char) Sample() byte {
	chars := "abcdefghijklmnopqrstuvwxyz"
	i := rand.Intn(len(chars))
	return chars[i]
}

var char = &Char{}

func (s *String) Sample() string {
	n := s.Length.Sample()
	acc := make([]byte, n)
	for i := range acc {
		acc[i] = char.Sample()
	}
	return string(acc)
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

func Arrayify(x interface{}) interface{} {
	switch vv := x.(type) {
	case map[string]interface{}:
		acc := make(map[string]interface{})
		for k, v := range vv {
			acc[k] = Arrayify(v)
		}
		return acc
	case []interface{}:
		if len(vv) == 0 {
			return nil
		}
		// If there's a map in here, just return it.
		for _, v := range vv {
			if m, is := v.(map[string]interface{}); is {
				return Arrayify(m)
			}
		}
		// If there's an array in here, ignore it.
		// Return a subset of the array of atoms.
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
		return []interface{}{vv}
	}
}
