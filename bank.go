package chinaid

func calculateLUHNCheckDigit(cardNo string) int {
	sum := 0
	for i := len(cardNo) - 1; i >= 0; i-- {
		digit := int(cardNo[i] - '0')
		if (len(cardNo)-i)%2 == 1 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return (10 - sum%10) % 10
}

// ValidateLUHN validates a bank card number using the LUHN algorithm.
func ValidateLUHN(cardNo string) bool {
	if len(cardNo) < 13 || len(cardNo) > 19 {
		return false
	}

	sum := 0
	for i := len(cardNo) - 1; i >= 0; i-- {
		digit := int(cardNo[i] - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		if (len(cardNo)-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
