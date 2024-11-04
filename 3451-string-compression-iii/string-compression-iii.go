import (
    "strings"
    "strconv"
)

func compressedString(word string) string {
    if word == "" {
        return ""
    }
    var curChar rune = []rune(word)[0]
    var curFreq int = 1
    var result strings.Builder
    for index, char := range(word) {
        if index == 0 {
            continue
        } else {
            if(char == curChar){
                if curFreq < 9{
                    curFreq++
                }  else {
                    result.WriteString(strconv.Itoa(curFreq))
                    result.WriteString(string(curChar))
                    curFreq = 1
                }
            } else {
                result.WriteString(strconv.Itoa(curFreq))
                result.WriteString(string(curChar))
                curFreq = 1
                curChar = char
            }
        }
    }
    result.WriteString(strconv.Itoa(curFreq))
    result.WriteString(string(curChar))
    return result.String()
}