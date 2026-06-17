package bafflingbirthdays

import (
    "math/rand"
    "time"
)

func SharedBirthday(dates []time.Time) bool {
    seen := make(map[string]bool)
    
	for _, date := range dates {
		key := date.Format("01-02")
		if seen[key] {
			return true
		}
		seen[key] = true
	}
    
	return false
}

func RandomBirthdates(size int) []time.Time {
    dates := make([]time.Time, size)
    
	for i := 0; i < size; i++ {
		dayOfYear := rand.Intn(365)
		dates[i] = time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, dayOfYear)
	}
    
	return dates
}

func EstimatedProbability(size int) float64 {
	if size > 365 {
		return 1.0
	}
    
	if size < 2 {
		return 0.0
	}

	probNoMatch := 1.0
	for i := 0; i < size; i++ {
		probNoMatch *= float64(365-i) / 365.0
	}

	return (1.0 - probNoMatch) * 100.0
}
