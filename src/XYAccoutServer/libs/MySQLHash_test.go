
package libs

import (	
	"testing"
	"fmt"
)



func TestMySQLHash(t *testing.T) {
	h := NewMySQLHash()
	h.DBCountMaste = 2
	h.DBCountSlave = 5
	for i := 0 ; i < 100 ; i++ {
		uid := fmt.Sprintf("chenzhangwei_%06d",i)
		t.Log(uid)
		
		res := h.GetMySQLHashByUID(1,uid)
		t.Log(res)
		t.Log("----master---")
		
		res = h.GetMySQLHashByUID(0,uid)
		t.Log(res)
		t.Log("----slave---")
	}
}





