package core

import (
	"fmt"
	"time"
)

const NOTE_ON = 0x90
const NOTE_OFF = 0x80

type NoteEvent struct {
	isPause bool

	noteValue int64
	ccValue   int64

	downDuration time.Duration

	timer *time.Timer
	isOn  bool
}

func (e *NoteEvent) trigger(playHead time.Duration, s *Sequencer) {
	fmt.Printf("%s %#v\n", playHead, e)
	if e.isOn {
		return
	}

	if !e.isPause {
		triggerEvent(s, NOTE_ON, *e)
	}
	e.isOn = true

	e.timer = time.AfterFunc(e.downDuration, func() {
		if !e.isPause {
			triggerEvent(s, NOTE_OFF, *e)
		}
		e.timer = nil
		e.isOn = false
	})
}

func triggerEvent(s *Sequencer, onOffFlag int64, event NoteEvent) {
	s.Stream.WriteShort(onOffFlag, event.noteValue, event.ccValue)
}
