package main

type minHeap struct {
	k    int   //k容量
	heap []int //heap数组
}

func createMinHeap(k int, nums []int) *minHeap {
	heap := &minHeap{k: k, heap: []int{}}
	for _, num := range nums {
		heap.add(num)
	}
	return heap
}
func (h *minHeap) add(num int) {
	if len(h.heap) < h.k { //heap 数组长度还不够k
		h.heap = append(h.heap, num) //将num加入heap数组
		h.up(len(h.heap) - 1)        //执行上浮，上浮到合适位置
	} else if num > h.heap[0] { // 如果num比堆顶数字要大
		h.heap[0] = num //堆顶更换
		h.down(0)
	}
}
func (h *minHeap) up(i int) { //位置i上的元素，上浮到合适位置
	for i > 0 { //上浮到索引0就停止上浮，0是堆顶位置
		parent := (i - 1) >> 1 //找到父节点在heap数组中的位置
		if h.heap[parent] > h.heap[i] {
			h.heap[parent], h.heap[i] = h.heap[i], h.heap[parent] // 交换
			i = parent                                            //更新i
		} else { //不需要交换位置跳出循环
			break
		}
	}
}

func (h *minHeap) down(i int) { //下沉到合适位置
	for 2*i+1 < len(h.heap) { //左子节点的索引如果已经越界,终止下沉
		child := 2*i + 1 //左子节点在heap数组中的位置
		if child+1 < len(h.heap) && h.heap[child+1] < h.heap[child] {
			//如果右子节点存在且值更小，则用右子节点比较
		}
		if h.heap[i] > h.heap[child] {
			h.heap[child], h.heap[i] = h.heap[i], h.heap[child]
			i = child
		} else {
			break
		}
	}
}
