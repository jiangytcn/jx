// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	clients "github.com/jenkins-x/jx/pkg/jx/cmd/clients"
)

func AnyClientsFactory() clients.Factory {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(clients.Factory))(nil)).Elem()))
	var nullValue clients.Factory
	return nullValue
}

func EqClientsFactory(value clients.Factory) clients.Factory {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue clients.Factory
	return nullValue
}