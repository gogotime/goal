package st

import "testing"

func TestST(t *testing.T) {
	st := NewST([]int{4, 5, 3, 1, 2, 9, 7})
	if st.Query(1, 2) != 5 ||
		st.Query(3, 6) != 9 {
		t.Error("not correct")
		return
	}
}
