package properties

import "strings"

func ConvertToRoman(number int) string {
	roman := strings.Builder{}

	for i := number; i > 0; i-- {
		if i == 4 {
			roman.WriteString("IV")
			break
		}
		roman.WriteString("I")
	}

	return roman.String()
}
