package private

import (
	"context"
	"github.com/go-kit/log/level"
	"gitlab.globerce.com/freedom-business/libs/shared-libs/errors"
	"gitlab.globerce.com/freedom-business/libs/shared-libs/utils/logger"
	"regexp"
	"strconv"
	"time"
)

const (
	IIN = iota + 1
	BIN
)

// Функция для проверки ИИН или БИН
func IINBINCheck(iin string) bool {
	logger := logger.NewServiceLogger("shared libs utils IINBINCheck")
	// Проверяем наличие ИИН или БИН
	if iin == "" || len(iin) != 12 || !iinLengthRegex.MatchString(iin) {
		_ = level.Info(logger).Log("iin validation error not matching required params len < 12 or regex mismatch - ", iin)
		return false
	}

	bd, err := GetBirthDayFromIIN(iin)
	if err != nil {
		_ = level.Info(logger).Log("error on parsing birthdate from iin - ", iin)
		return false
	}

	if bd.After(time.Now()) {
		_ = level.Info(logger).Log("iin birthdate appears to be after current date (not yet born) - ", iin)
		return false
	}

	centuryBornCodes := map[int]int{1: 19, 2: 19, 3: 20, 4: 20, 5: 21, 6: 21}

	day, err := strconv.Atoi(string(iin[4]))
	if err != nil {
		return false
	}

	//Определение ИИН или БИН если 5 разряд больше 3 то это скорее всего БИН
	clientType := IIN
	if day > 3 {
		clientType = BIN
	}

	switch clientType {
	case 1:
		gender, err := strconv.Atoi(string(iin[6]))
		if err != nil {
			return false
		}

		if _, ok := centuryBornCodes[gender]; !ok {
			return false
		}
	case 2:
		// Юридическое лицо (БИН)
		month, err := strconv.Atoi(iin[2:4])
		if err != nil || month > 12 {
			return false
		}

		residentDigit, err := strconv.Atoi(string(iin[4]))
		if err != nil {
			return false
		}

		// Проверяем признак резидентства
		if residentDigit < 4 || residentDigit > 6 {
			return false
		}
	}

	// Проверяем контрольную сумму
	b1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	b2 := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2}
	var sum int
	for i := 0; i < 11; i++ {
		digit, err := strconv.Atoi(string(iin[i]))
		if err != nil {
			return false
		}
		sum += digit * b1[i]
	}

	remainder := sum % 11
	if remainder == 10 {
		sum = 0
		for i := 0; i < 11; i++ {
			digit, err := strconv.Atoi(string(iin[i]))
			if err != nil {
				return false
			}
			sum += digit * b2[i]
		}
		remainder = sum % 11
	}

	controlDigit, err := strconv.Atoi(string(iin[11]))
	if err != nil {
		return false
	}

	return remainder == controlDigit
}

var iinLengthRegex = regexp.MustCompile(`^\d{12}$`)

func GetBirthDayFromIIN(iin string) (*time.Time, error) {
	// Проверяем наличие ИИН или БИН
	if iin == "" || len(iin) != 12 || !iinLengthRegex.MatchString(iin) {
		return nil, errors.IncorrectRequest.SetDevMessage("incorrect IIN")
	}

	day, err := strconv.Atoi(string(iin[4]))
	if err != nil {
		return nil, err
	}

	if day > 3 {
		return nil, errors.IncorrectRequest.SetDevMessage("client type should be iin not bin")
	}

	gender, err := strconv.Atoi(string(iin[6]))
	if err != nil {
		return nil, err
	}

	centuryBornCodes := map[int]int{1: 19, 2: 19, 3: 20, 4: 20, 5: 21, 6: 21}
	century, ok := centuryBornCodes[gender]

	if !ok {
		return nil, errors.IncorrectRequest.SetDevMessage("gender type should be iin not bin")
	}

	yearPrefix := strconv.Itoa(century - 1)

	iinBirthDate := yearPrefix + iin[0:6]

	birthDate, err := time.Parse("20060102", iinBirthDate) // Первые 6 цифр ИИН - дата рождения ГГММДД
	if err != nil {
		return nil, err
	}

	if birthDate.After(time.Now()) {
		return nil, errors.IncorrectRequest.SetDevMessage("birth date is in the future").
			AddClientError(context.Background(), "IB_IIN_002", nil)
	}

	return &birthDate, nil
}
