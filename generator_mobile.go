package chinaid

import (
	"fmt"
	"github.com/mritd/chinaid/metadata"
	"math/rand"
)

// Mobile 返回中国大陆地区手机号
func Mobile() string {
	return metadata.MobilePrefix[rand.Intn(len(metadata.MobilePrefix))] + fmt.Sprintf("%0*d", 8, randInt(0, 100000000))
}
