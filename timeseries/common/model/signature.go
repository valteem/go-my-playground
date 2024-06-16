package model

import "sort"

// Separates label names, label values and other strings used in calculating their combined
// hash value (= signature = fingerprint)
const SeparatorByte byte = 255 // // Bytes between 0xF8 (248) and 0xFF (255) are not used in UTF-8

// Signature of an empty label set
var emptyLabelSignature = hashNew() // should be emptyLabelSetSignature, not just `label`

// Looks like all stuff below (labels, LabelSet, signature, fingerprint etc) duplicate each other, on one hand,
// and are not used where they should be
// LabelSet, for example, is not used anywhere outside `model` package, and is replaced with `map[string]string`

// Returns a signature (fingerprint) for a given label set
// Collisions are possible but unlikely if the number of label sets is small (thank you, CO:) )
func LabelsToSignature(labels map[string]string) uint64 { // labels => labelset.LabelSets (?)
	if len(labels) == 0 {
		return emptyLabelSignature
	}
	labelNames := make([]string, 0, len(labels)) // labelNames => labels.LabelNames
	for labelName := range labels {
		labelNames = append(labelNames, labelName)
	}
	sort.Strings(labelNames)
	sum := hashNew()
	for _, labelName := range labelNames {
		sum = hashAdd(sum, labelName)
		sum = hashAddByte(sum, SeparatorByte)
		sum = hashAdd(sum, labels[labelName])
		sum = hashAddByte(sum, SeparatorByte) // leaves SeparatorByte at the end of `labels`
	}
	return sum
}

// Same as LabelsToSignature() but takes a LabelSet as an argument and returns a Fingerprint instead of uint64
func labelSetToFingerprint(ls LabelSet) Fingerprint {
	if len(ls) == 0 {
		return Fingerprint(emptyLabelSignature)
	}
	labelNames := make(LabelNames, 0, len(ls))
	for labelName := range ls {
		labelNames = append(labelNames, labelName)
	}
	sort.Sort(labelNames) // replaces sort.Strings(), as LabelNames implements Len(), Less(), Swap()
	sum := hashNew()
	for _, labelName := range labelNames {
		sum = hashAdd(sum, string(labelName)) // converting from LabelName type
		sum = hashAddByte(sum, SeparatorByte)
		sum = hashAdd(sum, string(ls[labelName])) // converting from LabelValue
		sum = hashAddByte(sum, SeparatorByte)
	}
	return Fingerprint(sum)
}

// Same as labelSetToFingerprint, but using a faster and less allocation heavy hash function
func labelSetToFastFingerprint(ls LabelSet) Fingerprint {
	if len(ls) == 0 {
		return Fingerprint(emptyLabelSignature)
	}
	var result uint64
	for labelName, labelValue := range ls {
		sum := hashNew()
		sum = hashAdd(sum, string(labelName))
		sum = hashAddByte(sum, SeparatorByte)
		sum = hashAdd(sum, string(labelValue))
		result ^= sum
	}
	return Fingerprint(result)
}
