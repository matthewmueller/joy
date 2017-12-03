package paymentaddress

import "github.com/matthewmueller/joy/macro"

// PaymentAddress struct
// js:"PaymentAddress,omit"
type PaymentAddress struct {
}

// ToJSON fn
// js:"toJSON"
func (*PaymentAddress) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

// AddressLine prop
// js:"addressLine"
func (*PaymentAddress) AddressLine() (addressLine []string) {
	macro.Rewrite("$_.addressLine")
	return addressLine
}

// City prop
// js:"city"
func (*PaymentAddress) City() (city string) {
	macro.Rewrite("$_.city")
	return city
}

// Country prop
// js:"country"
func (*PaymentAddress) Country() (country string) {
	macro.Rewrite("$_.country")
	return country
}

// DependentLocality prop
// js:"dependentLocality"
func (*PaymentAddress) DependentLocality() (dependentLocality string) {
	macro.Rewrite("$_.dependentLocality")
	return dependentLocality
}

// LanguageCode prop
// js:"languageCode"
func (*PaymentAddress) LanguageCode() (languageCode string) {
	macro.Rewrite("$_.languageCode")
	return languageCode
}

// Organization prop
// js:"organization"
func (*PaymentAddress) Organization() (organization string) {
	macro.Rewrite("$_.organization")
	return organization
}

// Phone prop
// js:"phone"
func (*PaymentAddress) Phone() (phone string) {
	macro.Rewrite("$_.phone")
	return phone
}

// PostalCode prop
// js:"postalCode"
func (*PaymentAddress) PostalCode() (postalCode string) {
	macro.Rewrite("$_.postalCode")
	return postalCode
}

// Recipient prop
// js:"recipient"
func (*PaymentAddress) Recipient() (recipient string) {
	macro.Rewrite("$_.recipient")
	return recipient
}

// Region prop
// js:"region"
func (*PaymentAddress) Region() (region string) {
	macro.Rewrite("$_.region")
	return region
}

// SortingCode prop
// js:"sortingCode"
func (*PaymentAddress) SortingCode() (sortingCode string) {
	macro.Rewrite("$_.sortingCode")
	return sortingCode
}
