package service

import (
	"bazi-analyzer/internal/engine/step1_scan"
	"bazi-analyzer/internal/engine/step2_rebuild"
	"bazi-analyzer/internal/engine/step3_wangshuai"
	"bazi-analyzer/internal/engine/step4_xiji"
	"bazi-analyzer/internal/engine/step5_analyze"
	"bazi-analyzer/internal/engine/step6_conclusion"
	"bazi-analyzer/internal/model/base"
	"bazi-analyzer/pkg/utils"
	"fmt"
	"strings"
)

type BaZiFullAnalysis struct {
	Input       string                                  `json:"input"`
	Chart       *base.BaZiChart                         `json:"chart"`
	Step1Result *step1_scan.Step1ScanResult             `json:"step1_result"`
	Step2Result *step2_rebuild.Step2RebuildResult       `json:"step2_result"`
	Step3Result *step3_wangshuai.Step3WangShuaiResult   `json:"step3_result"`
	Step4Result *step4_xiji.Step4XiJiResult             `json:"step4_result"`
	Step5Result *step5_analyze.Step5AnalyzeResult       `json:"step5_result"`
	Step6Result *step6_conclusion.Step6ConclusionResult `json:"step6_result"`
}

func AnalyzeBaZi(input string) (*BaZiFullAnalysis, error) {
	chart, err := utils.ParseBaZi(input)
	if err != nil {
		return nil, err
	}

	step1Result := step1_scan.ScanBaZi(&chart)
	step2Result := step2_rebuild.RebuildBaZi(&chart, step1Result)
	step3Result := step3_wangshuai.AnalyzeWangShuai(&chart, step2Result)
	step4Result := step4_xiji.AnalyzeXiJi(&chart, step3Result)
	step5Result := step5_analyze.AnalyzeInteractions(&chart, step1Result, step2Result, step4Result)
	step6Result := step6_conclusion.AnalyzeConclusion(&chart, step1Result, step2Result, step3Result, step4Result, step5Result)

	return &BaZiFullAnalysis{
		Input:       input,
		Chart:       &chart,
		Step1Result: step1Result,
		Step2Result: step2Result,
		Step3Result: step3Result,
		Step4Result: step4Result,
		Step5Result: step5Result,
		Step6Result: step6Result,
	}, nil
}

func joinWuXing(wx []base.WuXing) string {
	strs := make([]string, len(wx))
	for i, w := range wx {
		strs[i] = string(w)
	}
	return strings.Join(strs, "、")
}

