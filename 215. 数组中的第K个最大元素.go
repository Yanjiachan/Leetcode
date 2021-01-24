func findKthLargestHeap(nums []int, k int) int {
	heap := make([]int, len(nums))
	idx := 0
	for _, val := range nums {
		heap[idx] = val
		i := idx
		for (i-1)/2 >= 0 && heap[(i-1)/2] < heap[i] {
			heap[(i-1)/2], heap[i] = heap[i], heap[(i-1)/2]
			i = (i - 1) / 2
		}
		idx++
	}
	for i := 1; i < k; i++ {
		//交换
		heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
		heap = heap[:len(heap)-1]
		//堆化
		cur := 0
		for {
			left, right, idx := cur*2+1, cur*2+2, 0
			if left < len(heap) && right < len(heap) {
				if heap[left] < heap[right] {
					idx = right
				} else {
					idx = left
				}
			} else if left < len(heap) {
				idx = left
			} else if right < len(heap) {
				idx = right
			} else {
				break
			}
			heap[cur], heap[idx] = heap[idx], heap[cur]
			cur = idx
		}
	}
	return heap[0]
}
