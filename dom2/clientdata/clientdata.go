package clientdata

type ClientData struct {
	challenge    string
	extensions   *webauthnextensions.WebAuthnExtensions
	hashAlg      interface{}
	origin       string
	rpId         string
	tokenBinding *string
}
