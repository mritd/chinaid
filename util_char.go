package chinaid

import (
	"bytes"
	"math/rand"
	"time"
)

// 指定长度随机中文字符(包含复杂字符)
func fixedLengthChineseChars(length int) string {
	var buf bytes.Buffer
	for i := 0; i < length; i++ {
		buf.WriteRune(rune(randInt(19968, 40869)))
	}
	return buf.String()
}

// 指定范围随机中文字符
func randomLengthChineseChars(start, end int) string {
	return fixedLengthChineseChars(randInt(start, end))
}

// 随机英文小写字母
func randStr(len int) string {
	data := make([]byte, len)
	for i := 0; i < len; i++ {
		data[i] = byte(rand.Intn(26) + 97)
	}
	return string(data)
}

// 指定范围随机 int
func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// 指定范围随机 int64
func randInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}