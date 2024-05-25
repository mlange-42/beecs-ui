package registry

import (
	"fmt"
	"reflect"

	"github.com/mlange-42/beecs-ui/game/plot"
	"github.com/mlange-42/beecs/obs"
)

var drawersRegistry = map[string]reflect.Type{}
var observerRegistry = map[string]reflect.Type{}

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
	tp := reflect.TypeOf((*T)(nil)).Elem()
	if _, ok := observerRegistry[tp.String()]; ok {
		panic(fmt.Sprintf("there is already an observer with type name '%s' registered", tp.String()))
	}
	observerRegistry[tp.String()] = tp
}

func RegisterDrawer[T any]() {
	tp := reflect.TypeOf((*T)(nil)).Elem()
	if _, ok := drawersRegistry[tp.String()]; ok {
		panic(fmt.Sprintf("there is already a drawer with type name '%s' registered", tp.String()))
	}
	drawersRegistry[tp.String()] = tp
}

func GetObserver(name string) (reflect.Type, bool) {
	t, ok := observerRegistry[name]
	return t, ok
}

func GetDrawer(name string) (reflect.Type, bool) {
	t, ok := drawersRegistry[name]
	return t, ok
}
