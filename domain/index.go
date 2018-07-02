package domain

type File struct{
	ID         string     `json:"id"`
	FileName   string     `json:"fileName"`
	Path       string     `json:"path"`
	Type       string     `json:"type"`
	CreateDate int64      `json:"createDate"`
}