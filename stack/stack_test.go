package stack_test

import (
	"testing"

	"github.com/mvp-mogila/vk-test-task/stack"
	"github.com/stretchr/testify/require"
)

func TestNewStack(t *testing.T) {
	st := stack.NewStack[int]()
	require.NotEmpty(t, st)
}

func TestPush(t *testing.T) {
	testCases := []struct {
		inputData    []int
		expectedSize int
	}{
		{
			inputData:    []int{1, 2, 3},
			expectedSize: 3,
		},
		{
			inputData:    []int{1},
			expectedSize: 1,
		},
		{
			inputData:    []int{},
			expectedSize: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run("Test push", func(t *testing.T) {
			st := stack.NewStack[int]()
			for _, val := range testCase.inputData {
				st.Push(val)
			}
			size := st.Size()
			require.Equal(t, testCase.expectedSize, size)
		})
	}
}

func TestSize(t *testing.T) {
	testCases := []struct {
		inputData    []int
		expectedSize int
	}{
		{
			inputData:    []int{1, 2, 3},
			expectedSize: 3,
		},
		{
			inputData:    []int{1},
			expectedSize: 1,
		},
		{
			inputData:    []int{},
			expectedSize: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run("Test size", func(t *testing.T) {
			st := stack.NewStack[int]()
			for _, val := range testCase.inputData {
				st.Push(val)
			}
			size := st.Size()
			require.Equal(t, testCase.expectedSize, size)
		})
	}
}

func TestTop(t *testing.T) {
	testCases := []struct {
		inputData    []int
		expectedSize int
		expectedErr  error
		expectedItem int
	}{
		{
			inputData:    []int{1, 2, 3},
			expectedSize: 3,
			expectedErr:  nil,
			expectedItem: 3,
		},
		{
			inputData:    []int{1},
			expectedSize: 1,
			expectedErr:  nil,
			expectedItem: 1,
		},
		{
			inputData:    []int{},
			expectedSize: 0,
			expectedErr:  stack.ErrEmptyStack,
			expectedItem: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run("Test top", func(t *testing.T) {
			st := stack.NewStack[int]()
			for _, val := range testCase.inputData {
				st.Push(val)
			}

			item, err := st.Top()
			size := st.Size()
			require.Equal(t, testCase.expectedSize, size)
			require.Equal(t, testCase.expectedItem, item)
			require.Equal(t, err, testCase.expectedErr)
		})
	}
}

func TestPop(t *testing.T) {
	testCases := []struct {
		inputData    []int
		expectedSize int
		expectedErr  error
		expectedItem int
	}{
		{
			inputData:    []int{1, 2, 3},
			expectedSize: 2,
			expectedErr:  nil,
			expectedItem: 3,
		},
		{
			inputData:    []int{1},
			expectedSize: 0,
			expectedErr:  nil,
			expectedItem: 1,
		},
		{
			inputData:    []int{},
			expectedSize: 0,
			expectedErr:  stack.ErrEmptyStack,
			expectedItem: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run("Test top", func(t *testing.T) {
			st := stack.NewStack[int]()
			for _, val := range testCase.inputData {
				st.Push(val)
			}

			item, err := st.Pop()
			size := st.Size()
			require.Equal(t, testCase.expectedSize, size)
			require.Equal(t, testCase.expectedItem, item)
			require.Equal(t, err, testCase.expectedErr)
		})
	}
}
