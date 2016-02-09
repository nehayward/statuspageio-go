package statuspageio

type ComponentListResponse struct {
	CreatedAt   string      `json:"created_at"`
	Description interface{} `json:"description"`
	GroupID     interface{} `json:"group_id"`
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	PageID      string      `json:"page_id"`
	Position    int         `json:"position"`
	Status      string      `json:"status"`
	UpdatedAt   string      `json:"updated_at"`
}

type ComponentUpdateResponse struct {
	CreatedAt   string      `json:"created_at"`
	Description interface{} `json:"description"`
	GroupID     interface{} `json:"group_id"`
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	PageID      string      `json:"page_id"`
	Position    int         `json:"position"`
	Status      string      `json:"status"`
	UpdatedAt   string      `json:"updated_at"`
}
