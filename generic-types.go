package gosatellite

type genericReference struct {
	Description *string `json:"description"`
	ID          *int    `json:"id"`
	Name        *string `json:"name"`
	Title       *string `json:"title"`
}

type genericIDReference struct {
	ID *int `json:"id"`
}

type sort struct {
	By    *string `json:"by"`
	Order *string `json:"order"`
}

type searchResults struct {
	Page     *int    `json:"page"`
	PerPage  *int    `json:"per_page"`
	Search   *string `json:"search"`
	Sort     *sort   `json:"sort"`
	Subtotal *int    `json:"subtotal"`
	Total    *int    `json:"total"`
}
