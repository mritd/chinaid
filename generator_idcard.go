package chinaid

import (
	"fmt"
	"github.com/mritd/chinaid/v2/metadata"
	"math/rand"
	"strconv"
	"time"
)

// IssueOrg 返回身份证签发机关(eg: XXX公安局/XX区分局)
func IssueOrg() string {
	return metadata.CityName[rand.Intn(len(metadata.CityName))] + "公安局某某分局"
}

// ValidPeriod 返回身份证有效期限(eg: 20150906-20350906)，有效期限固定为 20 年
func ValidPeriod() string {
	begin := RandDate()
	end := begin.AddDate(20, 0, 0)
	return begin.Format("20060102") + "-" + end.Format("20060102")
}

// IDNo 返回中国大陆地区身份证号
func IDNo() string {
	// AreaCode 随机一个+4位随机数字(不够左填充0)
	areaCode := metadata.AreaCode[rand.Intn(len(metadata.AreaCode))] +
		fmt.Sprintf("%0*d", 4, randInt(1, 9999))
	birthday := RandDate().Format("20060102")
	randomCode := fmt.Sprintf("%0*d", 3, randInt(0, 999))
	prefix := areaCode + birthday + randomCode
	return prefix + VerifyCode(prefix)
}

// VerifyCode 通过给定的身份证号生成最后一位的 VerifyCode
func VerifyCode(cardId string) string {
	tmp := 0
	for i, v := range metadata.Wi {
		t, _ := strconv.Atoi(string(cardId[i]))
		tmp += t * v
	}
	return metadata.ValCodeArr[tmp%11]
}

// RandDate 返回随机时间，时间区间从 1970 年 ~ 2020 年
func RandDate() time.Time {
	begin, _ := time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
	end, _ := time.Parse("2006-01-02 15:04:05", "2020-01-01 00:00:00")
	return time.Unix(randInt64(begin.Unix(), end.Unix()), 0)
}
