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

type shortOrg struct {
	ID    *int    `json:"id"`
	Label *string `json:"label"`
	Name  *string `json:"name"`
}

type genericCompRes struct {
	ID                   *int    `json:"id"`
	Name                 *string `json:"name"`
	Provider             *string `json:"provider"`
	ProviderFriendlyName *string `json:"provider_friendly_name"`
}

type genericTemplate struct {
	ID               *int    `json:"id"`
	Name             *string `json:"name"`
	TemplateKindID   *int    `json:"template_kind_id"`
	TemplateKindName *string `json:"template_kind_name"`
}

type genericShortRef struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type genericPtables struct {
	CreatedAt *string `json:"created_at"`
	ID        *int    `json:"id"`
	Name      *string `json:"name"`
	OsFamily  *string `json:"os_family"`
	UpdatedAt *string `json:"updated_at"`
}

type genericSmartProxy struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
	URL  *string `json:"url"`
}

type genericSubnet struct {
	Description    *string `json:"description"`
	ID             *int    `json:"id"`
	Name           *string `json:"name"`
	NetworkAddress *string `json:"network_address"`
}

type genericUser struct {
	Description *string `json:"description"`
	ID          *int    `json:"id"`
	Login       *string `json:"login"`
}

type genericRole struct {
	Name        *string `json:"name"`
	ID          *int    `json:"id"`
	Description *string `json:"description"`
	Origin      *string `json:"origin"`
}

type genericUserGroup struct {
	Name      *string `json:"name"`
	ID        *int    `json:"id"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type genericAuthSourceLDAP struct {
	ID   *int    `json:"id"`
	Type *string `json:"type"`
	Name *string `json:"name"`
}
