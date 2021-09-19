package talk

type User struct {
  ID            int32
  Name          string
  Email         string
  Avatar        string
  QRCode        string
  DisplayName   string
  StatusMessage string
}

func (p *User) LeaveGroup() {}

func (p *User) TableName() string {
  return "users"
}

func (p *User) Paths() []string {
  return []string{
    "/user",
    "/groups/{id}/leave",
    "/groups/{id}/favorite",
  }
}
