package tago

import (
	"testing"
	"time"
)

func TestWithAndWithoutAgo(t *testing.T) {
	format := "2006/01/02 15:04"
	latest := time.Date(2019, time.August, 8, 22, 0, 0, 0, time.Local)
	tests := map[string]struct {
		tago     TAgo
		input    time.Time
		expected string
	}{
		"with": {
			tago: &With{
				Without{
					duration:    DefaultDuration,
					format:      format,
					latest:      latest,
					isLatestSet: true,
				},
			},
			input:    latest.Add(-5 * time.Second),
			expected: "5s ago",
		},
		"without": {
			tago: &Without{
				duration:    DefaultDuration,
				format:      format,
				latest:      latest,
				isLatestSet: true,
			},
			input:    latest.Add(-5 * time.Second),
			expected: "5s",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := test.tago.Ago(test.input)
			if actual != test.expected {
				t.Errorf("unexpected ago string: got %s, expect %s\n", actual, test.expected)
			}
		})
	}
}

func TestAgo(t *testing.T) {
	format := "2006/01/02 15:04"
	latest := time.Date(2019, time.August, 8, 22, 0, 0, 0, time.Local)
	tests := map[string]struct {
		duration time.Duration
		input    time.Time
		expected string
	}{
		"second": {
			duration: DefaultDuration,
			input:    latest.Add(-5 * time.Second),
			expected: "5s",
		},
		"minute": {
			duration: DefaultDuration,
			input:    latest.Add(-5 * time.Minute),
			expected: "5m",
		},
		"hour": {
			duration: DefaultDuration,
			input:    latest.Add(-5 * time.Hour),
			expected: "5h",
		},
		"day": {
			duration: DefaultDuration,
			input:    latest.Add(-5 * 24 * time.Hour),
			expected: "5d",
		},
		"exceeded time": {
			duration: DefaultDuration,
			input:    latest.Add(-8 * 24 * time.Hour),
			expected: "2019/07/31 22:00",
		},
		"day in not-default duration": {
			duration: 4 * 24 * time.Hour,
			input:    latest.Add(-5 * 24 * time.Hour),
			expected: "2019/08/03 22:00",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Logf("latest: %s\n", latest.Format(format))
			t.Logf("input: %s\n", test.input.Format(format))
			actual := ago(test.duration, format, latest, test.input)
			if actual != test.expected {
				t.Errorf("unexpected ago string: got %s, expect %s\n", actual, test.expected)
			}
		})
	}
}
