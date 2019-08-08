package tago

import (
	"fmt"
	"time"
)

const (
	DefaultDuration = 7 * 24 * time.Hour
)

type TAgo interface {
	Ago(time.Time) string
}

func NewWith(d time.Duration, format string) *With {
	return &With{
		*NewWithout(d, format),
	}
}

type With struct {
	Without
}

func (w *With) Ago(t time.Time) string {
	return fmt.Sprintf("%s ago", w.Without.Ago(t))
}

func NewWithout(d time.Duration, format string) *Without {
	return &Without{
		duration: d,
		format:   format,
	}
}

type Without struct {
	duration    time.Duration
	format      string
	latest      time.Time
	isLatestSet bool
}

func (w *Without) Ago(t time.Time) string {
	l := time.Now()
	if w.isLatestSet {
		l = w.latest
	}

	return ago(w.duration, w.format, l, t)
}

func ago(d time.Duration, format string, latest, t time.Time) string {
	if latest.Add(-d).After(t) {
		return t.Format(format)
	}

	diff := latest.Sub(t)
	if 24 <= diff.Hours() {
		return fmt.Sprintf("%dd", int(diff.Hours()/24))
	}
	if 1 <= diff.Hours() {
		return fmt.Sprintf("%dh", int(diff.Hours()))
	}
	if 1 <= diff.Minutes() {
		return fmt.Sprintf("%dm", int(diff.Minutes()))
	}

	return fmt.Sprintf("%ds", int(diff.Seconds()))
}
