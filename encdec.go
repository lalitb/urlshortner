package main
import (
	"bytes"
	"strings"
	"math"
)

var encodeStr string = "0123456789abcdefghijklmnopqrstuvwxyz"

func convertToBase36(value int64) string {
	var buffer bytes.Buffer;
	for (value > 0 ) {
		var v = value % 36;
		buffer.WriteString(string(encodeStr[v]));
		value = value / 36;
	}
	return reverseStr(buffer.String());
}

func convertFromBase36(str string) int64 {
	var result int64 = 0;
	pow := 0;
	for i := len(str) -1; i >= 0; i-- {
		index := strings.Index(encodeStr, string(str[i]));
		result += int64(index) * int64(math.Pow(float64(36), float64( pow )));
		pow ++;
	}
	return result;
}

func reverseStr ( str string) string {
    runes := []rune(str);
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

