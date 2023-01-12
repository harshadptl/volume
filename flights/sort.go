package flights

import (
	"github.com/harshadptl/volume_assignment/types"
	"github.com/harshadptl/volume_assignment/utils"
)

// FlightSort sorts a list of flights into a continuous Flight path.
// The flights are iterated and each flight is placed connecting its source,
// destination or both in collection of LinkedLists. Two lists are joined if
// possible.
// Finally we expect a unique linked list which has a valid flight path, if
// such a unique list is not found then an error is returned.
func FlightSort(fd *types.FlightsData) (*utils.LinkedList, error) {

	collections := make(map[int]*utils.LinkedList)
	collectionsCount := 0
	for _, fl := range *fd {
		node := utils.NewNode(fl.Data())

		prevFlightList := -1
		nextFlightList := -1
		for i, list := range collections {
			if list == nil {
				continue
			}
			if node.Data["src"] == list.Tail.Data["dest"] {
				prevFlightList = i
			} else if node.Data["dest"] == list.Head.Data["src"] {
				nextFlightList = i
			}
		}

		if prevFlightList != -1 && nextFlightList != -1 {
			collections[prevFlightList].InsertEnd(node)
			joinLinkedLists(
				collections[prevFlightList],
				collections[nextFlightList],
			)
			collections[nextFlightList] = nil
		} else if prevFlightList != -1 {
			collections[prevFlightList].InsertEnd(node)
		} else if nextFlightList != -1 {
			collections[nextFlightList].InsertStart(node)
		} else {
			collections[collectionsCount] = utils.NewLinkedList(node)
			collectionsCount += 1
		}
	}

	var finalList *utils.LinkedList
	// if more than one non nil list in collection return error, otherwise collect the final list
	for _, list := range collections {
		if list != nil && finalList != nil {
			return nil, types.ErrInvalidFlightList
		} else if list != nil {
			finalList = list
		}
	}

	return finalList, nil
}

// joinLinkedLists appends the second list to the first list
func joinLinkedLists(first, second *utils.LinkedList) {
	first.Tail.Next = second.Head

	first.Tail = second.Tail
}