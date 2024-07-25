package protocol

import "context"

// ServiceInfo is meta info of a service
type ServiceInfo struct {
	InterfaceName string
	ServiceType   interface{}
	Methods       []MethodInfo
	Meta          map[string]interface{}
}

type MethodInfo struct {
	Name           string
	Type           string
	ReqInitFunc    func() interface{}
	StreamInitFunc func(baseStream interface{}) interface{}
	MethodFunc     func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error)
	Meta           map[string]interface{}
}
