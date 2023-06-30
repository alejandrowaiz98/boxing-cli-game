package orchestrator

import "github.com/alejandrowaiz98/boxing-cli-game/entitys"

func introduction() {

	logger.Printf("----------¡Bienvenido a la arena de batalla!----------")
	logger.Printf("---------- Antes de comenzar necesitamos algunos datos ----------")

}
func createMainBoxer() *entitys.Boxer {

	logger.Print("Por favor dinos tu nombre:")
	scanner.Scan()

	logger.Print("Ahora, ¿Prefieres agilidad o fuerza?:")

	for scanner.Scan() {

		if election := scanner.Text(); election == "agilidad" || election == "fuerza" {

			

			break

		} else {


		}

	}

	logger.Printf("Listo %v, nos vemos en la arena de combate...")

	return nil

}
