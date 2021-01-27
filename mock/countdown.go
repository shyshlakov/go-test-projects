package mock

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper ...
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper ...
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown ...
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

// Sleep ...
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
