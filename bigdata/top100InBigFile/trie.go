package top100InBigFile

type Trie struct {
	Letter   byte
	Children [26]*Trie
	Cnt      int
}

func NewTrie(letter byte) Trie {
	return Trie{
		Letter:   letter,
		Children: [26]*Trie{},
		Cnt:      0,
	}
}

func GenTrie() {
	var words []string

	root := NewTrie(0)
	for _, word := range words {
		addWord(&root, word)
	}
}

func addWord(root *Trie, word string) {
	for i := range word {
		letter := word[i]

	}
}
