package BST

import (
	"Play-with-Data-Structures/Utils/Interfaces"
)

type Node struct {
	e     interface{}
	left  *Node
	right *Node
}

func generateNode(e interface{}) *Node {
	return &Node{e: e}
}

type BST struct {
	root *Node
	size int
}

func Constructor() *BST {
	return &BST{}
}

func (this *BST) GetSize() int {
	return this.size
}

func (this *BST) IsEmpty() bool {
	return this.size == 0
}

// 向二分搜索树中添加新的元素 e
func (this *BST) Add(e interface{}) {
	this.root = this.add(this.root, e)
}

// 向以 Node 为跟的二分搜索树中插入元素 e，递归算法
// 返回插入新节点后二分搜索树的根
func (this *BST) add(n *Node, e interface{}) *Node {
	if n == nil {
		this.size++
		return generateNode(e)
	}

	// 递归调用
	if Interfaces.Compare(e, n.e) < 0 {
		n.left = this.add(n.left, e)
	} else if Interfaces.Compare(e, n.e) > 0 {
		n.right = this.add(n.right, e)
	}
	return n
}
