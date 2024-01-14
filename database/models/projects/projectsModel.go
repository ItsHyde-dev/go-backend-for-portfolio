package projectsModel;

type Project struct {
  Id int `db:"id"`
  Title string `db:"title"`
  Description string `db:"description"`
  Url string `db:"url"`
  Technologies []string `db:"technologies"`
}
