package ctoai

import (
	"github.com/cto-ai/sdk-go/internal/daemon"
)

// Ux is the object that contains the UX methods
type Ux struct{}

func NewUx() *Ux {
	return &Ux{}
}

// Print prints text to the output interface (i.e. terminal/slack).
//
// Example:
//
//  u := ctoai.NewUx()
//  err := u.Print("testing")
//  if err != nil {
//      panic(err)
//  }
//
// Output:
//
// testing
func (*Ux) Print(text string) error {
	return daemon.SimpleRequest("print", daemon.PrintBody{Text: text})
}

// SpinnerStart presents a spinner on the output interface
// (i.e. terminal or slack) that to spin until the SpinnerStop method
// is called.
//
// Example:
//
//  u := ctoai.NewUx()
//  err := u.SpinnerStart("Starting process...")
//  if err != nil {
//      panic(err)
//  }
//
// Output:
// [spinner emoji w/ spinner animation here] Starting process...
func (*Ux) SpinnerStart(text string) error {
	return daemon.SimpleRequest("start-spinner", daemon.SpinnerStartBody{Text: text})
}

// SpinnerStop stops a spinner that has been previously started on the
// interface (i.e. terminal or slack).
//
// Example:
//
//  ... //previous spinner started here
//
//  err := u.SpinnerStop("Done!")
//  if err != nil {
//      panic(err)
//  }
//
// Output:
// [spinner completed completed here] Done!
func (*Ux) SpinnerStop(text string) error {
	return daemon.SimpleRequest("stop-spinner", daemon.SpinnerStopBody{Text: text})
}

// ProgressBarStart presents a progressbar on the output interface
// (i.e. terminal or slack) that will stay present until the
// progressbar stop method is called.
//
// The input length is the total length of the progress bar, e.g.
// if you have 5 steps in your logic, then a unit length of 5 might be
// and appropriate length.
//
// The initial length indicates the unit length (out of total length) that is initially
// filled at the start.
//
// Example:
//
//  u := ctoai.NewUx()
//  err := u.ProgressBarStart(5, 1, "Downloading...")
//  if err != nil {
//      panic(err)
//  }
//
// Output:
// [progressbar animation with 1/5 of the bar filled here] Downloading...
func (*Ux) ProgressBarStart(length, initial int, message string) error {
	return daemon.SimpleRequest("progress-bar/start", daemon.ProgressBarStartBody{Length: length, Initial: initial, Text: message})
}

// ProgressBarAdvance adds onto a progressbar that is already present
// on the interface (i.e. terminal or slack).
//
// The increment indicates the additional length (out of total length)
// that will be filled.
//
// Example:
//
//  ...
//  [progressbar animation with 1/5 of the bar filled here] Downloading...
//  err := u.ProgressBarAdvance(1)
//  if err != nil {
//      panic(err)
//  }
//
// Output:
// [progressbar animation with 2/5 of the bar filled here] Downloading...
func (*Ux) ProgressBarAdvance(increment int) error {
	return daemon.SimpleRequest("progress-bar/advance", daemon.ProgressBarAdvanceBody{Increment: increment})
}

// ProgressBarStop completes a progressbar that is already present on
// the interface (i.e. terminal or slack).
//
// The text will change the initial text (set from the ux.ProgressBarStart method).
//
// Example:
//
//  ...
//  [progressbar animation with 2/5 of the bar filled here] Downloading...
//  err := u.ProgressBarStop("Done!")
//  if err != nil {
//      panic(err)
//  }
//
// Output:
// [progressbar animation with 5/5 of the bar filled here] Done!
func (*Ux) ProgressBarStop(message string) error {
	return daemon.SimpleRequest("progress-bar/stop", daemon.ProgressBarStopBody{Text: message})
}
