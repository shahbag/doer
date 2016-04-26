package doer

import "testing"

func TestDoer(t *testing.T) {
	t.Fatal("WTF")
  d, err := NewDoer()
  if err!=nil {
    t.Fatal(err)
  }
  d.Close()
}
