package diff

import (
	"unicode/utf8"

	"github.com/mikeschinkel/go-lib"
)

type pair struct {
	*tree
	og1  string
	og2  string
	s1   string
	s2   string
	opts *CompareOpts
}

func newPair(s1, s2 string, opts *CompareOpts) *pair {
	return &pair{
		og1:  s1,
		s1:   s1,
		og2:  s2,
		s2:   s2,
		opts: opts,
		tree: newTree(opts),
	}
}

func (p *pair) compareStrings() (s string) {
	var ok bool

	if s, ok = p.handleEmptyString(); !ok {
		goto end
	}
	p = p.findPrefixes()
	p = p.findSuffixes()
	p = p.findInfixes()
end:
	return p.String()
}

// findPrefixes finds the initial suffixes. This could be handled by logic in
// findInfixes, but then the logic for trimming the prefixes to pad length
// becomes much more complicated
func (p *pair) findPrefixes() *pair {

	n1, n2 := 0, 0
	s1 := p.s1
	s2 := p.s2
	prefix := newNode(p.opts)
	for {
		if len(s1) == 0 {
			goto end
		}
		if len(s2) == 0 {
			goto end
		}
		r1, sz1 := utf8.DecodeRuneInString(s1)
		if r1 == utf8.RuneError {
			lib.Panicf("ERROR: Attempting to retrieve last rune in '%s'", s1)
		}
		r2, sz2 := utf8.DecodeRuneInString(s2)
		if r2 == utf8.RuneError {
			lib.Panicf("ERROR: Attempting to retrieve last rune in '%s'", s2)
		}
		if r1 != r2 {
			goto end
		}
		prefix.AddBoth(string(r1))
		s1 = s1[sz1:]
		n1 += sz1
		s2 = s2[sz2:]
		n2 += sz2
	}
end:
	p.s1 = s1
	p.s2 = s2

	pad := p.opts.MatchingPadLen.Value
	// Trim the prefix if longer than the pad amount.
	if pad > 0 && len(prefix.both) > pad {
		prefix.both = prefix.both[len(prefix.both)-pad:]
	}

	p.prefix = prefix
	return p
}

// findSuffixes finds the initial suffixes. This could be handled by logic in
// findInfixes, but then the logic for trimming the suffixes to pad length
// becomes much more complicated
func (p *pair) findSuffixes() *pair {
	n1, n2 := 0, 0
	s1 := p.s1
	s2 := p.s2
	suffix := newNode(p.opts)
	for {
		if len(s1) == 0 {
			goto end
		}
		if len(s2) == 0 {
			goto end
		}
		r1, sz1 := utf8.DecodeLastRuneInString(s1)
		if r1 == utf8.RuneError {
			lib.Panicf("ERROR: Attempting to retrieve last rune in '%s'", s1)
		}
		r2, sz2 := utf8.DecodeLastRuneInString(s2)
		if r2 == utf8.RuneError {
			lib.Panicf("ERROR: Attempting to retrieve last rune in '%s'", s2)
		}
		if r1 != r2 {
			goto end
		}
		suffix.InsertBoth(string(r1))
		s1 = s1[:len(s1)-sz1]
		n1 += sz1
		s2 = s2[:len(s2)-sz2]
		n2 += sz2
	}
end:
	p.s1 = p.s1[:len(p.s1)-n1]
	p.s2 = p.s2[:len(p.s2)-n2]

	pad := p.opts.MatchingPadLen.Value
	// Trim the prefix if longer than the pad amount.
	if pad > 0 && len(suffix.both) > pad {
		suffix.both = suffix.both[:pad]
	}

	p.suffix = suffix
	return p
}

func (p *pair) findInfixes() *pair {
	ft := p.opts.findInfixes(p.s1, p.s2)
	p.infix = ft
	return p
}

// handleEmptyString upfront handles empty strings on left, right or both. This
// could be handled by logic in findInfixes, but then the logic for trimming the
// suffixes to pad length becomes much more complicated.
func (p *pair) handleEmptyString() (s string, ok bool) {
	switch {
	case len(p.s1) == 0 && len(p.s2) == 0:
		goto end
	case len(p.s1) == 0:
		p.prefix.(*node).AddRight(p.s2)
		goto end
	case len(p.s2) == 0:
		p.prefix.(*node).AddLeft(p.s1)
		goto end
	default:
		ok = true
	}
end:
	return s, ok
}
