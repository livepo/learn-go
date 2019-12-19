package main
import "fmt"


type Trie struct {
	Root map[rune]interface{}
}



/** Initialize your data structure here. */
func Constructor() Trie {
	root := make(map[rune]interface{})
	return Trie{Root: root}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	r := this.Root
	for _, w := range word {
		if v, ok := r[w]; ok {
			r = v.(map[rune]interface{})
		} else {
			r[w] = make(map[rune]interface{})
			r = r[w].(map[rune]interface{})
		}
	}
	r['#'] = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this.Root
	for _, w := range word {
		if _, ok := root[w]; !ok {
			return false
		}
		root = root[w].(map[rune]interface{})
	}
	if _, ok := root['#']; !ok {
		return false
	}
	return true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this.Root
	for _, w := range prefix {
		if _, ok := root[w]; !ok {
			return false
		} else {
			root = root[w].(map[rune]interface{})
		}
	}
	return true
}


func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
