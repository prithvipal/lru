package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutFirstNode(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	assert.Len(t, lruN.bucket, 1)
	assert.Equal(t, lruN.l.String(), "A=Apple,")
}

func TestPutKeyTwise(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("A", "Apple")
	assert.Len(t, lruN.bucket, 1)
	assert.Equal(t, lruN.l.String(), "A=Apple,")
}

func TestPutKeyTwiseDifferentVal(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("A", "ApplE")
	assert.Len(t, lruN.bucket, 1)
	assert.Equal(t, lruN.l.String(), "A=ApplE,")
}

func TestPutHalf(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	assert.Len(t, lruN.bucket, 2)
	assert.Equal(t, lruN.l.String(), "A=Apple,B=Ball,")
}

func TestPutHitFirstEleWithTwoElements(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("A", "ApplE")
	assert.Len(t, lruN.bucket, 2)
	assert.Equal(t, lruN.l.String(), "B=Ball,A=ApplE,")
}

func TestPutHitLastEleWithTwoElements(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("B", "Ball")
	assert.Len(t, lruN.bucket, 2)
	assert.Equal(t, lruN.l.String(), "A=Apple,B=Ball,")
}

func TestPutFull(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("C", "Cat")
	lruN.put("D", "Dog")
	assert.Len(t, lruN.bucket, 4)
	assert.Equal(t, lruN.l.String(), "A=Apple,B=Ball,C=Cat,D=Dog,")
}

func TestPutHitMiddleEle(t *testing.T) {
	lruN := NewLRU(5)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("C", "Cat")
	lruN.put("B", "BaLL")
	assert.Len(t, lruN.bucket, 3)
	assert.Equal(t, lruN.l.String(), "A=Apple,C=Cat,B=BaLL,")
}

func TestPutRemoveFirstRecentlyUsed(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("C", "Cat")
	lruN.put("D", "Dog")
	lruN.put("E", "Elephent")
	assert.Len(t, lruN.bucket, 4)
	assert.Equal(t, lruN.l.String(), "B=Ball,C=Cat,D=Dog,E=Elephent,")
}

func TestPutHitLastElement(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("C", "Cat")
	lruN.put("D", "Dog")
	lruN.put("D", "DOG")
	assert.Len(t, lruN.bucket, 4)
	assert.Equal(t, lruN.l.String(), "A=Apple,B=Ball,C=Cat,D=DOG,")
}

func TestPutRemoveSecondRecentlyUsed(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("A", "Apple")
	lruN.put("C", "Cat")
	lruN.put("D", "Dog")
	lruN.put("E", "Elephent")
	assert.Len(t, lruN.bucket, 4)
	assert.Equal(t, lruN.l.String(), "A=Apple,C=Cat,D=Dog,E=Elephent,")
}
