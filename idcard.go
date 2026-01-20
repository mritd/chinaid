// idcard.go
package chinaid

var idCardWeights = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var idCardCheckCodes = []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

func calculateCheckCode(idNo17 string) string {
	sum := 0
	for i, w := range idCardWeights {
		sum += int(idNo17[i]-'0') * w
	}
	return idCardCheckCodes[sum%11]
}

// ValidateIDNo validates whether the ID number is valid.
func ValidateIDNo(idNo string) bool {
	if len(idNo) != 18 {
		return false
	}

	for i := 0; i < 17; i++ {
		if idNo[i] < '0' || idNo[i] > '9' {
			return false
		}
	}

	expectedCheck := calculateCheckCode(idNo[:17])
	actualCheck := string(idNo[17])
	if actualCheck == "x" {
		actualCheck = "X"
	}

	return expectedCheck == actualCheck
}
