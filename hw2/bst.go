package main

import "fmt"

type bst_elem struct {
	val   int
	left  *bst_elem
	right *bst_elem
}

type bst struct {
	root *bst_elem
}

func new_bst_elem(v int) *bst_elem {
	new := &bst_elem{
		val:   v,
		left:  nil,
		right: nil,
	}
	return new
}

func (b *bst) Add(v int) {
	b.root = b.add(b.root, v)
}

func (b *bst) add(elem *bst_elem, v int) *bst_elem {
	if elem == nil {
		return new_bst_elem(v)
	}
	switch {
	case v > elem.val:
		elem.right = b.add(elem.right, v)
	case v < elem.val:
		elem.left = b.add(elem.left, v)
	}
	return elem
}
func (b *bst) Delete(v int) {
	b.root = b.delete(b.root, v)
}

func (b *bst) delete(elem *bst_elem, v int) *bst_elem {
	if elem == nil {
		return nil
	}

	if v < elem.val {
		elem.left = b.delete(elem.left, v)
	} else if v > elem.val {
		elem.right = b.delete(elem.right, v)
	} else {
		switch {
		case elem.left == nil && elem.right == nil:
			elem = nil
		case elem.left == nil:
			elem = elem.right
		case elem.right == nil:
			elem = elem.left
		default:
			temp := min_tree_elem(elem)
			if temp == elem.right {
				elem.right.left = elem.left
				elem = elem.right
			} else {
				temp.left = temp.right
				elem = temp
			}

		}

	}
	return elem
}
func min_tree_elem(elem *bst_elem) *bst_elem {
	for elem.left != nil {
		elem = elem.left
	}
	return elem
}

func (b *bst) IsExist(v int) bool {
	return b.isexist(b.root, v)
}

func (b *bst) isexist(elem *bst_elem, v int) bool {
	if elem == nil {
		return false
	}
	if elem.val == v {
		return true
	}
	if v < elem.val {
		return b.isexist(elem.left, v)
	}
	return b.isexist(elem.right, v)
}

func main() {
	q := []int{2, 3, -5, 7, -11, 13}
	tree := bst{}
	for _, i := range q {
		tree.Add(i)
	}

	f := 7
	fmt.Printf("Is %d have been found? %t\n", f, tree.IsExist(f))

	tree.Delete(7)
	fmt.Printf("Is %d have been found? %t\n", f, tree.IsExist(f))

}
