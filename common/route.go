package common

import "reflect"

const (
	RouteMethod = "Routes"
)

func RegisterRoutes(controllers ...any) {
	for _, controller := range controllers {
		refType := reflect.TypeOf(controller)
		_, ok := refType.MethodByName(RouteMethod)
		if ok {
			reflect.ValueOf(controller).MethodByName(RouteMethod).Call([]reflect.Value{})
		}
	}
}
