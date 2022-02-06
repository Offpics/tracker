package utils

import "fmt"

func getTimeString(value int) string {
	if value < 10 {
		return appendZero(value)
	} else {
		return fmt.Sprintf("%v", value)
	}
}

func appendZero(value int) string {
	return fmt.Sprintf("0%v", value)
}

func SecondsToTime(seconds int) string {
	minutes := seconds / 60
	remainingSeconds := seconds % 60

	return fmt.Sprintf("%v:%v", getTimeString(minutes), getTimeString(remainingSeconds))
}
