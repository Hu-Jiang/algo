package medium

// Design a data structure that supports the following two operations:
//
// void addWord(word)
// bool search(word)
//
// search(word) can search a literal word or a regular expression string
// containing only letters a-z or '.'. A '.' means it can represent any
// one letter.
//
// Example:
// addWord("bad")
// addWord("dad")
// addWord("mad")
// search("pad") -> false
// search("bad") -> true
// search(".ad") -> true
// search("b..") -> true
//
// Note:
// You may assume that all words are consist of lowercase letters a-z.

type WordDictionary struct {
	next   map[rune]*WordDictionary
	isWord bool
}

func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{next: make(map[rune]*WordDictionary), isWord: false}
}

func (this *WordDictionary) AddWord(word string) {
	for _, v := range word {
		if this.next[v] == nil {
			this.next[v] = &WordDictionary{next: make(map[rune]*WordDictionary), isWord: false}
		}
		this = this.next[v]
	}
	this.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	for k, v := range word {
		if v == '.' {
			for _, v := range this.next {
				if v.Search(word[k+1:]) {
					return true
				}
			}
			return false
		} else {
			if this.next[v] == nil {
				return false
			}
			this = this.next[v]
		}
	}
	return this.isWord
}

// A Java soloution from discuss in leetcode:
//
// public class WordDictionary {
//
//     public class TrieNode {
//         public TrieNode[] children = new TrieNode[26];
//         public boolean isWord;
//     }
//
//     private TrieNode root = new TrieNode();
//
//     // Adds a word into the data structure.
//     public void addWord(String word) {
//         TrieNode node = root;
//         for (char c : word.toCharArray()) {
//             if (node.children[c - 'a'] == null) {
//                 node.children[c - 'a'] = new TrieNode();
//             }
//             node = node.children[c - 'a'];
//         }
//         node.isWord = true;
//     }
//
//     // Returns if the word is in the data structure. A word could
//     // contain the dot character '.' to represent any one letter.
//     public boolean search(String word) {
//         return match(word.toCharArray(), 0, root);
//     }
//
//     private boolean match(char[] chs, int k, TrieNode node) {
//         if (k == chs.length) {
//             return node.isWord;
//         }
//         if (chs[k] == '.') {
//             for (int i = 0; i < node.children.length; i++) {
//                 if (node.children[i] != null && match(chs, k + 1, node.children[i])) {
//                     return true;
//                 }
//             }
//         } else {
//             return node.children[chs[k] - 'a'] != null && match(chs, k + 1, node.children[chs[k] - 'a']);
//         }
//         return false;
//     }
// }
