package algorithms

import (
	"errors"
	"fmt"
)

type Set struct {
	list []string
}

func(s *Set) add(param []string) error {

	if s.list == nil || param == nil {
		return errors.New("Empty list or param.")
	}

	for i := 0; i < len(param); i++ {
		for j:=0; j < len(s.list); j++ {
			if param[i] != s.list[j] {
				s.list = append(s.list, param[j])
			}
		}
	}
	return nil
}

func(s *Set) get() []string {
	result := make([]string, 0)
	for i := 0; i < len(s.list); i++ {
		result = append(result, s.list[i])
	}
	return result
}

func NewSet() *Set {
	return &Set{}
}

func GreedyCovering() {
	// You pass an array in, and it gets converted to a set.
	statesNeeded := [...]string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}

	stations := make(map[string][]string, 0)

	stations["kone"] = [3]string{"id", "nv", "ut"}
	stations["ktwo"] = [3]string{"wa", "id", "mt"}
	stations["kthree"] = [3]string{"or", "nv", "ca"}
	stations["kfour"] = [3]string{"nv", "ut"}
	stations["kfive"] = [3]string{"ca", "az"}

	finalStations := make([]string, 0)
	covered := NewSet()

	for len(statesNeeded) > 0 {
		bestStation := ""
		statesCovered := make([]string, 0)
		for station, states := range stations {
			covered.add(statesNeeded)
			covered.add(states)
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered.get()
			}
		}
		//states_needed -= states_covered
		statesNeeded = updateStatesNeeded(statesNeeded, statesCovered)
		finalStations = append(finalStations, bestStation)
	}
	fmt.Println(finalStations)
}

func updateStatesNeeded(statesNeeded []string, stCovered []string) []string {
	updated := make([]string, 0)
	for i := 0; i < len(stCovered); i++ {
		for _, val := range statesNeeded {
			if val != stCovered[i] {
				updated = append(updated, val)
			}
		}
	}
	return updated
}