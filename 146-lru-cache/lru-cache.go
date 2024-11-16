type LRUCache struct {
    capacity int
    hashMap  map[int]*LinkedListNode
    head     *LinkedListNode
    tail     *LinkedListNode
}

type LinkedListNode struct {
    Next *LinkedListNode
    Prev *LinkedListNode
    Key  int
    Val  int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity: capacity,
        hashMap:  make(map[int]*LinkedListNode),
        head:     nil,
        tail:     nil,
    }
}

func (this *LRUCache) Get(key int) int {
    node, exists := this.hashMap[key]
    if !exists {
        return -1
    }
    // Move the accessed node to the head
    this.moveToHead(node)
    return node.Val
}

func (this *LRUCache) Put(key int, value int) {
    node, exists := this.hashMap[key]
    if exists {
        // Update the value and move to head
        node.Val = value
        this.moveToHead(node)
    } else {
        // Create a new node
        newNode := &LinkedListNode{
            Key: key,
            Val: value,
        }
        this.hashMap[key] = newNode
        this.addNode(newNode)

        if len(this.hashMap) > this.capacity {
            // Remove the least recently used node
            tail := this.popTail()
            delete(this.hashMap, tail.Key)
        }
    }
}

// Helper to add a new node at the head
func (this *LRUCache) addNode(node *LinkedListNode) {
    node.Next = this.head
    node.Prev = nil

    if this.head != nil {
        this.head.Prev = node
    }
    this.head = node

    if this.tail == nil {
        this.tail = node
    }
}

// Helper to remove a node
func (this *LRUCache) removeNode(node *LinkedListNode) {
    if node.Prev != nil {
        node.Prev.Next = node.Next
    } else {
        this.head = node.Next
    }

    if node.Next != nil {
        node.Next.Prev = node.Prev
    } else {
        this.tail = node.Prev
    }
}

// Helper to move a node to the head
func (this *LRUCache) moveToHead(node *LinkedListNode) {
    this.removeNode(node)
    this.addNode(node)
}

// Helper to remove the tail node
func (this *LRUCache) popTail() *LinkedListNode {
    tail := this.tail
    this.removeNode(tail)
    return tail
}
