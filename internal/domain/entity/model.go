package entity

type Model struct {
	Name      string
	MaxTokens int
}

func NewModel(name string, maxtokens int) *Model {
	return &Model{
		Name:      name,
		MaxTokens: maxtokens,
	}
}

func (m *Model) GetMaxTokens() int {
	return m.MaxTokens
}

func (m *Model) GetModelName() string {
	return m.Name
}
