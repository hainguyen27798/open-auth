package dto

type SearchDTO struct {
	Search string `form:"search"`
	By     string `form:"by"`
	Take   int    `form:"take"`
	Page   int    `form:"page"`
}
