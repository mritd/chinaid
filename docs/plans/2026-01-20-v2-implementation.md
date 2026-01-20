# chinaid v2 重构实现计划

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** 将 chinaid 重构为 v2 版本，采用 Builder 链式调用模式，修复现有问题并增强功能。

**Architecture:** 以 Person 为核心生成器，各字段（身份证、姓名、地址等）保持一致性。每个 Builder 持有独立的随机源实现并发安全。使用 TDD 方式逐步实现。

**Tech Stack:** Go 1.21+, 标准库 math/rand

---

## Phase 1: 基础设施

### Task 1: 更新 go.mod 到 v2

**Files:**
- Modify: `go.mod`

**Step 1: 更新 module 路径**

```go
module github.com/mritd/chinaid/v2

go 1.21
```

**Step 2: 验证 module**

Run: `go mod tidy`
Expected: 无错误

**Step 3: Commit**

```bash
git add go.mod
git commit -m "chore: upgrade module to v2"
```

---

### Task 2: 创建并发安全随机数封装

**Files:**
- Create: `rand.go`
- Create: `rand_test.go`

**Step 1: 写测试**

```go
// rand_test.go
package chinaid

import (
    "sync"
    "testing"
)

func TestRngIntn(t *testing.T) {
    rng := NewRng()
    for i := 0; i < 100; i++ {
        n := rng.Intn(10)
        if n < 0 || n >= 10 {
            t.Errorf("Intn(10) = %d, want [0, 10)", n)
        }
    }
}

func TestRngWithSeed(t *testing.T) {
    rng1 := NewRngWithSeed(12345)
    rng2 := NewRngWithSeed(12345)

    for i := 0; i < 10; i++ {
        a := rng1.Intn(1000)
        b := rng2.Intn(1000)
        if a != b {
            t.Errorf("Same seed produced different results: %d vs %d", a, b)
        }
    }
}

func TestRngChoice(t *testing.T) {
    rng := NewRng()
    items := []string{"a", "b", "c"}

    for i := 0; i < 100; i++ {
        choice := rng.Choice(items)
        found := false
        for _, item := range items {
            if choice == item {
                found = true
                break
            }
        }
        if !found {
            t.Errorf("Choice returned %s, not in slice", choice)
        }
    }
}

func TestRngConcurrentSafety(t *testing.T) {
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            rng := NewRng()
            _ = rng.Intn(100)
        }()
    }
    wg.Wait()
}
```

**Step 2: 运行测试验证失败**

Run: `go test -run TestRng -v`
Expected: FAIL (undefined: NewRng)

**Step 3: 实现 rand.go**

```go
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
```

**Step 4: 运行测试验证通过**

Run: `go test -run TestRng -v`
Expected: PASS

**Step 5: Commit**

```bash
git add rand.go rand_test.go
git commit -m "feat: add concurrent-safe random number generator"
```

---

### Task 3: 创建 Gender 枚举和基础类型

**Files:**
- Create: `types.go`
- Create: `types_test.go`

**Step 1: 写测试**

```go
// types_test.go
package chinaid

import "testing"

func TestGenderString(t *testing.T) {
    tests := []struct {
        g    Gender
        want string
    }{
        {GenderRandom, "random"},
        {GenderMale, "male"},
        {GenderFemale, "female"},
    }

    for _, tt := range tests {
        if got := tt.g.String(); got != tt.want {
            t.Errorf("Gender.String() = %s, want %s", got, tt.want)
        }
    }
}
```

**Step 2: 运行测试验证失败**

Run: `go test -run TestGender -v`
Expected: FAIL (undefined: Gender)

**Step 3: 实现 types.go**

```go
// types.go
package chinaid

// Gender 性别枚举
type Gender int

const (
    GenderRandom Gender = iota // 随机
    GenderMale                 // 男
    GenderFemale               // 女
)

// String 返回性别的字符串表示
func (g Gender) String() string {
    switch g {
    case GenderMale:
        return "male"
    case GenderFemale:
        return "female"
    default:
        return "random"
    }
}

// IsMale 是否为男性
func (g Gender) IsMale() bool {
    return g == GenderMale
}

// IsFemale 是否为女性
func (g Gender) IsFemale() bool {
    return g == GenderFemale
}
```

**Step 4: 运行测试验证通过**

Run: `go test -run TestGender -v`
Expected: PASS

**Step 5: Commit**

```bash
git add types.go types_test.go
git commit -m "feat: add Gender enum type"
```

---

## Phase 2: 元数据更新

### Task 4: 更新地区码数据结构

**Files:**
- Modify: `metadata/area_code.go`

**Step 1: 创建新的数据结构**

