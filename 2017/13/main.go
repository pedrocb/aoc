package main

import (
	"fmt"
)

type wall struct {
	position, _range int
}

func CalcSeverity(initialDelay int, firewall []wall, returnOnHit bool) int {
	severity := 0
	for _, _wall := range firewall {
		wallPos := (initialDelay + _wall.position) % ((_wall._range - 1) * 2)
		if wallPos == 0 {
			if returnOnHit {
				return -1
			}
			severity += (_wall.position * _wall._range)
		}
	}
	return severity
}

func FindZeroSeverity(firewall []wall) int {
	delay := 0
	for {
		severity := CalcSeverity(delay, firewall, true)
		if severity != -1 {
			return delay
		}
		delay++
	}

}

func main() {
	firewall := []wall{}
	for {
		inputWall := wall{}
		_, err := fmt.Scanf("%d: %d", &inputWall.position, &inputWall._range)
		if err != nil {
			break
		}
		firewall = append(firewall, inputWall)
	}
	fmt.Println(CalcSeverity(0, firewall, false))
	fmt.Println(FindZeroSeverity(firewall))
}
