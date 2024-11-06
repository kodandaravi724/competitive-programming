import "strconv"

func evalRPN(tokens []string) int {
    stack := []int{}
    for _, token := range(tokens) {
        if token != "+" && token != "-" && token != "*" && token != "/" {
            num, _ := strconv.Atoi(token)
            stack = append(stack, num)
        } else {
            r := stack[len(stack)-1]
            l := stack[len(stack)-2]
            stack = stack[:len(stack)-2]
            res := 0
            if token == "+" {
                res+=(l+r)
            } else if token == "-" {
                res+=(l-r)
            } else if token == "*" {
                res+=(l*r)
            } else {
                res+=(l/r)
            } 
            stack = append(stack, res)
        }
    }
    return stack[0]
}