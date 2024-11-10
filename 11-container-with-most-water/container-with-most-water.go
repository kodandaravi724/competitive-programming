func maxArea(height []int) int {
    var i, j int = 0, len(height) - 1
    var maxArea int = -1
    for i < j {
        maxArea = max(maxArea, (j-i)*min(height[i], height[j]))
        if height[i] < height[j] {
            i+=1
        } else {
            j-=1
        }
    }
    return maxArea   
}

func min(i, j int) int {
    if i < j {
        return i
    }
    return j
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}
