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
	} else {
		var nn *rbnode
		if k < n.k {
			nn, old = n.left.insert(k, v)
			connectLeft(n, nn)
		} else { //k > n.k
			nn, old = n.right.insert(k, v)
			connectRight(n, nn)
		}
	}
	n = checkToFlipColorOrRotate(n)
	return n, old
}

func (n *rbnode) remove(k int) (*rbnode, interface{}) {
	var old interface{}
	if n != nil {
		if k == n.k {
			old = n.v
			s := swapInOrderSuccessorRb(n)
			if s == n && n.parent == nil { //n is single root
				n = nil
			} else {
				if s.c == black {
					borrowDownwardRb(s)
				} else { //red minimum
					s.parent.left, s.parent = nil, nil //disconnect
				}
			}
		} else {
			var nn *rbnode
			if k < n.k {
				nn, old = n.left.remove(k)
				connectLeft(n, nn)
			} else { //k > n.k
				nn, old = n.right.remove(k)
				connectRight(n, nn)
			}
		}
	}

	return n, old
}

/* ----------------- utils ----------------- */

/*      p          p
        |          |
       1b         2b
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
	connectRight(n, r.left)
	replaceChild(n, r)
	connectLeft(r, n)
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
	connectLeft(n, l.right)
	replaceChild(n, l)
	connectRight(l, n)
	n.c, l.c = l.c, n.c
	return l
}

func (n *rbnode) isBlack() bool {
	return n == nil || n.c == black
}

func connectLeft(p, c *rbnode) {
	p.left = c
	if c != nil {
		c.parent = p
	}
}
func connectRight(p, c *rbnode) {
	p.right = c
	if c != nil {
		c.parent = p
	}
}

// replace c1 with c2 in words of parent relationship
func replaceChild(c1, c2 *rbnode) {
	if p := c1.parent; p != nil {
		if p.left == c1 {
			connectLeft(p, c2)
		} else {
			connectRight(p, c2)
		}
	}
}

func disconnectRb(p, c *rbnode) {
	if p.left == c {
		p.left = nil
	} else if p.right == c {
		p.right = nil
	}
	if c != nil {
		c.parent = nil
	}
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

func borrowDownwardRb(n *rbnode) {
	p := n.parent
	if p == nil {

	} else {
		if n.c == black {
			if p.c == red {
				disconnectRb(p, n)
				rotateLeft(p)
			} else {

			}
		} else {

		}
	}
}
