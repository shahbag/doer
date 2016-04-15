package doer

import (
 	"fmt"
	"time"
)
type Stats struct{
  said int
}
type Doer struct {
	stats Stats
  done   chan bool // Signals that the say routine to stop
}

func NewDoer()(*Doer, error) {
  d := &Doer{
    stats: Stats{said: 0},
  }
  go d.say("hello")
  return d, nil
}

func (d *Doer) say(s string) {

  ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
					fmt.Println(s)
          d.stats.said++
		case <-d.done:
			return
		}
	}
}

func (d *Doer) Status() Stats {
   return d.stats
}
func (d * Doer) Close() {
  close(d.done)
}
