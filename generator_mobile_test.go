package chinaid

import "testing"

func TestMobile(t *testing.T) {
	t.Log(Mobile())
}

func BenchmarkMobile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(Mobile())
	}
}
