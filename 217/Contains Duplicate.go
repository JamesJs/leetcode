package contains_duplicate

//Tempo: O(n)
//Espaço: O(n)

//Implementação 1

//A ideia foi percorrer o array e guardar
//cada valor do array como chave de uma hash
//auxiliar.
//Essa forma usa mais memória mas é mais rápida

//Implementação 2

//A ideia é ordenar o vetor antes de percorre-lo
//e comparar os valores adjacentes para encontrar duplicados.
//Essa forma usa menos memória mas é mais lenta pois depende
//de um algoritmo ordenação.

func ContainsDuplicate(nums []int) bool {
	aux := make(map[int]int)
	size := len(nums)
	for i := 0; i < size; i++ {
		_, hasKeyValue := aux[nums[i]]
		if hasKeyValue {
			return true
		}
		aux[nums[i]] = 1
	}
	return false
}
