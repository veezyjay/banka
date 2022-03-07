package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"10001", "Vee", "Kigali", "110011", "2000-01-01", "1"},
		{"10002", "Tee", "Kigali", "110011", "2005-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
