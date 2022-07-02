package event

import "reflect"

func resolveTypeID(t any) string {
	return reflect.TypeOf(t).String()
}
