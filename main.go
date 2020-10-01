package main

import "fmt"

func main() {
	var mes int
	var dia int
	var signo string

	fmt.Scanf("%d\n", &dia)
	fmt.Scanf("%d\n", &mes)

	switch {
	case mes == 1:
		if dia > 0 && dia <= 19 {
			signo = "capricornio"
		} else if dia >= 20 && dia <= 31 {
			signo = "acuario"
		} else {
			signo = "No es un día válido"
		}
	case mes == 2:
		if dia > 0 && dia <= 18 {
			signo = "acuario"
		} else if dia >= 19 && dia <= 29 {
			signo = "piscis"
		} else {
			signo = "No es un día válido"
		}
	case mes == 3:
		if dia > 0 && dia <= 20 {
			signo = "piscis"
		} else if dia >= 21 && dia <= 31 {
			signo = "aries"
		} else {
			signo = "No es un día válido"
		}
	case mes == 4:
		if dia > 0 && dia <= 19 {
			signo = "aries"
		} else if dia >= 20 && dia <= 30 {
			signo = "tauro"
		} else {
			signo = "No es un día válido"
		}
	case mes == 5:
		if dia > 0 && dia <= 20 {
			signo = "tauro"
		} else if dia >= 21 && dia <= 31 {
			signo = "geminis"
		} else {
			signo = "No es un día válido"
		}
	case mes == 6:
		if dia > 0 && dia <= 20 {
			signo = "geminis"
		} else if dia >= 21 && dia <= 30 {
			signo = "cancer"
		} else {
			signo = "No es un día válido"
		}
	case mes == 7:
		if dia > 0 && dia <= 22 {
			signo = "cancer"
		} else if dia >= 23 && dia <= 31 {
			signo = "leo"
		} else {
			signo = "No es un día válido"
		}
	case mes == 8:
		if dia > 0 && dia <= 22 {
			signo = "leo"
		} else if dia >= 23 && dia <= 31 {
			signo = "virgo"
		} else {
			signo = "No es un día válido"
		}
	case mes == 9:
		if dia > 0 && dia <= 22 {
			signo = "virgo"
		} else if dia >= 23 && dia <= 30 {
			signo = "libra"
		} else {
			signo = "No es un día válido"
		}
	case mes == 10:
		if dia > 0 && dia <= 22 {
			signo = "libra"
		} else if dia >= 23 && dia <= 31 {
			signo = "escorpio"
		} else {
			signo = "No es un día válido"
		}
	case mes == 11:
		if dia > 0 && dia <= 21 {
			signo = "escorpio"
		} else if dia >= 22 && dia <= 30 {
			signo = "sagitario"
		} else {
			signo = "No es un día válido"
		}
	case mes == 12:
		if dia > 0 && dia <= 21 {
			signo = "sagitario"
		} else if dia >= 22 && dia <= 31 {
			signo = "capricornio"
		} else {
			signo = "No es un día válido"
		}
	default:
		signo = "No es un mes válido"
	}

	fmt.Println(signo)
}