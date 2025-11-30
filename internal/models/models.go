package models

type LinksResponse struct {
	Links []string `json:"links"`
}

type LinksRequest struct {
	Links    map[string]string `json:"links"`
	LinkNums int               `json:"link_num"`
}
