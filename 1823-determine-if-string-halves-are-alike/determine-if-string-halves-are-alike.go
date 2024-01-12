func halvesAreAlike(s string) bool {
    half := len(s)/2
    count2:=0
    count1:=0

    for i:=0; i<=half-1; i++ {
        if s[i]=='a' || s[i]=='e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i]=='A' || s[i]=='E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
            count1++
        }
        if s[i+half]=='a' || s[i+half]=='e' || s[i+half] == 'i' || s[i+half] == 'o' || s[i+half] == 'u' || s[i+half]=='A' || s[i+half]=='E' || s[i+half] == 'I' || s[i+half] == 'O' || s[i+half] == 'U'{
            count2++
        }
    }
    if count1 == count2 {
        return true
    }
    return false
}