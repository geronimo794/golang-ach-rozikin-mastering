package web

type ResponseStandard struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    any             `json:"data"`
	Errors  []ErrorResponse `json:"errors"`
	// Meta       any `json:"meta"`
	// JsonApi    any `json:"jsonapi"`
	// Links      any `json:"links"`
	// Included   any `json:"included"`
	// Self       any `json:"self"`
	// Related    any `json:"related"`
	// Pagination any `json:"pagination"`
}

type ErrorResponse struct {
	Title    string   `json:"title"`
	Messages []string `json:"messages"`
}
