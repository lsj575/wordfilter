package datasource

import (
	"bufio"
	"github.com/lsj575/wordfilter/models"
	"io"
	"os"
	"strings"
	"sync"
)

const SensitiveFile string = "words/sensitive_words.txt"

var trieLock sync.Mutex
var sensitiveWordTrieInstance *models.Trie

func InstanceSensitiveWord() *models.Trie {
	if sensitiveWordTrieInstance != nil {
		return sensitiveWordTrieInstance
	}
	trieLock.Lock()
	defer trieLock.Unlock()

	if sensitiveWordTrieInstance != nil {
		return sensitiveWordTrieInstance
	}

	return newSensitiveWordTrie()
}

func newSensitiveWordTrie() *models.Trie {
	sensitiveWordTrieInstance = models.NewTrie()
	err := importSensitiveWords(sensitiveWordTrieInstance, SensitiveFile)
	if err != nil {
		return nil
	}
	return sensitiveWordTrieInstance
}

//导入过滤词库
func importSensitiveWords(T *models.Trie, file string) error {
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