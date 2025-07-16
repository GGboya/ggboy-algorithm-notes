package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

// Problem 结构体用于解析原始JSON数据
type Problem struct {
	Rating    float64 `json:"Rating"`
	ID        int     `json:"ID"`
	Title     string  `json:"Title"`
	TitleZH   string  `json:"TitleZH"`
	TitleSlug string  `json:"TitleSlug"`
}

// SimplifiedProblem 结构体用于输出简化JSON
type SimplifiedProblem struct {
	Title       string  `json:"title"`
	Score       float64 `json:"score"`
	Difficulty  string  `json:"difficulty"`
	LeetcodeURL string  `json:"leetcode_url"`
}

// 根据分数确定难度等级
func getDifficulty(rating float64) string {
	switch {
	case rating >= 3000:
		return "Hard"
	case rating >= 2000:
		return "Medium"
	case rating >= 1500:
		return "Easy"
	default:
		return "Easy"
	}
}

func main() {
	// 读取JSON文件
	file, err := os.Open("leetcode_data.json")
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	// 读取文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		return
	}

	// 解析JSON数据
	var problems []Problem
	err = json.Unmarshal(content, &problems)
	if err != nil {
		fmt.Printf("解析JSON失败: %v\n", err)
		return
	}

	// 转换为简化格式
	var simplifiedProblems []SimplifiedProblem
	for _, p := range problems {
		// 生成LeetCode链接
		leetcodeURL := fmt.Sprintf("https://leetcode.cn/problems/%s/description/", p.TitleSlug)

		simplifiedProblem := SimplifiedProblem{
			Title:       p.Title,
			Score:       p.Rating,
			Difficulty:  getDifficulty(p.Rating),
			LeetcodeURL: leetcodeURL,
		}
		simplifiedProblems = append(simplifiedProblems, simplifiedProblem)
	}

	// 按分数排序（从高到低）
	sort.Slice(simplifiedProblems, func(i, j int) bool {
		return simplifiedProblems[i].Score > simplifiedProblems[j].Score
	})

	// 保存为JSON文件
	outputData, err := json.MarshalIndent(simplifiedProblems, "", "  ")
	if err != nil {
		fmt.Printf("生成JSON失败: %v\n", err)
		return
	}

	err = os.WriteFile("leetcode_simplified.json", outputData, 0644)
	if err != nil {
		fmt.Printf("保存文件失败: %v\n", err)
		return
	}

	fmt.Printf("成功生成简化JSON文件，包含 %d 道题目\n", len(simplifiedProblems))

	// 显示前10道题目的示例
	fmt.Println("\n前10道题目的示例:")
	fmt.Println("排名 | 分数 | 难度 | 题目 | 链接")
	fmt.Println("-----|------|------|------|------")
	for i := 0; i < 10 && i < len(simplifiedProblems); i++ {
		p := simplifiedProblems[i]
		// 截断过长的标题
		title := p.Title
		if len(title) > 30 {
			title = title[:27] + "..."
		}
		fmt.Printf("%4d | %6.1f | %6s | %s | %s\n",
			i+1, p.Score, p.Difficulty, title, p.LeetcodeURL)
	}

	// 统计难度分布
	difficultyCount := make(map[string]int)
	for _, p := range simplifiedProblems {
		difficultyCount[p.Difficulty]++
	}

	fmt.Println("\n难度分布:")
	for difficulty, count := range difficultyCount {
		percentage := float64(count) / float64(len(simplifiedProblems)) * 100
		fmt.Printf("%s: %d 题 (%.1f%%)\n", difficulty, count, percentage)
	}
}
