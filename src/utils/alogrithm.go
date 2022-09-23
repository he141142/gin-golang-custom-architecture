package utils

func MergeObject(SetOfObject ...map[string]interface{}) map[string]interface{} {
	objectMerged := make(map[string]interface{})
	for _, v1 := range SetOfObject {
		for k, v := range v1 {
			if _, ok := v1[k]; ok {
				objectMerged[k] = v
			}
		}
	}
	return objectMerged
}
