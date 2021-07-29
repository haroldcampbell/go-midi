package core

import (
	"fmt"
	"log"
	"time"

	"github.com/haroldcampbell/go_utils/utils"
	"github.com/rakyll/portmidi"
)

type NoteLength int

const (
	N0  NoteLength = iota // Whole Note
	N2                    // Half Note
	N4                    // Quarter Note
	N8                    // Eight Note
	N16                   // Sixteenth Note
)

type NoteEvents struct {
	events []NoteEvent
}

type Sequencer struct {
	Stream        *portmidi.Stream
	noteTable     map[string]int64
	durationTable map[NoteLength]time.Duration

	recordHead    time.Duration
	triggerEvents map[time.Duration]*NoteEvents

	tempo               int
	framerate           time.Duration
	quarterNote         int
	quarterNoteDuration time.Duration
}

type ScoreEntry struct {
	Note   string
	Octave int64
	Length NoteLength
}

type S ScoreEntry

func NewSequencer(tempo int, quarterNote int) *Sequencer {
	err := portmidi.Initialize()
	if err != nil {
		fmt.Printf("Error initializing portmidi: %v\n", err)
		return nil
	}

	_, outputDeviceID := getDeviceInfo()

	out, err := portmidi.NewOutputStream(outputDeviceID, 1024, 0)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	s := &Sequencer{
		Stream:              out,
		triggerEvents:       make(map[time.Duration]*NoteEvents),
		tempo:               tempo,
		framerate:           time.Duration(tempo) * time.Millisecond,
		quarterNote:         quarterNote,
		quarterNoteDuration: time.Duration(quarterNote) * time.Millisecond,
	}

	initSequencerNotes(s)
	initNoteDurations(s)

	return s
}

func getDeviceInfo() (portmidi.DeviceID, portmidi.DeviceID) {
	var numDevices = portmidi.CountDevices() // returns the number of MIDI devices
	fmt.Printf("Found %v devices\n", numDevices)

	var inputDeviceID = portmidi.DefaultInputDeviceID()   // returns the ID of the system default input
	var outputDeviceID = portmidi.DefaultOutputDeviceID() // returns the ID of the system default output

	var deviceInfo *portmidi.DeviceInfo

	deviceInfo = portmidi.Info(inputDeviceID) // returns info about a MIDI device
	fmt.Printf("Input Device: %v\n", utils.PrettyMongoString(deviceInfo))

	deviceInfo = portmidi.Info(outputDeviceID) // returns info about a MIDI device
	fmt.Printf("Output Device: %v\n", utils.PrettyMongoString(deviceInfo))

	return inputDeviceID, outputDeviceID
}

func (s *Sequencer) WaitToEnd() {
	var option int

	fmt.Println("Press enter to end.")
	fmt.Scanf("%c", &option)
	fmt.Printf("Stopping all notes.\nDone.\n")

	StopAll(s)
	s.Stream.Close()
}

func addNoteEvent(s *Sequencer, notes []S) {
	if s.triggerEvents[s.recordHead] == nil {
		s.triggerEvents[s.recordHead] = &NoteEvents{events: make([]NoteEvent, 0, 0)}
	}
	noteEvents := s.triggerEvents[s.recordHead]

	for _, note := range notes {
		eventOn := NoteEvent{
			noteValue:    s.toNoteValue(note.Note, note.Octave),
			ccValue:      20,
			downDuration: s.durationTable[note.Length]}

		if note.Note == "_" {
			eventOn.isPause = true
		}
		noteEvents.events = append(noteEvents.events, eventOn)
	}
}

func (s *Sequencer) Add(notes []S) {
	addNoteEvent(s, notes)
	s.recordHead += s.framerate
}

func (s *Sequencer) Play() {
	go func() {
		var playHead time.Duration = 0
		for _ = range time.Tick(s.framerate) {
			noteEvents, ok := s.triggerEvents[playHead]

			if !ok {
				fmt.Printf("Exiting at playHead: %v\n", playHead)
				return
			}

			for _, event := range noteEvents.events {
				event.trigger(playHead, s)
			}

			playHead += s.framerate
		}
	}()
}

func (s *Sequencer) ReatPlay() {
	// go func() {
	// 	var playHead time.Duration = 0
	// 	for _ = range time.Tick(s.framerate) {
	// 		noteEvents, ok := s.triggerEvents[playHead]

	// 		if !ok {
	// 			playHead = 0
	// 			noteEvents, ok = s.triggerEvents[playHead]
	// 			if !ok {
	// 				return
	// 			}
	// 		}

	// 		go func() {
	// 			for _, event := range noteEvents.events {
	// 				event.trigger(playHead, s)
	// 			}
	// 		}()

	// 		playHead += s.framerate
	// 	}
	// }()
	go func() {
		for {
			s.Play()
		}
	}()
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

	s.noteTable["_"] = -1 // Pause
}

func initNoteDurations(s *Sequencer) {
	s.durationTable = make(map[NoteLength]time.Duration)

	s.durationTable[N0] = s.quarterNoteDuration * 4
	s.durationTable[N2] = s.quarterNoteDuration * 2
	s.durationTable[N4] = s.quarterNoteDuration
	s.durationTable[N8] = s.quarterNoteDuration / 2.0
	s.durationTable[N16] = s.quarterNoteDuration / 4.0
}

func (s *Sequencer) toNoteValue(note string, octave int64) int64 {
	return s.noteTable[note] + octave*12
}

func StopAll(s *Sequencer) {
	for i := 0; i <= 127; i++ {
		s.Stream.WriteShort(NOTE_OFF, int64(i), 10)
	}
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
