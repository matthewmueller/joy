package paymentaddress

import "github.com/matthewmueller/golly/js"

// PaymentAddress struct
// js:"PaymentAddress,omit"
type PaymentAddress struct {
}

// ToJSON
func (*PaymentAddress) ToJSON() (i interface{}) {
	js.Rewrite("$<.ToJSON()")
	return i
}

// AddressLine
func (*PaymentAddress) AddressLine() (addressLine []string) {
	js.Rewrite("$<.AddressLine")
	return addressLine
}

// City
func (*PaymentAddress) City() (city string) {
	js.Rewrite("$<.City")
	return city
}

// Country
func (*PaymentAddress) Country() (country string) {
	js.Rewrite("$<.Country")
	return country
}

// DependentLocality
func (*PaymentAddress) DependentLocality() (dependentLocality string) {
	js.Rewrite("$<.DependentLocality")
	return dependentLocality
}

// LanguageCode
func (*PaymentAddress) LanguageCode() (languageCode string) {
	js.Rewrite("$<.LanguageCode")
	return languageCode
}

// Organization
func (*PaymentAddress) Organization() (organization string) {
	js.Rewrite("$<.Organization")
	return organization
}

// Phone
func (*PaymentAddress) Phone() (phone string) {
	js.Rewrite("$<.Phone")
	return phone
}

// PostalCode
func (*PaymentAddress) PostalCode() (postalCode string) {
	js.Rewrite("$<.PostalCode")
	return postalCode
}

// Recipient
func (*PaymentAddress) Recipient() (recipient string) {
	js.Rewrite("$<.Recipient")
	return recipient
}

// Region
func (*PaymentAddress) Region() (region string) {
	js.Rewrite("$<.Region")
	return region
}

// SortingCode
func (*PaymentAddress) SortingCode() (sortingCode string) {
	js.Rewrite("$<.SortingCode")
	return sortingCode
}
