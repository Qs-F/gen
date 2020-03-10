package gen

// Tree is tree of sum of depth.
type Tree []*Depth

func (list List) GenTree(root *Node) Tree {
	im := make(imported)
	zeroDepth := &Depth{
		Nodes: []*Node{root},
	}

	ret := Tree{zeroDepth}

	for {
		depth := resolve(list, im, zeroDepth)
		if depth == nil {
			break
		}
		ret = append(ret, depth)
	}

	return ret
}

func (tree Tree) Reduce() Variables {
	variables := make(Variables)

	for _, depth := range tree {
		for _, node := range depth.Nodes {
			for k, v := range node.Content {
				if k == ImportIdent {
					continue
				}
				if _, ok := variables[k]; ok {
					continue
				}
				variables[k] = v
			}
		}
	}

	return variables
}

type imported map[string]bool

// Depth is nodes slice for the same depth.
type Depth struct {
	Nodes []*Node
}

// resolve returns next depth.
// resolve must be called from root to leaves, left to right
func resolve(list List, im imported, depth *Depth) *Depth {
	nodes := []*Node{}
	for _, node := range depth.Nodes {
		// if node is already imported, then continue
		if b, ok := im[node.Key]; ok && b {
			continue
		}

		if !node.resolved {
			node.link(list)
		}

		nodes = append(nodes, node.children...)
		im[node.Key] = true
	}

	if len(nodes) == 0 {
		return nil
	}

	return &Depth{
		Nodes: nodes,
	}
}
