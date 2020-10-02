package spinner

import (
	"sync"
	"time"
)

type SpinnerBuilder struct {
	types map[SpinnerType]SpinnerData

	stype    *SpinnerType
	text     *string
	frames   *[]string
	interval *time.Duration
}

func NewSpinnerBuilder() *SpinnerBuilder {
	types := make(map[SpinnerType]SpinnerData)
	types[Dots] = SpinnerData{[]string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}, 80 * time.Millisecond}
	types[Dots2] = SpinnerData{[]string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}, 80 * time.Millisecond}
	types[Dots3] = SpinnerData{[]string{"⠋", "⠙", "⠚", "⠞", "⠖", "⠦", "⠴", "⠲", "⠳", "⠓"}, 80 * time.Millisecond}
	types[Dots4] = SpinnerData{[]string{"⠄", "⠆", "⠇", "⠋", "⠙", "⠸", "⠰", "⠠", "⠰", "⠸", "⠙", "⠋", "⠇", "⠆"}, 80 * time.Millisecond}
	types[Dots5] = SpinnerData{[]string{"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"}, 80 * time.Millisecond}
	types[Dots6] = SpinnerData{[]string{"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁"}, 80 * time.Millisecond}
	types[Dots7] = SpinnerData{[]string{"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"}, 80 * time.Millisecond}
	types[Dots8] = SpinnerData{[]string{"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"}, 80 * time.Millisecond}
	types[Dots9] = SpinnerData{[]string{"⢹", "⢺", "⢼", "⣸", "⣇", "⡧", "⡗", "⡏"}, 80 * time.Millisecond}
	types[Dots10] = SpinnerData{[]string{"⢄", "⢂", "⢁", "⡁", "⡈", "⡐", "⡠"}, 80 * time.Millisecond}
	types[Dots11] = SpinnerData{[]string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"}, 80 * time.Millisecond}
	types[Dots12] = SpinnerData{[]string{"⢀⠀", "⡀⠀", "⠄⠀", "⢂⠀", "⡂⠀", "⠅⠀", "⢃⠀", "⡃⠀", "⠍⠀", "⢋⠀", "⡋⠀", "⠍⠁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⢈⠩", "⡀⢙", "⠄⡙", "⢂⠩", "⡂⢘", "⠅⡘", "⢃⠨", "⡃⢐", "⠍⡐", "⢋⠠", "⡋⢀", "⠍⡁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⠈⠩", "⠀⢙", "⠀⡙", "⠀⠩", "⠀⢘", "⠀⡘", "⠀⠨", "⠀⢐", "⠀⡐", "⠀⠠", "⠀⢀", "⠀⡀"}, 80 * time.Millisecond}
	types[Line] = SpinnerData{[]string{"-", "\\", "|", "/"}, 130 * time.Millisecond}
	types[Line2] = SpinnerData{[]string{"⠂", "-", "–", "—", "–", "-"}, 100 * time.Millisecond}
	types[Pipe] = SpinnerData{[]string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}, 100 * time.Millisecond}
	types[SimpleDots] = SpinnerData{[]string{".  ", ".. ", "...", "   "}, 400 * time.Millisecond}
	types[SimpleDotsScrolling] = SpinnerData{[]string{".  ", ".. ", "...", " ..", "  .", "   "}, 200 * time.Millisecond}
	types[Star] = SpinnerData{[]string{"✶", "✸", "✹", "✺", "✹", "✷"}, 70 * time.Millisecond}
	types[Star2] = SpinnerData{[]string{"+", "x", "*"}, 80 * time.Millisecond}
	types[Flip] = SpinnerData{[]string{"_", "_", "_", "-", "`", "`", "'", "´", "-", "_", "_", "_"}, 70 * time.Millisecond}
	types[Hamburger] = SpinnerData{[]string{"☱", "☲", "☴"}, 100 * time.Millisecond}
	types[GrowVertical] = SpinnerData{[]string{"▁", "▃", "▄", "▅", "▆", "▇", "▆", "▅", "▄", "▃"}, 120 * time.Millisecond}
	types[GrowHorizontal] = SpinnerData{[]string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "▊", "▋", "▌", "▍", "▎"}, 120 * time.Millisecond}
	types[Balloon] = SpinnerData{[]string{" ", ".", "o", "O", "@", "*", " "}, 140 * time.Millisecond}
	types[Balloon2] = SpinnerData{[]string{".", "o", "O", "°", "O", "o", "."}, 120 * time.Millisecond}
	types[Noise] = SpinnerData{[]string{"▓", "▒", "░"}, 100 * time.Millisecond}
	types[Bounce] = SpinnerData{[]string{"⠁", "⠂", "⠄", "⠂"}, 120 * time.Millisecond}
	types[BoxBounce] = SpinnerData{[]string{"▖", "▘", "▝", "▗"}, 120 * time.Millisecond}
	types[BoxBounce2] = SpinnerData{[]string{"▌", "▀", "▐", "▄"}, 100 * time.Millisecond}
	types[Triangle] = SpinnerData{[]string{"◢", "◣", "◤", "◥"}, 50 * time.Millisecond}
	types[Arc] = SpinnerData{[]string{"◜", "◠", "◝", "◞", "◡", "◟"}, 100 * time.Millisecond}
	types[Circle] = SpinnerData{[]string{"◡", "⊙", "◠"}, 120 * time.Millisecond}
	types[SquareCorners] = SpinnerData{[]string{"◰", "◳", "◲", "◱"}, 100 * time.Millisecond}
	types[CircleQuarters] = SpinnerData{[]string{"◴", "◷", "◶", "◵"}, 120 * time.Millisecond}

	return &SpinnerBuilder{
		types: types,

		stype:    nil,
		text:     nil,
		frames:   nil,
		interval: nil,
	}
}

