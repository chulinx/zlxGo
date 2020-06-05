package array

func RemoveNil(arr []string) (newArr []string) {
	for i :=range arr {
		if len(arr[i]) >1 {
			newArr= append(newArr, arr[i])
		}
	}
	return newArr
}