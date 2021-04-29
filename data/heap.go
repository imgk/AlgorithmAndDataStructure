package data

type Heap struct {
	data []int
}

func (h *Heap) ShiftUp(i int) {
}

func (h *Heap) ShiftDown(i int) {
}

func (h *Heap) Add(x int) {
	h.data = append(data, x)
	h.ShiftUp(len(h.data) - 1)
}

func (h *Heap) Del(x int) {
	for i := 1; i < len(h.data); {

	}
}
