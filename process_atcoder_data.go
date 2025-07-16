package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Problem struct {
	ContestID int    `json:"contestId"`
	Index     string `json:"index"`
	Name      string `json:"name"`
	Rating    int    `json:"rating"`
}

type Result struct {
	Problems []Problem `json:"problems"`
}

type APIResponse struct {
	Status string `json:"status"`
	Result Result `json:"result"`
}

type SimplifiedProblem struct {
	Name   string `json:"title"`
	Rating int    `json:"score"`
	URL    string `json:"url"`
}

func main() {
	resp, err := http.Get("https://codeforces.com/api/problemset.problems")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		panic(err)
	}

	var problems []SimplifiedProblem
	for _, p := range apiResp.Result.Problems {
		// 只保留有分数的题目
		if p.Rating == 0 {
			continue
		}
		url := fmt.Sprintf("https://codeforces.com/problemset/problem/%d/%s", p.ContestID, p.Index)
		problems = append(problems, SimplifiedProblem{
			Name:   p.Name,
			Rating: p.Rating,
			URL:    url,
		})
	}

	// 保存为 JSON 文件
	file, err := os.Create("codeforces_problems.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(problems); err != nil {
		panic(err)
	}

	fmt.Printf("共保存 %d 道题目到 codeforces_problems.json\n", len(problems))
}
