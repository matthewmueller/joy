package paymentaddress

import "github.com/matthewmueller/golly/js"

type PaymentAddress struct {
}

func (*PaymentAddress) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*PaymentAddress) GetAddressLine() (addressLine []string) {
	js.Rewrite("$<.addressLine")
	return addressLine
}

func (*PaymentAddress) GetCity() (city string) {
	js.Rewrite("$<.city")
	return city
}

func (*PaymentAddress) GetCountry() (country string) {
	js.Rewrite("$<.country")
	return country
}

func (*PaymentAddress) GetDependentLocality() (dependentLocality string) {
	js.Rewrite("$<.dependentLocality")
	return dependentLocality
}

func (*PaymentAddress) GetLanguageCode() (languageCode string) {
	js.Rewrite("$<.languageCode")
	return languageCode
}

func (*PaymentAddress) GetOrganization() (organization string) {
	js.Rewrite("$<.organization")
	return organization
}

func (*PaymentAddress) GetPhone() (phone string) {
	js.Rewrite("$<.phone")
	return phone
}

func (*PaymentAddress) GetPostalCode() (postalCode string) {
	js.Rewrite("$<.postalCode")
	return postalCode
}

func (*PaymentAddress) GetRecipient() (recipient string) {
	js.Rewrite("$<.recipient")
	return recipient
}

func (*PaymentAddress) GetRegion() (region string) {
	js.Rewrite("$<.region")
	return region
}

func (*PaymentAddress) GetSortingCode() (sortingCode string) {
	js.Rewrite("$<.sortingCode")
	return sortingCode
}
