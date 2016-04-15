package doer

import "testing"

func TestDoer(t *testing.T) {
  d, err := NewDoer()
  if err!=nil {
    t.Fatal(err)
  }
  d.Close()
}
