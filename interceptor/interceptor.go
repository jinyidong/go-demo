package interceptor

type IModel interface {
	New() IModel
	Query(id int) int
}

type Account struct {
	Id int
	Name string
}

func (a *Account) Query(id int) int {
	return a.Id
}

func (a *Account) New() IModel {
	return &Account{a.Id,a.Name}
}


type Proxy struct {
	IModel
	TraceId string
}

func (p *Proxy) Query(id int) int {
	return p.IModel.Query(id)
}

func (p *Proxy) New() IModel {
	return p.IModel.New()
}




