package model

import (
	"fmt"
	"sort"
)

type LabelSet map[LabelName]LabelValue

// Return `nil` if all label names and values are valid
func (ls LabelSet) Validate() error {
	for ln, lv := range ls {
		if !ln.IsValid() {
			return fmt.Errorf("invalid label name %q", ln) // %q escapes strings and adds quotes
		}
		if !lv.IsValid() {
			return fmt.Errorf("invalid label value %q", lv)
		}
	}
	return nil
}

// Returns `true` if all key/value pairs in both label sets are the same
func (ls LabelSet) Equal(ls1 LabelSet) bool {
	if len(ls) != len(ls1) {
		return false
	}
	for ln, lv := range ls {
		lv1, ok := ls1[ln]
		if !ok {
			return false
		}
		if lv1 != lv {
			return false
		}
	}
	return true
}

// Compares two label sets
func (ls LabelSet) Before(ls1 LabelSet) bool {
	if len(ls) < len(ls1) {
		return true
	} else if len(ls) > len(ls1) {
		return false
	} else {
		lns := make(LabelNames, 0, len(ls)+len(ls1))
		for ln := range ls {
			lns = append(lns, ln)
		}
		for ln := range ls1 {
			lns = append(lns, ln)
		}
		sort.Sort(lns)
		for _, ln := range lns {
			lv, ok := ls[ln]
			if !ok { // label name is not in ls, ls is before ls1
				return true
			}
			lv1, ok := ls1[ln]
			if !ok { // label name is in ls, and is not in ls1, ls is not before ls1
				return false
			}
			if lv < lv1 {
				return true
			}
			if lv > lv1 {
				return false
			}
		}
	}
	return false // ls is not before ls1 if all names and values are equal
}

// Returns a copy of the label set
func (ls LabelSet) Clone() LabelSet {
	ls1 := make(LabelSet, len(ls))
	for ln, lv := range ls {
		ls1[ln] = lv
	}
	return ls1
}

// Merges two label sets
func (ls LabelSet) Merge(ls1 LabelSet) LabelSet {
	lsnew := make(LabelSet, len(ls))
	for ln, lv := range ls {
		lsnew[ln] = lv
	}
	for ln, lv := range ls1 {
		lsnew[ln] = lv // value from ls1 overrides value for the same name in ls
	}
	return lsnew
}
