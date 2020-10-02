package spinner

import "time"

const (
	defaultText     = "computing"
	defaultInterval = 80 * time.Millisecond
)

var (
	defaultFrames      = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	defaultSpinnerData = SpinnerData{[]string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}, 80 * time.Millisecond}
)
