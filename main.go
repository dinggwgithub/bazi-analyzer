package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"bazi-analyzer/internal/engine"
)

func main() {
	fmt.Println()
	fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    八 字 分 析 系 统 v1.0                      ║")
	fmt.Println("║                    Bazi Analyzer System                        ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
	fmt.Println()
	
	eng := engine.NewEngine()
	
	if len(os.Args) > 1 {
		input := strings.Join(os.Args[1:], " ")
		analyzeAndPrint(eng, input)
		return
	}
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Println()
		fmt.Print("请输入八字（格式：壬戌 壬寅 庚午 丙戌），输入 q 退出：")
		
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入失败：", err)
			continue
		}
		
		input = strings.TrimSpace(input)
		
		if input == "q" || input == "quit" || input == "exit" {
			fmt.Println()
			fmt.Println("感谢使用，再见！")
			break
		}
		
		if input == "" {
			continue
		}
		
		analyzeAndPrint(eng, input)
	}
}

func analyzeAndPrint(eng *engine.Engine, input string) {
	result, err := eng.Analyze(input)
	if err != nil {
		fmt.Println()
		fmt.Printf("错误：%s\n", err)
		fmt.Println("请检查输入格式，正确格式如：壬戌 壬寅 庚午 丙戌")
		return
	}
	
	report := eng.Reporter.FormatReport(result.Report)
	fmt.Println(report)
}
