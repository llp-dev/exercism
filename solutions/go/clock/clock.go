package clock

import "fmt"

type Clock struct {
    m int
}

const DayMinutes = 1440

func New(h, m int) Clock {
    totalMinutes := (h*60 + m) % DayMinutes

	if totalMinutes < 0 {
		totalMinutes += DayMinutes
	}

	return Clock{m: totalMinutes}
}

func (c Clock) Add(m int) Clock {
    return New(0, (c.m + m))
}

func (c Clock) Subtract(m int) Clock {
	return New(0, (c.m - m))
}

func (c Clock) String() string {
    return fmt.Sprintf("%02d:%02d", (c.m / 60), (c.m % 60))
}
