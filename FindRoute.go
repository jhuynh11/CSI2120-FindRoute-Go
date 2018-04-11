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
	"io/ioutil"
	"encoding/json"
	"os"
	"strconv"
)

type Node struct {

}

type Tree struct {

}

type Edge struct {
	poolA Pool
	poolB Pool
	distance float64
}

type coordinates[2]float64

type geometry struct {
	//coordinates []float64
	Coordinates coordinates
}

type properties struct {
	NAME string	
	Geometry geometry	
}

type Pool struct {
	Properties properties
	Geometry geometry	
}

/*func findRoute (filename string, num int) (route []Edge){

}

func saveRoute(route []Edge, filename string)(bool){

}*/

func FloatToString(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func main () {
	raw, err := ioutil.ReadFile("wading-pools-min.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	
	var p []Pool
	json.Unmarshal(raw, &p)
	for i:= range p{
		fmt.Println(p[i].Properties.NAME + " [" +  FloatToString(p[i].Geometry.Coordinates[0]) + ", " + FloatToString(p[i].Geometry.Coordinates[1]) + "]") 
	}
	

}

func euclidDistance(lat1, lon1, lat2, lon2 float64)(dRadians float64){
	rLat1 := lat1 * math.Pi / 180
	rLat2 := lat2 * math.Pi / 180
	rLon1 := lon1 * math.Pi / 180
	rLon2 := lon2 * math.Pi / 180
	
	return 6371.0 * 2*math.Asin(math.Sqrt(math.Pow((math.Sin((rLat1 - rLat2)/2)), 2) + (math.Cos(rLat1) * math.Cos(rLat2) * math.Pow(math.Sin((rLon1 - rLon2)/2),2))))
	

}
