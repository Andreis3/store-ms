package store

type Store struct {
	StoreKey    string
	CompanyName string
	Status      string
	CNPJ        string
	Domain      string
	Contact
	GroupCOD string
}

type Contact struct {
	ContactName  string
	ContactEmail string
	ContactPhone string
	ContactRamal string
}
