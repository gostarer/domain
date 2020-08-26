package infra

var apiInitializerRegister *InitailizerRegister = new(InitailizerRegister)

//注册web对象  API初始化对象
func RegisterApi(ai Initailizer) {
	apiInitializerRegister.Register(ai)
}

//获取注册web api初始化对象
func GetApiInitializers() []Initailizer {
	return apiInitializerRegister.Initailizers
}

type WebApiStarer struct {
	BaseGoStarer
}

func (w *WebApiStarer) Setup(ctx GoStarerContext) {
	for _, v := range GetApiInitializers() {
		v.Init()
	}
}
