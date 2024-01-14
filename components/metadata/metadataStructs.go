package metadata;

type AddProject struct {
  Name string `validate:"required"`
  Url string `validate:"required"`
  Description string `validate:"required"`
  Technologies []string `validate:"required"`
}
