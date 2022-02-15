package controllers

import (
	"time"

	tzcronjobv1 "github.com/kkothule/cronjobber/api/v1"
)

func getCurrentTimeInZone(sj *tzcronjobv1.TZCronjob) (time.Time, error) {
	if sj.Spec.TimeZone == "" {
		return time.Now(), nil
	}

	loc, err := time.LoadLocation(sj.Spec.TimeZone)
	if err != nil {
		return time.Now(), err
	}

	return time.Now().In(loc), nil
}
