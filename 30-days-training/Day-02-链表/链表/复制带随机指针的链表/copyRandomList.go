package main

var cachedNode map[*Node]*Node

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// https://leetcode-cn.com/problems/copy-list-with-random-pointer/solution/fu-zhi-dai-sui-ji-zhi-zhen-de-lian-biao-rblsf/

func deepCopy(node *Node) *Node {

	// 哈希的key是旧链表，有指向关系；哈希的value是新链表，只是结构。
	// 借助旧链表的指向关系，把新链表的指针连接上就行了。
	if node == nil {
		return nil
	}
	if n, exist := cachedNode[node]; exist {
		return n
	}
	newNode := &Node{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode

}

func copyRandomList(head *Node) *Node {
	//
	cachedNode = map[*Node]*Node{}
	return deepCopy(head)
}

func copyRandomListSplit(head *Node) *Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeTmp := node.Next
		node.Next = node.Next.Next
		if nodeTmp.Next != nil {
			nodeTmp.Next = nodeTmp.Next.Next
		}
	}
	return headNew

}

func copyRandomListNew(head *Node) *Node {
	//leetcode-cn.com/problems/copy-list-with-random-pointer/solution/138-fu-zhi-dai-sui-ji-zhi-zhen-de-lian-b-vf2x/
	cachedNode := map[*Node]*Node{}
	for node := head; node != nil; node = node.Next {
		cachedNode[node] = &Node{node.Val, nil, nil}
	}
	for node := head; node != nil; node = node.Next {
		cachedNode[node].Next = cachedNode[node.Next]
		cachedNode[node].Random = cachedNode[node.Random]
	}
	return cachedNode[head]
}
