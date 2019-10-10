package hard

// Given a 2D board and a list of words from the dictionary,
// find all words in the board.
//
// Each word must be constructed from letters of sequentially
// adjacent cell, where "adjacent" cells are those horizontally
// or vertically neighboring. The same letter cell may not be used
// more than once in a word.
//
// Example:
// Input:
// board = [
//   ['o','a','a','n'],
//   ['e','t','a','e'],
//   ['i','h','k','r'],
//   ['i','f','l','v']
// ]
// words = ["oath","pea","eat","rain"]
//
// Output: ["eat","oath"]
//
// Note:
// All inputs are consist of lowercase letters a-z.
// The values of words are distinct.

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(words) == 0 {
		return nil
	}

	trie := constructor()
	for _, word := range words {
		trie.insert(word)
	}

	newBoard := pretreatBoard(board)

	ss := NewSearchSet(trie, newBoard)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			ss.dfs(i+1, j+1)
		}
	}

	return ss.result()
}

// generateVisit returns a 2'D array which flag a position whether was visited.
func generateVisit(len int) [][]bool {
	visit := make([][]bool, len)
	for i := range visit {
		visit[i] = make([]bool, len)
	}
	return visit
}

// pretreatBoard returns a newBoard used 0 to surround board.
func pretreatBoard(board [][]byte) (newBoard [][]byte) {
	row := len(board)
	col := len(board[0])

	newBoard = make([][]byte, row+2)
	for i := range newBoard {
		newBoard[i] = make([]byte, col+2)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			newBoard[i+1][j+1] = board[i][j]
		}
	}

	return newBoard
}

type searchSet struct {
	trie    *trie    // trie store all words which will be searched
	board   [][]byte // a board should surround with 0
	visited [][]bool // visited flag whether current character was visited.

	// resultWords is result words which can be found in board.
	// Here use resultWords as a set, because maybe there has
	// many path to find an identical word.
	resultWords map[string]struct{}

	dynamicArr [1 << 10]rune // dynamicArr record one path. Suppose a word length less than 1024.
	index      int           // index record index of dynamicArr
}

func NewSearchSet(t *trie, board [][]byte) *searchSet {
	return &searchSet{
		trie:        t,
		board:       board,
		visited:     generateVisit(len(board)),
		resultWords: make(map[string]struct{}),
	}
}

func (s *searchSet) dfs(row, col int) {
	s.dynamicArr[s.index] = rune(s.board[row][col])
	if !s.trie.startsWith(string(s.dynamicArr[:s.index+1])) {
		return
	}

	if s.trie.search(string(s.dynamicArr[:s.index+1])) {
		s.resultWords[string(s.dynamicArr[:s.index+1])] = struct{}{}
		// NOTE: can not return, because maybe prefix in tree is a word
	}

	s.visited[row][col] = true
	s.index++

	if s.board[row+1][col] != 0 && !s.visited[row+1][col] {
		s.dfs(row+1, col)
	}
	if s.board[row-1][col] != 0 && !s.visited[row-1][col] {
		s.dfs(row-1, col)
	}
	if s.board[row][col+1] != 0 && !s.visited[row][col+1] {
		s.dfs(row, col+1)
	}
	if s.board[row][col-1] != 0 && !s.visited[row][col-1] {
		s.dfs(row, col-1)
	}

	s.index--
	s.visited[row][col] = false
}

// result returns all words which can be found in board.
func (s *searchSet) result() []string {
	var words []string
	for word := range s.resultWords {
		words = append(words, word)
	}
	return words
}

type trie struct {
	next   map[rune]*trie
	isWord bool
}

func constructor() *trie {
	return &trie{next: make(map[rune]*trie), isWord: false}
}

func (t *trie) insert(word string) {
	for _, v := range word {
		if t.next[v] == nil {
			t.next[v] = &trie{next: make(map[rune]*trie), isWord: false}
		}
		t = t.next[v]
	}
	t.isWord = true
}

func (t *trie) search(word string) bool {
	node := t.prefixNode(word)
	return node != nil && node.isWord
}

func (t *trie) startsWith(prefix string) bool {
	node := t.prefixNode(prefix)
	return node != nil
}

// prefixNode returns node related with prefix or nil if this node is not exist.
func (t *trie) prefixNode(prefix string) *trie {
	for _, v := range prefix {
		if t.next[v] == nil {
			return nil
		}
		t = t.next[v]
	}
	return t
}
