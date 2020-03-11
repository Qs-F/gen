package gen

// Node is the deps tree node.
type Node struct {
	Key     string
	Content Variables
	Deps    []string

	children []*Node
	resolved bool
}

// ToNode converts Variable to Node.
func (v Variables) ToNode(key string) *Node {
	return &Node{
		Key:      key,
		Content:  v,
		Deps:     v.ToImports(),
		children: []*Node{},
		resolved: false,
	}
}

// resolve resolves 1 incremental depth for node.
// resolve will changes node.children and node.resolved.
func (n *Node) link(list List) {
	// if already resolved, then no more resolve
	if n.resolved {
		return
	}

	for _, dep := range n.Deps {
		v, err := list.Get(dep)

		// if dep is not found in list, then continue
		if err != nil {
			continue
		}

		n.children = append(n.children, v.ToNode(dep))
	}
	n.resolved = true
}
