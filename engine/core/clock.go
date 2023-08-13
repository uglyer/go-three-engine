package core

import (
	"time"
)

type Clock struct {
	autoStart   bool
	startTime   int64
	oldTime     int64
	elapsedTime float64
	running     bool
}

func NewClock(autoStart bool) *Clock {
	clock := &Clock{
		autoStart:   autoStart,
		startTime:   0,
		oldTime:     0,
		elapsedTime: 0,
		running:     false,
	}
	return clock
}

func (c *Clock) Start() {
	c.startTime = now()
	c.oldTime = c.startTime
	c.elapsedTime = 0
	c.running = true
}

func (c *Clock) Stop() {
	c.GetElapsedTime()
	c.running = false
	c.autoStart = false
}

func (c *Clock) GetElapsedTime() float64 {
	c.GetDelta()
	return c.elapsedTime
}

func (c *Clock) GetDelta() float64 {
	diff := 0.0
	if c.autoStart && !c.running {
		c.Start()
		return 0
	}
	if c.running {
		newTime := now()
		diff = float64(newTime-c.oldTime) / 1000
		c.oldTime = newTime
		c.elapsedTime += diff
	}
	return diff
}

func now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
