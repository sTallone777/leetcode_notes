func lengthOfLastWord(s string) int {
    s = strings.Trim(s, " ")
    return len(s)-1-strings.LastIndex(s," ")
}