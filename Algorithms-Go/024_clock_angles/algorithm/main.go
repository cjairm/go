// Package algorithm contains different algorithms
package algorithm

import "fmt"

// ClockAngle checks the angle betwen hours hand and minutes hand. PM and AM format
func ClockAngle(h, m int) (int, error) {
	// Check if hours are valid
	if h < 0 || h > 12 {
		return 0, fmt.Errorf("Hours not admited. Has to be greater equal than 0 and lower equal than 12")
	}

	// Check if minutes are valid
	if m < 0 || m > 60 {
		return 0, fmt.Errorf("Minutes not admited. Has to be greater equal than 0 and lower equal than 60")
	}

	// If hours are equal to 12 PM/AM just has to be reset to zero.
	if h == 12 {
		h = 0
	}

	// If minutes are equal to 60 minutes just has to be reset to zero.
	// But if we reset we have to increment one hour and fotmat it to AM/PM.
	if m == 60 {
		m = 0
		h++
		if h > 12 {
			h = h - 12
		}
	}

	// MINUTES: The clock rotates 360 in 60 minutes so each minute represents 6 degrees.
	// HOURS: The clock rotates 360 in 12 hours so each ohur represent 30 degrees.
	// PAST HOURS: current minute divided by total minutes in a hours, then multiplied for 30 degrees (hour degrees)
	// (m / 60) * 30 = (m / 30) * 15 = (m / 6) * 3 = (m / 2) * 1 = m / 2
	bAngle := ((h * 30) + (m / 2)) - (m * 6)

	return bAngle, nil
}
