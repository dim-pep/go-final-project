package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func NextDate(now time.Time, dstart, repeat string) (string, error) {

	if len(repeat) == 0 {
		return "", fmt.Errorf("правило пустое")
	}

	dateStart, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("ошибка преобразования даты, dstart: %s, err: %v", dstart, err)
	}

	repeatSplit := strings.Split(repeat, " ")
	rule := repeatSplit[0]

	switch rule { //possible rules: y, d, w, m
	case "y":
		for {
			dateStart = dateStart.AddDate(1, 0, 0)
			if afterDate(dateStart, now) {
				return dateStart.Format(dateFormat), nil
			}
		}
	case "d":
		if len(repeatSplit) != 2 {
			return "", fmt.Errorf("неверное количество параметров для правила d")
		}

		days, err := strconv.Atoi(repeatSplit[1])
		if err != nil {
			return "", fmt.Errorf("ожидался число для параметра d, получено: %s, err: %v", repeatSplit[1], err)
		}

		if days < 1 || days > 400 {
			return "", fmt.Errorf("допустимый диапозон для параметра d: от 1 до 400")
		}

		for {
			dateStart = dateStart.AddDate(0, 0, days)
			if afterDate(dateStart, now) {
				return dateStart.Format(dateFormat), nil
			}
		}
	case "w":
		if len(repeatSplit) < 2 {
			return "", fmt.Errorf("неверное количество параметров для правила w")
		}
		days := strings.Split(repeatSplit[1], ",")
		daysInt := make([]int, len(days))
		for _, day := range days {
			dayInt, err := strconv.Atoi(day)
			if err != nil {
				return "", fmt.Errorf("ожидался число для параметра w, получено: %s, err: %v", day, err)
			}
			if dayInt < 1 || dayInt > 7 {
				return "", fmt.Errorf("допустимый диапозон для параметра w: от 1 до 7")
			}
			daysInt = append(daysInt, dayInt)
		}

		//now.Weekday()
		for {
			dateStart = dateStart.AddDate(0, 0, 1)
			for _, day := range daysInt {
				weekDay := int(dateStart.Weekday())
				if weekDay == 0 {
					weekDay = 7
				}
				if weekDay == day {
					if afterDate(dateStart, now) {
						return dateStart.Format(dateFormat), nil
					}
				}
			}
		}
	case "m":
		var month [13]bool
		if len(repeatSplit) < 2 {
			return "", fmt.Errorf("нету параметров для правила m")
		}
		day, daysOnEnd, err := getDayTrue(repeatSplit[1])
		if err != nil {
			return "", err
		}
		switch len(repeatSplit) {
		case 2:
			for i := range month {
				month[i] = true
			}
		default:
			month, err = getMonthTrue(repeatSplit[2])
			if err != nil {
				return "", err
			}
		}
		for {
			dateStart = dateStart.AddDate(0, 0, 1)
			if afterDate(dateStart, now) {
				if len(daysOnEnd) == 0 {
					if month[int(dateStart.Month())] && day[int(dateStart.Day())] {
						return dateStart.Format(dateFormat), nil
					}
				} else {
					lastDayInMonth := time.Date(dateStart.Year(), dateStart.Month()+1, 0, 0, 0, 0, 0, time.UTC)
					for _, dayOnEnd := range daysOnEnd {
						if (int(lastDayInMonth.Day())+dayOnEnd+1) == int(dateStart.Day()) || day[int(dateStart.Day())] && month[int(dateStart.Month())] {
							return dateStart.Format(dateFormat), nil
						}
					}
				}
			}
		}
	default:
		return "", fmt.Errorf("неизвестное правило повторения: %s", rule)
	}
}

func afterDate(dateAfter, dateBefore time.Time) bool {
	return dateAfter.After(dateBefore)
}

func getDayTrue(dayStr string) ([32]bool, []int, error) {
	var day [32]bool
	var daysOnEnd []int
	daysTrue := strings.Split(dayStr, ",")
	for _, dayTrue := range daysTrue {
		dayTrueInt, err := strconv.Atoi(dayTrue)
		if err != nil {
			return day, daysOnEnd, fmt.Errorf("ожидался число для параметра m, получено: %s, err: %v", dayTrue, err)
		}
		if dayTrueInt < -2 || dayTrueInt > 31 || dayTrueInt == 0 {
			return day, daysOnEnd, fmt.Errorf("допустимый диапозон для параметра m: от -2 до 31, без нуля")
		}
		if dayTrueInt > 0 {
			day[dayTrueInt] = true
		} else {
			daysOnEnd = append(daysOnEnd, dayTrueInt)
		}
	}
	return day, daysOnEnd, nil
}

func getMonthTrue(monthStr string) ([13]bool, error) {
	var month [13]bool
	monthsTrue := strings.Split(monthStr, ",")
	for _, monthTrue := range monthsTrue {
		monthTrueInt, err := strconv.Atoi(monthTrue)
		if err != nil {
			return month, fmt.Errorf("ожидался число для параметра m, получено: %s, err: %v", monthTrue, err)
		}
		if monthTrueInt < 1 || monthTrueInt > 12 {
			return month, fmt.Errorf("допустимый диапозон для параметра m: от 1 до 12")
		}
		month[monthTrueInt] = true
	}
	return month, nil
}
