package rappelconso

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/xefiji/rappelconso/pkg/datagouv"
)

func fetchHandler(repo *Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		offset, err := repo.countDBRecords()
		if err != nil {
			log.Error().Caller().Err(err).Msg("error while counting records")
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error while counting records"})
			return
		}

		if err := saveRecords(repo, offset, 100); err != nil {
			log.Error().Err(err).Msg("error while saving records")
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error while saving records"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func saveRecords(repo *Repository, offset, limit int) error {
	totalAPI, err := datagouv.CountRecords()
	if err != nil {
		log.Error().Caller().Err(err).Msg("failed counting records")
		return err
	}

	maxBatches := math.Ceil(float64(totalAPI) / float64(100))
	for i := 0; float64(i) < maxBatches; i++ {
		if offset >= totalAPI {
			return nil
		}
		records, err := datagouv.GetRecords(limit, offset)
		if err != nil {
			log.Error().Caller().Err(err).Msg("failed fetching records")
			return err
		}

		for _, recordData := range records.RecordData {
			re := newRecordEntity(recordData)
			err := repo.save(re)
			if err != nil {
				log.Error().Caller().Err(err).Interface("record", re).Msg("failed inserting record")
			}
		}
		offset += limit
	}

	return nil
}
