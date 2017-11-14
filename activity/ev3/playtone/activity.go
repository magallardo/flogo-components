package playtone

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/ev3go/ev3dev"
	"log"
	"time"
)

// SoundPath is the path to the ev3 sound events
const SoundPath = "/dev/input/by-path/platform-snd-legoev3-event"

var speaker = ev3dev.NewSpeaker(SoundPath)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the activity data from the context
	//tone := context.GetInput("tone").(int)
	//duration := context.GetInput("duration").(int)
	must(speaker.Init())
	defer speaker.Close()

	// Play tone for given duration
	must(speaker.Tone(400))
	time.Sleep(200 * time.Millisecond)

	// Then stop tone playback
	must(speaker.Tone(0))

	// Set the output as part of the context
	context.SetOutput("output", "Success")

	return true, nil
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
