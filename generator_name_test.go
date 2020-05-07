package chinaid

import "testing"

func TestName(t *testing.T) {
	t.Log(Name())
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(Name())
	}
}
