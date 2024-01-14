func minimumSwap(s1 string, s2 string) int {
    n := len(s1)
    x , y := 0,0
    for i:= 0; i< n ; i++ {
        if s1[i]!=s2[i] {
            if s1[i] == 'x' {
                x++
            } else {
                y++
            }
        }
    }
    if (x+y)%2 != 0 {
        return -1
    }

    return x / 2 + y / 2 + x % 2 + y % 2
}