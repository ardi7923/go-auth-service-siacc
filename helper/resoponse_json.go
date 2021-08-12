package helper

import "encoding/json"

/*
fungsi ini untuk merubag response default dari struct model ke struct yang kita inginkan dengan paramter berikut :
	1. data interface ( struct hasil insert dari gorm )
	2. response struct ( struct reponse )
*/
func SaveJsonResponse(data interface{}, response_sturct interface{}) interface{} {
	struct_to_json, _ := json.Marshal(data)
	json.Unmarshal(struct_to_json, response_sturct)
	return response_sturct
}
