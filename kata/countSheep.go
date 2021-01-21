package kata

import (
	"strconv"
	"strings"
)

func CountSheep(num int) string {
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
