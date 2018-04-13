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
	"bufio"
)

type Node struct {
	parent *Node
	info Pool
	children []Node
}

func newNode(p Pool)(n Node){
	n.info = p
	//fmt.Println("ADDED NODE : " + n.info.toString())
	return n
}

func (n *Node) addChild(child Node){
	child.parent = n
	n.children = append(n.children, child)
}


type Tree struct {
	root Node
	//route []Pool
	route Edge
}

func (t *Tree) addRoot(p Pool){
	t.root = newNode(p)
	t.root.parent = nil
}

func (t *Tree) addNode(u *Node, p Pool){
	child := newNode(p)
	u.addChild(child)
}

func (t *Tree) preOrder(root *Node) {
	if root != nil {
		fmt.Println("Added " + root.info.toString() + " to route")
		t.route.pools = append(t.route.pools, root.info)
		
		for i := range root.children {
			t.preOrder(&root.children[i])
		}
	}
}

func (t *Tree) addEdge(root *Node, closestPool Pool, newPool Pool) bool{
	if root != nil {
		if root.info == closestPool {
			fmt.Println("EDGE CREATED from: " + root.info.toString() + " to " + newPool.toString())
			t.addNode(root, newPool)
			return true
		}
		for i := 0; i < len(root.children); i++{//for i := range root.children {
			//fmt.Println("Child is : " + root.children[i].info.toString()) 
			t.addEdge(&root.children[i], closestPool, newPool)
		}
	}
	return true
}


type Edge struct {
	pools []Pool
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

func findRoute (filename string, num int) (Edge){

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
		return p[i].Geometry.Coordinates[0] < p[j].Geometry.Coordinates[0]
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
	for i :=  1; i < len(p); i++ {
		for j := 0; j <= mostRecentPool; j++{
			currentDistance = euclidDistance(p[i].Geometry.Coordinates[1],
											 p[i].Geometry.Coordinates[0],
											 p[j].Geometry.Coordinates[1],
											 p[j].Geometry.Coordinates[0])
			
			if currentDistance < closestDistance {
				closestDistance = currentDistance
				closestPool = p[j]
			}
		}
			//Create an edge between the closest pool and the new pool
			ok := t.addEdge(&t.root, closestPool, p[i])
			if ok == true {}
			closestDistance = 9999.9
			mostRecentPool++
		}
	t.preOrder(&t.root)

	num += 1
	return t.route
}

func saveRoute(route Edge, filename string)(bool){
	totalDistance := 0.0
	
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Cannot create file to write")
	}
	defer file.Close()
	
	w := bufio.NewWriter(file)
	
	fmt.Fprintln(w, route.pools[0].Properties.NAME + " " + FloatToString(totalDistance) + "\r\n") //Root pool, distance is 0 by default
	for i := 1; i < len(route.pools); i++{
		totalDistance += euclidDistance(route.pools[i - 1].Geometry.Coordinates[1],
										route.pools[i - 1].Geometry.Coordinates[0],
										route.pools[i].Geometry.Coordinates[1],
										route.pools[i].Geometry.Coordinates[0])
			
		fmt.Fprintln(w, route.pools[i].Properties.NAME + " " + FloatToString(totalDistance) + "\r\n")
	}
	w.Flush()
	return true
	
}

func FloatToString(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func main () {
	fmt.Println("Enter the number of go routines to use:")
	var num int
	fmt.Scanf("%d\n", &num)
	
	route := findRoute("wading-pools-min.json", num)
	fmt.Printf("%v", saveRoute(route, "solution.txt"))
}

func euclidDistance(lat1, lon1, lat2, lon2 float64)(dRadians float64){
	rLat1 := lat1 * math.Pi / 180
	rLat2 := lat2 * math.Pi / 180
	rLon1 := lon1 * math.Pi / 180
	rLon2 := lon2 * math.Pi / 180
	
	return 6371.0 * 2*math.Asin(math.Sqrt(math.Pow((math.Sin((rLat1 - rLat2)/2)), 2) + (math.Cos(rLat1) * math.Cos(rLat2) * math.Pow(math.Sin((rLon1 - rLon2)/2),2))))
	

}
