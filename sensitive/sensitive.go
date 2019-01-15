package sensitive

import (
	"strings"
)

type Sensitive struct {
	contentLength int
	WordTree      *HashMap
}

func SetTreeByFile(filedir string) *Sensitive {
	contents := ReadByLine(filedir)
	return buildTrees(contents)
}

func SetTreeByArrs(contents []string) *Sensitive {
	return buildTrees(contents)
}

func buildTrees(contents []string) *Sensitive {
	s := &Sensitive{contentLength: 0, WordTree: NewHashMap()}
	if len(contents) == 0 {
		return s
	}
	for _, content := range contents {
		s.WordTree.buildWordToTree(content)
	}
	return s
}

func (s *Sensitive) GetBadWord(content string, matchType, wordNum uint8) (badwordlist []string) {
	c := []rune(content)
	s.contentLength = len(c)
	if s.contentLength == 0 {
		return badwordlist
	}
	for i := 0; i < s.contentLength; i++ {
		matchFlag := 0
		flag := false
		tempMap := s.WordTree
		for j := i; j < s.contentLength; j++ {
			key := string(c[j])
			nowMap := tempMap.Get(key)
			if nowMap == nil {
				break
			}
			tempMap = nowMap.(*HashMap)
			matchFlag++
			ending := tempMap.Get("ending").(bool)
			if false == ending {
				continue
			}
			flag = true
			if 1 == matchType {
				break
			}
		}

		if !flag {
			matchFlag = 0
		}
		if matchFlag <= 0 {
			continue
		}
		badwordlist = append(badwordlist, string(c[i:(i+matchFlag)]))
		length := len(badwordlist)
		if wordNum > 0 && uint8(length) == wordNum {
			return badwordlist
		}
		i = i + matchFlag - 1
	}
	return badwordlist
}

func (s *Sensitive) Replace(content, replace string, matchType uint8) string {
	badwordlist := s.GetBadWord(content, matchType, 0)
	if len(badwordlist) == 0 {
		return content
	}
	var oldnew []string
	for _, bad := range badwordlist {
		oldnew = append(oldnew, bad, replace)
	}
	return strings.NewReplacer(oldnew...).Replace(content)
}
