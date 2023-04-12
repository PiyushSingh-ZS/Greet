package name

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Greet(name string) (string, error) {
	path := "http://localhost:8081/hello?name=" + name
	res, _ := http.Get(path)
	var b interface{}
	body, _ := io.ReadAll(res.Body)
	err1 := json.Unmarshal(body, &b)
	return fmt.Sprint(b), err1
}
