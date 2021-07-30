package main

import (
	"fmt"
	c "go-midi/core"
)

func main() {
	const framerate = 800

	seq := c.NewSequencer(framerate, framerate/2.0)
	if seq == nil {
		fmt.Printf("Failed to start sequencer\n")
		return
	}

	c.StopAll(seq)
	seq.Add([]c.S{{"C", 1, c.N4}})

	seq.RepeatPlay()
	seq.WaitToEnd()
}
