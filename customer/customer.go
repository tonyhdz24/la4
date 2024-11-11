package customer

// Customer Interface:  Defines what a Customer has to have to be a Customer
type ICustomer interface{
		// String: Returns a string reprsentation of the Customer
	String() string
}
// Defines the Customer struct
type Customer struct{
	name string
}

// Constructor
func NewCustomer(name string) (c *Customer){
	c = new(Customer)
	c.Init(name)
	return
}

// Customer Methods
func (c *Customer) Init(name string){
	c.name = name
}

func (c *Customer) String() string{
	return c.name;
}