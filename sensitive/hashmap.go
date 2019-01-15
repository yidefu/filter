package sensitive

import (
	"strings"
)

type tree interface{}

type HashMap struct {
	Trees map[string]tree
}

func NewHashMap() *HashMap {
	return &HashMap{make(map[string]tree)}
}

func (h *HashMap) buildWordToTree(word string) {
	word = strings.TrimSpace(word)
	if word == "" {
		return
	}
	runes := []rune(word)
	for i, r := range runes {
		value := h.Get(string(r))
		if value != nil {
			h = value.(*HashMap)
		} else {
			t := NewHashMap()
			t.Put("ending",false)
			h.Put(string(r),t)
			h = t
		}
		if i == (len(runes) - 1) {
			h.Put("ending",true)
		}
	}
	return
}

func (h *HashMap) Get(key string) interface{} {
	var a interface{} = nil
	t, err := h.Trees[key]
	if err {
		return t
	}
	return a
}

 func (h *HashMap) Put(key string, value interface{}) interface{} {
 	 tmp, err := h.Trees[key]
 	 var v interface{}
 	 switch value.(type) {
 	     case bool:v = value.(bool)
	     case *HashMap:v = value.(*HashMap)
	     default:
 	 }
 	 if err == false {
 		 h.Trees[key] = v
		 return nil
 	 }
	 h.Trees[key] = v
 	 return tmp
 }
