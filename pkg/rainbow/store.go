package rainbow

type Store interface {
	InsertOptions(options []Option) error
	Options(args StoreArgs) ([]Option, error)
}

type StoreArgs struct {
	Assets    []string
	Expiries  []string
	Providers []string
}
