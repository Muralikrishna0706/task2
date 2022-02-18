package main

func main() {
	

	a := App{}
	a.Init()


	a.Run(":8000")
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/customer/create", a.CreateCustomer).Methods("POST")
	a.Router.HandleFunc("/customer/{id}", a.GetCustomer).Methods("GET")
	a.Router.HandleFunc("/customer", a.GetAllCustomers).Methods("GET")
	a.Router.HandleFunc("/customer/{id}", a.UpdateCustomer).Methods("PATCH")
	a.Router.HandleFunc("/customer/delete/{id}", a.DeleteCustomer).Methods("DELETE")
}
