package paymentaddress

import "github.com/matthewmueller/golly/js"

// PaymentAddress struct
// js:"PaymentAddress,omit"
type PaymentAddress struct {
}

// ToJSON fn
func (*PaymentAddress) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

// AddressLine prop
func (*PaymentAddress) AddressLine() (addressLine []string) {
	js.Rewrite("$<.addressLine")
	return addressLine
}

// City prop
func (*PaymentAddress) City() (city string) {
	js.Rewrite("$<.city")
	return city
}

// Country prop
func (*PaymentAddress) Country() (country string) {
	js.Rewrite("$<.country")
	return country
}

// DependentLocality prop
func (*PaymentAddress) DependentLocality() (dependentLocality string) {
	js.Rewrite("$<.dependentLocality")
	return dependentLocality
}

// LanguageCode prop
func (*PaymentAddress) LanguageCode() (languageCode string) {
	js.Rewrite("$<.languageCode")
	return languageCode
}

// Organization prop
func (*PaymentAddress) Organization() (organization string) {
	js.Rewrite("$<.organization")
	return organization
}

// Phone prop
func (*PaymentAddress) Phone() (phone string) {
	js.Rewrite("$<.phone")
	return phone
}

// PostalCode prop
func (*PaymentAddress) PostalCode() (postalCode string) {
	js.Rewrite("$<.postalCode")
	return postalCode
}

// Recipient prop
func (*PaymentAddress) Recipient() (recipient string) {
	js.Rewrite("$<.recipient")
	return recipient
}

// Region prop
func (*PaymentAddress) Region() (region string) {
	js.Rewrite("$<.region")
	return region
}

// SortingCode prop
func (*PaymentAddress) SortingCode() (sortingCode string) {
	js.Rewrite("$<.sortingCode")
	return sortingCode
}
