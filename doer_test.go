package doer

import "testing"

func TestDoer(t *testing.T) {
  if d, err := NewDoer(); err!=nil {
    t.Fatal(err)
  }
  d.Close()
}
