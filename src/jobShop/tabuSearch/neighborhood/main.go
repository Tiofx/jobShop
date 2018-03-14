package neighborhood

type Neighborhood interface {
	Generator() (iterator <-chan Move)
}

type Empty struct{}
