package main

import (
	"fmt"
)

type route struct {
	from     string
	to       string
	distance int
}

func calcPermDistance(routeToCalc []int, routesMap map[string]map[string]int, cities []string) int {
	distance := 0
	for index, _ := range routeToCalc {
		if index == 0 {
			continue
		}
		fromCity := cities[routeToCalc[index-1]]
		toCity := cities[routeToCalc[index]]
		fromCityRoutes, exists := routesMap[fromCity]
		if !exists {
			return 0
		}
		citiesDistance, exists := fromCityRoutes[toCity]
		if !exists {
			return 0
		}
		distance += citiesDistance
	}
	return distance
}

func perm(currentPerm []int, filter int, minimum *int, n int, routesMap map[string]map[string]int, cities []string) {
	if len(currentPerm) == n {
		candidate := calcPermDistance(currentPerm, routesMap, cities)
		if candidate > 0 && (candidate < *minimum || *minimum == -1) {
			*minimum = candidate
			return
		}
	} else {
		for i := 0; i < n; i++ {
			// Check ith bit is 1 or 0
			if filter&(1<<i) == 0 {
				perm(append(currentPerm, i), filter|(1<<i), minimum, n, routesMap, cities)
			}
		}
	}
}

func permTwo(currentPerm []int, filter int, maximum *int, n int, routesMap map[string]map[string]int, cities []string) {
	if len(currentPerm) == n {
		candidate := calcPermDistance(currentPerm, routesMap, cities)
		if candidate > *maximum {
			*maximum = candidate
			return
		}
	} else {
		for i := 0; i < n; i++ {
			// Check ith bit is 1 or 0
			if filter&(1<<i) == 0 {
				permTwo(append(currentPerm, i), filter|(1<<i), maximum, n, routesMap, cities)
			}
		}
	}

}

func partOne(inputList []route) int {
	routesMap := make(map[string]map[string]int)
	cities := []string{}

	for _, currentRoute := range inputList {
		_, exists := routesMap[currentRoute.from]
		if !exists {
			routesMap[currentRoute.from] = make(map[string]int)
			cities = append(cities, currentRoute.from)
		}

		_, exists = routesMap[currentRoute.to]
		if !exists {
			routesMap[currentRoute.to] = make(map[string]int)
			cities = append(cities, currentRoute.to)
		}
		routesMap[currentRoute.from][currentRoute.to] = currentRoute.distance
		routesMap[currentRoute.to][currentRoute.from] = currentRoute.distance
	}

	var minimum int
	minimum = -1
	perm([]int{}, 0, &minimum, len(cities), routesMap, cities)

	return minimum
}

func partTwo(inputList []route) int {
	routesMap := make(map[string]map[string]int)
	cities := []string{}

	for _, currentRoute := range inputList {
		_, exists := routesMap[currentRoute.from]
		if !exists {
			routesMap[currentRoute.from] = make(map[string]int)
			cities = append(cities, currentRoute.from)
		}

		_, exists = routesMap[currentRoute.to]
		if !exists {
			routesMap[currentRoute.to] = make(map[string]int)
			cities = append(cities, currentRoute.to)
		}
		routesMap[currentRoute.from][currentRoute.to] = currentRoute.distance
		routesMap[currentRoute.to][currentRoute.from] = currentRoute.distance
	}

	var maximum int
	maximum = 0
	permTwo([]int{}, 0, &maximum, len(cities), routesMap, cities)

	return maximum
}

func main() {
	inputList := []route{}
	for {
		var inputRoute route
		_, err := fmt.Scanf("%s to %s = %d", &inputRoute.from, &inputRoute.to, &inputRoute.distance)
		if err != nil {
			break
		}
		inputList = append(inputList, inputRoute)
	}
	fmt.Println(partOne(inputList))
	fmt.Println(partTwo(inputList))
}
