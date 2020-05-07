package chinaid

import "testing"

func TestLUHNProcess(t *testing.T) {
	t.Log(LUHNProcess("623190380371814"))
}

func TestBankNo(t *testing.T) {
	t.Log(BankNo())
}

func BenchmarkLUHNProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(LUHNProcess("623190380371814"))
	}
}

func BenchmarkBankNo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(BankNo())
	}
}
