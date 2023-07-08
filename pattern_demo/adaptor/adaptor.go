package adaptor

// Target 是适配的目标接口
type Target interface {
	Request() string
}

// ==========实际的方法 =============

type Adaptee interface {
	SpecificRequest() string
}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

type adapteeImpl struct {
}

func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// ==================================

// ==================转换器===================

type Adapter struct {
	Adaptee
}

func NewAdapter(adaptee Adaptee) Target {
	return &Adapter{
		Adaptee: adaptee,
	}
}

func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

// ===========================================
