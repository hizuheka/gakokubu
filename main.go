package main

import (
	"bufio"
	"fmt"
	"os"
)

// 2つの地域が連続しているかをチェックする関数
func CheckContinuity(r1, r2 Record) bool {
	// TODO
	return true
	// return AddressToString(r1.End) == AddressToString(r2.Start)
}

// 2つの地域の間の欠けている地域を見つける関数
func FindMissingRegions(r1, r2 Region) []Region {
	// TODO
	var missingRegions []Region
	// 欠けている地域のチェックをシミュレーション。この部分は正確な住所範囲ロジックに基づいて実装する必要があります。
	return missingRegions
}

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

	syogakuRecords := make([]Record, 0, 3000)
	chugakuRecords := make([]Record, 0, 3000)

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

	// if !isFirstLine {
	// if !CheckContinuity(previousRecord, currentRecord) {
	// missingRegions := FindMissingRegions(previousRecord, currentRecord)
	// for _, region := range missingRegions {
	// outFile.WriteString(fmt.Sprintf("%s %s\n", AddressToString(region.Start), AddressToString(region.End)))
	// }
	// }
	// }
	//
	// previousRecord = currentRecord
	// isFirstLine = false
}
