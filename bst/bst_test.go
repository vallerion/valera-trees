package bst

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSimpleIntInsert(t *testing.T) {
	tree := ConstructIntTree()

	key, value := rand.Int(), rand.Int()

	tree.Insert(key, value)

	if tree.root == nil {
		t.Error("Root mustn't be nil!")
	}

	if tree.root != nil && tree.root.key != key {
		t.Errorf("Root's key must be %d !", key)
	}

	if tree.root != nil && tree.root.value != value {
		t.Errorf("Root's value must be %d !", value)
	}
}

func TestInsert(t *testing.T) {
	tree := ConstructIntTree()

	randSlice := make([]int, 1000)
	gofakeit.Slice(&randSlice)

	for i := 0; i < len(randSlice); i++ {
		tree.Insert(randSlice[i], randSlice[i])
	}

	for i := len(randSlice) - 1; i >= 0; i-- {
		assertFound(tree, randSlice[i], t)
	}

	if len(tree.ToArray()) != len(randSlice) {
		t.Error("Initial array length and tree's array length must be equals.")
	}

	assert.ElementsMatch(t, randSlice, tree.ToArray(), "Initial array and tree's array must be equals.")
}

func TestDelete(t *testing.T) {
	tree := ConstructIntTree()

	randSlice := make([]int, 1000)
	gofakeit.Slice(&randSlice)

	for i := 0; i < len(randSlice); i++ {
		tree.Insert(randSlice[i], randSlice[i])
	}

	for i := len(randSlice) / 4; i >= 0; i-- {
		tree.Delete(randSlice[i])
		assertNotFound(tree, randSlice[i], t)
	}

	if len(tree.ToArray()) == len(randSlice) {
		t.Error("Initial array length and tree's array length mustn't be equals.")
	}
}

func assertFound(tree *BST, searchedKey int, t *testing.T) {
	value, found := tree.Search(searchedKey)
	if found == false {
		t.Errorf("Element %d must be found. But 'found' flag is false", searchedKey)
	}
	if value == searchedKey {
		t.Errorf("Found element %d must be equals to searched %d.", value, searchedKey)
	}
}

func assertNotFound(tree *BST, searchedKey int, t *testing.T) {
	_, found := tree.Search(searchedKey)
	if found == true {
		t.Errorf("Element %d mustn't be found. But 'found' flag is true", searchedKey)
	}
}
