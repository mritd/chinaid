package chinaid

import (
	"github.com/mritd/chinaid/v2/metadata"
	"math/rand"
	"strconv"
)

// ProvinceAndCity 返回随机省/城市
func ProvinceAndCity() string {
	return metadata.ProvinceCity[rand.Intn(len(metadata.ProvinceCity))]
}

// Address 返回随机地址
func Address() string {
	return ProvinceAndCity() +
		randomLengthChineseChars(2, 3) + "路" +
		strconv.Itoa(randInt(1, 8000)) + "号" +
		randomLengthChineseChars(2, 3) + "小区" +
		strconv.Itoa(randInt(1, 20)) + "单元" +
		strconv.Itoa(randInt(101, 2500)) + "室"
}
