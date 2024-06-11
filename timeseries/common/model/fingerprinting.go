package model

import (
	"fmt"
	"strconv"
)

// Hash-capable (?) representation of a Metric (?)
// FNV-1a 64bit
type Fingerprint uint64

// Transforms string representation of a number to a Fingerprint
// Returns error if the string does not represent a uint number
func FingerprintFromString(s string) (Fingerprint, error) {
	n, e := strconv.ParseUint(s, 16, 64) // implies hex-as-astring input
	return Fingerprint(n), e
}

// Parses string representation of a number into a Fingerprint
func ParseFingerprint(s string) (Fingerprint, error) {
	n, e := strconv.ParseUint(s, 16, 64)
	if e != nil {
		return 0, e
	}
	return Fingerprint(n), nil
}

// Returns hex-as-string representation of a Fingerprint
func (f Fingerprint) String() string {
	return fmt.Sprintf("%016x", uint64(f))
}

// Represents a collection of Fingerprints.
// Implements sort.Interface{}
type Fingerprints []Fingerprint

func (f Fingerprints) Len() int {
	return len(f)
}

func (f Fingerprints) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f Fingerprints) Swap(i, j int) {
	// It seems that sort.Sort() implementation takes care of validating i, j
	if i < 0 || i >= len(f) || j < 0 || j >= len(f) {
		return
	}
	f[i], f[j] = f[j], f[i]
}

// Represents as set of Fingerprints
type FingerprintSet map[Fingerprint]struct{}

// Returns true if both sets are of the same size and contain same elements
func (fs FingerprintSet) Equal(fs1 FingerprintSet) bool {
	if len(fs) != len(fs1) {
		return false
	}
	for key := range fs {
		if _, ok := fs1[key]; !ok {
			return false
		}
	}
	return true
}

// Returns new set of elements contained in both input sets
func (fs FingerprintSet) Intersection(fs1 FingerprintSet) FingerprintSet {
	l, l1 := len(fs), len(fs1)
	if l == 0 || l1 == 0 {
		return FingerprintSet{}
	}
	subSet, superSet := fs, fs1
	if l > l1 {
		subSet, superSet = fs1, fs

	}
	output := FingerprintSet{}
	for key := range subSet {
		if _, ok := superSet[key]; ok {
			output[key] = struct{}{}
		}
	}
	return output
}
