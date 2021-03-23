func isPalindrome(s string) bool {
    if len(s) == 0 {
        return true
    }
    l := 0
    r := len(s) - 1
    for l < r {
        lc := isTarget(int(s[l]))
        lr := isTarget(int(s[r]))
        if lc == -1 {
            l++
            continue
        }
        if lr == -1 {
            r--
            continue
        }
        if lc != lr {
            return false
        }
        l++
        r--
    }
    return true
}

func isTarget(r int) int{
    if (r >= 48 && r <=57) || (r >= 97 && r <= 122) {
        return r
    }
    if (r >= 65 && r <= 90){
        return r + 32
    }
    return -1
}