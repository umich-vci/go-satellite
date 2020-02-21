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
