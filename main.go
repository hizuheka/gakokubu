package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("使用方法: gakokubu.exe <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("ファイルを開く際のエラー: %v\n", err)
		return
	}
	defer inFile.Close()

	outputFile := os.Args[2]
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("出力ファイルの作成エラー: %v\n", err)
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	// var previousRecord Record
	// var isFirstLine = true

	syogakuRecords := createRecords()
	chugakuRecords := createRecords()

	for scanner.Scan() {
		line := scanner.Text()
		record, err := createRecord(line)
		if err != nil {
			fmt.Printf("レコード生成エラー: %v\n", err)
			return
		}

		// 小学校と中学校のスライスに振り分ける
		if record.IsShogaku() {
			syogakuRecords = append(syogakuRecords, record)
		} else {
			chugakuRecords = append(chugakuRecords, record)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("ファイル読み取りエラー: %v\n", err)
	}

	// 小学校のダミーレコードを補完
	syogakuRecords.FillDummyRecords()

	// 中学校のダミーレコードを補完
	chugakuRecords.FillDummyRecords()

	// 小学校と中学校を結合してソート
	allRecords := append(syogakuRecords, chugakuRecords...)
	allRecords.Sort()

	w := bufio.NewWriter(outFile)
	for _, r := range allRecords {
		if _, err := w.WriteString(r.ToString() + "\r\n"); err != nil {
			fmt.Printf("ファイル出力エラー: %v\n", err)
			return
		}
	}
	w.Flush()
}
