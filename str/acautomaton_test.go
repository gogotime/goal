package str

import (
	"fmt"
	"testing"
)

func TestStreamChecker(t *testing.T) {
	//sc := Constructor([]string{"cd", "f", "kl"})
	//fmt.Println(sc.Query('a'))
	//fmt.Println(sc.Query('b'))
	//fmt.Println(sc.Query('c'))
	//fmt.Println(sc.Query('d'))
	//fmt.Println(sc.Query('e'))
	//fmt.Println(sc.Query('f'))
	//fmt.Println(sc.Query('g'))
	//fmt.Println(sc.Query('h'))
	//fmt.Println(sc.Query('i'))
	//fmt.Println(sc.Query('j'))
	//fmt.Println(sc.Query('k'))
	//fmt.Println(sc.Query('l'))

	sc := Constructor([]string{"ab", "ba", "abab"})
	fmt.Println(sc.Query('a'))
	fmt.Println(sc.Query('a'))
	fmt.Println(sc.Query('b'))
	fmt.Println(sc.Query('a'))
	fmt.Println(sc.Query('a'))

}
