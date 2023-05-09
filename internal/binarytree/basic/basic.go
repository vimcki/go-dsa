package basic

type node struct {
	value  int
	left   *node
	right  *node
	parent *node
}

type BinTree struct {
	root *node
}

func New() *BinTree {
	return &BinTree{}
}

func (b *BinTree) Search(value int) bool {
	for current := b.root; current != nil; {
		if value < current.value {
			current = current.left
		} else if value > current.value {
			current = current.right
		} else {
			return true
		}
	}
	return false
}

func (b *BinTree) minimum(x *node) *node {
	for current := x; current != nil; {
		if current.left == nil {
			return current
		}
		current = current.left
	}
	return nil
}

func (b *BinTree) Insert(value int) {
	if b.root == nil {
		b.root = &node{value: value}
		return
	}
	current := b.root
	for {
		if value < current.value {
			if current.left == nil {
				current.left = &node{value: value, parent: current}
				return
			}
			current = current.left
		} else if value > current.value {
			if current.right == nil {
				current.right = &node{value: value, parent: current}
				return
			}
			current = current.right
		}
	}
}

func (b *BinTree) transplant(u, v *node) {
	if u.parent == nil {
		b.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (b *BinTree) Delete(value int) {
	var z *node
	for z = b.root; z != nil; {
		if value < z.value {
			z = z.left
		} else if value > z.value {
			z = z.right
		} else {
			break
		}
	}

	if z.left == nil {
		b.transplant(z, z.right)
	} else if z.right == nil {
		b.transplant(z, z.left)
	} else {
		y := b.minimum(z.right)
		if y.parent != z {
			b.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		b.transplant(z, y)
		y.left = z.left
		y.left.parent = y
	}
}
