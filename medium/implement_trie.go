package medium

// Implement a trie with insert, search, and startsWith methods.
//
// Example:
// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple");   // returns true
// trie.search("app");     // returns false
// trie.startsWith("app"); // returns true
// trie.insert("app");
// trie.search("app");     // returns true
//
// Note:
//
// You may assume that all inputs are consist of lowercase letters a-z.
// All inputs are guaranteed to be non-empty strings.

type Trie struct {
	next   map[rune]*Trie
	isWord bool
}

func Constructor() Trie {
	return Trie{next: make(map[rune]*Trie), isWord: false}
}

func (t *Trie) Insert(word string) {
	for _, v := range word {
		if t.next[v] == nil {
			t.next[v] = &Trie{next: make(map[rune]*Trie), isWord: false}
		}
		t = t.next[v]
	}
	t.isWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.prefixNode(word)
	return node != nil && node.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.prefixNode(prefix)
	return node != nil
}

// prefixNode returns node related with prefix or nil if this node is not exist.
func (t *Trie) prefixNode(prefix string) *Trie {
	for _, v := range prefix {
		if t.next[v] == nil {
			return nil
		}
		t = t.next[v]
	}
	return t
}

// The official soloution:
//
// Trie node structure:
// Image location: [/image/solutions/implement_trie.jpeg]
//
// class TrieNode {
//
//     // R links to node children
//     private TrieNode[] links;
//
//     private final int R = 26;
//
//     private boolean isEnd;
//
//     public TrieNode() {
//         links = new TrieNode[R];
//     }
//
//     public boolean containsKey(char ch) {
//         return links[ch -'a'] != null;
//     }
//     public TrieNode get(char ch) {
//         return links[ch -'a'];
//     }
//     public void put(char ch, TrieNode node) {
//         links[ch -'a'] = node;
//     }
//     public void setEnd() {
//         isEnd = true;
//     }
//     public boolean isEnd() {
//         return isEnd;
//     }
// }
//
// Insertion of a key to a trie:
//
// class Trie {
//     private TrieNode root;
//
//     public Trie() {
//         root = new TrieNode();
//     }
//
//     // Inserts a word into the trie.
//     public void insert(String word) {
//         TrieNode node = root;
//         for (int i = 0; i < word.length(); i++) {
//             char currentChar = word.charAt(i);
//             if (!node.containsKey(currentChar)) {
//                 node.put(currentChar, new TrieNode());
//             }
//             node = node.get(currentChar);
//         }
//         node.setEnd();
//     }
// }
//
// Search for a key in a trie:
//
// class Trie {
//     ...
//
//     // search a prefix or whole key in trie and
//     // returns the node where search ends
//     private TrieNode searchPrefix(String word) {
//         TrieNode node = root;
//         for (int i = 0; i < word.length(); i++) {
//            char curLetter = word.charAt(i);
//            if (node.containsKey(curLetter)) {
//                node = node.get(curLetter);
//            } else {
//                return null;
//            }
//         }
//         return node;
//     }
//
//     // Returns if the word is in the trie.
//     public boolean search(String word) {
//        TrieNode node = searchPrefix(word);
//        return node != null && node.isEnd();
//     }
// }
//
// Search for a key prefix in a trie:
//
// class Trie {
//     ...
//
//     // Returns if there is any word in the trie
//     // that starts with the given prefix.
//     public boolean startsWith(String prefix) {
//         TrieNode node = searchPrefix(prefix);
//         return node != null;
//     }
// }
