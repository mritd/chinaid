package chinaid

import (
	"github.com/mritd/chinaid/metadata"
)

// Name 返回中国姓名，姓名已经尽量返回常用姓氏和名字
func Name() string {
	return metadata.LastName[randInt(0, len(metadata.LastName))] + metadata.FirstName[randInt(0, len(metadata.LastName))]
}
