package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var (
		s = flag.String("f", "sample1.env", "env file name")
	)

	flag.Parse()
	filename := *s

	// ファイルオープン
	fp, err := os.Open(filename)
	if err != nil {
		// エラー処理
		fmt.Println("env file not found")
	}
	defer fp.Close()

	m := map[string]string{}

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		line := scanner.Text()

		// 空白行だったらスキップ
		if line == "" {
			continue
		}

		// 先頭が#で始まっていたらスキップ
		if strings.HasPrefix(line, "#") {
			continue
		}

		slice := strings.Split(line, "=")
		k := slice[0]
		v := slice[1]

		m[k] = v
	}

	count := 1
	for key, value := range m {
		fmt.Println("{")
		fmt.Printf("	\"Name\": \"%s\",\n", key)
		if numberCheck(value) {
			fmt.Printf("	\"Value\": %s\n", value)
		} else {
			fmt.Printf("	\"Value\": \"%s\"\n", value)
		}
		if count == len(m) {
			fmt.Println("}")
		} else {
			fmt.Println("},")
		}

		count++
	}

	if err = scanner.Err(); err != nil {
		// エラー処理
	}
}

func numberCheck(str string) bool {
	return regexp.MustCompile(`[0-9]$`).Match([]byte(str))
}

func keys(m map[string]string) []string {
	ks := []string{}
	for _, k := range m {
		ks = append(ks, k)
	}
	return ks
}

func values(m map[string]string) []string {
	vs := []string{}
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}
