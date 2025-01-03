package gasStation

import (
	utils "leetcode/utils"
)

func CanCompleteCircuit(gas []int, cost []int) int {
	i := 0
	localSum := 0
	position := 0
	var aux int
	globalSum := 0
	for i < len(gas) {

		aux = localSum + (gas[i] - cost[i])
		globalSum = globalSum + (gas[i] - cost[i])
		utils.Log(aux, globalSum, localSum)
		if localSum > 0 && aux < 0 {
			localSum = 0
		}
		if localSum <= 0 && aux >= 0 {

			position = i
		}
		if aux >= 0 {

			localSum = aux
		}
		i = i + 1
	}
	utils.Log(position)

	if globalSum >= 0 {
		return position
	}
	return -1
}
