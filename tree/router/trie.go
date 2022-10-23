package trie

import (
	"fmt"
	"strings"
	"sync"
)

type Trie struct {
	MatchNext   map[string]*Trie
	DynamicNext map[string]*Trie
	End         bool
}

func NewTrie() *Trie {
	return &Trie{
		MatchNext:   map[string]*Trie{},
		DynamicNext: map[string]*Trie{},
		End:         false,
	}
}

func (t *Trie) Insert(s string) {
	node := t
	pre := -1
	d := -100
	n := len(s)
	for idx := 0; idx <= n; idx++ {
		ch := byte('/')
		if idx != n {
			ch = s[idx]
		}
		if ch != '/' {
			if ch == ':' && idx+1 < len(s) && s[idx+1] != ':' && s[idx+1] != '/' {
				d = idx
			}
			continue
		}
		if pre+1 >= idx {
			pre = idx
			d = idx
			continue
		}

		if d > pre {
			prefix := s[pre+1 : d]
			if node.DynamicNext[prefix] == nil {
				node.DynamicNext[prefix] = NewTrie()
			}
			node = node.DynamicNext[prefix]
		} else {
			ss := s[pre+1 : idx]
			if ss == "" {
				continue
			}
			if node.MatchNext[ss] == nil {
				node.MatchNext[ss] = NewTrie()
			}
			node = node.MatchNext[ss]
		}
		pre = idx
		d = idx
	}
	node.End = true
}

func (t *Trie) Search(s string) bool {
	node := t
	pre := -1
	n := len(s)
	for idx := 0; idx <= n; idx++ {
		ch := byte('/')
		if idx != n {
			ch = s[idx]
		}
		if ch == '?' {
			if idx > 0 && s[idx-1] == '/' {
				return node.End
			}
			ch = '/'
			s = s[:idx]
		}
		if ch != '/' {
			continue
		}
		if pre+1 >= idx {
			pre = idx
			continue
		}
		ss := s[pre+1 : idx]
		if next := node.MatchNext[ss]; next != nil {
			if next.Search(s[idx:]) {
				return true
			}
		}

		for ds, next := range node.DynamicNext {
			if len(ss) > len(ds) && ss[:len(ds)] == ds {
				if next.Search(s[idx:]) {
					return true
				}
			}
		}

		return false
	}

	return node.End
}

func (t *Trie) Print() {
	var helper func(node *Trie, layer int)
	helper = func(node *Trie, layer int) {
		for prefix, next := range node.MatchNext {
			fmt.Println(strings.Repeat("  ", layer), "m", prefix)
			helper(next, layer+1)
		}

		for prefix, next := range node.DynamicNext {
			fmt.Println(strings.Repeat("  ", layer), "d", prefix)
			helper(next, layer+1)
		}
	}

	helper(t, 0)
}

type Router struct {
	root *Trie
	lock sync.RWMutex
}

func NewRouter(arr []string) *Router {
	router := &Router{}
	router.ReBuild(arr)
	return router
}

func (r *Router) ReBuild(arr []string) {
	root := NewTrie()
	for _, s := range arr {
		root.Insert(s)
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	r.root = root
}

func (r *Router) Match(s string) bool {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.root.Search(s)
}

func (r *Router) Print() {
	r.lock.RLock()
	defer r.lock.RUnlock()
	r.root.Print()
}
