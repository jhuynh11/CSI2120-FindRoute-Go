/* Justin Huynh
 * 7745112
 * CSI2120
 * Comprehensive Assignment Part 3
 */
 
package main

import (
	"fmt"
	//"errors"
	"math"
	//"os"
)

type node struct {

}

type tree struct {

}

type edge struct {
	poolA pool
	poolB pool
	distance float64
}

type pool struct {
	name string
	lat float64
	lon float64
}

/*func findRoute (filename string, num int) (route []Edge){

}

func saveRoute(route []Edge, filename string)(bool){

}*/

func main () {

res := euclidDistance(45.421016, -75.690018, 45.4222, -75.6824)

fmt.Printf("%v", res)

}

func euclidDistance(lat1, lon1, lat2, lon2 float64)(dRadians float64){
	rLat1 := lat1 * math.Pi / 180
	rLat2 := lat2 * math.Pi / 180
	rLon1 := lon1 * math.Pi / 180
	rLon2 := lon2 * math.Pi / 180
	
	return 6371.0 * 2*math.Asin(math.Sqrt(math.Pow((math.Sin((rLat1 - rLat2)/2)), 2) + (math.Cos(rLat1) * math.Cos(rLat2) * math.Pow(math.Sin((rLon1 - rLon2)/2),2))))
	

}
