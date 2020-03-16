package service

import (
	"math"
	"strings"
)

var tenToAny = map[int64]string{
	0:  "D",
	1:  "j",
	2:  "p",
	3:  "S",
	4:  "w",
	5:  "0",
	6:  "V",
	7:  "i",
	8:  "t",
	9:  "E",
	10: "c",
	11: "Z",
	12: "J",
	13: "F",
	14: "Y",
	15: "k",
	16: "h",
	17: "5",
	18: "u",
	19: "W",
	20: "g",
	21: "f",
	22: "q",
	23: "R",
	24: "L",
	25: "y",
	26: "N",
	27: "3",
	28: "x",
	29: "l",
	30: "a",
	31: "P",
	32: "4",
	33: "m",
	34: "e",
	35: "o",
	36: "9",
	37: "n",
	38: "C",
	39: "M",
	40: "G",
	41: "H",
	42: "X",
	43: "7",
	44: "2",
	45: "z",
	46: "O",
	47: "T",
	48: "A",
	49: "b",
	50: "I",
	51: "1",
	52: "r",
	53: "6",
	54: "U",
	55: "s",
	56: "B",
	57: "8",
	58: "Q",
	59: "v",
	60: "K",
	61: "d",
}

// 10进制转任意进制
func DecimalToAny(num, n int64) string {
	if num == 0 {
		return "0"
	}
	newNumStr := ""
	var remainder int64
	var remainderString string
	for num != 0 {
		remainder = num % n
		if 76 > remainder && remainder > 9 {
			remainderString = tenToAny[remainder]
		} else {
			remainderString = tenToAny[remainder]
			// remainderString = strconv.FormatInt(remainder, 10)
		}
		newNumStr = remainderString + newNumStr
		num = num / n
	}
	return newNumStr
}

func findKey(in string) int64 {
	var result int64 = -1
	for k, v := range tenToAny {
		if in == v {
			result = k
		}
	}
	return result
}

// 任意进制转10进制
func AnyToDecimal(num string, n int64) int64 {
	var newNum float64
	newNum = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(findKey(value))
		if tmp != -1 {
			newNum = newNum + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int64(newNum)
}
