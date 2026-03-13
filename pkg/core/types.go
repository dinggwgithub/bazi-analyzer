package core

import (
	"fmt"
	"strings"
)

// 基础枚举类型

// TianGan 天干枚举: 甲=0, 乙=1, 丙=2, 丁=3, 戊=4, 己=5, 庚=6, 辛=7, 壬=8, 癸=9
type TianGan int

const (
	Jia  TianGan = iota // 甲
	Yi                  // 乙
	Bing                // 丙
	Ding                // 丁
	Wu                  // 戊
	Ji                  // 己
	Geng                // 庚
	Xin                 // 辛
	Ren                 // 壬
	Gui                 // 癸
)

// DiZhi 地支枚举: 子=0, 丑=1, 寅=2, 卯=3, 辰=4, 巳=5, 午=6, 未=7, 申=8, 酉=9, 戌=10, 亥=11
type DiZhi int

const (
	Zi     DiZhi = iota // 子
	Chou                // 丑
	YinZhi              // 寅
	Mao                 // 卯
	Chen                // 辰
	Si                  // 巳
	WuZhi               // 午
	Wei                 // 未
	Shen                // 申
	You                 // 酉
	Xu                  // 戌
	Hai                 // 亥
)

// WuXing 五行枚举: 木=0, 火=1, 土=2, 金=3, 水=4
type WuXing int

const (
	Wood  WuXing = iota // 木
	Fire                // 火
	Earth               // 土
	Metal               // 金
	Water               // 水
)

// YinYang 阴阳枚举: 阴=0, 阳=1
type YinYang int

const (
	Yin  YinYang = iota // 阴
	Yang                // 阳
)

// ZhuWei 柱位枚举: 年=0, 月=1, 日=2, 时=3
type ZhuWei int

const (
	Year  ZhuWei = iota // 年柱
	Month               // 月柱
	Day                 // 日柱
	Hour                // 时柱
)

// String 方法实现

func (tg TianGan) String() string {
	names := []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	if tg < 0 || tg > 9 {
		return "未知"
	}
	return names[tg]
}

func (dz DiZhi) String() string {
	names := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	if dz < 0 || dz > 11 {
		return "未知"
	}
	return names[dz]
}

func (wx WuXing) String() string {
	names := []string{"木", "火", "土", "金", "水"}
	if wx < 0 || wx > 4 {
		return "未知"
	}
	return names[wx]
}

func (yy YinYang) String() string {
	if yy == Yin {
		return "阴"
	}
	return "阳"
}

func (zw ZhuWei) String() string {
	names := []string{"年柱", "月柱", "日柱", "时柱"}
	if zw < 0 || zw > 3 {
		return "未知"
	}
	return names[zw]
}

// GanZhi 干支组合
type GanZhi struct {
	Gan TianGan
	Zhi DiZhi
}

func (gz GanZhi) String() string {
	return gz.Gan.String() + gz.Zhi.String()
}

// Bazi 八字四柱
type Bazi struct {
	Year  GanZhi // 年柱
	Month GanZhi // 月柱
	Day   GanZhi // 日柱
	Hour  GanZhi // 时柱
}

// GetDayGan 获取日主（日柱天干）
func (b *Bazi) GetDayGan() TianGan {
	return b.Day.Gan
}

// GetDayZhi 获取日支（日柱地支）
func (b *Bazi) GetDayZhi() DiZhi {
	return b.Day.Zhi
}

// GetMonthZhi 获取月令（月柱地支）
func (b *Bazi) GetMonthZhi() DiZhi {
	return b.Month.Zhi
}

// GetGanZhiByZhuWei 根据柱位获取干支
func (b *Bazi) GetGanZhiByZhuWei(zhuWei ZhuWei) GanZhi {
	switch zhuWei {
	case Year:
		return b.Year
	case Month:
		return b.Month
	case Day:
		return b.Day
	case Hour:
		return b.Hour
	default:
		return GanZhi{}
	}
}

