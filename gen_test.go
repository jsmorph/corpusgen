package corpusgen

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDims(t *testing.T) {
	s := `{"drhtuaiz":-39,"xiylvbnxhkpmmidcp":{"loeohqs":[92,-80,"wyuolcwcbdixvddtwgqoyukjrvsk",-49],"lvnd":[["ffjugczdbuciuubbqbokrzuow"],[-94,"wxyftndvcsmjrvp",0,-35],[[92]]],"nafaouhyiyawrpjam":-2}}`

	fmt.Printf("%#v\n", ComputeDims(s))

}

func TestGen(t *testing.T) {
	s := DefaultValue.Copy()

	x := s.Sample(nil)
	js, err := json.MarshalIndent(&x, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("event %s\n", js)

	trimmer := &Trimmer{
		Map:   0.5,
		Array: 0.5,
	}
	x = trimmer.Trim(x)
	js, err = json.MarshalIndent(&x, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("pruned %s\n", js)

	x = Arrayify(x)
	js, err = json.MarshalIndent(&x, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("pattern %s\n", js)

}

func TestValueCopy(t *testing.T) {
	c0 := &Value{}
	c1 := c0.Copy()
	c1.Ints = Int{
		Max: 1,
	}
	if c0.Ints.Max == c1.Ints.Max {
		t.Fatal(c0.Ints, c1.Ints)
	}
}
