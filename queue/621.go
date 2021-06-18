func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	charlist := make([]int, 26)
	for k := range tasks {
		//统计相同的任务数
		charlist[tasks[k]-'A']++
	}
	maxExec, maxCount := 0, 0
	for i := range charlist {
		if charlist[i] > maxExec {
			//选出最多次的任务
			maxExec, maxCount = charlist[i], 1

		} else if maxExec == charlist[i] {
			//最后的任务数+1，因为charlist中统计的最后一轮多了一个要处理的任务
			maxCount++
		}
	}
	time := (maxExec-1)*(n+1) + maxCount
	if time < len(tasks) {
		return len(tasks)
	}
	return time

}