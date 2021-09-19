package story

type Comment struct {
  ID       int32
  UserID   int32
  StatusID int32
  Content  string
}

func (p *Comment) TableName() string {
  return "comments"
}

func (p *Comment) Paths() []string {
  return []string{"/statuses/{id}/comments"}
}
