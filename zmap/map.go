package zmap

func MergeMap(mObj ...map[interface{}]interface{}) map[interface{}]interface{} {
	newObj := map[interface{}]interface{}{}
	for _, m := range mObj {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}