// String 返回八字的字符串表示
func (b *Bazi) String() string {
	return fmt.Sprintf("%s %s %s %s", b.Year.String(), b.Month.String(), b.Day.String(), b.Hour.String())
}

// ParseBazi 从字符串解析八字
// 输入格式: "壬戌 壬寅 庚午 丙戌"
func ParseBazi(input string) (*Bazi, error) {
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)

	if len(parts) != 4 {
		return nil, fmt.Errorf("八字输入格式错误，应为4个柱，用空格分隔，如: 壬戌 壬寅 庚午 丙戌")
	}

	bazi := &Bazi{}

	// 解析年柱
	gz, err := ParseGanZhi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("年柱解析错误: %w", err)
	}
	bazi.Year = *gz

	// 解析月柱
	gz, err = ParseGanZhi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("月柱解析错误: %w", err)
	}
	bazi.Month = *gz

	// 解析日柱
	gz, err = ParseGanZhi(parts[2])
	if err != nil {
		return nil, fmt.Errorf("日柱解析错误: %w", err)
	}
	bazi.Day = *gz

	// 解析时柱
	gz, err = ParseGanZhi(parts[3])
	if err != nil {
		return nil, fmt.Errorf("时柱解析错误: %w", err)
	}
	bazi.Hour = *gz

	return bazi, nil
}

// ParseGanZhi 从字符串解析干支
// 输入格式: "壬戌"
func ParseGanZhi(input string) (*GanZhi, error) {
	input = strings.TrimSpace(input)
	if len([]rune(input)) != 2 {
		return nil, fmt.Errorf("干支格式错误，应为2个字符，如: 壬戌")
	}

	runes := []rune(input)
	ganStr := string(runes[0])
	zhiStr := string(runes[1])

	gan, err := ParseTianGan(ganStr)
	if err != nil {
		return nil, fmt.Errorf("天干解析错误: %w", err)
	}

	zhi, err := ParseDiZhi(zhiStr)
	if err != nil {
		return nil, fmt.Errorf("地支解析错误: %w", err)
	}

	return &GanZhi{Gan: gan, Zhi: zhi}, nil
}

// ParseTianGan 从字符串解析天干
func ParseTianGan(input string) (TianGan, error) {
	switch input {
	case "甲":
		return Jia, nil
	case "乙":
		return Yi, nil
	case "丙":
		return Bing, nil
	case "丁":
		return Ding, nil
	case "戊":
		return Wu, nil
	case "己":
		return Ji, nil
	case "庚":
		return Geng, nil
	case "辛":
		return Xin, nil
	case "壬":
		return Ren, nil
	case "癸":
		return Gui, nil
	default:
		return -1, fmt.Errorf("无效的天干: %s", input)
	}
}

// ParseDiZhi 从字符串解析地支
func ParseDiZhi(input string) (DiZhi, error) {
	switch input {
	case "子":
		return Zi, nil
	case "丑":
		return Chou, nil
	case "寅":
		return YinZhi, nil
	case "卯":
		return Mao, nil
	case "辰":
		return Chen, nil
	case "巳":
		return Si, nil
	case "午":
		return WuZhi, nil
	case "未":
		return Wei, nil
	case "申":
		return Shen, nil
	case "酉":
		return You, nil
	case "戌":
		return Xu, nil
	case "亥":
		return Hai, nil
	default:
		return -1, fmt.Errorf("无效的地支: %s", input)
	}
}

// ParseWuXing 从字符串解析五行
func ParseWuXing(input string) (WuXing, error) {
	switch input {
	case "木":
		return Wood, nil
	case "火":
		return Fire, nil
	case "土":
		return Earth, nil
	case "金":
		return Metal, nil
	case "水":
		return Water, nil
	default:
		return -1, fmt.Errorf("无效的五行: %s", input)
	}
}
