package domain

type File struct{
	ID         string     `json:"id"`
	FileName   string     `json:"fileName"`
	Path       string     `json:"path"`
	Type       string     `json:"type"`
	CreateDate int64      `json:"createDate"`
}

type Mate struct {
	Total     int `json:"total"`
}

type Result struct {
	TextStatus string      `json:"textStatus"`
	Data       interface{} `json:"data"`
	Error      error       `json:"error"`
	Mate       Mate        `json:"mate"`
}