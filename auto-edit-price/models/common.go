package cm

// Paging ...
type Paging struct {
	Offset int64
	First  int64
	Last   int64
}

// Filter ...
type Filter struct {
	Q         string
	State     string
	CreatedAt string
	UpdatedAt string
}

// Error ...
type Error struct {
	Code    int64  `json:"c"`
	Message string `json:"m"`
}

// PageInfo ...
type PageInfo struct {
	Total int64 `json:"t"`
	Prev  bool  `json:"p"`
	Next  bool  `json:"n"`
}
