package tree

import (
	"fmt"
)

const (
	MAX_SIZE           = 500
	DEFAULT_ROOT_INDEX = 0
)

// Public interface function
func MakeHeap(b *BaseTree, isMaxHeap bool, heapSize int) *Heap {
	if heapSize <= 0 {
		heapSize = MAX_SIZE
	}

	return &Heap{
		BaseSequentialTree{
			*b,
			make([]*Node, heapSize),
			heapSize,
		},
		DEFAULT_ROOT_INDEX,
		isMaxHeap,
	}
}

// HEAP
type Heap struct {
	BaseSequentialTree
	nextInsertIndex int
	isMaxHeap       bool
}

func (self *Heap) String() string {
	heapType := "MinHeap"
	if self.isMaxHeap {
		heapType = "MaxHeap"
	}
	return fmt.Sprintf("Heap Size: %d\nHeap Type: %s\nTree:\n%v\n", self.maxSize, heapType, &self.BaseSequentialTree)
}

func (self *Heap) Insert(newVal interface{}) {
	if self.nextInsertIndex >= self.maxSize {
		panic("Heap size limit reached")
	}

	newNode := CreateTreeNode(&newVal)
	if self.root == nil {
		self.checkTypeForComparator(newNode)
		self.root = newNode
	}

	fmt.Println("Index - ", self.nextInsertIndex)
	fmt.Println("Value - ", newVal)

	// Inserting into the node arr
	self.nodeArr[self.nextInsertIndex] = newNode
	if self.nextInsertIndex != 0 {
		parentNode := self.nodeArr[self.getParentIndex(self.nextInsertIndex)]
		parentDirn := "right"
		if self.isLeftChild(self.nextInsertIndex) {
			parentDirn = "left"
		}
		parentNode.link[parentDirn] = newNode
		newNode.link["parent"] = parentNode

		// Reheap up from here
		self.reheapUp(newNode)

		fmt.Println("Dirn - ", parentDirn)
	}
	fmt.Println("*********************")
	self.len += 1
	self.nextInsertIndex += 1
}

func (self *Heap) reheapUp(node *Node) {
	// Compare the present node with it's parent
	parentNode := node.link["parent"]
	for parentNode != nil {
		if !self.isSizer(parentNode.data, node.data) {
			self.swapData(parentNode, node)
			node = parentNode
			parentNode = node.link["parent"]
		} else {
			break
		}
	}
}

func (self *Heap) swapData(node1, node2 *Node) {
	// Do not mess arround with the link pointers
	// Just change the data section
	container := node1.data
	node1.data = node2.data
	node2.data = container
}

// Used as a wrapper for the comparator for channelling in different heaps
func (self *Heap) isSizer(obj1, obj2 *interface{}) bool {
	if self.comparator != nil {
		compareResp := (*self.comparator)(obj1, obj2)
		if self.isMaxHeap {
			// obj1 >= obj2
			if compareResp >= 0 {
				return true
			}
		} else {
			if compareResp <= 0 {
				return true
			}
		}
	}
	return false
}

func (self *Heap) Pop() (*interface{}, bool) {
	panic("Not implemented yet!")
	// Bubble down
}

// Keeping this for the interface purpose.
// TODO: Think about how to navigate through this problem. Or just let it be
func (self *Heap) HasVal(*Node, interface{}) bool {
	panic("Heap does not understand the HasVal method. Get the root element")
}

func (self *Heap) Remove(interface{}) bool {
	panic("Heap does not understand the Remove method. Use pop()")
}
