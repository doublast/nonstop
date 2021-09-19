package story

type Status struct {
  ID        int32
  Public    bool
  Content   string
  Parent    *Status
  Emojis    []string
  HashTags  []string
  Mentions  []string
  MediaUrls []string
}

func (* Status) TableName() string {
  return "statuses"
}

func (* Status) Paths() []string {
  return []string{
    "/statuses",
    "/statuses/home",
    "/statuses/{id}",
    "/statuses/{id}/pin",
    "/statuses/{id}/unpin",
    "/statuses/{id}/report",
    "/statuses/{id}/bookmark",
    "/statuses/{id}/unbookmark",
    "/statuses/{id}/favourite",
    "/statuses/{id}/unfavourite",
  }
}