func (r *BaZiFullAnalysis) ToMarkdown() string {
	var sb strings.Builder

	sb.WriteString("# 八字分析报告\n\n")
	sb.WriteString(fmt.Sprintf("## 输入八字：%s\n\n", r.Input))

	if r.Chart != nil {
		sb.WriteString("### 基本信息\n\n")
		dm := r.Chart.GetDayMaster()
		sb.WriteString(fmt.Sprintf("- 日主：**%s(%s)**\n", dm.Name, dm.Element))
		sb.WriteString("\n")
	}

	if r.Step1Result != nil {
		sb.WriteString("## 第一步：扫描全局特殊关系\n\n")
		if len(r.Step1Result.TianGanWuHe) > 0 {
			sb.WriteString("### 天干五合\n")
			for _, wh := range r.Step1Result.TianGanWuHe {
				sb.WriteString(fmt.Sprintf("- %s-%s\n", wh.Elements[0].Name, wh.Elements[1].Name))
			}
			sb.WriteString("\n")
		}
		if len(r.Step1Result.TianGanSiChong) > 0 {
			sb.WriteString("### 天干四冲\n")
			for _, sc := range r.Step1Result.TianGanSiChong {
				sb.WriteString(fmt.Sprintf("- %s-%s\n", sc.Elements[0].Name, sc.Elements[1].Name))
			}
			sb.WriteString("\n")
		}
		if len(r.Step1Result.DiZhiSanHui) > 0 {
			sb.WriteString("### 地支三会\n")
			for _, sh := range r.Step1Result.DiZhiSanHui {
				sb.WriteString(fmt.Sprintf("- %s-%s-%s\n", sh.Elements[0].Name, sh.Elements[1].Name, sh.Elements[2].Name))
			}
			sb.WriteString("\n")
		}
		if len(r.Step1Result.DiZhiSanHe) > 0 {
			sb.WriteString("### 地支三合\n")
			for _, sh := range r.Step1Result.DiZhiSanHe {
				sb.WriteString(fmt.Sprintf("- %s-%s-%s\n", sh.Elements[0].Name, sh.Elements[1].Name, sh.Elements[2].Name))
			}
			sb.WriteString("\n")
		}
		if len(r.Step1Result.DiZhiLiuHe) > 0 {
			sb.WriteString("### 地支六合\n")
			for _, lh := range r.Step1Result.DiZhiLiuHe {
				sb.WriteString(fmt.Sprintf("- %s-%s\n", lh.Elements[0].Name, lh.Elements[1].Name))
			}
			sb.WriteString("\n")
		}

	}

	if r.Step2Result != nil && len(r.Step2Result.RebuiltDiZhi) > 0 {
		sb.WriteString("## 第二步：格局重构\n\n")
		sb.WriteString("| 原地支 | 原五行 | 重定五行 | 原因 |\n")
		sb.WriteString("|--------|--------|----------|------|\n")
		for _, rd := range r.Step2Result.RebuiltDiZhi {
			sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", rd.OriginalName, rd.OriginalWuXing, rd.RebuiltWuXing, rd.Reason))
		}
		sb.WriteString("\n")
	}

	if r.Step3Result != nil {
		sb.WriteString("## 第三步：日主旺衰分析\n\n")
		sb.WriteString(fmt.Sprintf("- **日主**：%s(%s)\n", r.Step3Result.DayMaster.Name, r.Step3Result.DayMasterWuXing))
		sb.WriteString(fmt.Sprintf("- **得令**：%t\n", r.Step3Result.IsDeLing))
		sb.WriteString(fmt.Sprintf("- **得地**：%t\n", r.Step3Result.IsDeDi))
		sb.WriteString(fmt.Sprintf("- **得势**：%t\n", r.Step3Result.IsDeShi))
		sb.WriteString(fmt.Sprintf("- **生扶力量**：%d\n", r.Step3Result.ShengFuCount))
		sb.WriteString(fmt.Sprintf("- **克泄耗力量**：%d\n", r.Step3Result.KeXieHaoCount))
		sb.WriteString(fmt.Sprintf("- **旺衰结论**：**%s**\n", r.Step3Result.WangShuaiType))
		sb.WriteString(fmt.Sprintf("- **判断依据**：%s\n\n", r.Step3Result.Reason))
	}

	if r.Step4Result != nil {
		sb.WriteString("## 第四步：喜忌用神分析\n\n")
		sb.WriteString(fmt.Sprintf("- **喜用五行**：%s\n", joinWuXing(r.Step4Result.FavorableWuXing)))
		sb.WriteString(fmt.Sprintf("- **喜用十神**：%s\n", strings.Join(r.Step4Result.FavorableShiShen, "、")))
		sb.WriteString(fmt.Sprintf("- **忌神五行**：%s\n", joinWuXing(r.Step4Result.UnfavorableWuXing)))
		sb.WriteString(fmt.Sprintf("- **忌神十神**：%s\n", strings.Join(r.Step4Result.UnfavorableShiShen, "、")))
		sb.WriteString(fmt.Sprintf("- **判断依据**：%s\n\n", r.Step4Result.Reason))
	}

	if r.Step5Result != nil {
		sb.WriteString("## 第五步：细析作用关系\n\n")
		if len(r.Step5Result.TianGanInteractions) > 0 {
			sb.WriteString("### 天干作用\n")
			for _, inter := range r.Step5Result.TianGanInteractions {
				sb.WriteString(fmt.Sprintf("- %s%s %s %s%s [%s]\n", inter.From.Name, inter.From.Element, inter.Type, inter.To.Name, inter.To.Element, inter.Strength))
			}
			sb.WriteString("\n")
		}
		if len(r.Step5Result.DiZhiInteractions) > 0 {
			sb.WriteString("### 地支作用\n")
			for _, inter := range r.Step5Result.DiZhiInteractions {
				sb.WriteString(fmt.Sprintf("- %s: %s\n", inter.Type, inter.Description))
			}
			sb.WriteString("\n")
		}
		if len(r.Step5Result.KeyPoints) > 0 {
			sb.WriteString("### 关键点\n")
			for _, kp := range r.Step5Result.KeyPoints {
				sb.WriteString(fmt.Sprintf("- %s\n", kp))
			}
			sb.WriteString("\n")
		}
	}

	if r.Step6Result != nil {
		sb.WriteString("## 第六步：综合论断\n\n")
		sb.WriteString(fmt.Sprintf("- **格局类型**：**%s**\n", r.Step6Result.GeJuType))
		sb.WriteString(fmt.Sprintf("- **格局层次**：%s\n", r.Step6Result.GeJuLevel))
		sb.WriteString(fmt.Sprintf("- **格局描述**：%s\n\n", r.Step6Result.GeJuDescription))

		if len(r.Step6Result.SuiYunAnalysis) > 0 {
			sb.WriteString("### 岁运分析\n")
			for _, sy := range r.Step6Result.SuiYunAnalysis {
				sb.WriteString(fmt.Sprintf("- %s\n", sy))
			}
			sb.WriteString("\n")
		}

		sb.WriteString(fmt.Sprintf("### 总体总结\n%s\n\n", r.Step6Result.OverallSummary))

		if len(r.Step6Result.KeyAdvice) > 0 {
			sb.WriteString("### 关键建议\n")
			for _, ka := range r.Step6Result.KeyAdvice {
				sb.WriteString(fmt.Sprintf("- %s\n", ka))
			}
		}
	}

	sb.WriteString("\n---\n*八字分析系统 v1.0*\n")

	return sb.String()
}
