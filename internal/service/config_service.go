package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/PNYwise/config-service/internal/domain"
	"github.com/spf13/viper"
)

type configService struct {
	conf *viper.Viper
}

func NewConfigServie(conf *viper.Viper) domain.IConfigService {
	return &configService{
		conf: conf,
	}
}

func (c *configService) Get(request *domain.ConfigRequest) (interface{}, error) {
	token := c.conf.Get(fmt.Sprintf("%s.token", request.ID))
	if token != request.Token {
		return nil, errors.New("invalid token")
	}
	jsonData, err := json.MarshalIndent(c.conf.Get(fmt.Sprintf("%s.config", request.ID)), "", "  ")
	if err != nil {
		log.Fatalf("Error converting to JSON: %s", err)
	}

	return string(jsonData), nil
}
