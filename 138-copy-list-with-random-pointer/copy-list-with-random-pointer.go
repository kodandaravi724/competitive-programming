/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {

    if head == nil {
        return nil
    }

    var listMap map[*Node]*Node = make(map[*Node]*Node)

    temp := head
    for temp != nil {
        listMap[temp] = &Node{
            temp.Val, nil, nil, 
        }
        temp = temp.Next
    }
    temp = head
    for temp != nil {
        nodeP := listMap[temp]
        if temp.Next != nil {
                nodeP.Next = listMap[temp.Next]
        } else {
            nodeP.Next = nil
        }
        if temp.Random != nil {
                nodeP.Random = listMap[temp.Random]
        } else {
            nodeP.Random = nil
        }
        temp = temp.Next
    }
    return listMap[head]
}