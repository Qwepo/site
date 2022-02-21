package db

type dbUsers interface {
	UserSave(u *User) error
	GetUserByID(id int64, ctx Context) (*User, error)
	GetUserByPhone(phone string, ctx Context) (*User, error)
}

type User struct {
	ID       int64  `reindex:"id,,pk"            json:"id"`
	Username string `reindex:"username"          json:"username"`
	Phone    string `reindex:"phone"             json:"phone"`
	Code     string `reindex:"code"              json:"code"`
}

func (c *client) UserSave(u *User) error {
	if u.ID == 0 {
		return c.Create(nsUser, u)
	}
	return c.Update(nsUser, u)
}

func (c *client) GetUserByID(id int64, ctx Context) (*User, error) {
	q := c.Query(nsUser).WhereInt64("id", EQ, id)
	return c.userGet(ctx, q)
}
func (c *client) GetUserByPhone(phone string, ctx Context) (*User, error) {
	q := c.Query(nsUser).WhereString("phone", EQ, phone)
	return c.userGet(ctx, q)
}


//=========================================================================================================================================
func (c *client) userGet(ctx Context, q Query) (*User, error) {
	item, err := c.Get(ctx, q)
	if item == nil {
		return nil, err
	}
	return item.(*User), nil

}
