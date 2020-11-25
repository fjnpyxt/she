package plaintext

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/zu1k/she/pkg/result"

	"github.com/zu1k/she/pkg/source"
)

type plaintext struct {
	name     string
	filePath string
}

func init() {
	source.Register(source.PlainText, newPlain)
}

func (p *plaintext) Name() string {
	return p.name
}

func (p *plaintext) Type() source.Type {
	return source.PlainText
}

// Search return result slice from source plaintext
func (p *plaintext) Search(key interface{}, resChan chan result.Result, wg *sync.WaitGroup) {
	str := key.(string)
	fmt.Println("Search plain text, key=", str)
	//开始搜索
	cmp := []byte(str)
	f, err := os.Open(p.filePath)
	if err != nil {
		wg.Done()
		return
	}
	sourceName := f.Name()
	defer f.Close()
	input := bufio.NewScanner(f)
	for input.Scan() {
		info := input.Bytes()
		if bytes.Contains(info, cmp) {
			result := result.Result{
				Source: sourceName,
				Score:  1,
				Hit:    str,
				Text:   string(info),
			}
			resChan <- result
		}
	}
	wg.Done()
}

func newPlain(name string, info interface{}) source.Source {
	path := info.(string)
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	return &plaintext{filePath: path, name: name}
}
