package conf

type File struct {
	Name string
	Type string
}

const (
	TypeDomain    = "domain"
	TypeClassical = "classical"
)

func (f File) IsDomain() bool {
	return f.Type == TypeDomain
}

func (f File) IsClassical() bool {
	return f.Type == TypeClassical
}
