package responses

type CartResponse struct {
	Status  int                    `json:"status" bson:"status"`
	Message string                 `json:"message" bson:"message"`
	Data    map[string]interface{} `json:"data" bson:"data"`
}
