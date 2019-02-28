package models

import (
	"log"
	"sync"
)

type Trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Children map[rune]*TrieNode
	End      bool
}

func NewTrie() *Trie {
	t := &Trie{}
	t.Root = t.NewTrieNode()

	return t
}

func (T *Trie) NewTrieNode() *TrieNode {
	n := &TrieNode{}
	n.Children = make(map[rune]*TrieNode)
	n.End = false

	return n
}

// 新增要过滤的词
func (T *Trie) Add(txt string) bool {
	if len(txt) < 1 {
		log.Println("trie:trie Add 字符串长度小于1")
		return true
	}
	flag := true	// 表示trie树中是否已经存在该字符串
	chars := []rune(txt)
	sLen := len(chars)
	node := T.Root
	locker := sync.Mutex{}
	locker.Lock()
	defer locker.Unlock()
	for i := 0; i < sLen; i++ {
		if _, exists := node.Children[chars[i]]; !exists {
			flag = false
			node.Children[chars[i]] = T.NewTrieNode()
		}
		node = node.Children[chars[i]]
	}
	node.End = true
	return flag
}

func (T *Trie) Find(s string) bool {
	chars := []rune(s)
	sLen := len(chars)
	node := T.Root
	for i := 0; i < sLen; i++ {
		if _, exists := node.Children[chars[i]]; exists {
			node = node.Children[chars[i]]
			for j := i + 1; j < sLen; j++ {
				if _, exists := node.Children[chars[j]]; !exists {
					break
				}
				node = node.Children[chars[j]]
				if node.End == true {
					return true
				}
			}
			node = T.Root
		}
	}
	return false
}
// 屏蔽字搜索替换
func (T *Trie) Replace(txt string, whiteTire *Trie) (string, []string) {
	chars := []rune(txt)
	result := []rune(txt)
	find := make([]string, 0, 10)
	sLen := len(chars)
	node := T.Root
	for i := 0; i < sLen; i++ {
		if _, exists := node.Children[chars[i]]; exists {
			node = node.Children[chars[i]]
			for j := i + 1; j < sLen; j++ {
				if _, exists := node.Children[chars[j]]; !exists {
					break
				}
				node = node.Children[chars[j]]
				pre := i - 4
				if pre < 0 {
					pre = 0
				}
				tail := j+1+4
				if tail > sLen {
					tail = sLen
				}
				if node.End == true && !whiteTire.Find(string(chars[pre:tail])) {
					for t := i; t <= j; t++ {
						result[t] = '*'
					}
					find = append(find, string(chars[i:j+1]))
					i = j
					node = T.Root
					break
				}
			}
			node = T.Root
		}
	}

	return string(result), find
}
