package dto

type SearchDTO struct {
	Search string `form:"search"`
	By     string `form:"by"`
	Take   int    `form:"take"`
	Page   int    `form:"page"`
}

func (s *SearchDTO) Skip() int {
	if s.Page <= 1 {
		return 0
	}
	return (s.Page - 1) * s.Limit()
}

func (s *SearchDTO) PageSelected() int {
	if s.Page <= 1 {
		return 1
	}
	return s.Page
}

func (s *SearchDTO) Limit() int {
	if s.Take == 0 {
		return 10
	}
	return s.Take
}
