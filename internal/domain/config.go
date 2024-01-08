package domain

type ConfigRequest struct {
	ID    string
	Token string
}
type IConfigService interface {
	Get(*ConfigRequest) (interface{}, error)
}
