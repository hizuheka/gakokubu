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

	// 小学校
	var addRecords Records
	var preRecord Record
	for i, record := range syogakuRecords {
		fmt.Printf("i=%d\n", i)
		// 最初のレコードの場合
		if i == 0 {
			fmt.Println("最初のレコード")
			if re := record.region.StartRegion(); re != nil {
				addRecord := createDummyRecord(record.gakuKubun, record.jichiCode)
				addRecord.region = *re
				addRecords = append(addRecords, addRecord)
			}
		} else if i == len(syogakuRecords)-1 { // 最後のレコード
			fmt.Println("最後のレコード")
			if re := record.region.EndRegion(); re != nil {
				addRecord := createDummyRecord(record.gakuKubun, record.jichiCode)
				addRecord.region = *re
				addRecords = append(addRecords, addRecord)
			}
		} else { // 2行目以降のレコードの場合
			fmt.Println("2番目以降のレコード")
			if fmr := record.FindMissingRegions(preRecord); fmr != nil {
				for _, region := range fmr {
					addRecord := createDummyRecord(record.gakuKubun, region.Start.JichiCode)
					re := region
					addRecord.region = re
					addRecords = append(addRecords, addRecord)
				}
			}
		}

		preRecord = record
	}
	syogakuRecords = append(syogakuRecords, addRecords...)

	// ソート
	syogakuRecords.Sort()

	// fmt.Println(syogakuRecords)

	w := bufio.NewWriter(outFile)
	for _, r := range syogakuRecords {
		if _, err := w.WriteString(r.ToString() + "\n"); err != nil {
			fmt.Printf("ファイル出力エラー: %v\n", err)
			return
		}
	}
	w.Flush()

	// 中学校
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
