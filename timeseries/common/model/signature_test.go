package model

import (
	"bytes"
	"hash"
	"hash/fnv"
	"sort"
	"testing"
)

// Computes fingerprint of map[string[string using FNV hash function from standard library
func labelsToSignatureStd(labels map[string]string) uint64 {
	if len(labels) == 0 {
		return emptyLabelSignature
	}
	labelNames := make([]string, 0, len(labels))
	for labelName := range labels {
		labelNames = append(labelNames, labelName)
	}
	sort.Strings(labelNames)
	var b bytes.Buffer
	var h hash.Hash64 = fnv.New64a()
	for _, labelName := range labelNames {
		b.WriteString(labelName)
		b.WriteByte(SeparatorByte)
		b.WriteString(string(labels[labelName]))
		b.WriteByte(SeparatorByte)
		h.Write(b.Bytes())
		b.Reset()
	}
	return h.Sum64()
}

func TestLabelsToSignature(t *testing.T) {
	tests := []struct {
		input map[string]string
	}{
		{
			input: map[string]string{"one ring": "to bring then all", "and": "in the darkness bind them"},
		},
	}
	for _, tc := range tests {
		signatureCustom := LabelsToSignature(tc.input)
		signatureStd := labelsToSignatureStd(tc.input)
		if signatureCustom != signatureStd {
			t.Errorf("should be equal: custom %d, std %d", signatureCustom, signatureStd)
		}
	}
}

func TestLabelsToSignatureSortOrder(t *testing.T) {
	tests := []struct {
		input1 map[string]string
		input2 map[string]string
	}{
		{
			input1: map[string]string{"first": "first", "second": "second"},
			input2: map[string]string{"second": "second", "first": "first"},
		},
		{
			input1: map[string]string{"_first": "first_", "second": "second"},
			input2: map[string]string{"second": "second", "_first": "first_"},
		},
		{
			input1: map[string]string{"fi_rst": "fir_st", "se_cond": "seco_nd"},
			input2: map[string]string{"se_cond": "seco_nd", "fi_rst": "fir_st"},
		},
	}
	for _, tc := range tests {
		signature1 := LabelsToSignature(tc.input1)
		signature2 := LabelsToSignature(tc.input2)
		if signature1 != signature2 {
			t.Errorf("%v to signature -%d, %v to signature - %d", tc.input1, signature1, tc.input2, signature2)
		}
	}
}

func labelSetToFingerprintStd(ls LabelSet) Fingerprint {
	if len(ls) == 0 {
		return Fingerprint(emptyLabelSignature)
	}
	labelNames := make(LabelNames, 0, len(ls))
	for labelName := range ls {
		labelNames = append(labelNames, labelName)
	}
	sort.Sort(labelNames)
	var b bytes.Buffer
	var h hash.Hash64 = fnv.New64a()
	for _, labelName := range labelNames {
		b.WriteString(string(labelName))
		b.WriteByte(SeparatorByte)
		b.WriteString(string(ls[labelName]))
		b.WriteByte(SeparatorByte)
		h.Write(b.Bytes())
		b.Reset()
	}
	return Fingerprint(h.Sum64())
}

func TestLabelSetToFingerprint(t *testing.T) {
	tests := []struct {
		input LabelSet
	}{
		{
			input: LabelSet{"one ring": "to bring them all", "and": "in the darkness bind them"},
		},
	}
	for _, tc := range tests {
		signatureCustom := labelSetToFingerprint(tc.input)
		signatureStd := labelSetToFingerprintStd(tc.input)
		if signatureCustom != signatureStd {
			t.Errorf("should be equal: custom %d, std %d", signatureCustom, signatureStd)
		}
	}
}

func TestLabelSetToFingerprintSortOrder(t *testing.T) {
	tests := []struct {
		input1 LabelSet
		input2 LabelSet
	}{
		{
			input1: LabelSet{"first": "first", "second": "second"},
			input2: LabelSet{"second": "second", "first": "first"},
		},
		{
			input1: LabelSet{"_first": "first_", "second": "second"},
			input2: LabelSet{"second": "second", "_first": "first_"},
		},
		{
			input1: LabelSet{"fi_rst": "fir_st", "se_cond": "seco_nd"},
			input2: LabelSet{"se_cond": "seco_nd", "fi_rst": "fir_st"},
		},
	}
	for _, tc := range tests {
		fingerprint1 := labelSetToFingerprint(tc.input1)
		fingerprint2 := labelSetToFingerprintStd(tc.input2)
		if fingerprint1 != fingerprint2 {
			t.Errorf("%v to fingerprint -%d, %v to fingerprint - %d", tc.input1, fingerprint1, tc.input2, fingerprint2)
		}
	}
}

func TestCompareFingerprints(t *testing.T) {
	tests := []struct {
		input LabelSet
	}{
		{
			LabelSet{"one ring": "to bring them all", "and in the darkness": "bind them"},
		},
	}
	for _, tc := range tests {
		fingerprintRegular := labelSetToFingerprint(tc.input)
		fingerprintFast := labelSetToFastFingerprint(tc.input)
		if fingerprintRegular == fingerprintFast {
			t.Errorf("%v: regular fingerprint %d, fast fingerprint %d (should not be equal)", tc.input, fingerprintRegular, fingerprintFast)
		}
	}
}
