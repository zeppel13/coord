/************************************
 * coord.go							*
 * by Sebastian Kind 12.09.2015     *
 ************************************/

package main

import (
	"fmt"
	"math"
	"strconv"
)

/******* Linked List *******/
type CoordList struct {
	value Coord
	next  *CoordList
}

func (list *CoordList) appendItem(item CoordList) {
	for ; list.next != nil; list = list.next {
	}
	list.next = &item
}

func (list *CoordList) length() int {
	var i int
	for ; list.next != nil; list = list.next {
		i++
	}
	return i
}

func (list *CoordList) printValue() {
	for ; list.next != nil; list = list.next {
		fmt.Println(list.value)
	}
}

func (list *CoordList) getIndexValue(index int) Coord {
	for i := 0; i < index; i++ {
		list = list.next
	}
	return list.value
}

func (list *CoordList) insertAfter(pValue Coord) {
	var newItem CoordList
	newItem.value = pValue
	newItem.next = list.next
	list.next = &newItem
}

func (list *CoordList) insertAfterIndex(index int, pValue Coord) {
	var newItem CoordList
	for i := 0; i < index; i++ {
		list = list.next
	}

	newItem.value = pValue
	newItem.next = list.next
	list.next = &newItem
}

/********* Coordinates ******/
type Coord struct {
	Lat, Long float64
}

func (coord Coord) getPos() (float64, float64) {
	return coord.Lat, coord.Long
}

func toCoord(pLat, pLong float64) Coord {
	var newCoord Coord
	newCoord.Lat = pLat
	newCoord.Long = pLong
	return newCoord
}

/***************** Track *****************/
type Track struct {
	coordNumber int
	distance    float64
	coordSlice  []float64
	list        CoordList
}

func (track Track) getCoordSlice() []float64 {
	return track.coordSlice[:]
}

func (track Track) getDistance() float64 {
	return track.distance
}

func (track *Track) appendCoord(plat, plong float64) {
	var (
		lCoord  Coord
		newItem CoordList
	)

	lCoord.Lat = plat
	lCoord.Long = plong
	newItem.value = lCoord
	track.list.appendItem(newItem)
	track.coordNumber++
	track.calcTrackDistance()

}

func (track *Track) insertCoord(index int, plat, plong float64) {
	var lCoord Coord
	lCoord.Lat = plat
	lCoord.Long = plong
	track.list.insertAfterIndex(index, lCoord)

}
func (track *Track) resetCoord() {
	track.list.next = nil
	track.coordNumber = 0
}

func (track *Track) calcTrackDistance() {
	track.distance = 0.0
	for i := 0; i+1 <= track.list.length(); i++ { // Some trouble is hidden behind this line
		alat, along := track.list.getIndexValue(i).getPos()
		blat, blong := track.list.getIndexValue(i + 1).getPos()
		track.distance += track.calcDistance(alat, along, blat, blong)
	}
}

func (track Track) calcDistance(alat, along, blat, blong float64) float64 {
	var (
		pi       float64 = 3.14159265359
		distance float64 = 0.0
		angle    float64 = 0.0

		cosDeltaLambda float64 = 0.0
		deltaLong      float64 = 0.0

		aSin = math.Sin(alat * (pi / 180))
		bSin = math.Sin(blat * (pi / 180))
		aCos = math.Cos(alat * (pi / 180))
		bCos = math.Cos(blat * (pi / 180))
	)

	if along < 0.0 || blong < 0.0 {
		if along < 0.0 {
			along *= -1
		}
		if blong < 0.0 {
			blong *= -1
		}
		deltaLong = along + blong
	} else {
		deltaLong = along - blong
	}

	cosDeltaLambda = math.Cos(deltaLong * (pi / 180))
	angle = math.Acos(aSin*bSin + aCos*bCos*cosDeltaLambda)
	distance = 2 * pi * 6371 * ((angle * 180 / pi) / 360)

	return distance
}

/**********************
 * Terminal Interface *
 **********************/
func inputloop(track Track) {
	var input = ""
	for {
		fmt.Printf("coord>")
		fmt.Scanf("%v", &input)
		if input == "q" {
			break
		}
		if input == "h" || input == "help" {
			printHelp()
		}
		if input == "a" || input == "append" {
			var x, y float64
			fmt.Printf("Latitude: ")
			fmt.Scanf("%v", &input)
			x, _ = strconv.ParseFloat(input, 64)
			fmt.Printf("Longitude: ")
			fmt.Scanf("%v", &input)
			y, _ = strconv.ParseFloat(input, 64)
			track.appendCoord(x, y)
			input = ""
		}

		if input == "i" || input == "insert" {
			var (
				x, y  float64
				index int
			)
			fmt.Printf("Index :")
			fmt.Scanf("%v", &input)
			index, _ = strconv.Atoi(input)
			fmt.Printf("Latitude :")
			fmt.Scanf("%v", &input)
			x, _ = strconv.ParseFloat(input, 64)

			fmt.Printf("Longitude :")
			fmt.Scanf("%v", &input)
			y, _ = strconv.ParseFloat(input, 64)
			track.insertCoord(index, x, y)

			track.list.insertAfterIndex(index, toCoord(x, y)) //missing function

		}
		if input == "l" || input == "list" {
			for i, v := range track.getCoordSlice() {
				if i%2 == 0 {
					fmt.Printf("%v° ", v)
				} else {
					fmt.Printf("%v°  |  ", v)
				}

			}
			fmt.Printf("\n")
			input = ""

		}
		if input == "p" || input == "print" {
			fmt.Printf("Track distance: %.2f km\t\t\tNumber of Coordinates: %d\n", track.getDistance(), track.coordNumber)
			input = ""
		}

		if input == "reset" {
			track.resetCoord()
			fmt.Println("Track was deleted.")
			input = ""
		}

		input = ""
		//Check all at the end
		track.calcTrackDistance()
	}
}

func printHelp() {
	fmt.Printf("Coord Help Message\n\nType either the word or the shortcut to the coord commandline.\n")
	fmt.Printf("\na - append\nl - list\np - print\n\nh - help\nq - quit\n")
}

/*****************
 * main function *
 *****************/

func main() {

	var (
		place   = make(map[string]Coord)
		myTrack Track
	)
	// get us some places
	place["Joseph"] = Coord{39.851666666667, -88.9441666666670}
	place["Sebastian"] = Coord{49, 10}
	place["Sydney"] = Coord{-34, 151}
	place["Wörth"] = Coord{49.051666666667, 8.2602777777778}
	inputloop(myTrack)

}
