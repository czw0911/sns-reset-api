/*
*@Description  令牌列表
*@Contact czw@outlook.com
*/
package libs

import (
	"sync"
)

func NewTokenList() *TokenList {
	return &TokenList{
		make(map[string]string),
		new(sync.Mutex),
	}
}

type TokenList struct {
	token map[string]string
	mux  *sync.Mutex
}

func (self *TokenList) Add(k,v string)  {
	self.mux.Lock()
	defer self.mux.Unlock()
	self.token[k] = v
	
}

func (self *TokenList) Del(k string) {
	self.mux.Lock()
	defer self.mux.Unlock()
	if _,ok := self.token[k]; ok {
		delete(self.token,k)
	}
}

func (self *TokenList) Get(k string) string {
	self.mux.Lock()
	defer self.mux.Unlock()
	if v,ok := self.token[k]; ok {
		return v
	}
	return ""
}