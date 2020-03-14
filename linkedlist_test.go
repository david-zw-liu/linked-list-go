package linkedlist

import (
	"math/rand"
	"reflect"
	"testing"
)

func assetLinkedList(t *testing.T, linkedList LinkedList, want []int) {
	t.Helper()

	got := []int{}
	for head := linkedList.Head; head != nil; head = head.Next {
		got = append(got, head.Val)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expect %v, but got %v", want, got)
	}
}

func TestNew(t *testing.T) {

	t.Run("Create empty LinkedList", func(t *testing.T) {
		linkedList := NewLinkedList([]int{})
		want := []int{}

		assetLinkedList(t, linkedList, want)
	})

	t.Run("Create LinkedList with initial numbers", func(t *testing.T) {
		linkedList := NewLinkedList([]int{1, 3, 5, 7, 9})
		want := []int{1, 3, 5, 7, 9}

		assetLinkedList(t, linkedList, want)
	})
}

func TestSort(t *testing.T) {
	t.Run("Sort empty list", func(t *testing.T) {
		linkedList := NewLinkedList([]int{})
		linkedList.Sort()
		want := []int{}

		assetLinkedList(t, linkedList, want)
	})

	t.Run("Sort numbers", func(t *testing.T) {
		linkedList := NewLinkedList([]int{9, 7, 5, 3, 1})
		linkedList.Sort()
		want := []int{1, 3, 5, 7, 9}

		assetLinkedList(t, linkedList, want)
	})
}

func BenchmarkSort(b *testing.B) {
	randInts := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		randInts[i] = rand.Int()
	}

	for i := 0; i < b.N; i++ {
		linkedList := NewLinkedList(randInts)
		linkedList.Sort()
	}
}
