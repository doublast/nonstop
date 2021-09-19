package talk

type Message struct {
  ID       int32
  Type     string
  Text     string
  Audio    string
  Image    string
  Video    string
  Status   string
  Emojis   []string
  Sticker  string
  GroupID  int32
  FromUser *User
  Mentions []*User
}

func (p *Message) TableName() string {
  return "messages"
}

func (p *Message) Paths() []string {
  return []string{
    "/groups/{group_id}/messages",
    "/groups/{group_id}/messages/{id}",
  }
}
