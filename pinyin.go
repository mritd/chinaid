package chinaid

import (
	"strings"

	"github.com/mritd/chinaid/v2/metadata"
)

// ConvertPinyin 将中文转换为拼音
func ConvertPinyin(chinese string) string {
	var result strings.Builder
	for _, r := range chinese {
		if py, ok := metadata.PinyinMap[r]; ok {
			result.WriteString(py)
		}
	}
	return result.String()
}

// ConvertPinyinFirst 将中文转换为拼音，只取第一个字的拼音
func ConvertPinyinFirst(chinese string) string {
	for _, r := range chinese {
		if py, ok := metadata.PinyinMap[r]; ok {
			return py
		}
	}
	return ""
}
