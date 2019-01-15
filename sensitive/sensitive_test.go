package sensitive_test

import (
	"fmt"
	"test/sensitive"
	"testing"
)

func TestSetTreeByFile(t *testing.T) {
	s := sensitive.SetTreeByFile("../sense.txt")
	tree := s.WordTree.Get("出")
	if tree == nil {
		t.Fatalf("fail to read hashmaptree")
	}
	tree = s.WordTree.Get("undefined")
	if tree != nil {
		t.Fatalf("fail to read hashmaptree")
	}
	t.Logf("success")
}

func TestSetTreeByArrs(t *testing.T) {
	s := sensitive.SetTreeByArrs([]string{"测试", "你好", "你们", "啥玩意"})
	tree := s.WordTree.Get("测")
	if tree == nil {
		t.Fatalf("fail to read hashmaptree")
	}
	tree = s.WordTree.Get("a")
	if tree != nil {
		t.Fatalf("fail to read hashmaptree")
	}
	t.Logf("success")
}

func TestGetBadWord(t *testing.T) {
	s := sensitive.SetTreeByArrs([]string{"测试", "你好", "你们", "啥玩意"})
	c := s.GetBadWord("测试grgrege哈哈gefewf啥玩意", 1, 0)
	fmt.Print(c)
	if len(c) != 2 {
		t.Fatalf("fail to get badword")
	}
	c = s.GetBadWord("fewdfnjwefewfwfew啥", 1, 0)
	if len(c) != 0 {
		t.Fatalf("fail to get badword")
	}
	t.Logf("success")
}

func TestReplace(t *testing.T) {
	s := sensitive.SetTreeByArrs([]string{"测试", "你好", "你们", "啥玩意"})
	c := s.Replace("测试grgrege哈哈gefewf啥玩意", "***", 1)
	if c != "***grgrege哈哈gefewf***" {
		t.Fatalf("fail to replace")
	}
	c = s.Replace("fewdfnjwefewfwfew啥", "***", 1)
	if c != "fewdfnjwefewfwfew啥" {
		t.Fatalf("fail to replace")
	}
	t.Logf("success")
}
