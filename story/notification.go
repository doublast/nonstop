package story

type Notification struct {
  ID             int32
  UserID         int32
  ActorID        int32
  NotifiableID   int32
  NotifiableType string
  Payload        string
  Unread         bool
}

func (p *Notification) Clear() {}

func (p *Notification) Dismiss() {}

func (p *Notification) TableName() string {
  return "notifications"
}

func (p *Notification) Paths() []string {
  return []string{
    "/notifications",
    "/notifications/{id}",
    "/notifications/{id}/dismiss",
  }
}
