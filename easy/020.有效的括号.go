func isValid(s string) bool {
    m := make(map[rune]rune)
    m[')'] = '('
    m['}'] = '{'
    m[']'] = '['
    e := -1
    arr := ""
    for _, v := range s{
        if v == '(' || v == '{' || v == '['{
            arr += string(v)
            e++
        }else{
            if e < 0 || len(arr) <= 0{
                return false
            }
            if rune(arr[e]) != m[v] {
                return false
            }
            arr = arr[:e]
            e--
        }
    }
    if e >= 0 {
        return false
    }
    return true
}