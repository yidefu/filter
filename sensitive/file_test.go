package sensitive_test

import (
	"testing"

	"github.com/yidefu/filter/sensitive"
)

func TestReadByLine(t *testing.T) {
	c := sensitive.ReadByLine("undefined.file")
	if len(c) != 0 {
		t.Fatalf("fail read file")
	}
	c = sensitive.ReadByLine("../sense.txt")
	if len(c) == 0 {
		t.Fatalf("fail read file")
	}
	t.Logf("success")
}
