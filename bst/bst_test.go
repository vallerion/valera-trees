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
		deleted := tree.Delete(randSlice[i])
		if deleted == false {
			t.Error("Deletion method returned false, expected true.")
		}
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

func BenchmarkInsert(b *testing.B) {
	randSlice := make([]int, 1000)
	gofakeit.Slice(&randSlice)

	for n := 0; n < b.N; n++ {
		tree := ConstructIntTree()
		for i := 0; i < len(randSlice); i++ {
			tree.Insert(randSlice[i], randSlice[i])
		}
	}
}

func BenchmarkSelect(b *testing.B) {
	length := 1000

	for n := 0; n < b.N; n++ {
		gofakeit.Seed(0)
		tree, randSlice := initRandTree(length)

		_, found := tree.Search(randSlice[gofakeit.Number(0, length-1)])

		if found == false {
			b.Error("[BenchmarkSelect]: Value must be found.")
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	length := 1000

	for n := 0; n < b.N; n++ {
		gofakeit.Seed(0)
		tree, randSlice := initRandTree(length)

		randItem := randSlice[gofakeit.Number(0, length-1)]

		deleted := tree.Delete(randItem)
		if deleted == false {
			b.Error("[BenchmarkSelect]: Method 'delete' returned false, but expected true.")
		}

		_, found := tree.Search(randItem)
		if found == true {
			b.Error("[BenchmarkSelect]: Search after deletion must returns false, but it's returned true.")
		}
	}
}

func initRandTree(length int) (*BST, []int) {
	randSlice := make([]int, length)
	gofakeit.Slice(&randSlice)

	tree := ConstructIntTree()
	for i := 0; i < len(randSlice); i++ {
		tree.Insert(randSlice[i], randSlice[i])
	}

	return tree, randSlice
}
