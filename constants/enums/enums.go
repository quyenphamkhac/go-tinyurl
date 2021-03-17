package enums

type Enviroment string
type DatabaseName string

const (
	Development Enviroment = "dev"
	Staging     Enviroment = "staging"
	Production  Enviroment = "production"
)
