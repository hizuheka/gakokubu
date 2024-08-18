package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 住所を表す構造体
type Address struct {
	MachiCode string // 町コード
	BanCode   string // 番地コード
	EdaCode   string // 枝番コード
	KoedaCode string // 小枝番コード
	Eda3Code  string // 枝番３コード
}

// 住所範囲を表す構造体
type Region struct {
	Start Address // 開始住所
	End   Address // 終了住所
}

// レコードを表す構造体
type Record struct {
	region       Region // 住所範囲
	gakuKubun    string // 学校区分
	gakuCode     string // 学校コード
	sakujoFlag   string // 削除フラグ
	updateYMD    string // 更新日
	jichiCode    string // 自治体コード
	updateYMDHMS string // 更新日
}

// 1行の入力を Record 構造体に変換する関数
func createRecord(line string) (Record, error) {
	items := strings.Split(line, ",")
	if len(items) == 16 {
		return Record{}, fmt.Errorf("入力ファイルの形式が誤っています。想定している項目数は 20 です。(len(line)=%d", len(line))
	}

	// 住所範囲
	r := Region{
		Start: Address{MachiCode: items[0], BanCode: items[1], EdaCode: items[2], KoedaCode: items[3], Eda3Code: items[4]},
		End:   Address{MachiCode: items[5], BanCode: items[6], EdaCode: items[7], KoedaCode: items[8], Eda3Code: items[9]},
	}

	record := Record{
		region:       r,
		gakuKubun:    items[10],
		gakuCode:     items[11],
		sakujoFlag:   items[12],
		updateYMD:    items[13],
		jichiCode:    items[14],
		updateYMDHMS: items[15],
	}

	return record, nil
}

// Address 構造体を文字列に変換する関数
func AddressToString(a Address) string {
	return a.machiCode + a.banCode + a.edaCode + a.koedaCode
}

// 2つの地域が連続しているかをチェックする関数
func CheckContinuity(r1, r2 Region) bool {
	// TODO
	return AddressToString(r1.End) == AddressToString(r2.Start)
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
	var previousRecord Record
	var isFirstLine = true

	for scanner.Scan() {
		line := scanner.Text()
		currentRecord, err := createRecord(line)
		if err != nil {
			fmt.Printf("ERR: %w\n", err)
			return
		}

		if !isFirstLine {
			if !CheckContinuity(previousRecord, currentRecord) {
				missingRegions := FindMissingRegions(previousRecord, currentRecord)
				for _, region := range missingRegions {
					outFile.WriteString(fmt.Sprintf("%s %s\n", AddressToString(region.Start), AddressToString(region.End)))
				}
			}
		}

		previousRecord = currentRecord
		isFirstLine = false
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("ファイル読み取りエラー: %v\n", err)
	}
}
