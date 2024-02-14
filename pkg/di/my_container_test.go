package di_test

import (
	"testing"

	"github.com/nitoba/poll-voting/pkg/di"
	"github.com/stretchr/testify/assert"
)

func TestMyContainer(t *testing.T) {
	t.Run("should create a container", func(t *testing.T) {
		di.CreateContainer()
		assert.NotNil(t, di.GetMyContainer())
	})

	t.Run("should add providers to a container", func(t *testing.T) {
		di.CreateContainer()

		di.GetMyContainer().RegisterModuleProviders([]di.Dep{
			{
				Name: "foo",
				Provide: func(ctn di.CTN) (interface{}, error) {
					return "foo", nil
				},
			},
			{
				Name: "bar",
				Provide: func(ctn di.CTN) (interface{}, error) {
					return "bar", nil
				},
			},
		})

		assert.Equal(t, "foo", di.GetMyContainer().Get("foo"))
		assert.Equal(t, "bar", di.GetMyContainer().Get("bar"))
	})

	t.Run("should add providers to a container with retrieval by name", func(t *testing.T) {
		type CustomStruct struct{}

		type CustomStruct2 struct {
			*CustomStruct
		}

		di.CreateContainer()

		di.GetMyContainer().RegisterModuleProviders([]di.Dep{
			{
				Name: "foo",
				Provide: func(ctn di.CTN) (interface{}, error) {
					return "foo", nil
				},
			},
			{
				Name: "bar",
				Provide: func(ctn di.CTN) (interface{}, error) {
					return "bar", nil
				},
			},
			{
				Name: "custom",
				Provide: func(ctn di.CTN) (interface{}, error) {
					return &CustomStruct{}, nil
				},
			},
			{
				Name: "custom2",
				Provide: func(ctn di.CTN) (interface{}, error) {
					return &CustomStruct2{
						CustomStruct: ctn.Get("custom").(*CustomStruct),
					}, nil
				},
			},
		})

		assert.Equal(t, "foo", di.GetMyContainer().Get("foo"))
		assert.Equal(t, "bar", di.GetMyContainer().Get("bar"))
		assert.IsType(t, &CustomStruct{}, di.GetMyContainer().Get("custom"))
		assert.IsType(t, &CustomStruct2{}, di.GetMyContainer().Get("custom2"))
		assert.NotNil(t, di.GetMyContainer().Get("custom2").(*CustomStruct2).CustomStruct)
	})
}
