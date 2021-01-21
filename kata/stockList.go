package kata

import (
	"fmt"
	"github.com/alex-dev-master/codewars/utils"
	"strconv"
	"strings"
)

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
	var counter = 1
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
