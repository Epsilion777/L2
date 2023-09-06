package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	name    string
	host    string
	wantErr bool
}

func TestGetTime(t *testing.T) {
	tests := []testStruct{
		{"ok", "0.beevik-ntp.pool.ntp.org", false},
		{"notok", "sdfsdfds.org", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeFromNtp, err := GetTime(tt.host)
			if !tt.wantErr {
				assert.NoError(t, err, "Fetching time from NTP server should not result in an error")
				assert.NotNil(t, timeFromNtp, "Fetched NTP time should not be nil")
			} else {
				assert.Error(t, err)
			}
		})
	}
}
