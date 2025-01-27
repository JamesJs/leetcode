package uniquePaths2

import (
	Utils "leetcode/utils"
)

/*
Clássico DP. A intuição foi utilizar dfs e
guardar a quantidade de possíveis caminhos
em uma matriz de cache. Cada posição da matriz de cache
guarda o valor do número de possibilidades de alcançar
o destino mxn. Ao passar pelo mesmo caminho duas vezes,
a matriz de cache terá o valor da quantidade de caminhos.
Com isso, não é necessário percorrer o mesmo caminho duas vezes.

Tempo: O(n*m)
Espaço: O(n*m) São usados uma matriz mxn e o valor da pilha de recurssão

* Uma forma mais inteligente de se resolver o problema é contar quantas formas de
se chegar na célula atual são possíveis. As formas são a soma da possíbilidade
de chegar na célula superior e na célula da esquerda pois, os únicos movimentos
possíveis são para baixo e para a direita. Assim reduz o espaço para O(1)


*/

type cacheType struct {
	pass  bool
	value int
}

func dfs(m [][]int, cache [][]cacheType, x int, y int) int {
	if x == len(m[0])-1 && y == len(m)-1 && m[y][x] != 1 {
		return 1
	}

	if x >= len(m[0]) || x < 0 {
		return 0
	}
	if y >= len(m) || y < 0 {
		return 0
	}
	if m[y][x] == 1 {
		return 0
	}
	if cache[y][x].pass {
		return cache[y][x].value
	}
	//go up
	rightPath := dfs(m, cache, x+1, y)
	//go down
	downPath := dfs(m, cache, x, y+1)
	cache[y][x].value = rightPath + downPath
	cache[y][x].pass = true
	return cache[y][x].value
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	cache := Utils.CreateArray[cacheType](len(obstacleGrid), len(obstacleGrid[0]))
	Utils.Log(cache)
	return dfs(obstacleGrid, cache, 0, 0)

}

func Evaluate() {
	Utils.Log(uniquePathsWithObstacles([][]int{{0, 0}, {0, 0}}))

}
