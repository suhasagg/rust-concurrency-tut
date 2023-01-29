type Node struct {
	left, right *Node
	start, end  net.IP
	blocked      bool
}

func (n *Node) build(ips []net.IP) {
	n.start = ips[0]
	n.end = ips[len(ips)-1]
	if len(ips) == 1 {
		return
	}

	mid := len(ips) / 2
	n.left = &Node{}
	n.left.build(ips[:mid])
	n.right = &Node{}
	n.right.build(ips[mid:])
}

func (n *Node) query(ip net.IP) bool {
	if bytes.Compare(ip, n.start) < 0 || bytes.Compare(ip, n.end) > 0 {
		return false
	}

	if n.left == nil {
		return n.blocked
	}

	if bytes.Compare(ip, n.left.end) <= 0 {
		return n.left.query(ip)
	}
	return n.right.query(ip)
}

func main() {
	ips := []net.IP{
		net.ParseIP("192.168.0.0"),
		net.ParseIP("192.168.0.5"),
		net.ParseIP("192.168.0.10"),
		net.ParseIP("192.168.0.255"),
	}

	root := &Node{}
	root.build(ips)

	fmt.Println(root.query(net.ParseIP("192.168.0.7")))
	fmt.Println(root.query(net.ParseIP("192.168.1.7")))
}
