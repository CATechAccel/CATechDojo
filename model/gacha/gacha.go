package gacha

type gachaInterface interface {
	Select() error
	Insert() error
}

type CharacterData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func New() gachaInterface {
	return &CharacterData{}
}

func (c *CharacterData) Select() error {
	panic("implement me")
}

func (c CharacterData) Insert() error {
	panic("implement me")
}
