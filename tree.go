// Package tree builds a tree of Nodes from Records.
package tree

import (
	"fmt"
)

type Record struct {
	ID, Parent int
}
type Node struct {
	ID       int
	Children []*Node
}

const (
	rootID = 0  // causes problems watching for duplicate records, so...
	rootXX = -1 // a nonzero, invalid record ID to substitute for rootID
)

// Build creates a tree from Records.
// Given: IDs in range 0..len(records)-1,
// Parent ID < ID for non-root, Parent ID == ID for root.
// Assumption: record IDs are unique. Implication: root ID will be zero
func Build(records []Record) (*Node, error) {
	var root *Node
	var err error
	if len(records) > 0 {
		nodes := make([]Node, len(records)) // direct index by ID
		for _, r := range records {
			if err = r.chk(len(records), nodes); err != nil {
				goto quit
			}
			parent := r.Parent
			if parent == rootID {
				parent = rootXX // temporary substitute for rootID
			}
			nodes[r.ID].ID = parent // save here for now
		}
		if nodes[rootID].ID != rootXX {
			err = fmt.Errorf("No valid root found")
		} else {
			root = &nodes[rootID]
			// the Children node lists are created in a second pass because
			// the test routine expects them to be in ascending ID order
			for i := range nodes {
				parent := nodes[i].ID
				if parent == rootXX {
					parent = rootID
				}
				nodes[i].ID = i // overwrite temp storage of parent
				if i != rootID {
					addChild(&nodes[parent].Children, &nodes[i])
				}
			}
		}
	}
quit:
	return root, err
}

// addChild adds child node n to the list of children
func addChild(children *[]*Node, n *Node) {
	*children = append(*children, n)
}

// chk validates a Record or returns an error
func (r Record) chk(reclen int, nodes []Node) error {
	var err error
	if r.ID >= reclen || r.ID < 0 {
		err = fmt.Errorf("Record ID %d out of range", r.ID)
	} else if r.Parent < 0 {
		err = fmt.Errorf("Invalid Parent ID %d", r.Parent)
	} else if nodes[r.ID].ID != 0 { // 0: unseen, valid ID or rootXX: seen
		err = fmt.Errorf("Duplicate record ID %d", r.ID)
	} else if r.ID > 0 && r.Parent >= r.ID {
		err = fmt.Errorf("Record %d has bad Parent ID %d", r.ID, r.Parent)
	} else if r.ID == 0 && r.Parent != 0 { // bad root ID
		err = fmt.Errorf("root has bad Parent %d", r.Parent)
	}
	return err
}
