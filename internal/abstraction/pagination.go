package abstraction

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PaginationInfo struct {
	Pagination
	Count     int64 `json:"count"`
	TotalPage int64 `json:"total_page"`
}
