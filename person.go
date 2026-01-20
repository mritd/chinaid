package chinaid

import (
	"fmt"
	"time"

	"github.com/mritd/chinaid/v2/metadata"
)

// Person represents generated person information.
type Person struct {
	idNo      string
	name      string
	lastName  string
	firstName string
	gender    Gender
	birthday  time.Time
	province  string
	city      string
	areaCode  string
	address   string
	mobile    string
	bankNo    string
	email     string
}

// Getter methods

// IDNo returns the ID card number.
func (p *Person) IDNo() string { return p.idNo }

// Name returns the full name.
func (p *Person) Name() string { return p.name }

// LastName returns the surname.
func (p *Person) LastName() string { return p.lastName }

// FirstName returns the given name.
func (p *Person) FirstName() string { return p.firstName }

// Gender returns the gender.
func (p *Person) Gender() Gender { return p.gender }

// Birthday returns the birthday.
func (p *Person) Birthday() time.Time { return p.birthday }

// Age returns the age.
func (p *Person) Age() int {
	now := time.Now()
	age := now.Year() - p.birthday.Year()
	if now.YearDay() < p.birthday.YearDay() {
		age--
	}
	return age
}

// Province returns the province name.
func (p *Person) Province() string { return p.province }

// City returns the city name.
func (p *Person) City() string { return p.city }

// AreaCode returns the 6-digit area code.
func (p *Person) AreaCode() string { return p.areaCode }

// Address returns the full address.
func (p *Person) Address() string { return p.address }

// Mobile returns the mobile phone number.
func (p *Person) Mobile() string { return p.mobile }

// BankNo returns the bank card number.
func (p *Person) BankNo() string { return p.bankNo }

// Email returns the email address.
func (p *Person) Email() string { return p.email }

// PersonBuilder is a builder for creating Person instances.
type PersonBuilder struct {
	rng      *Rng
	seed     int64
	hasSeed  bool
	province string
	gender   Gender
	minAge   int
	maxAge   int
}

// NewPerson creates a new PersonBuilder.
func NewPerson() *PersonBuilder {
	return &PersonBuilder{
		minAge: 18,
		maxAge: 60,
	}
}

// Province sets the province filter.
func (b *PersonBuilder) Province(province string) *PersonBuilder {
	b.province = province
	return b
}

// Gender sets the gender.
func (b *PersonBuilder) Gender(gender Gender) *PersonBuilder {
	b.gender = gender
	return b
}

// AgeRange sets the age range [min, max].
func (b *PersonBuilder) AgeRange(min, max int) *PersonBuilder {
	b.minAge = min
	b.maxAge = max
	return b
}

// Seed sets the random seed for reproducibility.
func (b *PersonBuilder) Seed(seed int64) *PersonBuilder {
	b.seed = seed
	b.hasSeed = true
	return b
}

// Build generates a single Person.
func (b *PersonBuilder) Build() *Person {
	if b.hasSeed {
		b.rng = NewRngWithSeed(b.seed)
	} else {
		b.rng = NewRng()
	}

	p := &Person{}

	b.generateLocation(p)
	b.generateGender(p)
	b.generateBirthday(p)
	b.generateIDNo(p)
	b.generateName(p)
	b.generateAddress(p)
	b.generateMobile(p)
	b.generateBankNo(p)
	b.generateEmail(p)

	return p
}

// BuildN generates multiple Person instances.
func (b *PersonBuilder) BuildN(n int) []*Person {
	persons := make([]*Person, n)
	for i := 0; i < n; i++ {
		builder := &PersonBuilder{
			province: b.province,
			gender:   b.gender,
			minAge:   b.minAge,
			maxAge:   b.maxAge,
		}
		if b.hasSeed {
			builder.seed = b.seed + int64(i)
			builder.hasSeed = true
		}
		persons[i] = builder.Build()
	}
	return persons
}

// generateLocation generates location information.
func (b *PersonBuilder) generateLocation(p *Person) {
	if b.province != "" {
		if prov, ok := metadata.ProvinceMap[b.province]; ok {
			p.province = prov.Short
			city := prov.Cities[b.rng.Intn(len(prov.Cities))]
			p.city = city.Name
			p.areaCode = city.AreaCodes[b.rng.Intn(len(city.AreaCodes))]
			return
		}
	}

	prov := metadata.Provinces[b.rng.Intn(len(metadata.Provinces))]
	p.province = prov.Short
	city := prov.Cities[b.rng.Intn(len(prov.Cities))]
	p.city = city.Name
	p.areaCode = city.AreaCodes[b.rng.Intn(len(city.AreaCodes))]
}

