package core

type CliCore interface {
	// Add creates a record. args is in the form of `alias1 value1 alias2 value2`
	Add(args []string) error
}
