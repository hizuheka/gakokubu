package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var version string

func main() {
	var (
		i, o, d string
		v       bool
	)
	flag.StringVar(&i, "i", "", "変換元ファイル")
	flag.StringVar(&o, "o", "", "変換先ファイル")
	flag.StringVar(&d, "d", "99999999", "ダミーレコードの更新年月日")
	flag.BoolVar(&v, "v", false, "version")
	flag.Parse()

	if v {
		fmt.Printf("gakokubu.exe version %s\r\n", version)
		return
	}

	if i == "" {
		fmt.Printf("-i が指定されていません。変換元ファイルのパスを指定してください。(-i=%s)", i)
		return
	}
	if o == "" {
		fmt.Printf("-o が指定されていません。変換先ファイルのパスを指定してください。(-o=%s)", o)
		return
	}

	inFile, err := os.Open(i)
	if err != nil {
		fmt.Printf("ファイルを開く際のエラー: %v\n", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(o)
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
	syogakuRecords.FillDummyRecords(d)

	// 中学校のダミーレコードを補完
	chugakuRecords.FillDummyRecords(d)

	// 小学校と中学校を結合してソート
	allRecords := append(syogakuRecords, chugakuRecords...)
	allRecords.Sort()

	w := bufio.NewWriter(outFile)
	// 最初にBOMを出力する
	w.Write([]byte{0xEF, 0xBB, 0xBF})
	// レコード出力
	for _, r := range allRecords {
		if _, err := w.WriteString(r.ToString() + "\r\n"); err != nil {
			fmt.Printf("ファイル出力エラー: %v\n", err)
			return
		}
	}
	w.Flush()
}
