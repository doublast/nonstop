package talk

type Group struct {
  ID       int32
  Name     string
  QRCode   string
  OwnerID  int32
  OneToOne bool
  Messages []*Message
}

type GroupOption struct {
  GroupID int32
  UserID  int32
  Pinned  bool
  Liked   bool
  Muted   bool
}

func (p *Group) TableName() string {
  return "groups"
}

func (p *Group) Paths() []string {
  return []string{
    "/groups",
    "/groups/{id}",
  }
}
