package clock

import (
	"fmt"
)

const minsInHr = 60
const hrsInDay = 24
const minsInDay = hrsInDay * minsInHr

type Clock struct {
	minutes int
}

func (c Clock) Add(minutes int) Clock {
	return New(0, c.minutes+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minutes-minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/minsInHr, c.minutes%minsInHr)
}

func New2(hours, minutes int) Clock {
	minutes = (hours*minsInHr + minutes) % minsInDay
	if minutes < 0 {
		minutes += minsInDay
	}
	return Clock{minutes}
}

func New1(hours, minutes int) Clock {
	minutes += hours * minsInHr
	minutes += (1 - minutes/minsInDay) * minsInDay
	minutes -= (minutes / minsInDay) * minsInDay
	return Clock{minutes}
}
