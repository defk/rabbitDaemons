package handlers

import (
	"fmt"
	"first-rabbit/structures"
)

func Meteo(i structures.Task, out chan<- string) {

	out <- fmt.Sprintf(
		"Start task: {\"alias\": %s, \"guid\": %s}",
		i.Alias,
		i.Content.GUID,
	)
}
