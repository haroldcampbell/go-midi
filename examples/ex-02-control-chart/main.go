package main

import (
	"fmt"
	k "go-midi/core"
	"math"
)

func main() {

	const framerate = 300

	seq := k.NewSequencer(framerate, framerate/2.0)
	if seq == nil {
		fmt.Printf("Failed to start sequencer\n")
		return
	}

	k.StopAll(seq)

	times := []float64{
		2.7,
		1.2,
		-0.1,
		-0.1,
		-0.1,
		-0.1,
		1.1,
		1.8,
		1.8,
		-0.6,
		-0.4,
		-0.4,
		-0.4,
		0.9,
		-0.4,
		-0.4,
		0.3,
		2.4,
		-1.1,
		0.5,
		1.2,
		1.2,
		2.0,
		-1.1,
		1.2,
		-1.1,
		1.0,
		1.0,
		1.9,
		-0.5,
		1.3,
		0.0,
		-1.1,
		1.6,
		-1.0,
		-0.6,
		-0.9,
		1.0,
		0.9,
		-0.9,
		0.8,
		-0.8,
		-1.0,
		-1.0,
		-0.9,
		-0.9,
		-1.0,
		-0.1,
		-0.4,
		-0.4,
		-0.4,
		-0.4,
		-0.4,
		-0.3,
		-0.3,
		-0.3,
		-0.3,
		-0.7,
		-0.6,
		-0.9,
		-0.9,
		-0.9,
		-0.9,
		-0.9,
		-0.9,
		-0.9,
	}
	// seq.Add([]k.S{{"E", 1, k.N4}})

	for _, ct := range times {
		addWithCycleTime(seq, ct)
	}

	seq.Play()
	seq.WaitToEnd()
}

func addWithCycleTime(seq *k.Sequencer, ct float64) {
	const baseOctave = 4
	newOctave := int64(math.Round(baseOctave + ct*2))
	seq.Add([]k.S{{"C", newOctave, k.N8}})
}

/*
2.7
1.2
-0.1
-0.1
-0.1
-0.1
1.1
1.8
1.8
-0.6
-0.4
-0.4
-0.4
0.9
-0.4
-0.4
0.3
2.4
-1.1
0.5
1.2
1.2
2.0
-1.1
1.2
-1.1
1.0
1.0
1.9
-0.5
1.3
0.0
-1.1
1.6
-1.0
-0.6
-0.9
1.0
0.9
-0.9
0.8
-0.8
-1.0
-1.0
-0.9
-0.9
-1.0
-0.1
-0.4
-0.4
-0.4
-0.4
-0.4
-0.3
-0.3
-0.3
-0.3
-0.7
-0.6
-0.9
-0.9
-0.9
-0.9
-0.9
-0.9
-0.9
*/
