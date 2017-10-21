package clock

import "strconv"

const testVersion = 4

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	for hour > 23 || hour < 0 {
		if hour < 0 {
			hour += 24
		} else {
			hour -= 24
		}
	}
	for minute > 59 || minute < 0 {
		if minute < 0 {
			if hour-1 < 0 {
				hour = 23
			} else {
				hour -= 1
			}
			minute += 60
		} else {
			if hour+1 > 23 {
				hour = 0
			} else {
				hour += 1
			}
			minute -= 60
		}
	}
	return Clock{hour, minute}
}

func (c Clock) String() string {
	var time string
	var hour string
	var minute string
	if c.hour < 10 {
		hour = "0" + strconv.Itoa(c.hour)
	} else {
		hour = strconv.Itoa(c.hour)
	}
	if c.minute < 10 {
		minute = "0" + strconv.Itoa(c.minute)
	} else {
		minute = strconv.Itoa(c.minute)
	}
	time = hour + ":" + minute
	return time
}

func (c Clock) Add(minutes int) Clock {
	change := c.minute + minutes
	hour := c.hour
	minute := change
	for minute > 59 || minute < 0 {
		if minute < 0 {
			if hour-1 < 0 {
				hour = 23
			} else {
				hour -= 1
			}
			minute += 60
		} else {
			if hour+1 > 23 {
				hour = 0
			} else {
				hour += 1
			}
			minute -= 60
		}
	}
	return Clock{hour, minute}
}
