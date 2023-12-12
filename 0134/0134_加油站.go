package gasStation

//leetcode submit region begin(Prohibit modification and deletion)
// 如果x到不了y+1（但能到y），那么从x到y的任一点出发都不可能到达y+1。
// 因为从其中任一点出发的话，相当于从0开始加油，而如果从x出发到该点则不一定是从0开始加油，可能还有剩余的油。
// 既然不从0开始都到不了y+1，那么从0开始就更不可能到达y+1了
func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}

		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}

	}

	return -1

}

//leetcode submit region end(Prohibit modification and deletion)
