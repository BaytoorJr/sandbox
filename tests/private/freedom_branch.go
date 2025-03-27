package private

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"net/http"
	"os"
)

type BranchInfo struct {
	Branch       string `json:"branch"`        // Филиал
	BranchCode   string `json:"branch_code"`   // Branch code
	Department   string `json:"department"`    // Отделение
	WorkSchedule string `json:"work_schedule"` // График работы
	AddressRu    string `json:"address_ru"`    // Мекенжайы (русский)
	AddressKz    string `json:"address_kz"`    // Жұмыс кестесі (казахский)
}

type FreedomBranches struct {
	Code           string            `json:"code" bson:"code"`
	City           string            `json:"city" bson:"city"`
	CityBranchCode string            `json:"cityBranchCode" bson:"cityBranchCode"`
	FullAddress    map[string]string `json:"fullAddress" bson:"fullAddress"`
	WorkSchedule   map[string]string `json:"workSchedule" bson:"workSchedule"`
}

const referenceName = "FreedomBranch"

type CreateReferenceRequest struct {
	ReferenceName string `json:"referenceName"` // Name of the reference (ex. KNP)
	Data          any    `json:"data"`
}

func SendFreedomBranches() {
	filePath := "./table.xlsx"
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()

	sheetName := f.GetSheetName(0)

	rows, err := f.GetRows(sheetName)
	if err != nil {
		panic(err)
	}

	var branchData []BranchInfo
	for i, row := range rows {
		if i == 0 || i == 1 {
			continue
		}

		if len(row) < 7 {
			log.Printf("Пропущена строка %d из-за недостатка колонок: %v", i+1, row)
			continue
		}

		// Создаем структуру из строки
		branch := BranchInfo{
			Branch:       row[2],
			BranchCode:   row[3],
			Department:   row[4],
			WorkSchedule: row[5],
			AddressRu:    row[6],
			AddressKz:    row[7],
		}
		branchData = append(branchData, branch)
	}

	branchCodes := make(map[string]int)

	for i, d := range branchData {
		code, ok := branchCodes[d.BranchCode]
		if !ok {
			branchCodes[d.BranchCode] = 1
			code = 1
		}

		reference := FreedomBranches{
			Code:           fmt.Sprintf("%s-%d", d.BranchCode, code),
			City:           d.Branch,
			CityBranchCode: d.BranchCode,
			FullAddress: map[string]string{
				"ru": d.Department,
				"kz": d.AddressRu,
			},
			WorkSchedule: map[string]string{
				"ru": d.WorkSchedule,
				"kz": d.AddressKz,
			},
		}

		req := CreateReferenceRequest{
			ReferenceName: referenceName,
			Data:          reference,
		}

		reqBytes, err := json.Marshal(req)
		if err != nil {
			panic(err)
		}

		request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "https://ibul.trafficwave.kz/directory-api/v1/references/", bytes.NewBuffer(reqBytes))
		if err != nil {
			panic(err)
		}
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("API-Key", "10891ad8-06a8-437f-a788-9d0c4310a338")

		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			panic(err)
		}

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			log.Println(fmt.Sprintf("got something went wrong %s", body))
			os.Exit(1)
		}

		log.Println(fmt.Sprintf("sent request #%d: %s", i, string(reqBytes)))
		branchCodes[d.BranchCode]++
	}
}
