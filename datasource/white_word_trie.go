package datasource

import (
	"bufio"
	"github.com/lsj575/wordfilter/models"
	"io"
	"os"
	"strings"
	"sync"
)

const WhiteFile string = "words/white_words.txt"

var sensitiveTrieLock sync.Mutex
var whiteWordTrieInstance *models.Trie

func InstanceWhiteWord() *models.Trie {
	if whiteWordTrieInstance != nil {
		return whiteWordTrieInstance
	}
	sensitiveTrieLock.Lock()
	defer sensitiveTrieLock.Unlock()

	if whiteWordTrieInstance != nil {
		return whiteWordTrieInstance
	}

	return newWhiteWordTrie()
}

func newWhiteWordTrie() *models.Trie {
	whiteWordTrieInstance = models.NewTrie()
	err := importWhiteWords(whiteWordTrieInstance, WhiteFile)
	if err != nil {
		return nil
	}
	return whiteWordTrieInstance
}

//导入过滤词库
func importWhiteWords(T *models.Trie, file string) error {
	rd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer rd.Close()
	r := bufio.NewReader(rd)
	for {
		line, isPrefix, e := r.ReadLine()
		if e != nil {
			if e != io.EOF {
				err = e
			}
			break
		}
		if isPrefix {
			continue
		}
		if word := strings.TrimSpace(string(line)); word != "" {
			T.Add(word)
		}
	}
	return nil
}