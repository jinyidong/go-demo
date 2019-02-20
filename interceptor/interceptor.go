package interceptor

type IModel interface {
	Query(id int) int
}

type Account struct {
	Id int
	Name string
}

func (a *Account) Query(id int) int {
	return a.Id
}

var New= func(id int,name string) IModel{
	return &Account{id,name}
}

type Proxy struct {
	IModel
	TraceId string
}

func (p *Proxy) Query(id int) int {
	return p.IModel.Query(id)
}




