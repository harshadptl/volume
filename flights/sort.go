package flights

import (
	"github.com/harshadptl/volume_assignment/types"
	"github.com/harshadptl/volume_assignment/utils"
)

// FlightSort sorts a list of flights into a continuous Flight path.
// The flights are iterated and each flight is placed connecting its source,
// destination or both in collection of LinkedLists. Two lists are joined if
// possible.
// Finally, we expect a unique linked list which has a valid flight path, if
// such a unique list is not found then an error is returned.
func FlightSort(fd *types.FlightsData) (*utils.LinkedList, error) {

	// make collection to hold the linked lists
	collections := make(map[int]*utils.LinkedList)
	collectionsCount := 0

	// Iterate over each flight in FlightsData
	for _, fl := range *fd {
		node := utils.NewNode(fl.Data())

		// try to find Previous and Next Flight matches with
		// first and last flights in each List
		prevFlightList := -1
		nextFlightList := -1
		for i, list := range collections {
			if list == nil {
				continue
			}

			// flag all matches
			if node.Data["src"] == list.Tail.Data["dest"] {
				prevFlightList = i
			} else if node.Data["dest"] == list.Head.Data["src"] {
				nextFlightList = i
			}
		}

		// if both a next flight and previous flight match is found
		// merge two lists with joining the current flight in middle
		// and set the second list as nil
		if prevFlightList != -1 && nextFlightList != -1 {
			collections[prevFlightList].InsertEnd(node)
			joinLinkedLists(
				collections[prevFlightList],
				collections[nextFlightList],
			)
			collections[nextFlightList] = nil
		} else if prevFlightList != -1 {
			// if only previous flight is found then add current flight
			// to the end of the list
			collections[prevFlightList].InsertEnd(node)
		} else if nextFlightList != -1 {
			// if only next flight is found then add current flight to
			// the beginning of the list
			collections[nextFlightList].InsertStart(node)
		} else {
			// else create a new list with single flight in the collection
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

	// return the finalList
	return finalList, nil
}

// joinLinkedLists appends the second list to the first list
func joinLinkedLists(first, second *utils.LinkedList) {
	first.Tail.Next = second.Head

	first.Tail = second.Tail
}