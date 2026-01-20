// rand.go
package chinaid

import (
	"math/rand"
	"time"
)

// Rng 并发安全的随机数生成器
// 每个实例持有独立的 rand.Rand，避免全局锁竞争
type Rng struct {
	r *rand.Rand
}

// NewRng 创建新的随机数生成器
func NewRng() *Rng {
	return &Rng{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// NewRngWithSeed 创建带种子的随机数生成器（可复现）
func NewRngWithSeed(seed int64) *Rng {
	return &Rng{
		r: rand.New(rand.NewSource(seed)),
	}
}

// Intn 返回 [0, n) 范围内的随机整数
func (rng *Rng) Intn(n int) int {
	if n <= 0 {
		return 0
	}
	return rng.r.Intn(n)
}

// Int63n 返回 [0, n) 范围内的随机 int64
func (rng *Rng) Int63n(n int64) int64 {
	if n <= 0 {
		return 0
	}
	return rng.r.Int63n(n)
}

// IntRange 返回 [min, max) 范围内的随机整数
func (rng *Rng) IntRange(min, max int) int {
	if min >= max {
		return min
	}
	return min + rng.r.Intn(max-min)
}

// Choice 从切片中随机选择一个元素
func (rng *Rng) Choice(slice []string) string {
	if len(slice) == 0 {
		return ""
	}
	return slice[rng.r.Intn(len(slice))]
}

// ChoiceRune 从 rune 切片中随机选择一个元素
func (rng *Rng) ChoiceRune(slice []rune) rune {
	if len(slice) == 0 {
		return 0
	}
	return slice[rng.r.Intn(len(slice))]
}
