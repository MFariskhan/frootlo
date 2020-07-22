package api_responses

const (
	InvalidProductIdErr = "Invalid product id. Please make sure the provided id should be an int value."
	InvalidDatabaseErr  = "Failed to get database connection."
	ProductNotExistErr  = "The requested product does not exist."
	ServerErr           = "Internal Server Error."
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CustomError struct {
	Status      int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
	Error       string `json:"error"`
}
