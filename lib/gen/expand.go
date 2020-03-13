package gen

type Expander interface {
	Ext() (from, to string)
	Expand(p []byte, v Variables) ([]byte, error)
}
