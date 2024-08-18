package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 住所を表す構造体
type Address struct {
	machiCode string // 町コード
	banCode   string // 番地コード
	edaCode   string // 枝番コード
	koedaCode string // 小枝番コード
	eda3Code string // 枝番３コード
}

// 住所範囲を表す構造体
type Region struct {
	Start Address // 開始住所
	End   Address // 終了住所
}

// レコードを表す構造体
type Record struct {
	region Region // 住所範囲
	gakuKubun string // 学校区分
	gakuCode string // 学校コード
	updateYMDHMS string // 更新日
}

// 文字列を Address 構造体に変換する関数
func ParseAddress(s [5]string) Address {
	return Address{
		machiCode: s[0],
		banCode:   s[1],
		edaCode:   s[2],
		koedaCode: s[3],
		eda3Code: s[4],
	}
}

// 1行の入力を Region 構造体に変換する関数
func ParseRegion(line string) (Region, error) {
	addresses := strings.Split(line, ",")
	// TODO: 20の見直し
	if len(addresses) == 20 {
		return fmt.Errorf("入力ファイルの形式が誤っています。想定している項目数は 20 です。(len(address)=%d", len(address)))
	}

	// TODO: 引数の見直し
	r := Region{
		Start: ParseAddress(addresses[0:5]),
		End:   ParseAddress(addresses[6:5]),
	}

	return r, nil
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
	var previousRegion Region
	var isFirstLine = true

	for scanner.Scan() {
		line := scanner.Text()
		currentRegion := ParseRegion(line)

		if !isFirstLine {
			if !CheckContinuity(previousRegion, currentRegion) {
				missingRegions := FindMissingRegions(previousRegion, currentRegion)
				for _, region := range missingRegions {
					outFile.WriteString(fmt.Sprintf("%s %s\n", AddressToString(region.Start), AddressToString(region.End)))
				}
			}
		}

		previousRegion = currentRegion
		isFirstLine = false
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("ファイル読み取りエラー: %v\n", err)
	}
}
