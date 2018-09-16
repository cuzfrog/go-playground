package searching

type rbnode struct {
	k      int
	v      interface{}
	c      color
	left   *rbnode
	right  *rbnode
	parent *rbnode
}

type color bool

const (
	red   color = true
	black color = false
)

/* ----------------- functions ----------------- */

// return top node and old value
func (n *rbnode) insert(k int, v interface{}) (*rbnode, interface{}) {
	var old interface{}
	if n == nil { //empty, create single black
		n = &rbnode{k: k, v: v, c: red}
	} else if k == n.k {
		old, n.v = n.v, v
	} else if k < n.k {
		n.left, old = n.left.insert(k, v)
	} else { //k > n.k
		n.right, old = n.right.insert(k, v)
	}
	n = checkToFlipColorOrRotate(n)
	return n, old
}

func (n *rbnode) remove(k int) (*rbnode, interface{}) {
	var old interface{}
	if n != nil {
		if k == n.k {
			s := swapInOrderSuccessorRb(n)
			if s.c == black {
				n = borrowDownwardRb(s)
			} else { //red minimum
				old = s.v
				s.parent.left = nil
			}
		} else if k < n.k {
			n.left, old = n.left.remove(k)
		} else { //k > n.k
			n.right, old = n.right.remove(k)
		}
	}

	return n, old
}

/* ----------------- utils ----------------- */

/*     1b         2b
      /  \        / \
     a  2r      1r   d
       /  \    /  \
      c   d    a  c
 */
func rotateLeft(n *rbnode) *rbnode {
	if n.right == nil {
		return n
	}
	r := n.right
	n.right, r.left = r.left, n
	n.c, r.c = r.c, n.c
	return r
}

/*        2b      1b
         / \     /  \
       1r   d   a   2r
       / \         / \
      a   c       c  d
 */
func rotateRight(n *rbnode) *rbnode {
	if n.left == nil {
		return n
	}
	l := n.left
	n.left, l.right = l.right, n
	n.c, l.c = l.c, n.c
	return l
}

func (n *rbnode) isBlack() bool {
	return n == nil || n.c == black
}

func checkToFlipColorOrRotate(n *rbnode) *rbnode {
	if n.left.isBlack() && n.right.isBlack() {
		//do nothing
	} else if n.right.isBlack() && n.left.c == red {
		if !n.left.left.isBlack() { // double red on left branch
			n = rotateRight(n)
		}
	} else if n.left.isBlack() && n.right.c == red {
		n = rotateLeft(n)
	}
	if !n.left.isBlack() && !n.right.isBlack() { //both l,r are red
		n.left.c, n.right.c = black, black
		n.c = red
	}
	return n
}

//todo:test
func swapInOrderSuccessorRb(n *rbnode) *rbnode {
	if n.right == nil {
		return n
	}
	f := floorRbTree(n.right)
	n.k, f.k = f.k, n.k
	n.v, f.v = f.v, n.v
	return f
}

func floorRbTree(n *rbnode) *rbnode {
	if n.left == nil {
		return n
	} else {
		return floorRbTree(n.left)
	}
}

func borrowDownwardRb(n *rbnode) *rbnode {
	panic("implement me")
	return n
}
