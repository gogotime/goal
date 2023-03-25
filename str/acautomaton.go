package str

// AC自动机
type Node struct {
	next [26]*Node
	fail *Node
	end  bool
}

// https://leetcode.cn/problems/stream-of-characters/
type StreamChecker struct {
	root *Node
	cur  *Node
}

func Constructor(words []string) StreamChecker {
	root := &Node{}

	for _, s := range words {
		node := root
		for i := range s {
			ch := s[i]
			if node.next[ch-'a'] == nil {
				node.next[ch-'a'] = &Node{}
			}
			node = node.next[ch-'a']
		}
		node.end = true
	}

	q := []*Node{root}

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			u := q[0]
			q = q[1:]
			for v, c := range u.next {
				if c == nil {
					continue
				}
				if u == root {
					c.fail = root
				} else {
					p := u.fail
					for p != nil {
						if p.next[v] != nil {
							c.fail = p.next[v]
							break
						}
						p = p.fail
					}
					if p == nil {
						c.fail = root
					}
				}
				q = append(q, c)
			}
		}
	}

	return StreamChecker{root: root, cur: root}
}

func (this *StreamChecker) Query(letter byte) bool {
	for this.cur.next[letter-'a'] == nil && this.cur != this.root {
		this.cur = this.cur.fail
	}
	if this.cur.next[letter-'a'] != nil {
		this.cur = this.cur.next[letter-'a']
	}
	temp := this.cur
	for temp != this.root {
		if temp.end {
			return true
		}
		temp = temp.fail
	}
	return false
}
