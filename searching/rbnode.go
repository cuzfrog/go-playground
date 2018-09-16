package searching

type rbnode struct {
	k     int
	v     interface{}
	c     color
	left  *rbnode
	right *rbnode
}

type color bool

const (
	red   color = true
	black color = false
)

/*     1b         2b
      /  \        / \
     a  2r      1r   d
       /  \    /  \
      c   d    a  c
 */
func rotateLeft(n *rbnode) *rbnode {
	if n.right == nil{
		return n
	}
	r := n.right
	n.right, r.left = r.left, n
	n.c, r.c = red, black
	return r
}

/*        2b      1b
         / \     /  \
       1r   d   a   2r
       / \         / \
      a   c       c  d
 */
func rotateRight(n *rbnode) *rbnode {
	if n.left == nil{
		return n
	}
	l := n.left
	n.left, l.right = l.right, n
	n.c, l.c = red, black
	return l
}