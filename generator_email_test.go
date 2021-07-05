package chinaid

import "testing"

func TestEmail(t *testing.T) {
	t.Log(Email())
}

func BenchmarkEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Email()
	}
}