```go
// metadata/area_code.go
package metadata

// Province 省份信息
type Province struct {
    Name   string // 省份全称："北京市"
    Short  string // 简称："北京"
    Code   string // 省级代码："11"
    Cities []City // 下属城市
}

// City 城市/区县信息
type City struct {
    Name      string   // 名称："朝阳区"
    AreaCodes []string // 6位地区码：["110105"]
}

// Provinces 全国省份数据
var Provinces = []Province{
    {
        Name:  "北京市",
        Short: "北京",
        Code:  "11",
        Cities: []City{
            {Name: "东城区", AreaCodes: []string{"110101"}},
            {Name: "西城区", AreaCodes: []string{"110102"}},
            {Name: "朝阳区", AreaCodes: []string{"110105"}},
            {Name: "丰台区", AreaCodes: []string{"110106"}},
            {Name: "石景山区", AreaCodes: []string{"110107"}},
            {Name: "海淀区", AreaCodes: []string{"110108"}},
            {Name: "门头沟区", AreaCodes: []string{"110109"}},
            {Name: "房山区", AreaCodes: []string{"110111"}},
            {Name: "通州区", AreaCodes: []string{"110112"}},
            {Name: "顺义区", AreaCodes: []string{"110113"}},
            {Name: "昌平区", AreaCodes: []string{"110114"}},
            {Name: "大兴区", AreaCodes: []string{"110115"}},
            {Name: "怀柔区", AreaCodes: []string{"110116"}},
            {Name: "平谷区", AreaCodes: []string{"110117"}},
            {Name: "密云区", AreaCodes: []string{"110118"}},
            {Name: "延庆区", AreaCodes: []string{"110119"}},
        },
    },
    {
        Name:  "天津市",
        Short: "天津",
        Code:  "12",
        Cities: []City{
            {Name: "和平区", AreaCodes: []string{"120101"}},
            {Name: "河东区", AreaCodes: []string{"120102"}},
            {Name: "河西区", AreaCodes: []string{"120103"}},
            {Name: "南开区", AreaCodes: []string{"120104"}},
            {Name: "河北区", AreaCodes: []string{"120105"}},
            {Name: "红桥区", AreaCodes: []string{"120106"}},
            {Name: "东丽区", AreaCodes: []string{"120110"}},
            {Name: "西青区", AreaCodes: []string{"120111"}},
            {Name: "津南区", AreaCodes: []string{"120112"}},
            {Name: "北辰区", AreaCodes: []string{"120113"}},
            {Name: "武清区", AreaCodes: []string{"120114"}},
            {Name: "宝坻区", AreaCodes: []string{"120115"}},
            {Name: "滨海新区", AreaCodes: []string{"120116"}},
            {Name: "宁河区", AreaCodes: []string{"120117"}},
            {Name: "静海区", AreaCodes: []string{"120118"}},
            {Name: "蓟州区", AreaCodes: []string{"120119"}},
        },
    },
    {
        Name:  "上海市",
        Short: "上海",
        Code:  "31",
        Cities: []City{
            {Name: "黄浦区", AreaCodes: []string{"310101"}},
            {Name: "徐汇区", AreaCodes: []string{"310104"}},
            {Name: "长宁区", AreaCodes: []string{"310105"}},
            {Name: "静安区", AreaCodes: []string{"310106"}},
            {Name: "普陀区", AreaCodes: []string{"310107"}},
            {Name: "虹口区", AreaCodes: []string{"310109"}},
            {Name: "杨浦区", AreaCodes: []string{"310110"}},
            {Name: "闵行区", AreaCodes: []string{"310112"}},
            {Name: "宝山区", AreaCodes: []string{"310113"}},
            {Name: "嘉定区", AreaCodes: []string{"310114"}},
            {Name: "浦东新区", AreaCodes: []string{"310115"}},
            {Name: "金山区", AreaCodes: []string{"310116"}},
            {Name: "松江区", AreaCodes: []string{"310117"}},
            {Name: "青浦区", AreaCodes: []string{"310118"}},
            {Name: "奉贤区", AreaCodes: []string{"310120"}},
            {Name: "崇明区", AreaCodes: []string{"310151"}},
        },
    },
    {
        Name:  "广东省",
        Short: "广东",
        Code:  "44",
        Cities: []City{
            {Name: "广州市", AreaCodes: []string{"440100", "440103", "440104", "440105"}},
            {Name: "深圳市", AreaCodes: []string{"440300", "440303", "440304", "440305"}},
            {Name: "珠海市", AreaCodes: []string{"440400"}},
            {Name: "汕头市", AreaCodes: []string{"440500"}},
            {Name: "佛山市", AreaCodes: []string{"440600"}},
            {Name: "东莞市", AreaCodes: []string{"441900"}},
            {Name: "中山市", AreaCodes: []string{"442000"}},
            {Name: "惠州市", AreaCodes: []string{"441300"}},
        },
    },
    // TODO: 补充其他省份数据（Task 5 中详细处理）
}

// ProvinceMap 省份名称到 Province 的映射（用于快速查找）
var ProvinceMap map[string]*Province

// AreaCodeMap 6位地区码到省市信息的映射
var AreaCodeMap map[string]struct {
    Province string
    City     string
}

func init() {
    ProvinceMap = make(map[string]*Province)
    AreaCodeMap = make(map[string]struct {
        Province string
        City     string
    })

    for i := range Provinces {
        p := &Provinces[i]
        ProvinceMap[p.Name] = p
        ProvinceMap[p.Short] = p

        for _, city := range p.Cities {
            for _, code := range city.AreaCodes {
                AreaCodeMap[code] = struct {
                    Province string
                    City     string
                }{Province: p.Short, City: city.Name}
            }
        }
    }
}
```

**Step 2: 验证编译通过**

Run: `go build ./...`
Expected: 无错误

**Step 3: Commit**

```bash
git add metadata/area_code.go
git commit -m "feat(metadata): add structured province/city data"
```

---

### Task 5: 补充完整省份数据

**Files:**
- Modify: `metadata/area_code.go`

