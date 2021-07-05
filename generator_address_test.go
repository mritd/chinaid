package chinaid

import "testing"

func TestProvinceAndCity(t *testing.T) {
	t.Log(ProvinceAndCity())
}

func TestAddress(t *testing.T) {
	t.Log(Address())
}

func BenchmarkProvinceAndCity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProvinceAndCity()
	}
}

func BenchmarkAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Address()
	}
}
