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

func TestTypeCast(t *testing.T)  {
	var object interface{}
	object = &Dog{}
	dog,ok := object.(*Dog)
	dog.Move()
	assert.True(t, ok)
	assert.NotNil(t, dog)
}

//func TestTypeCast2(t *testing.T)  {
//	var object interface{}
//	object = &Cat{}
//	dog,ok := object.(*Dog)
//	assert.False(t, ok)
//	assert.Nil(t, dog)
//}

func TestDog(t *testing.T){
	dog := NewBlankDog()
	dog.Move()
}

func TestWhiteDog(t *testing.T)  {
	dog := new(WhiteDog)
	dog.Move()
}