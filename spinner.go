package spinner

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ltoddy/spinner/internal"
)

type Spinner struct {
	lock      sync.RWMutex
	condition bool

	text     string
	frames   []string
	interval time.Duration
}

func NewSpinner() *Spinner {
	return &Spinner{
		condition: false,

		text:     defaultText,
		frames:   defaultFrames,
		interval: defaultInterval,
	}
}

func (s *Spinner) Start() {
	go func() {
		stdout := os.Stdout

		for v := range internal.Cycle(s.frames) {
			status := []byte(fmt.Sprintf("  %s %s.", v, s.text))
			_, _ = stdout.Write(status)
			_ = stdout.Sync()
			_, _ = stdout.Write([]byte(strings.Join(internal.Repeat(len(status), "\x08"), "")))

			time.Sleep(s.interval)

			s.lock.RLock()
			condition := s.condition
			s.lock.RUnlock()
			if condition {
				break
			}
		}
	}()
}

func (s *Spinner) Stop() {
	s.lock.Lock()
	s.condition = true
	s.lock.Unlock()
}
