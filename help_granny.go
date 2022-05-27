package main

// import (
// 	"fmt"
// 	"math"
// )

// func HGmain() {
// 	fmt.Println("Hello world")
// 	var friends1 = []string{"A1", "A2", "A3", "A4", "A5"}
// 	var fTowns1 = map[string]string{"A1": "X1", "A2": "X2", "A3": "X3", "A4": "X4"}
// 	var dist1 = map[string]float64{"X1": 100.0, "X2": 200.0, "X3": 250.0, "X4": 300.0}
// 	// friends1 := []string{"B1", "B2"}
// 	// fTowns1 := map[string]string{"B1": "Y1", "B2": "Y2", "B3": "Y3", "B4": "Y4", "B5": "Y5"}
// 	// dist1 := map[string]float64{"Y1": 50.0, "Y2": 70.0, "Y3": 90.0, "Y4": 110.0, "Y5": 150.0}

// 	res := Tour(friends1, fTowns1, dist1)
// 	fmt.Println(res)

// }
// func contains(s []string, str string) bool {
// 	for _, v := range s {
// 		if v == str {
// 			return true
// 		}
// 	}
// 	return false
// }

// func Tour(arrFriends []string, ftwns map[string]string, h map[string]float64) int {
// 	// your code
// 	dist := 0.0
// 	fkeys := []string{}
// 	tkeys := []string{}
// 	dkeys := []string{}

// 	for i, _ := range ftwns {
// 		tkeys = append(tkeys, i)
// 	}

// 	for i, _ := range h {
// 		dkeys = append(dkeys, i)
// 	}

// 	for i := 1; i < len(arrFriends); i++ {
// 		if (contains(tkeys, arrFriends[i]) && contains(dkeys, ftwns[arrFriends[i]])) &&
// 			contains(tkeys, arrFriends[i-1]) && contains(dkeys, ftwns[arrFriends[i-1]]) {
// 			a1 := math.Pow(h[ftwns[arrFriends[i]]], 2)
// 			a2 := math.Pow(h[ftwns[arrFriends[i-1]]], 2)
// 			r := math.Sqrt(a1 - a2)
// 			dist += r
// 			if i == 1 {
// 				fkeys = append(fkeys, arrFriends[i-1])
// 			}
// 			fkeys = append(fkeys, arrFriends[i])
// 		}
// 	}
// 	dist += h[ftwns[fkeys[0]]]
// 	dist += h[ftwns[fkeys[len(fkeys)-1]]]

// 	return int(dist)
// }
