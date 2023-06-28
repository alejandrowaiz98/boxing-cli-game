package orchestrator

import "github.com/alejandrowaiz98/boxing-cli-game/entitys"

func introduction() {

	logger.Printf("----------¡Bienvenido a la arena de batalla!----------")
	logger.Printf("---------- Antes de comenzar necesitamos algunos datos ----------")

}

func createMainBoxer() *entitys.Boxer {

	logger.Print("Por favor dinos tu nombre:")
	scanner.Scan()

	Player.Name = scanner.Text()

	logger.Print("Ahora, ¿Prefieres agilidad o fuerza?:")

	for scanner.Scan() {

		if election := scanner.Text(); election == "agilidad" || election == "fuerza" {

			Player.Style = election

			break

		} else {

			logger.Printf("Por favor %v, elige una opcion valida...", Player.Name)

		}

	}

	logger.Printf("Listo %v, nos vemos en la arena de combate...")

	return nil

}
