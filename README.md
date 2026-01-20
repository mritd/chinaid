# chinaid v2

中国大陆身份证等信息生成库，用于测试内部系统校验有效性。

## 注意

**本项目仅用于测试目的，比如开发人员测试自己的关键校验规则是否正确；对于使用本项目产生的任何后果，使用者应当自行承担法律风险与责任，一切后果与本项目无关。**

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
fmt.Println(person.Birthday()) // 生日
fmt.Println(person.Province()) // 省份
fmt.Println(person.City())    // 城市
fmt.Println(person.Address()) // 完整地址
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
- **性别分类**: 名字按性别分类，更真实

## API

### PersonBuilder

| 方法 | 说明 |
|------|------|
| `NewPerson()` | 创建构建器 |
| `Province(string)` | 设置省份（如 "北京"、"广东"） |
| `Gender(Gender)` | 设置性别（GenderMale / GenderFemale） |
| `AgeRange(min, max)` | 设置年龄范围，默认 18-60 |
| `Seed(int64)` | 设置随机种子 |
| `Build()` | 生成单个 Person |
| `BuildN(n)` | 批量生成 n 个 Person |

### Person

| 方法 | 返回类型 | 说明 |
|------|---------|------|
| `IDNo()` | string | 18位身份证号 |
| `Name()` | string | 姓名 |
| `LastName()` | string | 姓 |
| `FirstName()` | string | 名 |
| `Gender()` | Gender | 性别 |
| `Birthday()` | time.Time | 生日 |
| `Age()` | int | 年龄 |
| `Province()` | string | 省份 |
| `City()` | string | 城市 |
| `Address()` | string | 完整地址 |
| `Mobile()` | string | 11位手机号 |
| `BankNo()` | string | 银行卡号 |
| `Email()` | string | 邮箱 |

### 验证函数

| 函数 | 说明 |
|------|------|
| `ValidateIDNo(string)` | 验证身份证号校验码 |
| `ValidateLUHN(string)` | 验证银行卡 LUHN 校验 |

## 生成数据说明

- **姓名**: 使用常用姓氏 + 按性别分类的名字，约 10000+ 个名字
- **身份证号**: 采用标准身份证规则生成，校验码有效
- **手机号**: 常用运营商号段 + 随机数字
- **银行卡号**: 正确的银行卡 BIN + LUHN 算法校验
- **邮箱**: 姓名拼音或常用前缀 + 常用邮箱后缀
- **地址**: 真实省市区数据 + 路名/小区词库

## License

MIT
