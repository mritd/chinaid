# chinaid v2 重构设计

## 概述

chinaid 是一个中国大陆身份证等信息生成库，用于测试内部系统校验有效性。v2 版本进行完全重构，采用 Builder 链式调用模式，修复现有问题并增强功能。

## 现有问题

| 优先级 | 问题 | 影响 |
|--------|------|------|
| P0 | RandDate 固定到 2020 年 | 无法生成 2021+ 出生的人 |
| P0 | FirstName 数据有错误条目和重复 | 可能生成无效名字 |
| P1 | 地址生成使用全 Unicode 范围 | 生成生僻字地址 |
| P1 | 全局 rand 无并发保护 | 高并发下有竞态风险 |
| P1 | 测试只检查非空 | 无法发现数据有效性问题 |
| P2 | ProvinceCity 有重复 | 数据冗余 |

## 设计决策

| 项目 | 决策 |
|------|------|
| 架构 | Builder 链式调用模式 |
| 版本 | v2 大版本，不兼容旧 API |
| 核心 | Person 生成器，字段一致性 |
| 字段 | 身份证、姓名、地址、手机、银行卡、邮箱 |
| 手机归属地 | 不强制匹配省份 |
| 并发安全 | 需要，独立随机源 |
| 姓名数据 | 精简到 10000-15000 个，按性别分类 |
| 地址数据 | 词库(500+) + 常用字组合 |
| 邮箱 | 姓名拼音/词库 + 常用后缀 |
| Seed | 支持可复现随机 |
| 批量 | 支持 BuildN(n) |

## 包结构

```
chinaid/
├── go.mod                    # module github.com/mritd/chinaid/v2
├── person.go                 # PersonBuilder + Person
├── idcard.go                 # 身份证生成 + 校验码
├── name.go                   # 姓名生成
├── address.go                # 地址生成
├── mobile.go                 # 手机号生成
├── bank.go                   # 银行卡生成 + LUHN
├── email.go                  # 邮箱生成
├── pinyin.go                 # 汉字转拼音
├── rand.go                   # 并发安全随机数
├── validate.go               # 校验函数
├── metadata/
│   ├── area_code.go          # 省市区 + 6位地区码映射
│   ├── last_name.go          # 姓氏 ~130 个
│   ├── first_name.go         # 名字 男/女各 5000-7000 个
│   ├── street_name.go        # 路名 500+ 个
│   ├── community_name.go     # 小区名 500+ 个
│   ├── common_chars.go       # 常用汉字 500 个
│   ├── mobile_prefix.go      # 手机号段
│   ├── bank_bin.go           # 银行卡 BIN
│   ├── email_suffix.go       # 邮箱后缀 ~15 个
│   ├── email_prefix.go       # 邮箱前缀 100-200 个
│   └── pinyin_map.go         # 汉字拼音映射
├── person_test.go
├── idcard_test.go
├── name_test.go
├── address_test.go
├── mobile_test.go
├── bank_test.go
├── email_test.go
└── README.md
```

## 核心 API 设计

### Gender 枚举

```go
type Gender int

const (
    GenderRandom Gender = iota  // 默认随机
    GenderMale                  // 男
    GenderFemale                // 女
)
```

### Person Builder

```go
// 创建 Person 构建器
person := chinaid.NewPerson()

// 链式配置
person.
    Province("北京").          // 指定省份
    Gender(chinaid.GenderMale). // 指定性别
    AgeRange(18, 35).          // 年龄范围
    Seed(12345).               // 可复现随机
    Build()                    // 生成单个
    // 或
    BuildN(100)                // 批量生成

// 获取生成的数据
person.IDNo()        // 身份证号
person.Name()        // 姓名
person.Gender()      // 性别
person.Birthday()    // 生日 time.Time
person.Age()         // 年龄 int
person.Province()    // 省份
person.City()        // 城市
person.Address()     // 完整地址
person.Mobile()      // 手机号
person.BankNo()      // 银行卡号
person.Email()       // 邮箱
```