**说明:** 从 [modood/Administrative-divisions-of-China](https://github.com/modood/Administrative-divisions-of-China) 获取完整数据，补充所有 34 个省级行政区的区县数据。

**Step 1: 补充剩余省份数据**

（由于数据量大，实现时需要从外部数据源提取并格式化）

**Step 2: 验证数据完整性**

Run: `go test -run TestProvinceData -v`

**Step 3: Commit**

```bash
git add metadata/area_code.go
git commit -m "feat(metadata): complete all province/city data"
```

---

### Task 6: 创建分性别的姓名数据

**Files:**
- Modify: `metadata/first_name.go`

**Step 1: 重构姓名数据结构**

```go
// metadata/first_name.go
package metadata

// MaleFirstNames 男性常用名字
var MaleFirstNames = []string{
    // 单字名
    "伟", "强", "磊", "军", "勇", "杰", "涛", "明", "超", "华",
    "刚", "辉", "波", "斌", "鹏", "飞", "峰", "毅", "威", "浩",
    "亮", "健", "宁", "俊", "凯", "林", "龙", "雷", "兵", "锋",
    // 双字名
    "建国", "建华", "建军", "建平", "志强", "志伟", "志明", "志刚",
    "文华", "文明", "文军", "文斌", "国强", "国华", "国平", "国庆",
    "永强", "永明", "永华", "永刚", "海涛", "海波", "海军", "海龙",
    "浩然", "浩宇", "浩天", "子轩", "子涵", "子豪", "宇轩", "宇航",
    "俊杰", "俊豪", "俊熙", "晨曦", "晨阳", "天佑", "天翔", "天宇",
    // ... 继续添加至 5000-7000 个
}

// FemaleFirstNames 女性常用名字
var FemaleFirstNames = []string{
    // 单字名
    "芳", "娜", "敏", "静", "丽", "艳", "霞", "燕", "玲", "娟",
    "萍", "红", "梅", "琴", "英", "华", "慧", "莉", "蓉", "洁",
    "颖", "婷", "雪", "琳", "璐", "倩", "薇", "妍", "瑶", "蕾",
    // 双字名
    "秀英", "秀兰", "秀芳", "秀珍", "玉兰", "玉华", "玉珍", "玉梅",
    "桂英", "桂兰", "桂芳", "桂珍", "淑华", "淑珍", "淑芳", "淑兰",
    "欣怡", "欣妍", "欣悦", "紫萱", "紫涵", "紫琪", "梦琪", "梦瑶",
    "雨婷", "雨萱", "雨欣", "诗涵", "诗琪", "诗雨", "子涵", "子萱",
    "若曦", "若萱", "若琳", "思涵", "思琪", "思雨", "语嫣", "语涵",
    // ... 继续添加至 5000-7000 个
}
```

**Step 2: 验证数据**

Run: `go build ./metadata`
Expected: 无错误

**Step 3: Commit**

```bash
git add metadata/first_name.go
git commit -m "feat(metadata): split first names by gender"
```

---

### Task 7: 创建路名和小区名词库

**Files:**
- Create: `metadata/street_name.go`
- Create: `metadata/community_name.go`

**Step 1: 创建路名词库**

```go
// metadata/street_name.go
package metadata

// StreetNames 街道/路名词库
var StreetNames = []string{
    // 方位类
    "东大街", "西大街", "南大街", "北大街", "中大街",
    "东路", "西路", "南路", "北路", "中路",
    "东街", "西街", "南街", "北街", "中街",

    // 人民/解放类
    "人民路", "人民大道", "人民街", "解放路", "解放大道",
    "胜利路", "胜利街", "和平路", "和平街", "团结路",

    // 建设/发展类
    "建设路", "建设街", "建国路", "新华路", "新华街",
    "发展路", "振兴路", "复兴路", "创业路", "富强路",

    // 中山/中华类
    "中山路", "中山大道", "中山街", "中华路", "民族路",

    // 美好寓意类
    "幸福路", "幸福街", "康乐路", "安康路", "祥和路",
    "文明路", "文化路", "科技路", "学府路", "书院街",

    // 自然景观类
    "滨江路", "滨河路", "湖滨路", "江滨路", "海滨路",
    "山水路", "翠竹路", "梧桐路", "银杏路", "香樟路",
    "春风路", "朝阳路", "阳光路", "明月路", "星光路",

    // 花木类
    "桃花路", "梨花街", "荷花路", "菊花街", "玫瑰路",
    "牡丹路", "兰花街", "桂花路", "茉莉街", "百合路",

    // 数字类
    "一马路", "二马路", "三马路", "四马路", "五马路",
    "第一大街", "第二大街", "第三大街",

    // 工业/商业类
    "工业路", "工业大道", "商业街", "商贸路", "金融街",
    "经济开发区路", "高新路", "科园路", "产业路",

    // 交通类
    "火车站路", "机场路", "港口路", "码头街", "站前路",

    // ... 继续添加至 500+ 条
}
```

**Step 2: 创建小区名词库**

```go
// metadata/community_name.go
package metadata

// CommunityNames 小区/住宅名词库
var CommunityNames = []string{
    // 花园类
    "阳光花园", "世纪花园", "金色花园", "翠苑花园", "怡景花园",
    "温馨花园", "幸福花园", "和谐花园", "美丽花园", "锦绣花园",
    "春天花园", "夏日花园", "秋实花园", "冬青花园", "四季花园",

    // 小区类
    "幸福小区", "和谐小区", "康乐小区", "安居小区", "宜居小区",
    "温馨小区", "祥和小区", "平安小区", "民乐小区", "兴旺小区",

    // 苑/庭类
    "书香苑", "怡景苑", "锦绣苑", "翠竹苑", "丹桂苑",
    "金桂苑", "银杏苑", "香樟苑", "紫荆苑", "玉兰苑",
    "锦绣庭", "雅致庭", "静雅庭", "和美庭", "祥瑞庭",

    // 城/府类
    "国际城", "未来城", "理想城", "阳光城", "幸福城",
    "世纪府", "华府", "御府", "名府", "公馆",

    // 湾/岸类
    "江景湾", "河景湾", "湖景湾", "海景湾", "半岛湾",
    "滨江岸", "临河岸", "望湖岸", "听涛岸", "观澜岸",

    // 现代命名
    "中央公馆", "都市华庭", "时代广场", "名门世家", "壹号院",
    "尚品国际", "金地名苑", "万科城", "碧桂园", "恒大名都",

    // 雅致类
    "清雅居", "静雅居", "逸雅居", "和雅居", "文雅居",
    "兰亭", "竹园", "梅苑", "菊园", "荷塘月色",

    // ... 继续添加至 500+ 条
}
```

**Step 3: 验证编译**

Run: `go build ./metadata`
Expected: 无错误

**Step 4: Commit**

```bash
git add metadata/street_name.go metadata/community_name.go
git commit -m "feat(metadata): add street and community name vocabularies"
```

---

### Task 8: 创建邮箱相关词库

**Files:**
- Modify: `metadata/domain_prefix.go` -> `metadata/email_suffix.go`
- Create: `metadata/email_prefix.go`

**Step 1: 更新邮箱后缀**

```go
// metadata/email_suffix.go
package metadata

// EmailSuffixes 常用邮箱后缀
var EmailSuffixes = []string{
    // 国内主流
    "qq.com",
    "163.com",
    "126.com",
    "sina.com",
    "sohu.com",
    "139.com",
    "189.com",
    "foxmail.com",
    "aliyun.com",
    "yeah.net",

    // 国际主流
    "gmail.com",
    "outlook.com",
    "hotmail.com",
    "yahoo.com",
    "icloud.com",
}
```

**Step 2: 创建邮箱前缀词库**

```go
// metadata/email_prefix.go
package metadata

// EmailPrefixes 常用邮箱前缀（非拼音类）
var EmailPrefixes = []string{
    // 测试/管理类
    "test", "admin", "user", "demo", "dev", "test123",

    // 情感类
    "happy", "lucky", "love", "cool", "nice", "sweet",
    "sunny", "candy", "angel", "dream", "hope", "wish",

    // 自然类
    "sky", "sun", "moon", "star", "snow", "rain", "wind",
    "cloud", "ocean", "river", "forest", "mountain",

    // 动物类
    "tiger", "lion", "eagle", "wolf", "bear", "fox",
    "rabbit", "cat", "dog", "bird", "fish", "dragon",

    // 颜色类
    "blue", "red", "green", "gold", "silver", "black", "white",

    // 品质类
    "super", "best", "good", "great", "top", "pro", "vip",
    "king", "queen", "master", "expert",

    // 科技类
    "cyber", "tech", "code", "data", "web", "net", "digital",

    // 时间类
    "today", "forever", "always", "ever", "now",

    // 简单词
    "hello", "hi", "me", "my", "your", "one", "first",
}
```

**Step 3: 删除旧文件（如果存在）**

Run: `rm -f metadata/domain_prefix.go` (如果需要)

**Step 4: Commit**

```bash
git add metadata/email_suffix.go metadata/email_prefix.go
git rm -f metadata/domain_prefix.go 2>/dev/null || true
git commit -m "feat(metadata): add email suffix and prefix vocabularies"
```

---

### Task 9: 创建拼音映射表

**Files:**
- Create: `metadata/pinyin_map.go`

**Step 1: 创建拼音映射**

```go
// metadata/pinyin_map.go
package metadata

// PinyinMap 汉字到拼音的映射（覆盖所有姓氏和常用名字用字）
var PinyinMap = map[rune]string{
    // 常见姓氏
    '李': "li", '王': "wang", '张': "zhang", '刘': "liu", '陈': "chen",
    '杨': "yang", '黄': "huang", '赵': "zhao", '周': "zhou", '吴': "wu",
    '徐': "xu", '孙': "sun", '朱': "zhu", '马': "ma", '胡': "hu",
    '郭': "guo", '林': "lin", '何': "he", '高': "gao", '梁': "liang",
    '郑': "zheng", '罗': "luo", '宋': "song", '谢': "xie", '唐': "tang",
    '韦': "wei", '曹': "cao", '许': "xu", '邓': "deng", '萧': "xiao",
    '冯': "feng", '曾': "zeng", '程': "cheng", '蔡': "cai", '彭': "peng",
    '潘': "pan", '袁': "yuan", '于': "yu", '董': "dong", '余': "yu",
    '苏': "su", '叶': "ye", '吕': "lv", '魏': "wei", '蒋': "jiang",
    '田': "tian", '杜': "du", '丁': "ding", '沈': "shen", '姜': "jiang",
    '范': "fan", '江': "jiang", '傅': "fu", '钟': "zhong", '卢': "lu",
    '汪': "wang", '戴': "dai", '崔': "cui", '任': "ren", '陆': "lu",
    '廖': "liao", '姚': "yao", '方': "fang", '金': "jin", '邱': "qiu",
    '夏': "xia", '谭': "tan", '韩': "han", '贾': "jia", '邹': "zou",
    '石': "shi", '熊': "xiong", '孟': "meng", '秦': "qin", '阎': "yan",
    '薛': "xue", '侯': "hou", '雷': "lei", '白': "bai", '龙': "long",
    '段': "duan", '郝': "hao", '孔': "kong", '邵': "shao", '史': "shi",
    '毛': "mao", '常': "chang", '万': "wan", '顾': "gu", '赖': "lai",
    '武': "wu", '康': "kang", '贺': "he", '严': "yan", '尹': "yin",
    '钱': "qian", '施': "shi", '牛': "niu", '洪': "hong", '龚': "gong",

    // 复姓
    '欧': "ou", '阳': "yang", // 欧阳
    '司': "si",               // 司马、司徒
    '上': "shang", '官': "guan", // 上官
    '诸': "zhu", '葛': "ge",   // 诸葛

    // 常用名字用字 - 男
    '伟': "wei", '强': "qiang", '磊': "lei", '军': "jun", '勇': "yong",
    '杰': "jie", '涛': "tao", '明': "ming", '超': "chao", '华': "hua",
    '刚': "gang", '辉': "hui", '波': "bo", '斌': "bin", '鹏': "peng",
    '飞': "fei", '峰': "feng", '毅': "yi", '威': "wei", '浩': "hao",
    '亮': "liang", '健': "jian", '宁': "ning", '俊': "jun", '凯': "kai",
    '龙': "long", '兵': "bing", '锋': "feng", '翔': "xiang", '宇': "yu",
    '轩': "xuan", '豪': "hao", '天': "tian", '佑': "you", '航': "hang",
    '晨': "chen", '曦': "xi", '然': "ran", '睿': "rui", '博': "bo",

    // 常用名字用字 - 女
    '芳': "fang", '娜': "na", '敏': "min", '静': "jing", '丽': "li",
    '艳': "yan", '霞': "xia", '燕': "yan", '玲': "ling", '娟': "juan",
    '萍': "ping", '红': "hong", '梅': "mei", '琴': "qin", '英': "ying",
    '慧': "hui", '莉': "li", '蓉': "rong", '洁': "jie", '颖': "ying",
    '婷': "ting", '雪': "xue", '琳': "lin", '璐': "lu", '倩': "qian",
    '薇': "wei", '妍': "yan", '瑶': "yao", '蕾': "lei", '涵': "han",
    '萱': "xuan", '琪': "qi", '欣': "xin", '怡': "yi", '悦': "yue",
    '诗': "shi", '语': "yu", '嫣': "yan", '若': "ruo", '思': "si",

    // 常用字
    '国': "guo", '建': "jian", '文': "wen", '永': "yong", '海': "hai",
    '子': "zi", '小': "xiao", '大': "da", '中': "zhong", '新': "xin",
    '志': "zhi", '学': "xue", '成': "cheng", '平': "ping", '春': "chun",
    '秀': "xiu", '玉': "yu", '桂': "gui", '淑': "shu", '紫': "zi",
    '梦': "meng", '雨': "yu", '月': "yue", '心': "xin", '美': "mei",
    '爱': "ai", '兰': "lan", '珍': "zhen", '珠': "zhu", '云': "yun",
}
```

**Step 2: 验证编译**

Run: `go build ./metadata`
Expected: 无错误

**Step 3: Commit**

```bash
git add metadata/pinyin_map.go
git commit -m "feat(metadata): add pinyin mapping for name characters"
```

---

### Task 10: 创建常用汉字表

**Files:**
- Create: `metadata/common_chars.go`

**Step 1: 创建常用汉字表**

```go
// metadata/common_chars.go
package metadata

// CommonChars 地址生成用常用汉字
var CommonChars = []rune{
    // 方位
    '东', '西', '南', '北', '中', '前', '后', '左', '右', '上', '下',

    // 美好寓意
    '福', '禄', '寿', '喜', '财', '安', '康', '宁', '和', '顺',
    '祥', '瑞', '吉', '庆', '乐', '欢', '美', '善', '德', '仁',

    // 自然
    '山', '水', '河', '江', '湖', '海', '林', '森', '园', '苑',
    '花', '草', '木', '竹', '松', '柏', '梅', '兰', '菊', '荷',
    '春', '夏', '秋', '冬', '风', '云', '雨', '雪', '日', '月',

    // 建筑
    '楼', '阁', '亭', '台', '院', '庭', '堂', '馆', '城', '府',

    // 颜色
    '红', '黄', '蓝', '绿', '紫', '青', '白', '金', '银', '翠',

    // 品质
    '新', '华', '盛', '兴', '旺', '隆', '泰', '丰', '茂', '荣',
    '雅', '静', '清', '明', '光', '辉', '耀', '灿', '锦', '绣',

    // 数字相关
    '一', '二', '三', '四', '五', '六', '七', '八', '九', '十', '百', '千', '万',

    // 常用动词/形容词
    '大', '小', '长', '高', '远', '近', '通', '达', '利', '顺',
}
```

**Step 2: 验证编译**

Run: `go build ./metadata`
Expected: 无错误

**Step 3: Commit**

```bash
git add metadata/common_chars.go
git commit -m "feat(metadata): add common Chinese characters for address generation"
```

---

## Phase 3: 核心功能实现

### Task 11: 创建拼音转换模块

**Files:**
- Create: `pinyin.go`
- Create: `pinyin_test.go`

**Step 1: 写测试**

```go
// pinyin_test.go
package chinaid

import "testing"

func TestConvertPinyin(t *testing.T) {
    tests := []struct {
        input string
        want  string
    }{
        {"王", "wang"},
        {"李明", "liming"},
        {"张伟", "zhangwei"},
        {"欧阳", "ouyang"},
    }

    for _, tt := range tests {
        got := ConvertPinyin(tt.input)
        if got != tt.want {
            t.Errorf("ConvertPinyin(%s) = %s, want %s", tt.input, got, tt.want)
        }
    }
}
```

**Step 2: 运行测试验证失败**

Run: `go test -run TestConvertPinyin -v`
Expected: FAIL

**Step 3: 实现 pinyin.go**

```go
// pinyin.go
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

// ConvertPinyinFirst 将中文转换为拼音，只取第一个字
func ConvertPinyinFirst(chinese string) string {
    for _, r := range chinese {
        if py, ok := metadata.PinyinMap[r]; ok {
            return py
        }
    }
    return ""
}
```

**Step 4: 运行测试验证通过**

Run: `go test -run TestConvertPinyin -v`
Expected: PASS

**Step 5: Commit**

```bash
git add pinyin.go pinyin_test.go
git commit -m "feat: add pinyin conversion module"
```

---

### Task 12: 创建 Person 和 PersonBuilder

**Files:**
- Create: `person.go`
- Create: `person_test.go`

**Step 1: 写基础测试**

```go
// person_test.go
package chinaid

import (
    "strings"
    "testing"
)

func TestNewPerson(t *testing.T) {
    p := NewPerson().Build()

    if p.IDNo() == "" {
        t.Error("IDNo should not be empty")
    }
    if p.Name() == "" {
        t.Error("Name should not be empty")
    }
    if p.Address() == "" {
        t.Error("Address should not be empty")
    }
}

func TestPersonProvince(t *testing.T) {
    p := NewPerson().Province("北京").Build()

    if !strings.HasPrefix(p.IDNo(), "11") {
        t.Errorf("IDNo should start with 11 for Beijing, got %s", p.IDNo())
    }
    if !strings.HasPrefix(p.Address(), "北京") {
        t.Errorf("Address should start with 北京, got %s", p.Address())
    }
}

func TestPersonGender(t *testing.T) {
    p := NewPerson().Gender(GenderMale).Build()

    // 身份证第17位：奇数为男，偶数为女
    digit17 := int(p.IDNo()[16] - '0')
    if digit17%2 != 1 {
        t.Errorf("Male IDNo 17th digit should be odd, got %d", digit17)
    }
}

func TestPersonSeed(t *testing.T) {
    p1 := NewPerson().Seed(12345).Build()
    p2 := NewPerson().Seed(12345).Build()

    if p1.IDNo() != p2.IDNo() {
        t.Errorf("Same seed should produce same IDNo: %s vs %s", p1.IDNo(), p2.IDNo())
    }
    if p1.Name() != p2.Name() {
        t.Errorf("Same seed should produce same Name: %s vs %s", p1.Name(), p2.Name())
    }
}

func TestPersonBuildN(t *testing.T) {
    persons := NewPerson().Province("广东").BuildN(10)

    if len(persons) != 10 {
        t.Errorf("BuildN(10) should return 10 persons, got %d", len(persons))
    }

    for i, p := range persons {
        if !strings.HasPrefix(p.Address(), "广东") {
            t.Errorf("Person %d address should start with 广东, got %s", i, p.Address())
        }
    }
}
```

**Step 2: 运行测试验证失败**

Run: `go test -run TestPerson -v`
Expected: FAIL

**Step 3: 实现 person.go（框架）**

```go
// person.go
package chinaid

import (
    "time"
)

// Person 生成的人物信息
type Person struct {
    idNo      string
    name      string
    lastName  string
    firstName string
    gender    Gender
    birthday  time.Time
    province  string
    city      string
    areaCode  string
    address   string
    mobile    string
    bankNo    string
    email     string
}

// Getter 方法
func (p *Person) IDNo() string       { return p.idNo }
func (p *Person) Name() string       { return p.name }
func (p *Person) LastName() string   { return p.lastName }
func (p *Person) FirstName() string  { return p.firstName }
func (p *Person) Gender() Gender     { return p.gender }
func (p *Person) Birthday() time.Time { return p.birthday }
func (p *Person) Age() int {
    now := time.Now()
    age := now.Year() - p.birthday.Year()
    if now.YearDay() < p.birthday.YearDay() {
        age--
    }
    return age
}
func (p *Person) Province() string { return p.province }
func (p *Person) City() string     { return p.city }
func (p *Person) AreaCode() string { return p.areaCode }
func (p *Person) Address() string  { return p.address }
func (p *Person) Mobile() string   { return p.mobile }
func (p *Person) BankNo() string   { return p.bankNo }
func (p *Person) Email() string    { return p.email }

// PersonBuilder 人物构建器
type PersonBuilder struct {
    rng      *Rng
    seed     int64
    hasSeed  bool
    province string
    gender   Gender
    minAge   int
    maxAge   int
}

// NewPerson 创建人物构建器
func NewPerson() *PersonBuilder {
    return &PersonBuilder{
        minAge: 18,
        maxAge: 60,
    }
}

// Province 设置省份
func (b *PersonBuilder) Province(province string) *PersonBuilder {
    b.province = province
    return b
}

// Gender 设置性别
func (b *PersonBuilder) Gender(gender Gender) *PersonBuilder {
    b.gender = gender
    return b
}

// AgeRange 设置年龄范围
func (b *PersonBuilder) AgeRange(min, max int) *PersonBuilder {
    b.minAge = min
    b.maxAge = max
    return b
}

// Seed 设置随机种子（可复现）
func (b *PersonBuilder) Seed(seed int64) *PersonBuilder {
    b.seed = seed
    b.hasSeed = true
    return b
}

// Build 生成单个人物
func (b *PersonBuilder) Build() *Person {
    // 初始化随机数生成器
    if b.hasSeed {
        b.rng = NewRngWithSeed(b.seed)
    } else {
        b.rng = NewRng()
    }

    p := &Person{}

    // 按顺序生成各字段
    b.generateLocation(p)
    b.generateGender(p)
    b.generateBirthday(p)
    b.generateIDNo(p)
    b.generateName(p)
    b.generateAddress(p)
    b.generateMobile(p)
    b.generateBankNo(p)
    b.generateEmail(p)

    return p
}

// BuildN 批量生成人物
func (b *PersonBuilder) BuildN(n int) []*Person {
    persons := make([]*Person, n)
    for i := 0; i < n; i++ {
        // 每次创建新的 builder 副本，保持配置但使用不同随机数
        builder := &PersonBuilder{
            province: b.province,
            gender:   b.gender,
            minAge:   b.minAge,
            maxAge:   b.maxAge,
        }
        if b.hasSeed {
            builder.seed = b.seed + int64(i)
            builder.hasSeed = true
        }
        persons[i] = builder.Build()
    }
    return persons
}

// 内部生成方法（后续 Task 实现）
func (b *PersonBuilder) generateLocation(p *Person)  { /* Task 13 */ }
func (b *PersonBuilder) generateGender(p *Person)    { /* Task 13 */ }
func (b *PersonBuilder) generateBirthday(p *Person)  { /* Task 13 */ }
func (b *PersonBuilder) generateIDNo(p *Person)      { /* Task 14 */ }
func (b *PersonBuilder) generateName(p *Person)      { /* Task 15 */ }
func (b *PersonBuilder) generateAddress(p *Person)   { /* Task 16 */ }
func (b *PersonBuilder) generateMobile(p *Person)    { /* Task 17 */ }
func (b *PersonBuilder) generateBankNo(p *Person)    { /* Task 18 */ }
func (b *PersonBuilder) generateEmail(p *Person)     { /* Task 19 */ }
```

**Step 4: Commit 框架（测试暂不通过）**

```bash
git add person.go person_test.go
git commit -m "feat: add Person and PersonBuilder skeleton"
```

---

### Task 13: 实现位置、性别、生日生成

**Files:**
- Modify: `person.go`

**Step 1: 实现 generateLocation**

```go
func (b *PersonBuilder) generateLocation(p *Person) {
    if b.province != "" {
        // 用户指定了省份
        if prov, ok := metadata.ProvinceMap[b.province]; ok {
            p.province = prov.Short
            city := prov.Cities[b.rng.Intn(len(prov.Cities))]
            p.city = city.Name
            p.areaCode = city.AreaCodes[b.rng.Intn(len(city.AreaCodes))]
            return
        }
    }

    // 随机选择省份
    prov := metadata.Provinces[b.rng.Intn(len(metadata.Provinces))]
    p.province = prov.Short
    city := prov.Cities[b.rng.Intn(len(prov.Cities))]
    p.city = city.Name
    p.areaCode = city.AreaCodes[b.rng.Intn(len(city.AreaCodes))]
}
```

**Step 2: 实现 generateGender**

```go
func (b *PersonBuilder) generateGender(p *Person) {
    if b.gender != GenderRandom {
        p.gender = b.gender
        return
    }
    // 随机选择性别
    if b.rng.Intn(2) == 0 {
        p.gender = GenderMale
    } else {
        p.gender = GenderFemale
    }
}
```

**Step 3: 实现 generateBirthday**

```go
func (b *PersonBuilder) generateBirthday(p *Person) {
    now := time.Now()
    // 计算出生年份范围
    minYear := now.Year() - b.maxAge
    maxYear := now.Year() - b.minAge

    year := b.rng.IntRange(minYear, maxYear+1)
    month := b.rng.IntRange(1, 13)
    day := b.rng.IntRange(1, 29) // 简化：最多28天，避免月份天数问题

    p.birthday = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}
```

**Step 4: 运行测试**

Run: `go test -run TestPerson -v`
Expected: 部分通过

**Step 5: Commit**

```bash
git add person.go
git commit -m "feat: implement location, gender, birthday generation"
```

---

### Task 14: 实现身份证号生成

**Files:**
- Modify: `person.go`
- Create: `idcard.go`

**Step 1: 实现 generateIDNo**

```go
// person.go
func (b *PersonBuilder) generateIDNo(p *Person) {
    // 格式：6位地区码 + 8位出生日期 + 3位顺序码 + 1位校验码
    birthday := p.birthday.Format("20060102")

    // 顺序码：奇数为男，偶数为女
    var seqCode int
    if p.gender == GenderMale {
        seqCode = b.rng.IntRange(0, 500)*2 + 1 // 奇数: 1, 3, 5, ..., 999
    } else {
        seqCode = b.rng.IntRange(0, 500) * 2 // 偶数: 0, 2, 4, ..., 998
    }

    idNo17 := fmt.Sprintf("%s%s%03d", p.areaCode, birthday, seqCode)
    checkCode := calculateCheckCode(idNo17)

    p.idNo = idNo17 + checkCode
}
```

**Step 2: 创建 idcard.go**

```go
// idcard.go
package chinaid

// 身份证校验码权重
var idCardWeights = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

// 校验码对应值
var idCardCheckCodes = []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

// calculateCheckCode 计算身份证校验码
func calculateCheckCode(idNo17 string) string {
    sum := 0
    for i, w := range idCardWeights {
        sum += int(idNo17[i]-'0') * w
    }
    return idCardCheckCodes[sum%11]
}

// ValidateIDNo 验证身份证号是否有效
func ValidateIDNo(idNo string) bool {
    if len(idNo) != 18 {
        return false
    }

    // 验证前17位是否为数字
    for i := 0; i < 17; i++ {
        if idNo[i] < '0' || idNo[i] > '9' {
            return false
        }
    }

    // 验证校验码
    expectedCheck := calculateCheckCode(idNo[:17])
    actualCheck := string(idNo[17])
    if actualCheck == "x" {
        actualCheck = "X"
    }

    return expectedCheck == actualCheck
}
```

**Step 3: 添加需要的 import**

```go
import "fmt"
```

**Step 4: 运行测试**

Run: `go test -run TestPerson -v`

**Step 5: Commit**

```bash
git add person.go idcard.go
git commit -m "feat: implement ID card number generation with checksum"
```

---

### Task 15: 实现姓名生成

**Files:**
- Modify: `person.go`

**Step 1: 实现 generateName**

```go
func (b *PersonBuilder) generateName(p *Person) {
    p.lastName = b.rng.Choice(metadata.LastNames)

    if p.gender == GenderMale {
        p.firstName = b.rng.Choice(metadata.MaleFirstNames)
    } else {
        p.firstName = b.rng.Choice(metadata.FemaleFirstNames)
    }

    p.name = p.lastName + p.firstName
}
```

**Step 2: 运行测试**

Run: `go test -run TestPerson -v`

**Step 3: Commit**

```bash
git add person.go
git commit -m "feat: implement name generation with gender support"
```

---

### Task 16: 实现地址生成

**Files:**
- Modify: `person.go`

**Step 1: 实现 generateAddress**

```go
func (b *PersonBuilder) generateAddress(p *Person) {
    street := b.rng.Choice(metadata.StreetNames)
    community := b.rng.Choice(metadata.CommunityNames)

    houseNum := b.rng.IntRange(1, 201)
    unit := b.rng.IntRange(1, 9)
    floor := b.rng.IntRange(1, 31)
    room := b.rng.IntRange(1, 5)
    roomNo := floor*100 + room

    p.address = fmt.Sprintf("%s%s%s%d号%s%d单元%d室",
        p.province, p.city, street, houseNum, community, unit, roomNo)
}
```

**Step 2: 运行测试**

Run: `go test -run TestPerson -v`

**Step 3: Commit**

```bash
git add person.go
git commit -m "feat: implement address generation with vocabulary"
```

---

### Task 17: 实现手机号生成

**Files:**
- Modify: `person.go`

**Step 1: 实现 generateMobile**

```go
func (b *PersonBuilder) generateMobile(p *Person) {
    prefix := b.rng.Choice(metadata.MobilePrefix)
    suffix := b.rng.IntRange(10000000, 100000000) // 8位数字
    p.mobile = fmt.Sprintf("%s%d", prefix, suffix)
}
```

**Step 2: 运行测试**

Run: `go test -run TestPerson -v`

**Step 3: Commit**

```bash
git add person.go
git commit -m "feat: implement mobile number generation"
```

---

### Task 18: 实现银行卡号生成

**Files:**
- Modify: `person.go`
- Create: `bank.go`

**Step 1: 创建 bank.go**

```go
// bank.go
package chinaid

import (
    "strconv"

    "github.com/mritd/chinaid/v2/metadata"
)

// generateBankNo 生成银行卡号
func (b *PersonBuilder) generateBankNo(p *Person) {
    bank := metadata.CardBins[b.rng.Intn(len(metadata.CardBins))]
    prefix := bank.Prefixes[b.rng.Intn(len(bank.Prefixes))]

    // 生成卡号（不含校验位）
    prefixStr := strconv.Itoa(prefix)
    remainLen := bank.Length - len(prefixStr) - 1 // 减1是校验位

    // 生成中间的随机数字
    var cardNo string = prefixStr
    for i := 0; i < remainLen; i++ {
        cardNo += strconv.Itoa(b.rng.Intn(10))
    }

    // 计算 LUHN 校验位
    checkDigit := calculateLUHNCheckDigit(cardNo)
    p.bankNo = cardNo + strconv.Itoa(checkDigit)
}

// calculateLUHNCheckDigit 计算 LUHN 校验位
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

// ValidateLUHN 验证银行卡号 LUHN 校验
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
```

**Step 2: 运行测试**

Run: `go test -run TestPerson -v`

**Step 3: Commit**

```bash
git add person.go bank.go
git commit -m "feat: implement bank card number generation with LUHN"
```

---

### Task 19: 实现邮箱生成

**Files:**
- Modify: `person.go`

**Step 1: 实现 generateEmail**

```go
func (b *PersonBuilder) generateEmail(p *Person) {
    var prefix string

    if b.rng.Intn(2) == 0 {
        // 50%: 基于姓名拼音
        prefix = b.generatePinyinPrefix(p)
    } else {
        // 50%: 常用前缀词库
        prefix = b.rng.Choice(metadata.EmailPrefixes)
    }

    num := b.rng.Intn(10000)
    suffix := b.rng.Choice(metadata.EmailSuffixes)

    p.email = fmt.Sprintf("%s%d@%s", prefix, num, suffix)
}

func (b *PersonBuilder) generatePinyinPrefix(p *Person) string {
    switch b.rng.Intn(4) {
    case 0: // 全名拼音
        return ConvertPinyin(p.name)
    case 1: // 仅姓
        return ConvertPinyin(p.lastName)
    case 2: // 仅名
        return ConvertPinyin(p.firstName)
    default: // 名的部分
        py := ConvertPinyin(p.firstName)
        if len(py) > 4 {
            return py[:b.rng.IntRange(2, len(py))]
        }
        return py
    }
}
```

**Step 2: 运行测试**

Run: `go test -run TestPerson -v`
Expected: PASS

**Step 3: Commit**

```bash
git add person.go
git commit -m "feat: implement email generation with pinyin support"
```

---

## Phase 4: 测试完善

### Task 20: 添加完整测试

**Files:**
- Modify: `person_test.go`
- Create: `idcard_test.go`
- Create: `bank_test.go`

**Step 1: 完善 person_test.go**

```go
// 添加并发安全测试
func TestPersonConcurrentSafety(t *testing.T) {
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            p := NewPerson().Build()
            if p.IDNo() == "" {
                t.Error("IDNo should not be empty")
            }
        }()
    }
    wg.Wait()
}

func TestPersonAgeRange(t *testing.T) {
    p := NewPerson().AgeRange(20, 30).Build()
    age := p.Age()
    if age < 20 || age > 30 {
        t.Errorf("Age should be between 20 and 30, got %d", age)
    }
}
```

**Step 2: 创建 idcard_test.go**

```go
// idcard_test.go
package chinaid

import "testing"

func TestValidateIDNo(t *testing.T) {
    // 生成1000个身份证号，全部应该有效
    for i := 0; i < 1000; i++ {
        p := NewPerson().Build()
        if !ValidateIDNo(p.IDNo()) {
            t.Errorf("Generated IDNo is invalid: %s", p.IDNo())
        }
    }
}

func TestIDNoGender(t *testing.T) {
    // 测试性别编码
    for i := 0; i < 100; i++ {
        male := NewPerson().Gender(GenderMale).Build()
        if int(male.IDNo()[16]-'0')%2 != 1 {
            t.Errorf("Male IDNo should have odd 17th digit: %s", male.IDNo())
        }

        female := NewPerson().Gender(GenderFemale).Build()
        if int(female.IDNo()[16]-'0')%2 != 0 {
            t.Errorf("Female IDNo should have even 17th digit: %s", female.IDNo())
        }
    }
}
```

**Step 3: 创建 bank_test.go**

```go
// bank_test.go
package chinaid

import "testing"

func TestValidateLUHN(t *testing.T) {
    // 生成1000个银行卡号，全部应该有效
    for i := 0; i < 1000; i++ {
        p := NewPerson().Build()
        if !ValidateLUHN(p.BankNo()) {
            t.Errorf("Generated BankNo is invalid: %s", p.BankNo())
        }
    }
}
```

**Step 4: 运行所有测试**

Run: `go test ./... -v`
Expected: All PASS

**Step 5: Commit**

```bash
git add person_test.go idcard_test.go bank_test.go
git commit -m "test: add comprehensive tests for person, idcard, and bank"
```

---

## Phase 5: 清理和文档

### Task 21: 删除旧代码文件

**Files:**
- Delete: `generator_idcard.go`
- Delete: `generator_name.go`
- Delete: `generator_address.go`
- Delete: `generator_mobile.go`
- Delete: `generator_bank.go`
- Delete: `generator_email.go`
- Delete: `util_char.go`
- Delete: 旧测试文件

**Step 1: 删除旧文件**

```bash
rm -f generator_*.go util_char.go
rm -f generator_*_test.go
```

**Step 2: 验证编译和测试**

Run: `go build ./... && go test ./...`

**Step 3: Commit**

```bash
git add -A
git commit -m "chore: remove legacy v1 code"
```

---

### Task 22: 更新 README

**Files:**
- Modify: `README.md`

**Step 1: 更新 README**

```markdown
# chinaid

中国大陆身份证等信息生成库 v2，用于测试内部系统校验有效性。

## 安装

```bash
go get github.com/mritd/chinaid/v2
```

## 使用

### 生成完整人物信息

```go
import "github.com/mritd/chinaid/v2"

// 基本用法
person := chinaid.NewPerson().Build()

fmt.Println(person.IDNo())    // 身份证号
fmt.Println(person.Name())    // 姓名
fmt.Println(person.Gender())  // 性别
fmt.Println(person.Age())     // 年龄
fmt.Println(person.Address()) // 地址
fmt.Println(person.Mobile())  // 手机号
fmt.Println(person.BankNo())  // 银行卡号
fmt.Println(person.Email())   // 邮箱
```

### 链式配置

```go
// 指定省份和性别
person := chinaid.NewPerson().
    Province("北京").
    Gender(chinaid.GenderMale).
    AgeRange(25, 35).
    Build()

// 可复现随机（相同 seed 生成相同结果）
person := chinaid.NewPerson().
    Seed(12345).
    Build()

// 批量生成
persons := chinaid.NewPerson().
    Province("广东").
    BuildN(100)
```

### 数据验证

```go
// 验证身份证号
valid := chinaid.ValidateIDNo("110101199001011234")

// 验证银行卡号（LUHN 算法）
valid := chinaid.ValidateLUHN("6222021234567890123")
```

## 特性

- **Builder 模式**: 链式调用，灵活配置
- **数据一致性**: 身份证、地址、省份自动关联
- **并发安全**: 每个 Builder 独立随机源
- **可复现**: 支持 Seed 设置
- **批量生成**: 支持 BuildN(n)

## License

MIT
```

**Step 2: Commit**

```bash
git add README.md
git commit -m "docs: update README for v2"
```

---

### Task 23: 最终验证

**Step 1: 运行完整测试**

Run: `go test ./... -v -cover`

**Step 2: 运行 lint**

Run: `golangci-lint run`

**Step 3: 验证构建**

Run: `go build ./...`

---

## 执行检查点

| Phase | Tasks | 检查点 |
|-------|-------|--------|
| Phase 1 | Task 1-3 | go.mod v2, Rng, Gender 类型就绪 |
| Phase 2 | Task 4-10 | 所有 metadata 准备完成 |
| Phase 3 | Task 11-19 | Person Builder 完整实现 |
| Phase 4 | Task 20 | 测试覆盖完整 |
| Phase 5 | Task 21-23 | 清理完成，文档更新 |
