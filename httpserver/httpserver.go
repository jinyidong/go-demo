package httpserver

import "net/http"

//自定义httpserver，其实只需要get、post两种行为基本满足需求
type MethodMaps []handler
type handler map[string]HandlerMapped
type HandlerMapped struct {
	f HandlerFunc
}
type HttpServer struct {
	router MethodMaps
}
const (
	GET         = iota
	POST
)

func Default() *HttpServer {
	return &HttpServer{
		router:NewRouter(),
	}
}

func NewRouter() MethodMaps {
	return []handler{//注意该数组初始化方式，{1:XXX,2:YYY}，指定位置初始化
		GET:    make(handler),
		POST:   make(handler),
	}
}

func (m MethodMaps) GetMapping(url string) (HandlerMapped, bool) {
	if hm, ok := m[GET][url]; ok {
		return hm, true
	}
	return HandlerMapped{}, false
}
//映射路由，获取Post方法下对应的接口
func (m MethodMaps) PostMapping(url string) (HandlerMapped, bool) {
	if hm, ok := m[POST][url]; ok {
		return hm, true
	}
	return HandlerMapped{}, false
}

//增加Get方法下的接口
func (m MethodMaps) GetAdd(url string, mapped HandlerMapped) {
	if _, ok := m.GetMapping(url); ok {
		panic("duplicate url with get method")
	}
	m[GET].SetUrl(url,mapped)
}
//增加Post方法下的接口
func (m MethodMaps) PostAdd(url string, mapped HandlerMapped) {
	if _, ok := m.GetMapping(url); ok {
		panic("duplicate url with Post method")
	}
	m[POST].SetUrl(url,mapped)

}

func (h handler) SetUrl(url string, mapped HandlerMapped) {
	h[url] = mapped
}

func (httpServer *HttpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//转发给doHandler进行执行
	httpServer.doHandler(w,req)
}

func (httpServer *HttpServer) doHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		{
			if hm, ok := httpServer.router.GetMapping(req.URL.RequestURI()); ok {
				hm.f(w, req)
			}
		}
	case http.MethodPost:
		{
			if hm, ok := httpServer.router.PostMapping(req.URL.RequestURI()); ok {
				hm.f(w, req)
			}

		}
	default:{}
	}
}

func (httpServer *HttpServer) GET(url string, f HandlerFunc) {
	httpServer.router.GetAdd(url, HandlerMapped{f: f})
}
func (httpServer *HttpServer) POST(url string, f HandlerFunc) {
	httpServer.router.PostAdd(url, HandlerMapped{f: f})
}

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

type HttpAction interface {
	Get(url string,handler HandlerFunc)
	POST(url string,handler HandlerFunc)
}
