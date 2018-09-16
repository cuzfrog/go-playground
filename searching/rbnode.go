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
