package array

// Array 去除 空元素

func RemoveNil(arr []string) (newArr []string) {
	for i :=range arr {
		if len(arr[i]) >1 {
			newArr= append(newArr, arr[i])
		}
	}
	return newArr
}

// array 去重

func Set(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}