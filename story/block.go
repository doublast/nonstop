package story

type Block struct {}

func (p *Block) TableName() string {
  return "blocks"
}

func (p *Block) Paths() []string {
  return []string{
    "/users/{id}/block",
    "/users/{id}/unblock",
  }
}
