package tt

/*

Written by Jack Watts for SMPTE Standards TC 27C Community.

.*/

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var tcRegexp = regexp.MustCompile(`^(\d\d)[:;](\d\d)[:;](\d\d)[:.](\d+)$`)

// divMod returns the floating-point remainder of a/b
func divMod(a float64, b float64) (float64, float64) {
	q := a / b
	r := math.Mod(a, b)
	return q, float64(r)
}

// Timecode struct.
type Timecode struct {
	frameRate   float64
	hours       int
	minutes     int
	seconds     int
	frames      int
	totalFrames int
}

// NewTimecode initialises a new timecode type from a given framerate.
func NewTimecode(frameRate float64) (*Timecode, error) {
	if frameRate >= 0 {
		tc := &Timecode{
			frameRate:   frameRate,
			hours:       0,
			minutes:     0,
			seconds:     0,
			frames:      0,
			totalFrames: 0,
		}
		return tc, nil
	}
	return nil, fmt.Errorf("unsupported framerate")
}

// GetTimeCode method generates a SMPTE timcode when called against type Timecode.
func (tc *Timecode) GetTimeCode() string {
	h := tc.hours
	m := tc.minutes
	s := tc.seconds
	f := tc.frames
	// set 23Pulldown 1
	fr := tc.frameRate
	a, b := math.Modf(fr)
	if b != 0.0 {
		fr = a + 1
	}
	totalFrames := float64(h*3600+m*60+s)*fr + float64(f)
	seconds, frames := divMod(totalFrames, fr)
	minutes, seconds := divMod(seconds, 60)
	hours, minutes := divMod(minutes, 60)
	return fmt.Sprintf("%02d:%02d:%02d:%02d", int(hours), int(minutes), int(seconds), int(frames))
}

// SetFrames sets the frame count in type Timecode.
func (tc *Timecode) SetFrames(frameCount int) {
	tc.totalFrames = frameCount
	seconds, frames := divMod(float64(tc.totalFrames), tc.frameRate)
	minutes, seconds := divMod(float64(seconds), 60)
	hours, minutes := divMod(float64(minutes), 60)
	tc.frames = int(frames)
	tc.seconds = int(seconds)
	tc.minutes = int(minutes)
	tc.hours = int(hours)
}

// getFloat returns a value of type float64 from a given value of type string that contains a whole number.
func getFloat(s string) float64 {
	if s == "" {
		return 0.0
	}
	f, _ := strconv.ParseFloat(s, 64)
	return float64(f)
}
