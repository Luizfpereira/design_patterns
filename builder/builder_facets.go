package main

type Person struct {
	// address
	StreetAddress, Postcode, City string

	// job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (p *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*p}
}

func (p *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*p}
}

func (p *PersonBuilder) Build() *Person {
	return p.person
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pa *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	pa.person.StreetAddress = streetAddress
	return pa
}

func (pa *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pa.person.City = city
	return pa
}

func (pa *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	pa.person.Postcode = postcode
	return pa
}

func (pj *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pj.person.CompanyName = companyName
	return pj
}

func (pj *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pj.person.Position = position
	return pj
}

func (pj *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pj.person.AnnualIncome = annualIncome
	return pj
}
