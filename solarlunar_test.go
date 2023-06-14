package solarlunar_test

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jetfueltw/solarlunar-go"
	"github.com/stretchr/testify/assert"
)

func TestSolarToLunar(t *testing.T) {
	assert := assert.New(t)
	startTime := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC)

	for startTime.Before(endTime) {
		solar := solarlunar.Solar{
			Year:  startTime.Year(),
			Month: int(startTime.Month()),
			Day:   startTime.Day(),
		}

		lunar := solarlunar.SolarToLunar(solar)
		expectLunar, err := getSolarToLunar(solar)
		if !assert.NoError(err) {
			return
		}
		assert.Equal(expectLunar, lunar)

		solar = solarlunar.LunarToSolar(lunar)
		expectSolar, err := getLunarToSolar(lunar)
		if !assert.NoError(err) {
			return
		}
		assert.Equal(expectSolar, solar)

		startTime = startTime.AddDate(0, 0, 1)
	}
}

/**
 * https://github.com/isee15/Lunar-Solar-Calendar-Converter
 * node check.js in javascript folder
 */
func getSolarToLunar(solar solarlunar.Solar) (solarlunar.Lunar, error) {
	resp, err := http.Get("http://localhost:1337/?src=" + strconv.Itoa(solar.Year) + "," + strconv.Itoa(solar.Month) + "," + strconv.Itoa(solar.Day))
	if err != nil {
		return solarlunar.Lunar{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return solarlunar.Lunar{}, err
	}

	res := strings.Split(string(body), ",")

	year, err := strconv.Atoi(res[0])
	if err != nil {
		return solarlunar.Lunar{}, err
	}
	month, err := strconv.Atoi(res[1])
	if err != nil {
		return solarlunar.Lunar{}, err
	}
	day, err := strconv.Atoi(res[2])
	if err != nil {
		return solarlunar.Lunar{}, err
	}
	isLeap, err := strconv.Atoi(res[3])
	if err != nil {
		return solarlunar.Lunar{}, err
	}

	return solarlunar.Lunar{
		Year:   year,
		Month:  month,
		Day:    day,
		IsLeap: isLeap == 1,
	}, nil
}

/**
 * https://github.com/isee15/Lunar-Solar-Calendar-Converter
 * node check.js in javascript folder
 */
func getLunarToSolar(lunar solarlunar.Lunar) (solarlunar.Solar, error) {
	resp, err := http.Get("http://localhost:1337/?src=" + strconv.Itoa(lunar.Year) + "," + strconv.Itoa(lunar.Month) + "," + strconv.Itoa(lunar.Day) + "," + strconv.Itoa(boolToInt(lunar.IsLeap)))
	if err != nil {
		return solarlunar.Solar{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return solarlunar.Solar{}, err
	}

	res := strings.Split(string(body), ",")

	year, err := strconv.Atoi(res[0])
	if err != nil {
		return solarlunar.Solar{}, err
	}
	month, err := strconv.Atoi(res[1])
	if err != nil {
		return solarlunar.Solar{}, err
	}
	day, err := strconv.Atoi(res[2])
	if err != nil {
		return solarlunar.Solar{}, err
	}

	return solarlunar.Solar{
		Year:  year,
		Month: month,
		Day:   day,
	}, nil
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
