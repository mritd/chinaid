// rand_test.go
package chinaid

import (
	"sync"
	"testing"
)

func TestRngIntn(t *testing.T) {
	rng := NewRng()
	for i := 0; i < 100; i++ {
		n := rng.Intn(10)
		if n < 0 || n >= 10 {
			t.Errorf("Intn(10) = %d, want [0, 10)", n)
		}
	}
}

func TestRngWithSeed(t *testing.T) {
	rng1 := NewRngWithSeed(12345)
	rng2 := NewRngWithSeed(12345)

	for i := 0; i < 10; i++ {
		a := rng1.Intn(1000)
		b := rng2.Intn(1000)
		if a != b {
			t.Errorf("Same seed produced different results: %d vs %d", a, b)
		}
	}
}

func TestRngChoice(t *testing.T) {
	rng := NewRng()
	items := []string{"a", "b", "c"}

	for i := 0; i < 100; i++ {
		choice := rng.Choice(items)
		found := false
		for _, item := range items {
			if choice == item {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Choice returned %s, not in slice", choice)
		}
	}
}

func TestRngConcurrentSafety(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rng := NewRng()
			_ = rng.Intn(100)
		}()
	}
	wg.Wait()
}
