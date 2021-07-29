package main

import (
	"fmt"
	"time"

	"github.com/rakyll/portmidi"
)

const NOTE_ON = 0x90
const NOTE_OFF = 0x80

type NoteEvent struct {
	noteValue int64
	ccValue   int64
	onOffFlag int64

	eventTime time.Duration
}

type NoteEvents struct {
	events []NoteEvent
}

type Sequencer struct {
	Stream    *portmidi.Stream
	noteTable map[string]int64

	recordHead    time.Duration
	triggerEvents map[time.Duration]*NoteEvents
}

type NoteD struct {
	Note   string
	Octave int64
}

func NewSequencer(out *portmidi.Stream) *Sequencer {
	s := &Sequencer{
		Stream:        out,
		triggerEvents: make(map[time.Duration]*NoteEvents),
	}

	initSequencerNotes(s)

	return s
}

func addNoteEvent(s *Sequencer, notes []NoteD, onOffFlag int64) {
	if s.triggerEvents[s.recordHead] == nil {
		s.triggerEvents[s.recordHead] = &NoteEvents{events: make([]NoteEvent, 0, 0)}
	}
	noteEvents := s.triggerEvents[s.recordHead]

	for _, note := range notes {
		noteValue := s.toNoteValue(note.Note, note.Octave)

		eventOn := NoteEvent{
			noteValue: noteValue,
			ccValue:   10,
			onOffFlag: onOffFlag,
			eventTime: s.recordHead}

		noteEvents.events = append(noteEvents.events, eventOn)
	}
}

var framerate = 300 * time.Millisecond

func (s *Sequencer) Add(notes []NoteD) {
	addNoteEvent(s, notes, NOTE_ON)
	s.recordHead += framerate
	addNoteEvent(s, notes, NOTE_OFF)
}

func (s *Sequencer) Play() {
	var playHead time.Duration = 0

	for _ = range time.Tick(1 * time.Second) {
		noteEvents, ok := s.triggerEvents[playHead]

		if !ok {
			fmt.Printf("Exiting at playHead: %v\n", playHead)
			return
		}

		for _, event := range noteEvents.events {
			fmt.Printf("%s %#v\n", playHead, event)
			onNoteEvent(s, event)
		}

		playHead += time.Second
	}
}

func (s *Sequencer) ReatPlay() {
	var playHead time.Duration = 0

	for _ = range time.Tick(framerate) {
		noteEvents, ok := s.triggerEvents[playHead]

		if !ok {
			playHead = 0
			noteEvents, ok = s.triggerEvents[playHead]
			if !ok {
				return
			}
		}

		for _, event := range noteEvents.events {
			fmt.Printf("%s %#v\n", playHead, event)
			onNoteEvent(s, event)
		}

		playHead += framerate
	}
}

func initSequencerNotes(s *Sequencer) {
	s.noteTable = make(map[string]int64)

	s.noteTable["A"] = 21
	s.noteTable["B"] = 23
	s.noteTable["C"] = 24
	s.noteTable["D"] = 26
	s.noteTable["E"] = 28
	s.noteTable["F"] = 29
	s.noteTable["G"] = 31
}

func (s *Sequencer) toNoteValue(note string, octave int64) int64 {
	return s.noteTable[note] + octave*12
}

func stopAll(s *Sequencer) {
	for i := 0; i <= 127; i++ {
		s.Stream.WriteShort(NOTE_OFF, int64(i), 10)
	}
}

func onNoteEvent(s *Sequencer, event NoteEvent) {
	s.Stream.WriteShort(event.onOffFlag, event.noteValue, event.ccValue)
}

func (s *Sequencer) noteOn(note string, octave int64, ccValue int64) {
	noteValue := s.toNoteValue(note, octave)
	s.Stream.WriteShort(NOTE_ON, noteValue, ccValue)
}

func (s *Sequencer) A(octave int64, ccValue int64) {
	s.noteOn("A", octave, ccValue)
}

func (s *Sequencer) B(octave int64, ccValue int64) {
	s.noteOn("B", octave, ccValue)
}

func (s *Sequencer) C(octave int64, ccValue int64) {
	s.noteOn("C", octave, ccValue)
}

func (s *Sequencer) D(octave int64, ccValue int64) {
	s.noteOn("D", octave, ccValue)
}

func (s *Sequencer) E(octave int64, ccValue int64) {
	s.noteOn("E", octave, ccValue)
}

func (s *Sequencer) F(octave int64, ccValue int64) {
	s.noteOn("F", octave, ccValue)
}

func (s *Sequencer) G(octave int64, ccValue int64) {
	s.noteOn("G", octave, ccValue)
}
