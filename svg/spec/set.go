package spec

import (
	"bytes"
	"fmt"
)

// usize is either 32 or 64 depending upon the system, 32 or 64 bit.
const usize = 32 << (^uint(0) >> 63)

// An Set is a set of small non-negative integers. Its zero value
// represents the empty set.
type Set struct {
	words []uint
}

// New returns an empty bitvector set.
func New() *Set {
	return &Set{}
}

// Has reports whether the set contains the non-negative value x.
func (s *Set) Has(x svgType) bool {
	word, bit := x/usize, uint(x%usize)
	return uint(word) < uint(len(s.words)) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *Set) Add(x svgType) {
	word, bit := uint(x/usize), uint(x%usize)
	for word >= uint(len(s.words)) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *Set) UnionWith(t *Set) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *Set) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < usize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteString(", ")
				}
				fmt.Fprintf(&buf, "%d", usize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements within the set.
func (s *Set) Len() (count uint) {
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for i := 0; i < usize; i++ {
			if word&(1<<uint(i)) != 0 {
				count++
			}
		}
	}
	return
}

// Remove removes the given integer from the set.
func (s *Set) Remove(x uint) {
	word, bit := x/usize, uint(x%usize)
	s.words[word] ^= uint(1 << bit)
}

// Clear clears the set.
func (s *Set) Clear() {
	s.words = s.words[:0]
}

// Copy makes and returns a copy of the current set.
func (s *Set) Copy() *Set {
	return &Set{words: s.words}
}

// AddAll add all of the positive integers given to the set.
func (s *Set) AddAll(elems ...svgType) {
	for _, x := range elems {
		s.Add(x)
	}
}

// IntersectWith sets s to the intersection between s and t.
func (s *Set) IntersectWith(t *Set) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			break
		}
	}
	s.words = s.words[:len(t.words)]
}

// DifferenceWith sets s to the difference between s and t.
func (s *Set) DifferenceWith(t *Set) {
	for i, word := range s.words {
		if i < len(t.words) {
			s.words[i] = word &^ t.words[i]
		}
	}
}

// SymmetricDifferenceWith sets s to the symmetric difference between s
// and t.
func (s *Set) SymmetricDifferenceWith(t *Set) {
	if len(t.words) > len(s.words) {
		s.words, t.words = t.words, s.words
	}
	for i, tword := range t.words {
		s.words[i] ^= tword
	}
}

// Elems returns a slice containing the content of the set in increasing
// order of magnitude.
func (s *Set) Elems() (set []uint) {
	l := s.Len()
	if uint(cap(set)) < l {
		set = append(set, make([]uint, 0, l)...)
	}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < usize; j++ {
			if word&(1<<uint(j)) != 0 {
				set = append(set, usize*uint(i+j))
			}
		}
	}
	return
}
