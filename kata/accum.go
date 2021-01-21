package kata

import "strings"

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
