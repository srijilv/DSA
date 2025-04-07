package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	AddSecondsToTimes([]string{"11:25:25"}, 3600)
}

func AddSecondsToTimes(timePoints []string, addedSeconds int) []string {
	newTimes := []string{}

	for _, t := range timePoints {
		parts := strings.Split(t, ":")
		if len(parts) != 3 {
			newTimes = append(newTimes, t) // skip invalid format
			continue
		}

		hours, _ := strconv.Atoi(parts[0])
		minutes, _ := strconv.Atoi(parts[1])
		seconds, _ := strconv.Atoi(parts[2])

		// Convert everything to total seconds
		totalSeconds := hours*3600 + minutes*60 + seconds + addedSeconds

		// Normalize the time to 24-hour format
		totalSeconds = totalSeconds % (24 * 3600) // Keep it within 0-86399

		newH := totalSeconds / 3600
		totalSeconds %= 3600
		newM := totalSeconds / 60
		newS := totalSeconds % 60

		// Format back to "HH:MM:SS"
		newTime := fmt.Sprintf("%02d:%02d:%02d", newH, newM, newS)
		newTimes = append(newTimes, newTime)
	}

	return newTimes
}
