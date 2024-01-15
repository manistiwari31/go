import "sort"

func findWinners(matches [][]int) [][]int {
    result := make([][]int, 2, 2)
    winners := make([]int, 0, len(matches))
    losers := make([]int, 0, len(matches))

    // storing number of loses for every team 
    cnt := make(map[int]int, len(matches))
    for _, val := range matches {
        cnt[val[0]] += 0
        cnt[val[1]]++
    }

    // saving to losers or winners based on cnt map values 
    for key, val := range cnt {
        if val == 0 {
            winners = append(winners, key)
        } else if val == 1 {
            losers = append(losers, key)
        }
    }

    // sorting slices
    sort.Ints(winners)
    sort.Ints(losers)

    // return results 
    result[0] = winners
    result[1] = losers
    return result
}