package model

type ResponseStandard struct {
	Data  any `json:"data"`
	Error []any `json:"error"`
	// Meta       any `json:"meta"`
	// JsonApi    any `json:"jsonapi"`
	// Links      any `json:"links"`
	// Included   any `json:"included"`
	// Self       any `json:"self"`
	// Related    any `json:"related"`
	// Pagination any `json:"pagination"`
}
