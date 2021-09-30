package chinaid

import (
	"github.com/mritd/chinaid/metadata"
	"math/rand"
)

// Name 返回中国姓名，姓名已经尽量返回常用姓氏和名字
func Name() string {
	return metadata.LastName[rand.Intn(len(metadata.LastName))] + metadata.FirstName[rand.Intn(len(metadata.FirstName))]
}
