package main

type (
	ToolService struct {
		repo *ToolRepository
	}
)

func NewToolService(r *ToolRepository) *ToolService {
	return &ToolService{
		repo: r,
	}
}
