package story

type Follow struct {
  ID             int32
  FollowerID     int32
  FollowableID   int32
  FollowerType   string
  FollowableType string
}

func (p *Follow) TableName() string {
  return "follows"
}

func (p *Follow) Paths() []string {
  return []string{
    "/users/{id}/follow",
    "/users/{id}/unfollow",
  }
}
