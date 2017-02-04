package main
import (
	"fmt"
	"bytes"
	"strings"
	"math"
)

var encodeStr string = "abcdefghijklmnopqrstuvwxyz0123456789";

func main_test() {
	fmt.Println("Hello, playground");
	fmt.Println(get_shortUrl(2234));
	fmt.Println(convertFromBase36(get_shortUrl(2234)));
}

func get_shortUrl( value int64) string {
	return convertToBase36(value);
}

func get_longUrl ( str string) int64 {
	return convertFromBase36(str);
}

func convertToBase36(value int64) string {
	var buffer bytes.Buffer;
	for (value > 0 ) {
		var v = value % 36;
		fmt.Println(v);
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
		fmt.Println(index);
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

