func nextWord(i *int, sentence string) {
    n := len(sentence)
    for *i < n && sentence[*i] != ' ' {
        *i = *i + 1
    }
    *i = *i + 1
}

func isPrefix(i *int, sentence string, searchWord string) bool {
    n := len(sentence)
    m := len(searchWord)
    sI := 0
    for *i < n && sI < m && sentence[*i] == searchWord[sI] {
        *i = *i + 1
        sI = sI + 1
    } 
    if sI == m {
        return true
    }
    return false
}

func isPrefixOfWord(sentence string, searchWord string) int {
    i := 0
    wordIndex := 1
    for i < len(sentence) {
        if isPrefix(&i, sentence, searchWord) {
            return wordIndex
        }
        nextWord(&i, sentence)
        wordIndex++
    }
    return -1
}