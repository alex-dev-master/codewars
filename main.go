package main

import (
	"fmt"
	"strconv"
	"strings",
)

func main() {
	//countSheep
	//print(countSheep(1))

	//StockList
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}
	print(StockList(b,c))
}


func StockList(listArt []string, listCat []string) string {
	for _, value := range listArt {
		valueArr := strings.Split(value, "")
		fmt.Printf("%v\n", valueArr)
		//checkArr, indx := InArray(valueArr[0], listCat)
		fmt.Printf("%t , %d \n", checkArr, indx)
	}

	return ""
}

func countSheep(num int) string {
	// Your code here!
	if num == 0 {
		return ""
	}
	var str strings.Builder
	for i := 0; i < num; i++ {
		str.WriteString(strconv.Itoa(i+1))
		str.WriteString(" sheep...")
	}
	return str.String()
}

func FindUniq(arr []float32) float32 {
	for i := 0; i < len(arr); i++ {
		current := arr[i]
		if (i+1) == len(arr) {
			return current
		}
		next := arr[i+1]
		if current != next {
			if (i+2) >= len(arr) {
				return next
			}
			if current != arr[i+2]{
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
