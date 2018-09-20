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
	return recursiveRemoveRb(n, n, k)
}

// to retain r as root all the time
func recursiveRemoveRb(n, r *rbnode, k int) (*rbnode, interface{}) {
	var old interface{}
	if n == nil {
		return nil, nil
	}
	if k == n.k {
		old, n.v = n.v, nil
		s := swapSuccessorRb(n)
		if s == n && n.parent == nil { //n is single root
			n = nil
		} else {
			if s.c == black {
				borrowDownwardRb(s, r)
			} else { //red minimum
				s.parent.left, s.parent = nil, nil //disconnect
			}
		}
	} else {
		if k < n.k {
			_, old = recursiveRemoveRb(n.left, r, k)
		} else { //k > n.k
			_, old = recursiveRemoveRb(n.right, r, k)
		}
	}
	nn := checkToFlipColorOrRotate(n)
	return nn, old
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

func (n *rbnode) isLeaf() bool {
	return n.left == nil
}

func connectLeft(p, c *rbnode) {
	if p.left != nil && p.left.parent == p {
		p.left.parent = nil
	}
	p.left = c
	if c != nil {
		c.parent = p
	}
}
func connectRight(p, c *rbnode) {
	if p.right != nil && p.right.parent == p {
		p.right.parent = nil
	}
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
	c1.parent = nil
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

func dropRefRb(n *rbnode) {
	if n.left != nil && n.left.parent == n {
		n.left.parent = nil
	}
	if n.right != nil && n.right.parent == n {
		n.right.parent = nil
	}
	replaceChild(n, nil)
	n.left, n.right, n.parent = nil, nil, nil
}

func checkToFlipColorOrRotate(n *rbnode) *rbnode {
	if n == nil {
		return nil
	}
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

func swapSuccessorRb(n *rbnode) (f *rbnode) {
	if n.left == nil { //if leaf
		return n
	}
	if n.left.c == red && n.left.left == nil {
		f = n.left
	} else {
		f = floorRbTree(n.right)
	}
	n.k, f.k = f.k, n.k
	n.v, f.v = f.v, n.v
	return
}

func floorRbTree(n *rbnode) *rbnode {
	if n.left == nil {
		return n
	} else {
		return floorRbTree(n.left)
	}
}

// param:
//  h - hole
//  r - subtree root, maintained as root all time
func borrowDownwardRb(h, root *rbnode) {
	if h.v != nil {
		panic("h must be hole")
	}
	p := h.parent
	if h == root || p == nil { //replace reference
		if s := h.left; s != nil {
			connectLeft(h, s.left)
			connectRight(h, s.right)
			h.k, h.v, h.c = s.k, s.v, s.c
			dropRefRb(s)
		}
	} else {
		if h.c == black {
			if p.c == black { //black parent
				h.k, h.v, p.k, p.v = p.k, p.v, h.k, h.v
				if p.left == h {
					h.c = red
					r := p.right
					p.right = nil
					connectRight(h, r.left)
					connectLeft(r, h)
					connectLeft(p, r)
					borrowDownwardRb(p, r)
				} else { //black R/ black parent
					l := p.left
					if l.c == black {
						l.c = red
						connectRight(h, h.left)
						connectLeft(h, l)
						connectLeft(p, h)
						p.right = nil
						borrowDownwardRb(p, root)
					} else { //red L
						if l.left.c == black { // red L(black child)
							s := l.right
							p.k, p.v = s.k, s.v
							l.left.c, l.c = red, black
							connectRight(l, s.left)
							connectRight(h, h.left)
							connectLeft(h, s.right)
						} else { //red L(red left child)
							panic("double red links should've been rotated")
						}
					}
				}
			} else { //red parent
				if p.left == h { //black L/ red parent
					h.k, h.v = p.k, p.v
					h.c = red
					r := p.right
					p.k, p.v = r.k, r.v
					p.c = black
					connectRight(h, r.left)
					connectRight(p, r.right)
					connectLeft(p, h)
					dropRefRb(r)
				} else { //black R/ red parent
					l := p.left
					p.c, l.c = black, red
					connectRight(p, h.left)
					dropRefRb(h)
				}
			}
		} else { //hole is red
			if p.c == black {
				if p.left == h {
					connectLeft(p, h.left)
					dropRefRb(h)
				} else {
					panic("red link should be maintained at left branch")
				}
			} else {
				panic("double red links should've been rotated")
			}
		}
	}
}

func distanceRb(r, n *rbnode, d int) int {
	if n.c == black {
		d++
	}
	if n.parent == r {
		return d
	} else {
		return distanceRb(r, n.parent, d)
	}
}
