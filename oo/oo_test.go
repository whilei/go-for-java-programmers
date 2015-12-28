package oo
import (
"testing"
	"github.com/stretchr/testify/assert"
)

func TestMissMethodInterface(t *testing.T) {
	dog := Dog{}
	assert.True(t,doMove(&dog))
	assert.False(t,doMove2(&dog))
}

func doMove(object interface{}) bool {
	animal,ok := object.(Animal)
	if(!ok){
		return false
	}
	animal.Move()
	return true
}

func doMove2(object interface{}) bool {
	animal,ok := object.(Animal2)
	if(!ok){
		return false
	}
	animal.Move()
	return true
}