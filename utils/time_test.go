package utils

import (
	"testing"
)

func TestSecondsToTime(t *testing.T) {
	times := map[int]string{
		0:    "00:00",
		60:   "01:00",
		61:   "01:01",
		3050: "50:50",
		3600: "60:00",
	}

	for seconds, expectedOutputString := range times {
		if SecondsToTime(seconds) != expectedOutputString {
			t.Errorf("Wrong output encoding. Expected: %v, got: %v", expectedOutputString, SecondsToTime(seconds))
		}
	}
}
