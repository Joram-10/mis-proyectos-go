package main 
import "fmt"




func main() {
	var nombre string
	var edad int
	fmt.Print("Como te llamas? ")
	fmt.Scanln(&nombre)
	fmt.Print("Cuantos años tienes? ")
	fmt.Scanln(&edad)
	fmt.Println("Hola,", nombre)
	fmt.Println("Dentro de 10 años tendras", edad+10)
}
