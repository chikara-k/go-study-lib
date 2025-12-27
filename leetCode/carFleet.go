package leetcode

import "sort"

func carFleet(target int, position []int, speed []int) int {
    n := len(position)
    type Car struct {
        pos int
        speed int
    }
    cars := make([]Car, n)
    for i := 0; i < n; i++ {
        cars[i] = Car{pos: position[i], speed: speed[i]}
    }

    sort.Slice(cars, func(i, j int) bool { return cars[i].pos > cars[j].pos })

    stack := make([]float64, 0, n)

    for _, car := range cars {
        t := float64(target-car.pos) / float64(car.speed)

        if len(stack) == 0 || t > stack[len(stack)-1] {
            stack = append(stack, t)
        }
    }

    return len(stack)
}