package tree

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
)

type huffmanData struct {
	id      string
	dataVal string
	freq    int
	leaf    bool
	link    map[string]*huffmanData
}

type HuffmanTree struct {
	root          *huffmanData
	priorityQueue *Heap
	freqMap       map[string]int
	encodingMap   map[string]string
}

func CreateHuffmanTree(freqMap map[string]int) *HuffmanTree {
	comparatorFunc := func(obj1, obj2 *interface{}) int {
		new_obj1 := (*obj1).(huffmanData)
		new_obj2 := (*obj2).(huffmanData)

		// Never send 0 - Doesn't have any use
		if new_obj1.freq >= new_obj2.freq {
			return 1
		} else {
			return -1
		}
	}

	huffmanTree := &HuffmanTree{
		nil,
		CreateHeapWithComparator(&comparatorFunc, false, 500),
		freqMap,
		make(map[string]string),
	}

	huffmanTree.buildTree()

	return huffmanTree
}

func (self *HuffmanTree) buildTree() bool {
	if len(self.freqMap) == 0 {
		return false
	}

	fmt.Println("In buildtree")
	leavesMapPtr := make(map[string]*huffmanData)

	respStatus := true
	// Put everything inside the priority queue
	var interfaceData interface{}
	for keyStr, freq := range self.freqMap {
		uuid, _ := uuid.NewV4()
		newData := huffmanData{
			uuid.String(),
			keyStr,
			freq,
			true,
			map[string]*huffmanData{
				"0":      nil,
				"1":      nil,
				"parent": nil,
			},
		}
		interfaceData = newData
		respStatus = respStatus &&
			self.priorityQueue.Insert(interfaceData)
		leavesMapPtr[keyStr] = &newData
	}

	// Pop 2 elements at a time.
	for !self.priorityQueue.IsEmpty() {
		leftChildInt, isValidLChild := self.priorityQueue.Pop()
		rightChildInt, isValidRChild := self.priorityQueue.Pop()

		if isValidLChild {
			leftChild := (*leftChildInt).(huffmanData)
			if isValidRChild {
				rightChild := (*rightChildInt).(huffmanData)
				uuid, _ := uuid.NewV4()

				// Add the data from both the nodes
				newData := huffmanData{
					uuid.String(),
					"",
					leftChild.freq + rightChild.freq,
					false,
					map[string]*huffmanData{
						"1":      &rightChild,
						"0":      &leftChild,
						"parent": nil,
					},
				}
				rightChild.link["parent"] = &newData
				leftChild.link["parent"] = &newData

				respStatus = respStatus &&
					self.priorityQueue.Insert(newData)
			} else {
				// Only one child remains. Add it to the tree
				self.root = &leftChild
			}
		}
	}

	// Fill the encoding map
	// Traverse it upwards
	for keyStr, valPtr := range leavesMapPtr {
		fmt.Println(keyStr)
		leafNode := valPtr
		codedStr := ""
		for leafNode != nil {
			fmt.Println(leafNode.freq)
			fmt.Println(leafNode.link)
			//Find out what child the current node is
			parentNode := leafNode.link["parent"]
			if parentNode != nil {
				selectedCode := "1"
				if parentNode.link["0"].id == leafNode.id {
					selectedCode = "0"
				}
				fmt.Println("selected code - ", selectedCode)

				codedStr = selectedCode + codedStr

				leafNode = parentNode
			} else {
				break
			}
		}
		self.encodingMap[keyStr] = codedStr
	}
	return true
}

func (self *HuffmanTree) EncodeStr(ipStr string) string {
	return "Not yet implemented"
}