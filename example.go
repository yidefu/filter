package main

import (
	"fmt"
	"time"

	"github.com/yidefu/filter/sensitive"
)

func main() {
	s := sensitive.SetTreeByFile("sense.txt")
	now := time.Now()
	c := s.Replace("测试xxxx", "***", 0)
	fmt.Print(time.Since(now))
	fmt.Print(c)
}
