func halvesAreAlike(s string) bool {
    vowelsSet := make(map[byte]struct{})
    for _, b := range []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'} {
        vowelsSet[b] = struct{}{}
    }

    leftCount, rightCount := 0, 0
    offset := len(s) / 2
    for i := 0; i < len(s) / 2; i++ {
        if _, ok := vowelsSet[s[i]]; ok {
            leftCount++
        }

        if _, ok := vowelsSet[s[i+offset]]; ok {
            rightCount++
        }
    }

    return leftCount == rightCount
}