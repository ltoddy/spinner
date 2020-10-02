package spinner

import (
	"fmt"
	"github.com/ltoddy/spinner/internal"
	"os"
	"strings"
	"sync"
	"time"
)

type Spinner struct {
	lock      sync.RWMutex
	condition bool
}

func NewSpinner() *Spinner {
	return &Spinner{
		condition: false,
	}
}

func (s *Spinner) Start() {
	go func() {
		stdout := os.Stdout

		for v := range internal.Cycle([]string{"|", "/", "-", "\\"}) {
			status := []byte(fmt.Sprintf("%s computing.", v))
			_, _ = stdout.Write(status)
			_ = stdout.Sync()
			_, _ = stdout.Write([]byte(strings.Join(internal.Repeat(len(status), "\x08"), "")))
			time.Sleep(time.Millisecond * 100)

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
