package structs

type Region struct {
	Id   uint64 `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Type string `json:"type" xml:"type"`
}