### 单独生成器

```go
chinaid.NewIDCard().Province("广东").BirthYear(1990).Build()
chinaid.NewMobile().Build()
chinaid.NewBankNo().Build()
chinaid.NewEmail().Build()
```

## 数据结构

### Person 结构体

```go
type Person struct {
    idNo      string
    name      string
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
```

### PersonBuilder 构建器

```go
type PersonBuilder struct {
    rng       *rand.Rand  // 独立随机源
    seed      int64
    province  string
    gender    Gender
    minAge    int         // 默认 18
    maxAge    int         // 默认 60
}
```

### 一致性保证流程

```
Build() 执行顺序:
1. 确定省份 → 得到 province, city, areaCode
2. 确定性别 → 得到 gender
3. 确定生日 → 根据 ageRange 计算
4. 生成身份证 → 使用 areaCode + birthday + gender
5. 生成姓名 → 根据 gender 选择名字
6. 生成地址 → 使用 province + city + 词库
7. 生成手机 → 随机号段
8. 生成银行卡 → LUHN 算法
9. 生成邮箱 → 基于姓名拼音或词库
```

## 地区码映射

```go
type Province struct {
    Name   string   // "北京市"
    Short  string   // "北京"
    Cities []City
}

type City struct {
    Name      string   // "朝阳区"
    AreaCodes []string // ["110105"]
}

var Provinces = []Province{
    {
        Name:  "北京市",
        Short: "北京",
        Cities: []City{
            {Name: "东城区", AreaCodes: []string{"110101"}},
            {Name: "西城区", AreaCodes: []string{"110102"}},
            // ...
        },
    },
    // ...
}
```

