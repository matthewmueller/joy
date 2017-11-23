package msassertion

// js:"MSAssertion,omit"
type MSAssertion interface {

	// ID
	ID() (id string)

	// Type
	Type() (kind *mscredentialtype.MSCredentialType)
}
