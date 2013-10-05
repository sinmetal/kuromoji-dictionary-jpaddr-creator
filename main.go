package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	rfile, err := os.OpenFile("KEN_ALL.CSV", os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wfile, err := os.OpenFile("userdict.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		rfile.Close()
		wfile.Close()
	}()

	r := csv.NewReader(rfile)
	w := csv.NewWriter(wfile)

	langMap := make(map[string]string)
	for {
		record, err := r.Read()
		if err != nil {
			break
		}
		prefName := removeUnnecessaryString(record[6])
		prefPhonetic := removeUnnecessaryString(record[3])
		langMap[prefName] = prefPhonetic

		cityName := removeUnnecessaryString(record[7])
		cityPhonetic := removeUnnecessaryString(record[4])
		langMap[cityName] = cityPhonetic

		if record[8] != "以下に掲載がない場合" {
			chouikiNames := replace(removeUnnecessaryString(record[8]))
			//chouikiName := removeUnnecessaryString(record[8])
			chouikiPhonetic := removeUnnecessaryString(record[5])
			for i := 0; i < len(chouikiNames); i++ {
				langMap[chouikiNames[i]] = chouikiPhonetic
			}
			//langMap[chouikiName] = chouikiPhonetic
		}
	}

	for key, value := range langMap {
		w.Write([]string{key, key, value, "カスタム地域"})
	}
	w.Flush()
}

func removeUnnecessaryString(value string) string {
	endChars := [...]string{"(", "（", "、"}
	endIndex := -1
	for i := 0; i < len(endChars); i++ {
		index := strings.Index(value, endChars[i])
		if index != -1 && index > endIndex {
			fmt.Println(value)
			endIndex = index
		}
	}
	if endIndex > -1 {
		ret := value[:endIndex]
		fmt.Println(ret)
		return ret
	}
	return value
}

func replace(value string) [50]string {
	var ret [50]string
	ret[0] = value
	ret[1] = strings.Replace(value, "ケ", "ヶ", -1)
	ret[2] = strings.Replace(value, "ヶ", "ケ", -1)
	ret[3] = strings.Replace(value, "１", "一", -1)
	ret[4] = strings.Replace(value, "１", "1", -1)
	ret[5] = strings.Replace(value, "２", "二", -1)
	ret[6] = strings.Replace(value, "２", "2", -1)
	ret[7] = strings.Replace(value, "３", "三", -1)
	ret[8] = strings.Replace(value, "３", "3", -1)
	ret[9] = strings.Replace(value, "４", "四", -1)
	ret[10] = strings.Replace(value, "４", "4", -1)
	ret[11] = strings.Replace(value, "５", "五", -1)
	ret[12] = strings.Replace(value, "５", "5", -1)
	ret[13] = strings.Replace(value, "６", "六", -1)
	ret[14] = strings.Replace(value, "６", "6", -1)
	ret[15] = strings.Replace(value, "７", "七", -1)
	ret[16] = strings.Replace(value, "７", "7", -1)
	ret[17] = strings.Replace(value, "８", "八", -1)
	ret[18] = strings.Replace(value, "８", "8", -1)
	ret[19] = strings.Replace(value, "９", "九", -1)
	ret[20] = strings.Replace(value, "９", "9", -1)
	ret[21] = strings.Replace(value, "１０", "十", -1)
	ret[22] = strings.Replace(value, "１０", "10", -1)
	ret[23] = strings.Replace(value, "１１", "十一", -1)
	ret[24] = strings.Replace(value, "１１", "一一", -1)
	ret[25] = strings.Replace(value, "１１", "11", -1)
	ret[26] = strings.Replace(value, "１２", "十二", -1)
	ret[27] = strings.Replace(value, "１２", "一二", -1)
	ret[28] = strings.Replace(value, "１２", "12", -1)
	ret[29] = strings.Replace(value, "１３", "十三", -1)
	ret[30] = strings.Replace(value, "１３", "一三", -1)
	ret[31] = strings.Replace(value, "１３", "13", -1)
	ret[32] = strings.Replace(value, "１４", "十四", -1)
	ret[33] = strings.Replace(value, "１４", "一四", -1)
	ret[34] = strings.Replace(value, "１４", "14", -1)
	ret[35] = strings.Replace(value, "１５", "十五", -1)
	ret[36] = strings.Replace(value, "１５", "一五", -1)
	ret[37] = strings.Replace(value, "１５", "15", -1)
	ret[38] = strings.Replace(value, "１６", "十六", -1)
	ret[39] = strings.Replace(value, "１６", "一六", -1)
	ret[40] = strings.Replace(value, "１６", "16", -1)
	ret[41] = strings.Replace(value, "１７", "十七", -1)
	ret[42] = strings.Replace(value, "１７", "一七", -1)
	ret[43] = strings.Replace(value, "１７", "17", -1)
	ret[44] = strings.Replace(value, "１８", "十八", -1)
	ret[45] = strings.Replace(value, "１８", "一八", -1)
	ret[46] = strings.Replace(value, "１８", "18", -1)
	ret[47] = strings.Replace(value, "１９", "十九", -1)
	ret[48] = strings.Replace(value, "１９", "一九", -1)
	ret[49] = strings.Replace(value, "１９", "19", -1)
	if strings.Index(value, "１") > -1 {
		//fmt.Println(value)
	}
	return ret
}
