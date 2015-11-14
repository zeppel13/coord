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



type Coord struct {
	Lat, Long float64
}

func (coord Coord) getPos() (float64, float64) {
	return coord.Lat, coord.Long
}

/***************** Track *****************/
type Track struct {
	coordNumber int
	distance    float64
	coordSlice  []float64
}

func (track Track) getCoordSlice() []float64 {
	return track.coordSlice[:]
}

func (track Track) getDistance() float64 {
	return track.distance
}

func (track *Track) appendCoord(plat, plong float64) {

	track.coordSlice = append(track.coordSlice, plat, plong)
	track.coordNumber++
	track.calcTrackDistance()

}

func (track *Track) insertCoord(i int, plat, plong float64) {
	track.coordSlice = append(track.coordSlice[i:], plat, plong) + track.coordSlice[:i]
}
func (track *Track) resetCoord() {
	track.coordSlice = track.coordSlice[:0]
	track.coordNumber = 0
}

func (track *Track) calcTrackDistance() {
	track.distance = 0.0
	for i := 0; i+3 < len(track.coordSlice); i++ {
		alat := track.coordSlice[i]
		along := track.coordSlice[i+1]
		blat := track.coordSlice[i+2]
		blong := track.coordSlice[i+3]
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
				x, y float64
                index int
			)
			fmt.Printf("Index :")
			fmt.Scanf("%v", &input)
			index, _ = strconv.Atoi(input) // What to do here? -> find parameters

			/*** Do Something here ****/
			fmt.Printf("Latitude :")
			fmt.Scanf("%v", &input)
			x, _ = strconv.ParseFloat(input, 64)
			/*** Do Something here ****/

			fmt.Printf("Longitude :")
			fmt.Scanf("%v", &input)
			y, _ = strconv.ParseFloat(input, 64)
			track.insertCoord(index, x, y)

			/*** Do Something here ****/

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
