package models

type LinksRequest struct {
	Links []string `json:"links"`
}

type LinksResponse struct {
	Links    map[string]string `json:"links"`
	LinkNums int               `json:"link_num"`
}

type ReportLinksRequest struct {
	LinksList []int `json:"links_list"`
}
