package utils

import "fmt"

type Heap struct {
	Heap []int
	min  bool
}

func leftChild(position int) int {
	return (2 * position) + 1
}
func rightChild(position int) int {
	return (2 * position) + 2
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func childrenPosition(arr []int, pos int) (biggerPos int, smallerPos int) {
	leftPos := leftChild(pos)
	rightPos := rightChild(pos)
	if rightPos >= len(arr) {
		return leftPos, -1
	}
	left := (arr)[leftPos]
	right := (arr)[rightPos]
	if right > left {
		return rightPos, leftPos
	}
	return leftPos, rightPos
}

// Pop Function
// Primeiro remove o valor do nó
// Coloca o último valor do array(mais à direita do menor level da árvore)
// Realiza o heapify down do nó e para todos os seus nós filhos.
func (h *Heap) Pop() (int, error) {
	if len(h.Heap) == 0 {
		return 0, fmt.Errorf("Heap is empty")
	}
	root := (h.Heap)[0]
	lastIndex := len(h.Heap) - 1
	(h.Heap)[0] = h.Heap[lastIndex]
	h.Heap = (h.Heap)[:lastIndex]
	heapifyMaxHelp(h.Heap, 0, true)
	return root, nil
}

func (h *Heap) Add(value int) {
	h.Heap = append(h.Heap, value)

	var n int = len(h.Heap)
	if h.min {
		heapifyMinHelp(h.Heap, (n/2)-1, false)
		return
	}
	heapifyMaxHelp(h.Heap, (n/2)-1, false)

}

func heapifyMaxHelp(arr []int, pos int, heapifyDown bool) {
	//Primeira condição de parada
	if pos < 0 {
		return
	}
	//segunda condição de parada
	if pos > ((len(arr) / 2) - 1) {
		return
	}
	biggerChildPos, _ := childrenPosition(arr, pos)
	if (arr)[biggerChildPos] > (arr)[pos] {
		swap(arr, biggerChildPos, pos)
		heapifyMaxHelp(arr, biggerChildPos, true)
	}
	if !heapifyDown {
		heapifyMaxHelp(arr, pos-1, false)
	}
}
func heapify(arr []int, min bool) {
	var n int = len(arr)
	if min {
		heapifyMinHelp(arr, (n/2)-1, false)
	} else {
		heapifyMaxHelp(arr, (n/2)-1, false)
	}
}
func heapifyMinHelp(arr []int, pos int, heapifyDown bool) {
	//Primeira condição de parada
	if pos < 0 {
		return
	}
	//segunda condição de parada
	if pos > ((len(arr) / 2) - 1) {
		return
	}

	biggerChildPos, smallerChildPos := childrenPosition(arr, pos)
	//Não tem filho na direita
	if smallerChildPos == -1 {
		smallerChildPos = biggerChildPos
	}
	if (arr)[smallerChildPos] < (arr)[pos] {
		swap(arr, smallerChildPos, pos)

		heapifyMinHelp(arr, smallerChildPos, true)
	}
	if !heapifyDown {
		heapifyMinHelp(arr, pos-1, false)
	}
}

func (h *Heap) CreateHeap(arr []int, min bool) {
	(h.Heap) = arr
	(h.min) = min
	heapify(h.Heap, min)
}

func (h *Heap) Len() int {
	return len(h.Heap)
}