// generateGender generates gender.
func (b *PersonBuilder) generateGender(p *Person) {
	if b.gender != GenderRandom {
		p.gender = b.gender
		return
	}
	if b.rng.Intn(2) == 0 {
		p.gender = GenderMale
	} else {
		p.gender = GenderFemale
	}
}

// generateBirthday generates birthday.
func (b *PersonBuilder) generateBirthday(p *Person) {
	now := time.Now()
	minYear := now.Year() - b.maxAge
	maxYear := now.Year() - b.minAge

	year := b.rng.IntRange(minYear, maxYear+1)
	month := b.rng.IntRange(1, 13)
	day := b.rng.IntRange(1, 29)

	p.birthday = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

// generateIDNo generates the ID card number.
func (b *PersonBuilder) generateIDNo(p *Person) {
	birthday := p.birthday.Format("20060102")

	var seqCode int
	if p.gender == GenderMale {
		seqCode = b.rng.IntRange(0, 500)*2 + 1
	} else {
		seqCode = b.rng.IntRange(0, 500) * 2
	}

	idNo17 := fmt.Sprintf("%s%s%03d", p.areaCode, birthday, seqCode)
	checkCode := calculateCheckCode(idNo17)
	p.idNo = idNo17 + checkCode
}

// generateName generates the name.
func (b *PersonBuilder) generateName(p *Person) {
	if b.rng.Intn(100) < 97 { // 97% 单姓, 3% 复姓
		p.lastName = b.rng.Choice(metadata.SingleLastName)
	} else {
		p.lastName = b.rng.Choice(metadata.CompoundLastName)
	}

	if p.gender == GenderMale {
		p.firstName = b.rng.Choice(metadata.MaleFirstNames)
	} else {
		p.firstName = b.rng.Choice(metadata.FemaleFirstNames)
	}

	p.name = p.lastName + p.firstName
}

// generateAddress generates the address.
func (b *PersonBuilder) generateAddress(p *Person) {
	street := b.rng.Choice(metadata.StreetNames)
	community := b.rng.Choice(metadata.CommunityNames)

	houseNum := b.rng.IntRange(1, 201)
	unit := b.rng.IntRange(1, 9)
	var floor int
	if b.rng.Intn(10) < 8 { // 80%: 1-9 层 (3位房号)
		floor = b.rng.IntRange(1, 10)
	} else { // 20%: 10-25 层 (4位房号)
		floor = b.rng.IntRange(10, 26)
	}
	room := b.rng.IntRange(1, 5)
	roomNo := floor*100 + room

	p.address = fmt.Sprintf("%s%s%s%d号%s%d单元%d室",
		p.province, p.city, street, houseNum, community, unit, roomNo)
}

// generateMobile generates the mobile phone number.
func (b *PersonBuilder) generateMobile(p *Person) {
	prefix := b.rng.Choice(metadata.MobilePrefix)
	suffix := b.rng.IntRange(10000000, 100000000)
	p.mobile = fmt.Sprintf("%s%d", prefix, suffix)
}

// generateBankNo generates the bank card number.
func (b *PersonBuilder) generateBankNo(p *Person) {
	bank := metadata.CardBins[b.rng.Intn(len(metadata.CardBins))]
	prefix := bank.Prefixes[b.rng.Intn(len(bank.Prefixes))]

	prefixStr := fmt.Sprintf("%d", prefix)
	remainLen := bank.Length - len(prefixStr) - 1

	cardNo := prefixStr
	for i := 0; i < remainLen; i++ {
		cardNo += fmt.Sprintf("%d", b.rng.Intn(10))
	}

	checkDigit := calculateLUHNCheckDigit(cardNo)
	p.bankNo = cardNo + fmt.Sprintf("%d", checkDigit)
}

// generateEmail generates the email address.
func (b *PersonBuilder) generateEmail(p *Person) {
	var prefix string

	if b.rng.Intn(2) == 0 {
		prefix = b.generatePinyinPrefix(p)
	} else {
		prefix = b.rng.Choice(metadata.EmailPrefixes)
	}

	num := b.rng.IntRange(1000, 99999999) // 4-8 digits
	suffix := b.rng.Choice(metadata.EmailSuffixes)

	p.email = fmt.Sprintf("%s%d@%s", prefix, num, suffix)
}

func (b *PersonBuilder) generatePinyinPrefix(p *Person) string {
	switch b.rng.Intn(4) {
	case 0:
		return ConvertPinyin(p.name)
	case 1:
		return ConvertPinyin(p.lastName)
	case 2:
		return ConvertPinyin(p.firstName)
	default:
		py := ConvertPinyin(p.firstName)
		if len(py) > 4 {
			return py[:b.rng.IntRange(2, len(py))]
		}
		return py
	}
}
