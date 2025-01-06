func minOperations(boxes string) []int {
    var ml, mr, bl, br = 0, 0, 0, 0
    var ans []int = make([]int, len(boxes))

    i := 0
    j := len(boxes) - 1
    for i < len(boxes) {

        // left pass 
        ans[i] += ml
        if boxes[i] == '1' {
            bl += 1
        }
        ml += bl

        //right pass
        ans[j] += mr
        if boxes[j] == '1' {
            br += 1
        }
        mr += br
        i++
        j--
    }
    return ans
}