数据来源: [modood/Administrative-divisions-of-China](https://github.com/modood/Administrative-divisions-of-China)

## 并发安全随机数

```go
type Rng struct {
    r *rand.Rand
}

func NewRng() *Rng {
    return &Rng{
        r: rand.New(rand.NewSource(time.Now().UnixNano())),
    }
}

func NewRngWithSeed(seed int64) *Rng {
    return &Rng{
        r: rand.New(rand.NewSource(seed)),
    }
}

func (rng *Rng) Intn(n int) int
func (rng *Rng) Int63n(n int64) int64
func (rng *Rng) Choice(slice []string) string
func (rng *Rng) ChoiceN(slice []string, n int) []string
func (rng *Rng) Shuffle(slice []string)
```

## 姓名生成

```go
var MaleFirstNames = []string{
    "伟", "强", "磊", "军", "勇", "杰", "涛", "明",
    // ... 约 5000-7000 个
}

var FemaleFirstNames = []string{
    "芳", "娜", "敏", "静", "丽", "艳", "霞",
    // ... 约 5000-7000 个
}

func (b *PersonBuilder) generateName() string {
    lastName := b.rng.Choice(metadata.LastNames)

    var firstName string
    switch b.gender {
    case GenderMale:
        firstName = b.rng.Choice(metadata.MaleFirstNames)
    case GenderFemale:
        firstName = b.rng.Choice(metadata.FemaleFirstNames)
    default:
        if b.rng.Intn(2) == 0 {
            b.gender = GenderMale
            firstName = b.rng.Choice(metadata.MaleFirstNames)
        } else {
            b.gender = GenderFemale
            firstName = b.rng.Choice(metadata.FemaleFirstNames)
        }
    }
    return lastName + firstName
}
```

## 地址生成

```go
// 格式：{省}{市}{区}{街道}{门牌号}号{小区}{楼栋}单元{房号}室

var StreetNames = []string{
    "东大街", "西大街", "幸福路", "和平路", "建设路",
    // ... 共 500+ 条
}

var CommunityNames = []string{
    "阳光花园", "世纪花园", "幸福小区", "书香苑",
    // ... 共 500+ 条
}

func (b *PersonBuilder) generateAddress() string {
    street := b.rng.Choice(metadata.StreetNames)
    community := b.rng.Choice(metadata.CommunityNames)

    houseNum := b.rng.Intn(200) + 1
    building := b.rng.Intn(20) + 1
    unit := b.rng.Intn(8) + 1
    room := (b.rng.Intn(30)+1)*100 + b.rng.Intn(4) + 1

    return fmt.Sprintf("%s%s%s%d号%s%d单元%d室",
        b.province, b.city, street, houseNum, community, unit, room)
}
```

## 邮箱生成

```go
var EmailSuffixes = []string{
    "qq.com", "163.com", "126.com", "gmail.com", "outlook.com",
    // ...
}

var EmailPrefixes = []string{
    "test", "admin", "user", "hello", "happy", "lucky",
    // ... 共 100-200 个
}

func (b *PersonBuilder) generateEmail() string {
    var prefix string

    if b.rng.Intn(2) == 0 {
        // 50%: 姓名拼音（全部或部分）
        prefix = b.generatePinyinPrefix()
    } else {
        // 50%: 常用前缀词库
        prefix = b.rng.Choice(metadata.EmailPrefixes)
    }

    num := b.rng.Intn(10000)
    suffix := b.rng.Choice(metadata.EmailSuffixes)

    return fmt.Sprintf("%s%d@%s", prefix, num, suffix)
}

func (b *PersonBuilder) generatePinyinPrefix() string {
    switch b.rng.Intn(4) {
    case 0: // 全名拼音
        return pinyin.Convert(b.name)
    case 1: // 仅姓
        return pinyin.Convert(b.lastName)
    case 2: // 仅名
        return pinyin.Convert(b.firstName)
    default: // 名的部分
        py := pinyin.Convert(b.firstName)
        if len(py) > 4 {
            return py[:b.rng.Intn(len(py)-2)+2]
        }
        return py
    }
}
```

## 测试设计

```go
func TestPersonConsistency(t *testing.T) {
    p := NewPerson().Province("北京").Build()
    assert.True(t, strings.HasPrefix(p.IDNo(), "11"))
    assert.True(t, strings.HasPrefix(p.Address(), "北京"))
}

func TestPersonSeed(t *testing.T) {
    p1 := NewPerson().Seed(12345).Build()
    p2 := NewPerson().Seed(12345).Build()
    assert.Equal(t, p1.IDNo(), p2.IDNo())
    assert.Equal(t, p1.Name(), p2.Name())
}

func TestPersonGender(t *testing.T) {
    p := NewPerson().Gender(GenderMale).Build()
    assert.Equal(t, 1, int(p.IDNo()[16]-'0')%2)
}

func TestIDNoChecksum(t *testing.T) {
    for i := 0; i < 1000; i++ {
        id := NewPerson().Build().IDNo()
        assert.True(t, ValidateIDNo(id))
    }
}

func TestBankNoLUHN(t *testing.T) {
    for i := 0; i < 1000; i++ {
        bankNo := NewPerson().Build().BankNo()
        assert.True(t, ValidateLUHN(bankNo))
    }
}

func TestBuildN(t *testing.T) {
    persons := NewPerson().Province("广东").BuildN(100)
    assert.Len(t, persons, 100)
}

func TestConcurrentSafety(t *testing.T) {
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            _ = NewPerson().Build()
        }()
    }
    wg.Wait()
}
```

## 改动总结

| 模块 | v1 | v2 |
|------|-----|-----|
| 架构 | 函数式 | Builder 链式调用 |
| 版本 | v1 | v2，不兼容 |
| 核心 | 独立函数 | Person 聚合生成器 |
| 一致性 | 无 | 身份证、地址、省份关联 |
| 并发 | 全局 rand | 独立 Rng 实例 |
| 姓名 | 98000 混合 | 10000-15000，按性别分类 |
| 地址 | Unicode 随机 | 词库 + 常用字 |
| 邮箱 | 随机字母 | 拼音/词库 + 常用后缀 |
| 时间 | 固定 2020 | 动态当前时间 |
| 测试 | 非空检查 | 有效性验证 |
