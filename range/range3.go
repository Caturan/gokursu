package Range

import "fmt"

func Demo3() {
	dictionary := map[string]string{"Bottle": "Şişe", "Water": "Su"}

	for k, v := range dictionary {
		fmt.Println(k, " : ", v)
	}
}
