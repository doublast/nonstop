package story

type Preference struct {}

func (p *Preference) TableName() string {
  return "preferences"
}

func (p *Preference) Paths() []string {
  return []string{"/user/preference"}
}