func (sb *SpinnerBuilder) WithText(text string) *SpinnerBuilder {
	sb.text = &text
	return sb
}

func (sb *SpinnerBuilder) WithType(stype SpinnerType) *SpinnerBuilder {
	sb.stype = &stype
	return sb
}

func (sb *SpinnerBuilder) WithFrames(frames []string) *SpinnerBuilder {
	sb.frames = &frames
	return sb
}

func (sb *SpinnerBuilder) WithInterval(interval int64) *SpinnerBuilder {
	*sb.interval = time.Duration(interval * 1000 * 1000)
	return sb
}

func (sb *SpinnerBuilder) Build() *Spinner {
	text := "computing"
	data := defaultSpinnerData

	if sb.text != nil {
		text = *sb.text
	}
	if sb.stype != nil {
		data = sb.types[*sb.stype]
	}

	frames := data.frames
	interval := data.interval

	if sb.frames != nil {
		frames = *sb.frames
	}
	if sb.interval != nil {
		interval = *sb.interval
	}

	return &Spinner{
		lock:      sync.RWMutex{},
		condition: false,
		text:      text,
		frames:    frames,
		interval:  interval,
	}
}

type SpinnerType = int

const (
	Dots                SpinnerType = iota
	Dots2               SpinnerType = iota
	Dots3               SpinnerType = iota
	Dots4               SpinnerType = iota
	Dots5               SpinnerType = iota
	Dots6               SpinnerType = iota
	Dots7               SpinnerType = iota
	Dots8               SpinnerType = iota
	Dots9               SpinnerType = iota
	Dots10              SpinnerType = iota
	Dots11              SpinnerType = iota
	Dots12              SpinnerType = iota
	Line                SpinnerType = iota
	Line2               SpinnerType = iota
	Pipe                SpinnerType = iota
	SimpleDots          SpinnerType = iota
	SimpleDotsScrolling SpinnerType = iota
	Star                SpinnerType = iota
	Star2               SpinnerType = iota
	Flip                SpinnerType = iota
	Hamburger           SpinnerType = iota
	GrowVertical        SpinnerType = iota
	GrowHorizontal      SpinnerType = iota
	Balloon             SpinnerType = iota
	Balloon2            SpinnerType = iota
	Noise               SpinnerType = iota
	Bounce              SpinnerType = iota
	BoxBounce           SpinnerType = iota
	BoxBounce2          SpinnerType = iota
	Triangle            SpinnerType = iota
	Arc                 SpinnerType = iota
	Circle              SpinnerType = iota
	SquareCorners       SpinnerType = iota
	CircleQuarters      SpinnerType = iota
)

type SpinnerData struct {
	frames   []string
	interval time.Duration
}
