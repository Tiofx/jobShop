package neighborhood

type Neighborhood interface {
	Generate() []Move
}

type Empty struct{}
