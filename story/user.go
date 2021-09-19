package story

type User struct {
  ID             int32
  Name           string
  Email          string
  Avatar         string
  Locale         string
  DisplayName    string
  HeaderCover    string
  RefreshToken   string
  FailedAttempts int
}

func (p *User) SecurePassword() {}

func (p *User) TableName() string {
  return "users"
}

func (p *User) Paths() []string {
  return []string{
    "/user/blocks",
    "/user/statuses",
    "/user/bookmarks",
    "/user/followers",
    "/user/following",
    "/user/favourites",
    "/user/follow_requests",
  }
}
