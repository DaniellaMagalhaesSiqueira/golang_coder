package matematica

import (
	"fmt"
	"strconv"
)

// calcula média
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64) //mudar p/ .3f dá erro no teste
	return mediaArredondada

}
