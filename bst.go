package bst

import "golang.org/x/exp/constraints"

type BST[T constraints.Ordered] struct {
	root *node[T]
}

func (b *BST[T]) Insert(val T) {
	insert(&b.root, val)
}

func (b *BST[T]) Remove(val T) {
	remove(&b.root, val)
}

func (b *BST[T]) Contains(val T) bool {
	return contains(&b.root, val)
}
