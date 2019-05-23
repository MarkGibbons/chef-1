package chef

import "fmt"

type UserService struct {
	client *Client
}

// User represents the native Go version of the deserialized User type
type User struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	Password    string `json:"password"`
	PublicKey   string `json:"public_key"`
}

type UserResult struct {
	Uri string `json:"uri"`
}

// List lists the users in the Chef server.
//
// Chef API docs: https://docs.chef.io/api_chef_server.html#users
func (e *UserService) List() (userlist map[string]string, err error) {
	err = e.client.magicRequestDecoder("GET", "users", nil, &userlist)
	return
}

// Get gets a user from the Chef server.
//
// Chef API docs: http://docs.opscode.com/api_chef_server.html#id28
func (e *UserService) Get(name string) (user User, err error) {
	url := fmt.Sprintf("users/%s", name)
	err = e.client.magicRequestDecoder("GET", url, nil, &user)
	return
}

// Creates an User on the chef server
//
// Chef API docs: https://docs.chef.io/api_chef_server.html#users
func (e *UserService) Create(user User) (data *UserResult, err error) {
	body, err := JSONReader(user)
	if err != nil {
		return
	}

	err = e.client.magicRequestDecoder("POST", "users", body, &data)
	return
}

// Update a user on the Chef server.
//
// Chef API docs: https://docs.chef.io/api_chef_server.html#users
func (e *UserService) Update(g User) (user User, err error) {
	url := fmt.Sprintf("users/%s", g.Name)
	body, err := JSONReader(g)
	if err != nil {
		return
	}

	err = e.client.magicRequestDecoder("PUT", url, body, &user)
	return
}

// Delete removes a user on the Chef server
//
// Chef API docs: https://docs.chef.io/api_chef_server.html#users
func (e *UserService) Delete(name string) (err error) {
	err = e.client.magicRequestDecoder("DELETE", "users/"+name, nil, nil)
	return
}
