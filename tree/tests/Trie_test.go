package testSuite

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"gotree/tree"
	//"sort"
	"testing"
)

func compareIntSlices(arr1, arr2 []int) bool {
	for index, _ := range arr1 {
		if arr1[index] != arr2[index] {
			return false
		}
	}
	return true
}

func TestTrie_development(t *testing.T) {
	assert := assert.New(t)

	// Create a Trie
	trieObj := tree.CreateTrie()
	assert.NotEmpty(trieObj)

	// Insertion
	assert.True(trieObj.Insert("basic Str"))

	// Searching
	assert.True(trieObj.HasVal("basic Str"))
	assert.True(trieObj.HasVal("BASIC STR"))
	assert.True(trieObj.HasVal("basic"))
	assert.False(trieObj.HasVal("avengers"))

	// Create a trie with custom options
	// No substring matching. Only complete words.
	// Case sensitive
	trieObjOpt := tree.CreateTrieWithOptions(false, false)
	assert.NotEmpty(trieObjOpt)

	// Insertion
	assert.True(trieObjOpt.Insert("New String"))

	// Searching
	assert.True(trieObjOpt.HasVal("New String"))
	assert.False(trieObjOpt.HasVal("BASIC STRING"))
	assert.False(trieObjOpt.HasVal("New"))
}

func TestTrie_multiple(t *testing.T) {
	assert := assert.New(t)

	// Create a simple trie
	trieObj := tree.CreateTrie()

	// Inserting
	insertStatus := trieObj.Insert(
		"Letme",
		"tell",
		"you",
		"a",
		"story",
		"to",
		"chill",
		"your",
		"bones",
		"About",
		"a",
		"thing",
		"that",
		"I",
		"saw",
	)
	assert.True(insertStatus)

	assert.True(trieObj.HasVal("tel"))
	assert.True(trieObj.HasVal("let"))
	assert.True(trieObj.HasVal("a"))
	assert.True(trieObj.HasVal("i"))
	assert.True(trieObj.HasVal("chil"))
	assert.True(trieObj.HasVal("bones"))

	assert.False(trieObj.HasVal("tella"))
	assert.False(trieObj.HasVal(""))
	assert.False(trieObj.HasVal("wooot"))
	assert.False(trieObj.HasVal("chiller"))

}
