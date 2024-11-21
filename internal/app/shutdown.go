package app

import "app/internal/log"

func (app *App) onShutdown() {
	l := log.L().Tag(log.TagShutdown)

	if app.MongoClient != nil {
		log.S.Debug("Shutting down MongoDB client", l)
		err := app.MongoClient.Close()
		if err != nil {
			log.S.Debug("Error in MongoDB client shutdown", l.Error(err))
		} else {
			log.S.Debug("MongoDB client shutdown complete", l)
		}
	}

	if app.ClickHouseClient != nil {
		log.S.Debug("Shutting down ClickHouse client", l)
		app.ClickHouseClient.Close()
		log.S.Debug("ClickHouse client shutdown complete", l)
	}
}
