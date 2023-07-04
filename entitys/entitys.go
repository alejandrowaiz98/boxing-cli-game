package entitys

import "github.com/alejandrowaiz98/boxing-cli-game/config"

var logger = config.GetLogger()

type Entitys struct{}

func (e *Entitys) Init() error {

	err := createPunches()

	if err != nil {
		logger.Err(err).Msgf("initEntitys | createPunches: %v", err)
		return err
	}

	err = createTournaments()

	if err != nil {
		logger.Err(err).Msgf("initEntitys | createTournaments: %v", err)
		return err
	}

	return nil

}
