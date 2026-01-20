package chinaid

import (
	"strings"
	"sync"
	"testing"
)

func TestNewPerson(t *testing.T) {
	p := NewPerson().Build()

	if p.IDNo() == "" {
		t.Error("IDNo should not be empty")
	}
	if len(p.IDNo()) != 18 {
		t.Errorf("IDNo length should be 18, got %d", len(p.IDNo()))
	}
	if p.Name() == "" {
		t.Error("Name should not be empty")
	}
	if p.Address() == "" {
		t.Error("Address should not be empty")
	}
	if p.Mobile() == "" {
		t.Error("Mobile should not be empty")
	}
	if len(p.Mobile()) != 11 {
		t.Errorf("Mobile length should be 11, got %d", len(p.Mobile()))
	}
	if p.BankNo() == "" {
		t.Error("BankNo should not be empty")
	}
	if p.Email() == "" {
		t.Error("Email should not be empty")
	}
}

func TestPersonProvince(t *testing.T) {
	p := NewPerson().Province("北京").Build()

	if !strings.HasPrefix(p.IDNo(), "11") {
		t.Errorf("IDNo should start with 11 for Beijing, got %s", p.IDNo())
	}
	if !strings.Contains(p.Address(), "北京") {
		t.Errorf("Address should contain 北京, got %s", p.Address())
	}
}

func TestPersonGender(t *testing.T) {
	male := NewPerson().Gender(GenderMale).Build()
	digit17 := int(male.IDNo()[16] - '0')
	if digit17%2 != 1 {
		t.Errorf("Male IDNo 17th digit should be odd, got %d in %s", digit17, male.IDNo())
	}

	female := NewPerson().Gender(GenderFemale).Build()
	digit17 = int(female.IDNo()[16] - '0')
	if digit17%2 != 0 {
		t.Errorf("Female IDNo 17th digit should be even, got %d in %s", digit17, female.IDNo())
	}
}

func TestPersonSeed(t *testing.T) {
	p1 := NewPerson().Seed(12345).Build()
	p2 := NewPerson().Seed(12345).Build()

	if p1.IDNo() != p2.IDNo() {
		t.Errorf("Same seed should produce same IDNo: %s vs %s", p1.IDNo(), p2.IDNo())
	}
	if p1.Name() != p2.Name() {
		t.Errorf("Same seed should produce same Name: %s vs %s", p1.Name(), p2.Name())
	}
}

func TestPersonBuildN(t *testing.T) {
	persons := NewPerson().Province("广东").BuildN(10)

	if len(persons) != 10 {
		t.Errorf("BuildN(10) should return 10 persons, got %d", len(persons))
	}

	for i, p := range persons {
		if !strings.Contains(p.Address(), "广东") {
			t.Errorf("Person %d address should contain 广东, got %s", i, p.Address())
		}
	}
}

func TestPersonAgeRange(t *testing.T) {
	p := NewPerson().AgeRange(25, 30).Build()
	age := p.Age()
	if age < 25 || age > 30 {
		t.Errorf("Age should be between 25 and 30, got %d", age)
	}
}

func TestPersonConcurrentSafety(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p := NewPerson().Build()
			if p.IDNo() == "" {
				t.Error("IDNo should not be empty")
			}
		}()
	}
	wg.Wait()
}

func TestPersonIDNoValidation(t *testing.T) {
	// Generate multiple persons and validate their ID numbers
	for i := 0; i < 100; i++ {
		p := NewPerson().Build()
		if !ValidateIDNo(p.IDNo()) {
			t.Errorf("Generated IDNo should be valid: %s", p.IDNo())
		}
	}
}

func TestPersonBankNoValidation(t *testing.T) {
	// Generate multiple persons and validate their bank card numbers
	for i := 0; i < 100; i++ {
		p := NewPerson().Build()
		if !ValidateLUHN(p.BankNo()) {
			t.Errorf("Generated BankNo should pass LUHN validation: %s", p.BankNo())
		}
	}
}

func TestPersonGetters(t *testing.T) {
	p := NewPerson().Seed(54321).Build()

	// Test all getters return non-empty values
	if p.LastName() == "" {
		t.Error("LastName should not be empty")
	}
	if p.FirstName() == "" {
		t.Error("FirstName should not be empty")
	}
	if p.Name() != p.LastName()+p.FirstName() {
		t.Errorf("Name should be LastName + FirstName, got %s vs %s+%s",
			p.Name(), p.LastName(), p.FirstName())
	}
	if p.Province() == "" {
		t.Error("Province should not be empty")
	}
	if p.City() == "" {
		t.Error("City should not be empty")
	}
	if p.AreaCode() == "" {
		t.Error("AreaCode should not be empty")
	}
	if len(p.AreaCode()) != 6 {
		t.Errorf("AreaCode length should be 6, got %d", len(p.AreaCode()))
	}
	if p.Birthday().IsZero() {
		t.Error("Birthday should not be zero")
	}
	if p.Gender() != GenderMale && p.Gender() != GenderFemale {
		t.Errorf("Gender should be Male or Female, got %v", p.Gender())
	}
}
