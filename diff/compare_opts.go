package diff

import (
	"strings"
)

type CompareOpts struct {
	MatchingPadLen  *Len
	MinSubstrLen    *Len
	LeftRightFormat string
}

// findInfixes finds the strings and substrings after prefixes and suffixes are
// found. It creates a down-growth tree structure where differing prefixes and
// suffixes are found and common values stored in infix property of the `node`
// struct.
func (opts *CompareOpts) findInfixes(s1, s2 string) (ifx fixer) {
	var t *tree
	var pos2 int

	ss, pos1 := longestCommonSubstr(s1, s2)
	pos2 = strings.Index(s2, ss)
	if !opts.hasCommonSubstr(ss) {
		n := newNode(opts)
		n.AddLeft(s1)
		n.AddRight(s2)
		ifx = n
		goto end
	}
	//goland:noinspection GoAssignmentToReceiver
	t = newTree(opts)
	t.prefix = opts.findInfixes(s1[:pos1], s2[:pos2])
	t.infix.(*node).AddBoth(ss)
	if len(s1) > len(ss)+pos1 {
		t.suffix = opts.findInfixes(s1[len(ss)+pos1:], s2[len(ss)+pos2:])
	}
	ifx = t
end:
	return ifx
}

func (opts *CompareOpts) SetDefaults() {
	if opts.MinSubstrLen == nil {
		opts.MinSubstrLen = NewLen(MinSubstrLen)
	}
	if opts.LeftRightFormat == "" {
		opts.LeftRightFormat = LeftRightFormat
	}
	if opts.MatchingPadLen == nil {
		opts.MatchingPadLen = NewLen(0)
	}
}

// hasCommonSubstr returns true is a "common substring" — see `const
// MinSubstrLen` for definition of common substring. Used to decide is we capture
// two strings are left vs. right or attempt to subdivide them again.
func (opts *CompareOpts) hasCommonSubstr(ss string) (has bool) {
	if len(ss) <= opts.MinSubstrLen.Value {
		goto end
	}
	if ss == "" {
		goto end
	}
	if ss == " " {
		goto end
	}
	has = true
end:
	return has
}
