package paymentaddress

import "github.com/matthewmueller/golly/js"

// PaymentAddress struct
// js:"PaymentAddress,omit"
type PaymentAddress struct {
}

// ToJSON fn
// js:"toJSON"
func (*PaymentAddress) ToJSON() (i interface{}) {
	js.Rewrite("$_.toJSON()")
	return i
}

// AddressLine prop
// js:"addressLine"
func (*PaymentAddress) AddressLine() (addressLine []string) {
	js.Rewrite("$_.addressLine")
	return addressLine
}

// City prop
// js:"city"
func (*PaymentAddress) City() (city string) {
	js.Rewrite("$_.city")
	return city
}

// Country prop
// js:"country"
func (*PaymentAddress) Country() (country string) {
	js.Rewrite("$_.country")
	return country
}

// DependentLocality prop
// js:"dependentLocality"
func (*PaymentAddress) DependentLocality() (dependentLocality string) {
	js.Rewrite("$_.dependentLocality")
	return dependentLocality
}

// LanguageCode prop
// js:"languageCode"
func (*PaymentAddress) LanguageCode() (languageCode string) {
	js.Rewrite("$_.languageCode")
	return languageCode
}

// Organization prop
// js:"organization"
func (*PaymentAddress) Organization() (organization string) {
	js.Rewrite("$_.organization")
	return organization
}

// Phone prop
// js:"phone"
func (*PaymentAddress) Phone() (phone string) {
	js.Rewrite("$_.phone")
	return phone
}

// PostalCode prop
// js:"postalCode"
func (*PaymentAddress) PostalCode() (postalCode string) {
	js.Rewrite("$_.postalCode")
	return postalCode
}

// Recipient prop
// js:"recipient"
func (*PaymentAddress) Recipient() (recipient string) {
	js.Rewrite("$_.recipient")
	return recipient
}

// Region prop
// js:"region"
func (*PaymentAddress) Region() (region string) {
	js.Rewrite("$_.region")
	return region
}

// SortingCode prop
// js:"sortingCode"
func (*PaymentAddress) SortingCode() (sortingCode string) {
	js.Rewrite("$_.sortingCode")
	return sortingCode
}
