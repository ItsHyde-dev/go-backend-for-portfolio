package metadata;

type AddProject struct {
  Name string `json:"name" validate:"required"`
  Url string `json:"url" validate:"required"`
  Description string `json:"description" validate:"required"`
  Technologies []string `json:"technologies" validate:"required"`
}
