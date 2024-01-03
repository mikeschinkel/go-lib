package diff

type fixer interface {
	Fixer()
	String() string
}
