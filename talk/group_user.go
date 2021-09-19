package talk

type GroupUser struct {
  GroupID int32
  UserID  int32
}

func (p *GroupUser) ManyToMany() []string {
  return []string{"Group", "User"}
}

func (p *GroupUser) TableName() string {
  return "group_users"
}

func (p *GroupUser) Paths() []string {
  return []string{
    "/groups/{id}/users",
    "/groups/{id}/add_user",
    "/groups/{id}/remove_user",
  }
}
