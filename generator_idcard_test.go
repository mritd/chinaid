package chinaid

import "testing"

func TestIssueOrg(t *testing.T) {
	t.Log(IssueOrg())
}

func TestValidPeriod(t *testing.T) {
	t.Log(ValidPeriod())
}

func TestIDNo(t *testing.T) {
	t.Log(IDNo())
}

func TestVerifyCode(t *testing.T) {
	t.Log(VerifyCode("636706198006242277"))
}

func TestRandDate(t *testing.T) {
	t.Log(RandDate())
}

func BenchmarkIssueOrg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(IssueOrg())
	}
}

func BenchmarkValidPeriod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(ValidPeriod())
	}
}

func BenchmarkIDNo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(IDNo())
	}
}

func BenchmarkVerifyCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(VerifyCode("636706198006242277"))
	}
}
