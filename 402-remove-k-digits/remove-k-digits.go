func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		for len(stack)>0 && stack[len(stack)-1]>num[i] && k>0 {
            stack = stack[:len(stack)-1]
            k--
        }
        if num[i] != '0' || len(stack) != 0{
            stack = append(stack, num[i])
        }
	}
	if k >= len(stack) {
		return "0"
	}
	return string(stack[:len(stack)-k])
}