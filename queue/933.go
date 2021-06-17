type RecentCounter struct {
	Queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		Queue: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.Queue = append(this.Queue, t)
	limitBottom := t - 3000
	limitTop := t
	cnt := len(this.Queue)
	for i := 0; i < len(this.Queue); i++ {
		if this.Queue[i] < limitBottom || this.Queue[i] > limitTop {
			cnt--
		}
	}
	return cnt
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */