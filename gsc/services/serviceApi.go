package services

import (
	"fmt"
	"gae-cli/gsc"
)

type ServiceItem struct {
	Category     string `json:"category"`
	Config       string `json:"config"`
	DataModel    string `json:"datamodel"`
	Debug        int    `json:"debug"`
	Description  string `json:"desc"`
	Dev          string `json:"dev"`
	Api          string `json:"api"`
	DevContainer string `json:"dev_container"`
	DevSSH       string `json:"dev_ssh"`
	DockerImage  string `json:"dockerimage"`
	ID           int    `json:"id"`
	Kind         string `json:"kind"`
	MQConfig     string `json:"mq_config"`
	Name         string `json:"name"`
	Port         int    `json:"port"`
	Protocol     string `json:"protocol"`
	SDKID        string `json:"sdk_id"`
	Transfer     string `json:"transfer"`
	UserID       string `json:"userid"`
	Version      string `json:"version"`
}

type ServiceService struct {
	Base *BaseService[ServiceItem]
}

func NewServiceService() *ServiceService {
	return &ServiceService{
		Base: NewBaseService[ServiceItem](),
	}
}

func (s *ServiceService) loadImpl() *[]ServiceItem {
	// 分页方式获得数据
	idxStr := fmt.Sprintf("i:%d", s.Base.idx)
	maxStr := fmt.Sprintf("i:%d", s.Base.max)
	res := gsc.GetRpcMasterClient().SetEndpoint("system", "Services").Call("pageEx", []string{idxStr, maxStr, "ja:[]"})
	if !res.GetStatus() {
		println(res.GetMessage())
		return nil
	}

	// 获得数据
	result := res.AsArray()
	// 转换数据
	d := make([]ServiceItem, 0)
	for _, v := range result {
		item := gsc.MapToStruct[ServiceItem](v.(map[string]interface{}))
		if item != nil {
			d = append(d, *item)
		}
	}
	return &d
}

func (s *ServiceService) Load() *ServiceService {
	s.Base.data = s.loadImpl()
	return s
}

func (s *ServiceService) Next() int {
	if s.Base.data == nil {
		return 0
	}
	s.Base.idx++
	d := s.loadImpl()
	r := len(*d)
	if d == nil || r == 0 {
		if s.Base.idx > 1 {
			s.Base.idx--
		}
	} else {
		s.Base.data = d
	}
	return r
}

func (s *ServiceService) Prev() int {
	if s.Base.data == nil {
		return 0
	}
	if s.Base.idx <= 1 {
		return len(*s.Base.data)
	}

	s.Base.idx--
	s.Base.data = s.loadImpl()
	return len(*s.Base.data)
}

func (s *ServiceService) ClearPrint() *ServiceService {
	l := len(*s.Base.data)
	for i := 0; i < l; i++ {
		fmt.Print("\033[A\033[K")
	}
	return s
}

func (s *ServiceService) Select(number int) ServiceItem {
	if s.Base.data == nil {
		println("data is nil")
		return ServiceItem{}
	}
	if number < 0 || number >= len(*s.Base.data) {
		println("selected number is out of range")
		return ServiceItem{}
	}
	return (*s.Base.data)[number]
}

func (s *ServiceService) Print() {
	if s.Base.data == nil {
		return
	}

	// 从 start 开始打印
	for idx, item := range *s.Base.data {
		fmt.Printf("(%d) %s\n", idx, item.Name)
	}

}
