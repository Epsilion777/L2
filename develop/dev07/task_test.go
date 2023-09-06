package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	name         string
	workingTimes []time.Duration
	expectResult int
	timeType     string
}

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	tests := []testStruct{
		{"№1", []time.Duration{3 * time.Second, time.Hour, 2 * time.Second}, 2, "Second"},
		{"№2", []time.Duration{3 * time.Second, time.Second, 15 * time.Second}, 1, "Second"},
		{"№3", []time.Duration{5 * time.Minute, 4 * time.Minute, 3 * time.Minute}, 3, "Minute"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()

			<-Or(
				sig(tt.workingTimes[0]),
				sig(tt.workingTimes[1]),
				sig(tt.workingTimes[2]),
			)
			switch tt.timeType {
			case "Second":
				assert.Equal(t, tt.expectResult, int(time.Since(start).Seconds()))
			case "Minute":
				assert.Equal(t, tt.expectResult, int(time.Since(start).Minutes()))
			default:
				t.Errorf("Unexpected type of time")
			}
		})
	}
}
