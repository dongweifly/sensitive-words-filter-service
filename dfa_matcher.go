package main

type DFAMatcher struct {
	replaceChar rune
	root        *Node
}

//最长匹配： 你好，你好啊，优先匹配你好啊，如果没有你好啊，那么匹配你好

type Node struct {
	End  bool
	Next map[rune]*Node
}

func (n *Node) AddWord(word string) {
	node := n
	chars := []rune(word)
	for index, _ := range chars {
		node = node.AddChild(chars[index])
	}
	node.End = true
}

func (n *Node) AddChild(c rune) *Node {
	if n.Next == nil {
		n.Next = make(map[rune]*Node)
	}

	//如果已经存在了，就不再往里面添加了；
	if next, ok := n.Next[c]; ok {
		return next
	} else {
		n.Next[c] = &Node{
			End:  false,
			Next: nil,
		}
		return n.Next[c]
	}
}

func (n *Node) FindChild(c rune) *Node {
	if n.Next == nil {
		return nil
	}

	if _, ok := n.Next[c]; ok {
		return n.Next[c]
	}
	return nil
}

func NewDFAMather() *DFAMatcher {
	return &DFAMatcher{
		root: &Node{
			End: false,
		},
	}
}

func (d *DFAMatcher) Build(words []string) {
	for _, item := range words {
		d.root.AddWord(item)
	}
}

//Match 查找替换发现的敏感词
func (d *DFAMatcher) Match(text string) (sensitiveWords []string, replaceText string) {
	if d.root == nil {
		return nil, text
	}

	textChars := []rune(text)
	textCharsCopy := make([]rune, len(textChars))
	copy(textCharsCopy, textChars)

	length := len(textChars)
	for i := 0; i < length; i++ {
		//root本身是没有key的，root的下面一个节点，才算是第一个；
		temp := d.root.FindChild(textChars[i])
		if temp == nil {
			continue
		}
		j := i + 1
		for ; j < length && temp != nil; j++ {
			if temp.End {
				sensitiveWords = append(sensitiveWords, string(textChars[i:j]))
				replaceRune(textCharsCopy, '*', i, j)
			}
			temp = temp.FindChild(textChars[j])
		}

		if j == length && temp != nil && temp.End {
			sensitiveWords = append(sensitiveWords, string(textChars[i:length]))
			replaceRune(textCharsCopy, '*', i, length)
		}
	}
	return sensitiveWords, string(textCharsCopy)
}

func replaceRune(chars []rune, replaceChar rune, begin int, end int) {
	for i := begin; i < end; i++ {
		chars[i] = replaceChar
	}
}
