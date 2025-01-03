package maxCoins

//A idéia foi ordenar o array e, a partir do segundo
//maior, pegar de 2 em 2 pois, o valor que eu tenho que
//pegar da pilha é sempre o segundo maior.
//A Alice pega sempre o maior valor disponível,
//o bob pega sempre o menor valor disponível no array
//e eu sempre pego o segundo valor maior
//1,2,4,5,6,7 neste caso, na primeira iteração
//eu pego o valor 6, a alice pega o valor 7 e o bob pega o valor 1
//na segunda iteração eu pego o valor 4, a alice pega o valor 5
//e o bob pega o valor 2
//Tempo: O(nlogn)
//Espaço: 1 (usa o mesmo array para guardar o array ordenado)
import (
	"leetcode/utils"
	"sort"
)

func MaxCoins(piles []int) int {
	sort.Ints(piles)
	size := len(piles)
	i := size / 3
	myChoose := size - 2
	sum := 0
	for i > 0 {
		sum = sum + piles[myChoose]
		myChoose = myChoose - 2
		i--
	}
	utils.Log(sum)

	return 0
}
