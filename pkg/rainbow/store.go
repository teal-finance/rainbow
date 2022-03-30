package rainbow

type Store interface {
	InsertOptions(options []Option) error
	Options(args StoreArgs) ([]Option, error)
}

type StoreArgs struct {
	Asset     []string
	Expiries  []string
	Providers []string
}
