package domain

type Citations struct {
	Id       int    `json:"id"`
	Author   string `json:"author"`
	Citation string `json:"citation"`
}
