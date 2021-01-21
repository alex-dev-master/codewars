package main

import (
	"fmt"
	"github.com/alex-dev-master/codewars/utils"
	"strconv"
	"strings"
)

func main() {
	//countSheep
	//print(countSheep(1))

	//StockList
	var b = []string{"eBBAR 150", "eCDXE 515", "eBKWR 250", "eBTSQ 890", "eDRTY 600"}
	var c = []string{}
	print(StockList(b, c))
}

type Categories struct {
	name  string
	value int
}

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	var arrStrRes []Categories
	for _, value := range listCat {
		arrStrRes = append(arrStrRes, Categories{value, 0})
	}

	for _, value := range listArt {
		valueArr := strings.Split(value, "")
		checkArr, _ := utils.InArray(valueArr[0], listCat)
		if checkArr {
			valueNumArr := strings.Split(value, " ")
			valueNum, _ := strconv.Atoi(valueNumArr[1])
			for i := range arrStrRes {
				if arrStrRes[i].name == valueArr[0] {
					arrStrRes[i].value += valueNum
				}
			}
		}
	}

	var strRes strings.Builder
	var counter int = 1
	for _, item := range arrStrRes {
		if len(arrStrRes) == counter {
			strRes.WriteString(fmt.Sprintf("(%s : %d)", item.name, item.value))
			break
		}
		strRes.WriteString(fmt.Sprintf("(%s : %d) - ", item.name, item.value))
		counter++
	}
	return strRes.String()
}

func countSheep(num int) string {
	// Your code here!
	if num == 0 {
		return ""
	}
	var str strings.Builder
	for i := 0; i < num; i++ {
		str.WriteString(strconv.Itoa(i + 1))
		str.WriteString(" sheep...")
	}
	return str.String()
}

func FindUniq(arr []float32) float32 {
	for i := 0; i < len(arr); i++ {
		current := arr[i]
		if (i + 1) == len(arr) {
			return current
		}
		next := arr[i+1]
		if current != next {
			if (i + 2) >= len(arr) {
				return next
			}
			if current != arr[i+2] {
				return current
			} else {
				return next
			}
		}
	}
	return 0
}

func Accum(s string) string {
	// your code
	arrStr := strings.Split(s, "")
	var arrStrRes []string
	for indx, value := range arrStr {
		if indx == 0 {
			arrStrRes = append(arrStrRes, strings.ToUpper(value))
			continue
		}
		var str strings.Builder
		str.WriteString(strings.ToUpper(value))
		for i := 0; i < indx; i++ {
			str.WriteString(strings.ToLower(value))
		}
		arrStrRes = append(arrStrRes, str.String())
	}
	return strings.Join(arrStrRes, "-")
}
