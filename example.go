package main

import (
	"fmt"
	"time"

	"github.com/yidefu/filter/sensitive"
)

func main() {
	s := sensitive.SetTreeByFile("sense.txt")
	now := time.Now()
	//matchType: 1代表最小匹配规则 其它值为最大匹配
	//wordNum: 0代表不限制返回数量  其它为返回数量上限
	c := s.GetBadWord("测试***", 1, 0)
	fmt.Print(time.Since(now))
	fmt.Print(c)
}
