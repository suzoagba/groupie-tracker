package artists

import (
	"groupie-tracker/variables"
	"sort"
	"strconv"
	"strings"
)

// Links the relations API and artists API
// Creates inputs for filters
func LinkArtists() {
	var (
		membersLen       int
		membersMax       int
		membersMin       int
		creationDateThis int
		creationDateMax  int
		creationDateMin  int
		creationDates    []int
		firstAlbumMax    int
		firstAlbumMin    int
		firstAlbumDates  []int
		locations        []string
	)
	for artistIndex, artist := range variables.Artists {
		variables.Artists[artistIndex].LocationDates = make(map[string][]string) // Initialize the map
		for _, rel := range variables.Relations.Index {
			if rel.ID == artist.ID {
				for key, values := range rel.DatesLocations {
					key = keyToProperForm(key)
					variables.Artists[artistIndex].LocationDates[key] = values
					locations = locationList(locations, key)
				}
			}
		}
		// Compiles information for filters
		membersLen = len(variables.Artists[artistIndex].Members)
		creationDateThis = variables.Artists[artistIndex].CreationDate
		firstAlbumThis, _ := strconv.Atoi(variables.Artists[artistIndex].FirstAlbum[6:10])
		if artistIndex == 0 {
			membersMin = membersLen
			creationDateMin = creationDateThis
			firstAlbumMin = firstAlbumThis
		}
		if membersLen > membersMax {
			membersMax = membersLen
		}
		if membersLen < membersMin {
			membersMin = membersLen
		}
		if creationDateThis > creationDateMax {
			creationDateMax = creationDateThis
		}
		if creationDateThis < creationDateMin {
			creationDateMin = creationDateThis
		}
		creationDates = dateList(creationDates, creationDateThis)
		if firstAlbumThis > firstAlbumMax {
			firstAlbumMax = firstAlbumThis
		}
		if firstAlbumThis < firstAlbumMin {
			firstAlbumMin = firstAlbumThis
		}
		firstAlbumDates = dateList(firstAlbumDates, firstAlbumThis)
	}
	getCoordinatesFromAPI(locations)
	sort.Strings(locations)
	sort.Ints(creationDates)
	sort.Ints(firstAlbumDates)
	membersNrList := []int(toMembersNr(membersMin, membersMax))
	variables.ForFilters = variables.Filter{
		Members:         membersNrList,
		CreationDateMin: creationDateMin,
		CreationDateMax: creationDateMax,
		CreationDates:   creationDates,
		FirstAlbumMin:   firstAlbumMin,
		FirstAlbumMax:   firstAlbumMax,
		FirstAlbumDates: firstAlbumDates,
		Locations:       locations,
	}
	//add together filter info and artists list
	variables.NewPage = variables.ForPage{
		Filters: variables.ForFilters,
		Artists: variables.Artists,
		Search:  variables.Artists,
	}
}

// Creates list of number of members
func toMembersNr(a int, b int) []int {
	var c []int
	for i := a; i <= b; i++ {
		c = append(c, i)
	}
	return c
}

// Adds locations to one list without duplicates
func locationList(a []string, b string) []string {
	for _, loc := range a {
		if loc == b {
			return a
		}
	}
	return append(a, b)
}

// Adds dates to list without duplicates
func dateList(a []int, b int) []int {
	for _, loc := range a {
		if loc == b {
			return a
		}
	}
	return append(a, b)
}

// Formats the location.
func keyToProperForm(a string) string {
	a = strings.Replace(a, "_", " ", -1)
	a = strings.Replace(a, "-", ", ", -1)
	length := len(a)
	if a[length-3:] == "usa" || a[length-2:] == "uk" {
		return strings.Title(a[:length-3]) + strings.ToUpper(a[length-3:])
	} else {
		return strings.Title(a)
	}
}
