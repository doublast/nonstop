package story

type Search struct {
  Type string
}

func (p *Search) Paths() []string {
  return []string{"/search?type={accounts,hashtags,statuses}"}
}
