package model

import (
	"testing"
)

func TestMetricEqual(t *testing.T) {
	tests := []struct {
		name    string
		m1      Metric
		m2      Metric
		outcome bool
	}{
		{
			name:    "Equal non-empty metrics",
			m1:      Metric{"label1": "value1", "label2": "value2", "label3": "value3"},
			m2:      Metric{"label1": "value1", "label2": "value2", "label3": "value3"},
			outcome: true,
		},
		{
			name:    "Non-equal non-empty metrics",
			m1:      Metric{"label1": "value1", "label2": "value2", "label3": "value3"},
			m2:      Metric{"label1": "value1", "label2": "value2"},
			outcome: false,
		},
		{
			name:    "Equal empty metrics",
			m1:      Metric{},
			m2:      Metric{},
			outcome: true,
		},
	}
	for _, tc := range tests {
		if outcome := tc.m1.Equal(tc.m2); outcome != tc.outcome {
			t.Errorf("%s: Equal() should be %t, get %t", tc.name, tc.outcome, outcome)
		}
	}
}

func TestMetricBefore(t *testing.T) {
	tests := []struct {
		name    string
		a       Metric
		b       Metric
		outcome bool
	}{
		{
			name:    "a before b - different length",
			a:       Metric{"label1": "value1", "label2": "value2"},
			b:       Metric{"label1": "value1", "label2": "value2", "label3": "value3"},
			outcome: true, // len(a) < len(b)
		},
		{
			name:    "a before b - same length",
			a:       Metric{"label1": "value1", "label2": "value2"},
			b:       Metric{"label1": "value1", "label3": "value3"},
			outcome: false,
		},
		{
			name:    "Equal empty metrics",
			a:       Metric{},
			b:       Metric{},
			outcome: false,
		},
	}
	for _, tc := range tests {
		if outcome := tc.a.Before(tc.b); outcome != tc.outcome {
			t.Errorf("%s: Before() should be %t, get %t", tc.name, tc.outcome, outcome)
		}
	}
}

func TestMetricClone(t *testing.T) {
	tests := []struct {
		input Metric
	}{
		{
			input: Metric{"some labels": "include spaces", "some_": "do_not"},
		},
		{
			input: Metric{"singleLabel": "singleValue"},
		},
		{
			input: Metric{},
		},
	}
	for _, tc := range tests {
		if clone := tc.input.Clone(); !clone.Equal(tc.input) {
			t.Errorf("Clone(): clone %v should be equal input %v", clone, tc.input)
		}
	}
}

func TestMetricString(t *testing.T) {
	tests := []struct {
		name   string
		input  Metric
		output string
	}{
		{
			name:   "Metric with name",
			input:  Metric{"label1": "value1", "__name__": "very_important_metric", "label2": "value2"},
			output: "very_important_metric{label1=\"value1\", label2=\"value2\"}",
		},
		{
			name:   "Metric without name",
			input:  Metric{"label1": "value1", "label3": "value3", "label2": "value2"},
			output: "{label1=\"value1\", label2=\"value2\", label3=\"value3\"}",
		},
		{
			name:   "Metric with name only",
			input:  Metric{"__name__": "some_metric_name"},
			output: "some_metric_name",
		},
		{
			name:   "Empty metric",
			input:  Metric{},
			output: "{}",
		},
	}
	for _, tc := range tests {
		if output := tc.input.String(); output != tc.output {
			t.Errorf("String(): get %s, expect %s", output, tc.output)
		}
	}
}
