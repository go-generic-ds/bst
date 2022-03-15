package bst

import "golang.org/x/exp/constraints"

type node[T constraints.Ordered] struct {
	val   T
	left  *node[T]
	right *node[T]
}

func insert[T constraints.Ordered](n **node[T], val T) {
	if *n == nil {
		*n = &node[T]{val: val}
	} else if val < (*n).val {
		insert(&(*n).left, val)
	} else if val > (*n).val {
		insert(&(*n).right, val)
	}
}

func findAndRemoveSucc[T constraints.Ordered](n **node[T]) *node[T] {
	if (*n).left != nil {
		return findAndRemoveSucc(&(*n).left)
	}
	ret := *n
	*n = ret.right
	return ret
}

func remove[T constraints.Ordered](n **node[T], val T) {
	if *n == nil {
		return
	}
	if val < (*n).val {
		remove(&(*n).left, val)
	} else if val > (*n).val {
		remove(&(*n).right, val)
	} else { // val == (*n).val
		if (*n).right == nil {
			*n = (*n).left
		} else if (*n).left == nil {
			*n = (*n).right
		} else {
			*n = findAndRemoveSucc(&(*n).right)
		}
	}
}

func contains[T constraints.Ordered](n **node[T], val T) bool {
	if *n == nil {
		return false
	}
	if val < (*n).val {
		return contains(&(*n).left, val)
	}
	if val > (*n).val {
		return contains(&(*n).right, val)
	}
	return true
}
