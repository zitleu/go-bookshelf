package domain

type Pagination struct {
	Page       int
	Limit      int
	Total      int
	TotalPages int
}
