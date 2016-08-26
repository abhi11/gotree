package tree

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
)

type customData struct {
	Num int
}
type customMap map[string]*Node

type Node struct {
	id   string
	data *customData
	link customMap
}

// TODO: Stackoverflow when using cycles, as when using parent key
func (n *Node) String() string {
	return fmt.Sprintf("Data: %d\nMap: {\n%s\n}\n", n.data, n.link)
}

func (n *Node) AddChild(key string, childPtr *Node) {
	n.link[key] = childPtr
	// TODO: Think of what to do with this block
	/*
		_, isExists := childPtr.link["parent"]
		if !isExists {
			childPtr.link["parent"] = n
		}
	*/
}

/*
Creates a Tree node that can be added to the tree
TODO; This abstraction is needed now because, not sure
how to  make this more general. Probably use something like
an interface{}
*/
func CreateTreeNode(n int) *Node {
	customData := createTreeData(n)
	return makeNode(customData)
}

// CreateCustomData creates Tree data populated from the argument and
// returns a reference to the customData
func createTreeData(n int) *customData {
	return &customData{n}
}

// Creates a node
func makeNode(data *customData) *Node {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic("Error generating a new UUID.")
	}
	return &Node{
		id:   uuid.String(),
		data: data,
		link: make(map[string]*Node),
	}
}
