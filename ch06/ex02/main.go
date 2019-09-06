package main

import (
	"bytes"
	"fmt"
	"math"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}

	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				len++
			}
		}
	}
	return len
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		var removeBit uint64
		removeBit = math.MaxUint64 & ^(1 << bit)
		s.words[word] &= removeBit
	}
}

func (s *IntSet) Clear() {
	for i, _ := range s.words {
		s.words[i] &= 0
	}
}

func (s *IntSet) Copy() *IntSet {
	var dst IntSet
	for _, tword := range s.words {
		dst.words = append(dst.words, tword)
	}
	return &dst
}

func (s *IntSet) AddAll(args ...int) {
	for _, arg := range args {
		s.Add(arg)
	}
}

func main() {
	var x, y IntSet
	x.AddAll([]int{1, 63, 144, 9}...)
	x.Remove(1)
	fmt.Println(x.String())
	x.Clear()
	fmt.Println(x.String())
	fmt.Println(x.Len())
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())
	fmt.Println(y.Len())

	fmt.Println("hoge")
	z := x.Copy()
	fmt.Println(z.String())
	z = y.Copy()
	fmt.Println(z.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Len())

	fmt.Println(x.Has(9), x.Has(123))
}