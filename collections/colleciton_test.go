package collections
import (
	"testing"
	"fmt"
)

func TestMapLength(t *testing.T){
	m := make(map[int]int, 1)
	fmt.Println(len(m))
	m[0] = 0
	m[1] = 1
	fmt.Println(len(m))
}

func TestSliceLength(t *testing.T){
	m := make([]int, 10)
	fmt.Println(len(m))

}
