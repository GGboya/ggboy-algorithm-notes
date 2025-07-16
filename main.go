package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type SimplifiedProblem struct {
	Title string  `json:"title"`
	Score float64 `json:"score"`
	URL   string  `json:"url"`
}

func loadProblems(filename, source string) ([]SimplifiedProblem, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var problems []SimplifiedProblem
	if err := json.Unmarshal(data, &problems); err != nil {
		return nil, err
	}

	return problems, nil
}

func main() {
	// 1. 连接数据库
	dsn := "root:ggcode123@tcp(127.0.0.1:3306)/ggcode?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. 需要导入的文件和来源
	files := []struct {
		Filename string
		Source   string
	}{
		{"questions/leetcode_simplified.json", "leetcode"},
		{"questions/atcoder_simplified.json", "atcoder"},
		{"questions/nowcoder_simplified_scored.json", "nowcoder"},
		{"questions/codeforces_problems.json", "codeforces"},
	}

	// 3. 插入数据
	stmt, err := db.Prepare(`INSERT INTO contest_problems (title, score, difficulty, url, source) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, file := range files {
		problems, err := loadProblems(file.Filename, file.Source)
		if err != nil {
			fmt.Printf("读取 %s 失败: %v\n", file.Filename, err)
			continue
		}
		for _, p := range problems {
			_, err := stmt.Exec(p.Title, p.Score, p.URL, file.Source)
			if err != nil {
				fmt.Printf("插入题目失败: %s, err: %v\n", p.Title, err)
			}
		}
		fmt.Printf("导入 %s 完成，共 %d 道题\n", file.Filename, len(problems))
	}
	fmt.Println("所有数据导入完成！")
}
