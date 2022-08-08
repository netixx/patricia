// Code generated by automation. DO NOT EDIT

package complex64_tree

import (
	"fmt"

	"github.com/kentik/patricia"
)

// this is IPv6 tree code that's not very copy/paste friendly for when we transfer IPv4 code to IPv6

// create a new node in the tree, return its index
func (t *TreeV6) newNode(address patricia.IPv6Address, prefixLength uint) uint {
	availCount := len(t.availableIndexes)
	if availCount > 0 {
		index := t.availableIndexes[availCount-1]
		t.availableIndexes = t.availableIndexes[:availCount-1]
		t.nodes[index] = treeNodeV6{prefixLeft: address.Left, prefixRight: address.Right, prefixLength: prefixLength}
		return index
	}

	t.nodes = append(t.nodes, treeNodeV6{prefixLeft: address.Left, prefixRight: address.Right, prefixLength: prefixLength})
	return uint(len(t.nodes) - 1)
}

// Address returns the current IP address for the iterator.
func (iter *TreeIteratorV6) Address() patricia.IPv6Address {
	var prefixLeft, prefixRight uint64
	var prefixLength uint
	for _, i := range iter.nodeHistory {
		prefixLeft, prefixRight, prefixLength = patricia.MergePrefixes64(prefixLeft, prefixRight, prefixLength,
			iter.t.nodes[i].prefixLeft, iter.t.nodes[i].prefixRight, iter.t.nodes[i].prefixLength)
	}
	prefixLeft, prefixRight, prefixLength = patricia.MergePrefixes64(prefixLeft, prefixRight, prefixLength,
		iter.t.nodes[iter.nodeIndex].prefixLeft, iter.t.nodes[iter.nodeIndex].prefixRight,
		iter.t.nodes[iter.nodeIndex].prefixLength)
	return patricia.IPv6Address{
		Left:   prefixLeft,
		Right:  prefixRight,
		Length: prefixLength,
	}
}

//nolint
func (t *TreeV6) print() {
	buf := make([]complex64, 0)
	for i := range t.nodes {
		buf = buf[:0]
		fmt.Printf("%d: \tleft: %d, right: %d, prefix: %032b %032b (%d), tags: (%d): %v\n", i, int(t.nodes[i].Left), int(t.nodes[i].Right), int(t.nodes[i].prefixLeft), int(t.nodes[i].prefixRight), int(t.nodes[i].prefixLength), t.nodes[i].TagCount, t.tagsForNode(buf, uint(i), nil))
	}
}
