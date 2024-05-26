package registry

import (
	"fmt"
	"reflect"

	"github.com/mlange-42/beecs-ui/internal/plot"
	"github.com/mlange-42/beecs/obs"
	"github.com/mlange-42/beecs/registry"
)

var drawersRegistry = map[string]reflect.Type{}

func init() {
	RegisterObserver[obs.WorkerCohorts]()
	RegisterObserver[obs.ForagingPeriod]()
	RegisterObserver[obs.Stores]()
	RegisterObserver[obs.PatchNectar]()
	RegisterObserver[obs.PatchPollen]()
	RegisterObserver[obs.NectarVisits]()
	RegisterObserver[obs.PollenVisits]()

	RegisterObserver[obs.AgeStructure]()
	RegisterObserver[obs.ForagingStats]()

	RegisterDrawer[plot.Lines]()
	RegisterDrawer[plot.TimeSeries]()

}

func RegisterObserver[T any]() {
	registry.RegisterObserver[T]()
}

func RegisterDrawer[T any]() {
	tp := reflect.TypeOf((*T)(nil)).Elem()
	if _, ok := drawersRegistry[tp.String()]; ok {
		panic(fmt.Sprintf("there is already a drawer with type name '%s' registered", tp.String()))
	}
	drawersRegistry[tp.String()] = tp
}

func RegisterResource[T any]() {
	registry.RegisterResource[T]()
}

func RegisterSystem[T any]() {
	registry.RegisterSystem[T]()
}

func GetObserver(name string) (reflect.Type, bool) {
	return registry.GetObserver(name)
}

func GetDrawer(name string) (reflect.Type, bool) {
	t, ok := drawersRegistry[name]
	return t, ok
}

func GetResource(name string) (reflect.Type, bool) {
	return registry.GetResource(name)
}

func GetSystem(name string) (reflect.Type, bool) {
	return registry.GetSystem(name)
}
