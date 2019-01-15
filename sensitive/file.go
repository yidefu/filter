package sensitive

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

func ReadByLine(filedir string) (content []string) {
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		file, err := os.Open(filedir)
		defer wait.Done()
		if err != nil {
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			content = append(content, strings.Replace(scanner.Text(), "\r", "", -1))
		}
		if err := scanner.Err(); err != nil {
			return
		}
	}()
	wait.Wait()
	return content
}
