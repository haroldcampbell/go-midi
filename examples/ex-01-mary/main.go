package main

import (
	"fmt"
	k "go-midi/core"
)

func main() {
	const framerate = 300

	seq := k.NewSequencer(framerate, framerate/2.0)
	if seq == nil {
		fmt.Printf("Failed to start sequencer\n")
		return
	}

	k.StopAll(seq)

	seq.Add([]k.S{{"E", 1, k.N4}})
	seq.Add([]k.S{{"D", 1, k.N4}})
	seq.Add([]k.S{{"C", 1, k.N4}})
	seq.Add([]k.S{{"D", 1, k.N4}})
	seq.Add([]k.S{{"E", 1, k.N4}})
	seq.Add([]k.S{{"E", 1, k.N4}})
	seq.Add([]k.S{{"E", 1, k.N2}})

	seq.Add([]k.S{{"_", 1, k.N2}})

	seq.Add([]k.S{{"D", 1, k.N4}})
	seq.Add([]k.S{{"D", 1, k.N4}})
	seq.Add([]k.S{{"D", 1, k.N2}})

	seq.Add([]k.S{{"_", 1, k.N4}})

	seq.Add([]k.S{{"E", 1, k.N4}})
	seq.Add([]k.S{{"G", 1, k.N4}})
	seq.Add([]k.S{{"G", 1, k.N2}})

	seq.Add([]k.S{{"_", 1, k.N2}})

	seq.Add([]k.S{{"E", 1, k.N4}})
	seq.Add([]k.S{{"D", 1, k.N4}})
	seq.Add([]k.S{{"C", 1, k.N4}})
	seq.Add([]k.S{{"D", 1, k.N4}})
	seq.Add([]k.S{{"E", 1, k.N4}})
	seq.Add([]k.S{{"E", 1, k.N4}})
	seq.Add([]k.S{{"E", 1, k.N2}})

	seq.Add([]k.S{{"_", 1, k.N4}})

	seq.Add([]k.S{{"E", 1, k.N4}})
	seq.Add([]k.S{{"D", 1, k.N4}})
	seq.Add([]k.S{{"D", 1, k.N4}})
	seq.Add([]k.S{{"E", 1, k.N4}})
	seq.Add([]k.S{{"D", 1, k.N4}})
	seq.Add([]k.S{
		{"C", 1, k.N0},
		{"E", 1, k.N0},
		{"G", 1, k.N0},
	})

	seq.Add([]k.S{{"_", 1, k.N0}})
	seq.Add([]k.S{{"_", 1, k.N0}})

	seq.Play()
	seq.WaitToEnd()
}
