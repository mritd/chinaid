package chinaid

import "testing"

func TestGenderString(t *testing.T) {
	tests := []struct {
		g    Gender
		want string
	}{
		{GenderRandom, "random"},
		{GenderMale, "male"},
		{GenderFemale, "female"},
	}

	for _, tt := range tests {
		if got := tt.g.String(); got != tt.want {
			t.Errorf("Gender.String() = %s, want %s", got, tt.want)
		}
	}
}

func TestGenderIsMale(t *testing.T) {
	if !GenderMale.IsMale() {
		t.Error("GenderMale.IsMale() should be true")
	}
	if GenderFemale.IsMale() {
		t.Error("GenderFemale.IsMale() should be false")
	}
}

func TestGenderIsFemale(t *testing.T) {
	if !GenderFemale.IsFemale() {
		t.Error("GenderFemale.IsFemale() should be true")
	}
	if GenderMale.IsFemale() {
		t.Error("GenderMale.IsFemale() should be false")
	}
}
