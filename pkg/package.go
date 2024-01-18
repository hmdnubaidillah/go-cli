package pkg

type Product struct {
	Name, Category   string
	Id, Stock, Price int
}

var Products = []Product{}

type User struct {
	Id                 int
	Username, Password string
}

var UserCredentials = map[string]User{}
