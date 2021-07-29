package main

import (
	"fmt"

	"github.com/haroldcampbell/go_utils/utils"
	"github.com/rakyll/portmidi"
)

func main() {

	err := portmidi.Initialize()
	if err != nil {
		fmt.Printf("Error initializing portmidi: %v\n", err)
		return
	}

	getDeviceInfo()
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
