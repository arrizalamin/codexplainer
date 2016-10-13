package context

type Block struct {
	line  int
	value string
}

type BlockHeap struct {
	blocks []*Block
}

func newBlock(val string, line int) Block {
	return Block{
		line:  line,
		value: val,
	}
}

func (h *BlockHeap) Push(val string, line int) {
	index := 0
	for i, block := range h.blocks {
		if line > block.line {
			index = i
			break
		}
	}
	h.pushInto(val, index)
}

func (h *BlockHeap) pushInto(val string, index int) {
	block := newBlock(val, index)
	temp := append([]*Block{}, h.blocks[:index]...)
	temp = append(temp, &block)
	h.blocks = append(temp, h.blocks[index:]...)
}

func (h *BlockHeap) Pop() string {
	last := len(h.blocks) - 1
	temp := h.blocks[last]
	h.blocks = h.blocks[:last]
	return temp.value
}
