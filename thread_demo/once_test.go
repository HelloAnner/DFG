package thread_demo

import (
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	var once sync.Once

	f1 := func() {
		t.Logf("in f1")
	}

	f2 := func() {
		t.Logf("in f12")
	}

	once.Do(f1)
	once.Do(f2)
}
