package views

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type Todo struct {
	Name string `bson:"name"`
	Todo string `bson:"todo"`
}
