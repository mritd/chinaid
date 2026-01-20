package chinaid

// Gender 性别枚举
type Gender int

const (
	GenderRandom Gender = iota // 随机
	GenderMale                 // 男
	GenderFemale               // 女
)

// String 返回性别的字符串表示
func (g Gender) String() string {
	switch g {
	case GenderMale:
		return "male"
	case GenderFemale:
		return "female"
	default:
		return "random"
	}
}

// IsMale 是否为男性
func (g Gender) IsMale() bool {
	return g == GenderMale
}

// IsFemale 是否为女性
func (g Gender) IsFemale() bool {
	return g == GenderFemale
}
