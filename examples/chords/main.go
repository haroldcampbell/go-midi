package main

import (
	"fmt"
	k "go-midi/core"
	"time"
)

func main() {
	const framerate = 300

	seq := k.NewSequencer(framerate, framerate/2.0)
	if seq == nil {
		fmt.Printf("Failed to start sequencer\n")
		return
	}

	k.StopAll(seq)

	seq.C(1, 50)
	seq.E(1, 50)
	seq.G(1, 50)

	time.Sleep(2 * time.Second)

	seq.F(1, 50)
	seq.A(1, 50)
	seq.D(2, 50)

	time.Sleep(2 * time.Second)

	seq.G(2, 50)
	seq.B(2, 50)
	seq.D(2, 50)

	time.Sleep(2 * time.Second)

	k.StopAll(seq)
}
