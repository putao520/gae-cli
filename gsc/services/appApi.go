package services

import (
	"fmt"
	"gae-cli/gsc"
	"gae-cli/types"
)

type AppItem struct {
	Category          string         `json:"category"`
	Config            string         `json:"config"`
	CreateAt          types.DateTime `json:"createat"`
	Desc              string         `json:"desc"`
	DevSecret         string         `json:"dev_secret"`
	Domain            string         `json:"domain"`
	Entry             string         `json:"entry"`
	ID                string         `json:"id"`
	K8s               int            `json:"k8s"`
	LogsServiceID     int            `json:"logs_service_id"`
	Master            string         `json:"master"`
	Name              string         `json:"name"`
	Roles             string         `json:"roles"`
	Secret            string         `json:"secret"`
	SessionType       string         `json:"session_type"`
	State             int            `json:"state"`
	UpdateAt          types.DateTime `json:"updateat"`
	UserServiceID     int            `json:"user_service_id"`
	UserID            string         `json:"userid"`
	WorkflowServiceID int            `json:"workflow_service_id"`
}

type AppService struct {
	Base *BaseService[AppItem]
}

func NewAppService() *AppService {
	return &AppService{
		Base: NewBaseService[AppItem](),
	}
}

func (s *AppService) loadImpl() *[]AppItem {
	// 分页方式获得数据
	idxStr := fmt.Sprintf("i:%d", s.Base.idx)
	maxStr := fmt.Sprintf("i:%d", s.Base.max)
	res := gsc.GetRpcMasterClient().SetEndpoint("system", "Apps").Call("pageEx", []string{idxStr, maxStr, "ja:[]"})
	if !res.GetStatus() {
		println(res.GetMessage())
		return nil
	}

	// 获得数据
	result := res.AsArray()
	// 转换数据
	d := make([]AppItem, 0)
	for _, v := range result {
		item := gsc.MapToStruct[AppItem](v.(map[string]interface{}))
		if item != nil {
			d = append(d, *item)
		}
	}
	return &d
}

func (s *AppService) Load() *AppService {
	s.Base.data = s.loadImpl()
	return s
}

func (s *AppService) Next() int {
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

func (s *AppService) Prev() int {
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

func (s *AppService) ClearPrint() *AppService {
	l := len(*s.Base.data)
	for i := 0; i < l; i++ {
		fmt.Print("\033[A\033[K")
	}
	return s
}

func (s *AppService) Select(number int) AppItem {
	if s.Base.data == nil {
		println("data is nil")
		return AppItem{}
	}
	if number < 0 || number >= len(*s.Base.data) {
		println("selected number is out of range")
		return AppItem{}
	}
	return (*s.Base.data)[number]
}

func (s *AppService) Print() {
	if s.Base.data == nil {
		return
	}
	// 从 start 开始打印
	for idx, item := range *s.Base.data {
		fmt.Printf("(%d) %s\n", idx, item.Name)
	}

}
