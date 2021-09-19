package talk

type FriendRequest interface {
  Approve()
  Reject()
  Send()
}

type Contact struct {
  ContactID int32
  UserID    int32
}

func (* Contact) TableName() string {
  return "contacts"
}

func (* Contact) Paths() []string {
  return []string{
    "/contacts",
    "/friend_requests",
    "/friend_requests/{id}/reject",
    "/friend_requests/{id}/approve",
  }
}
