package slice_fn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tkaelbel/slice_fn"
)

type Test struct {
	FirstName  string
	SecondName string
	Age        int
}

func TestFilter(t *testing.T) {
	in := []int{1, 2, 2, 4, 2}
	expected := []int{2, 2, 2}

	actual := slice_fn.Filter(in, func(val int) bool { return val == 2 })

	assert.Equal(t, expected, actual)
}

func TestReduce(t *testing.T) {
	in := []int{1, 2, 2, 4, 2}
	expected := 11

	actual := slice_fn.Reduce(in, 0, func(previous int, current int) int { return previous + current })

	assert.Equal(t, expected, actual)
}

func TestMap(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	expected := []int{2, 3, 4, 5, 6}

	actual := slice_fn.Map(in, func(val int) int { return val + 1 })

	assert.Equal(t, expected, actual)
}

func TestFindInt(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := 4
	actual := slice_fn.Find(input, func(val int) bool { return val == 4 })

	assert.Equal(t, expected, actual)
}

func TestNotFindInt(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := 0
	actual := slice_fn.Find(input, func(val int) bool { return val == 7 })

	assert.Equal(t, expected, actual)

}

func TestFindStruct(t *testing.T) {

	input := []Test{{FirstName: "Karl", SecondName: "Otto", Age: 47}, {FirstName: "Donald", SecondName: "Duck", Age: 112}}
	expected := input[1]

	actual := slice_fn.Find(input, func(val Test) bool { return val.FirstName == "Donald" })
	assert.Equal(t, expected, actual)

}
func TestNotFindStruct(t *testing.T) {
	input := []Test{{FirstName: "Karl", SecondName: "Otto", Age: 47}, {FirstName: "Donald", SecondName: "Duck", Age: 112}}
	expected := Test{}

	actual := slice_fn.Find(input, func(val Test) bool { return val.FirstName == "Hans" })
	assert.Equal(t, expected, actual)
}

func TestFindIndex(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := 2
	actual := slice_fn.FindIndex(input, func(val int) bool { return val == 3 })

	assert.Equal(t, expected, actual)
}

func TestFindStructIndex(t *testing.T) {
	input := []Test{{FirstName: "Karl", SecondName: "Otto", Age: 47}, {FirstName: "Donald", SecondName: "Duck", Age: 112}}
	expected := 1
	actual := slice_fn.FindIndex(input, func(val Test) bool { return val.FirstName == "Donald" })

	assert.Equal(t, expected, actual)
}

func TestFindNoIndex(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := -1
	actual := slice_fn.FindIndex(input, func(val int) bool { return val == 47 })

	assert.Equal(t, expected, actual)
}

func TestFindLast(t *testing.T) {
	input := []int{1, 5, 3, 4, 5}
	expected := 5
	actual := slice_fn.FindLast(input, func(val int) bool { return val == 5 })

	assert.Equal(t, expected, actual)
}

func TestNoFindLast(t *testing.T) {
	input := []int{1, 5, 3, 4, 5}
	expected := 0
	actual := slice_fn.FindLast(input, func(val int) bool { return val == 4771 })

	assert.Equal(t, expected, actual)
}

func TestFindLastStruct(t *testing.T) {
	input := []Test{{FirstName: "Karl", SecondName: "Otto", Age: 47}, {FirstName: "Donald", SecondName: "Duck", Age: 112}, {FirstName: "Donald", SecondName: "Duck", Age: 115}}
	expected := Test{FirstName: "Donald", SecondName: "Duck", Age: 112}
	actual := slice_fn.FindLast(input, func(val Test) bool { return val.Age == 112 })

	assert.Equal(t, expected, actual)
}

func TestFindLastIndexStruct(t *testing.T) {
	input := []Test{{FirstName: "Karl", SecondName: "Otto", Age: 47}, {FirstName: "Donald", SecondName: "Duck", Age: 112}, {FirstName: "Donald", SecondName: "Duck", Age: 112}}
	expected := 2
	actual := slice_fn.FindLastIndex(input, func(val Test) bool { return val.Age == 112 })

	assert.Equal(t, expected, actual)
}

func TestFindNoLastIndexStruct(t *testing.T) {
	input := []Test{{FirstName: "Karl", SecondName: "Otto", Age: 47}, {FirstName: "Donald", SecondName: "Duck", Age: 112}, {FirstName: "Donald", SecondName: "Duck", Age: 112}}
	expected := -1
	actual := slice_fn.FindLastIndex(input, func(val Test) bool { return val.Age == 5000 })

	assert.Equal(t, expected, actual)
}

func TestForEachStruct(t *testing.T) {
	input := []Test{{FirstName: "Karl", SecondName: "Otto", Age: 47}, {FirstName: "Donald", SecondName: "Duck", Age: 112}, {FirstName: "Donald", SecondName: "Duck", Age: 112}}
	actual := 0
	expected := 271
	slice_fn.ForEach(input, func(val Test) { actual += val.Age })

	assert.Equal(t, expected, actual)
}
