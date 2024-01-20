package global

type Package struct {
	Name        string
	Version     string
	Arch        string
	Description string
	Repository  string
}
type Packages struct {
	Packages []Package
	Token    string
	Type     int32
}
