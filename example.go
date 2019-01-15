package main

import (
	"fmt"
	"test/sensitive"
	"time"
)

func main() {
	s := sensitive.SetTreeByFile("sense.txt")
	now := time.Now()
	c := s.Replace("测试xxxx", "***", 0)
	fmt.Print(time.Since(now))
	fmt.Print(c)
}
