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
	"sort"
)

type Node struct {
	parent *Node
	info Pool
	children []Node
}

func newNode(p Pool)(n Node){
	n.info = p
	fmt.Println("ADDING NEW NODE : " + n.info.toString())
	return n
}

func (n *Node) addChild(child Node){
	child.parent = n
	n.children = append(n.children, child)
}


type Tree struct {
	root Node
	route []Pool
}

func (t *Tree) addRoot(p Pool){
	t.root = newNode(p)
	t.root.parent = nil
}

func (t *Tree) addNode(u Node, p Pool){
	child := newNode(p)
	u.addChild(child)
}

func (t *Tree) addEdge(root Node, closestPool Pool, newPool Pool){
	if root != nil {
		if root.info == closestPool {
			fmt.Println("EDGE CREATED from: " + root.info.toString() + " to " + newPool.toString)
			t.addNode(root, newPool)
		}
		for i := range root.children {
			t.addEdge(root.children[i], closestPool, newPool)
		}
	}
}




type Edge struct {
	poolA Pool
	poolB Pool
	distance float64
}

type coordinates[2]float64

//For parsing the JSON file
type geometry struct {
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

func (p Pool) toString() string {
	return p.Properties.NAME + " [" +  FloatToString(p.Geometry.Coordinates[0]) + ", " + FloatToString(p.Geometry.Coordinates[1]) + "]"
}

func findRoute (filename string, num int) (route []Edge){

	//Read and convert the JSON file
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var p []Pool
	json.Unmarshal(raw, &p)
	
	//Sort the pools from west to east
	sort.Slice(p, func(i, j int) bool{
		return p[i].Geometry.Coordinates[1] < p[j].Geometry.Coordinates[1]
	})
	fmt.Println("POOLS SORTED FROM WEST TO EAST")
	for i:= range p{
		fmt.Println(p[i].toString())
	}	
	
	//Store the most Western pool as the root node in tree
	var t Tree
	t.addRoot(p[0])

	//Add each pool from west to east into the tree, where edges are made between the closest pools
	closestPool := p[0]
	closestDistance := 9999.9
	mostRecentPool := 0 //Index of the most recently added pool to the tree
	var currentDistance float64
	
	//Find the closest pool
	for i := range p {
		for j := 0; j <= mostRecentPool; j++){
			currentDistance = euclidDistance(p[i].Geometry.Coordinates[1],
											 p[i].Geometry.Coordinates[0],
											 p[j].Geometry.Coordinates[1],
											 p[j].Geometry.Coordinates[0])
			
			if currentDistance < closestDistance
			{
				closestDistance = currentDistance
				closestPool = p[j]
			}
			//Create an edge between the closest pool and the new pool
			t.addEdge(t.root, closestPool, p[i])
			closestDistance = 9999.9
			mostRecentPool++
		}
		edges := t.preOrder(t.root)
		//Create an edge between the closest pool and the new pool
		
											 
											 
	//Connect the closest pool as a child of the root
	/*t.root.addChild(newNode(p[1]))
	fmt.Println("THE CHILD IS : " + t.root.children[0].info.toString())*/
	
	//For each pool from West to East, connect the node for the pool
	//with an edge as the child of the closest node in the tree
	
	num += 1
	return route
}

/*func saveRoute(route []Edge, filename string)(bool){

}*/

func FloatToString(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func main () {
	findRoute("wading-pools-min.json", 1)
	

}

func euclidDistance(lat1, lon1, lat2, lon2 float64)(dRadians float64){
	rLat1 := lat1 * math.Pi / 180
	rLat2 := lat2 * math.Pi / 180
	rLon1 := lon1 * math.Pi / 180
	rLon2 := lon2 * math.Pi / 180
	
	return 6371.0 * 2*math.Asin(math.Sqrt(math.Pow((math.Sin((rLat1 - rLat2)/2)), 2) + (math.Cos(rLat1) * math.Cos(rLat2) * math.Pow(math.Sin((rLon1 - rLon2)/2),2))))
	

}
