package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("kubectl", "get", "deployment", "hello-deployment", "-o=jsonpath='{.status.availableReplicas}'").Output()
	deployment_available_replicas_out := string(out)
	if err != nil {
		fmt.Println("error en la comprobación de las availableReplicas")
		log.Fatal(err)
	}

	out, err = exec.Command("kubectl", "get", "deployment", "hello-deployment", "-o=jsonpath='{.status.readyReplicas}'").Output()
	deployment_ready_replicas_out := string(out)
	if err != nil {
		fmt.Println("error en la comprobación de las readyReplicas")
		log.Fatal(err)
	}

	if deployment_available_replicas_out != "'2'" {
		fmt.Println("El número de réplicas disponibles no es el correcto")
	} else {
		fmt.Println("El número de réplicas disponibles es el correcto")
	}

	if deployment_ready_replicas_out != "'2'" {
		fmt.Println("El número de réplicas que están listas no es el correcto")
	} else {
		fmt.Println("El número de réplicas que están listas es el correcto")
	}
}