package arraygen

func GenerateRange(min,max int) []int {
	generatedArray := make([]int, max - min + 1)
	for i,_ := range generatedArray {
		generatedArray[i] = min+i
	}
	return generatedArray
}

func GenerateProtoArray(min,max int) map[int]bool {
	generatedArray := make([]int, max - min + 1)
	generatedMap := make(map[int]bool)
	for i,_ := range generatedArray {
		generatedMap[min+i] = true
	}
	return generatedMap
}