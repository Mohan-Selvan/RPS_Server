package main

import (
	"context"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/xuri/excelize/v2"
)

type SignTable [NUMBER_OF_SIGNS][NUMBER_OF_SIGNS]float64

var SIGN_TABLE SignTable

const SIGN_TABLE_PATH string = "./data/modules/files/sign_table.xlsx"
const SIGN_TABLE_SHEET_MAIN string = "sign_table"
const SIGN_TABLE_SHEET_STRIPPED string = "stripped"

func InitializeSignTable(ctx context.Context, logger runtime.Logger) {

	logger.Info("Initializing Sign table..")

	f, err := excelize.OpenFile(SIGN_TABLE_PATH)
	if err != nil {
		PrintError(logger, err, "Error initializing Sign Table!")
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Printing table
	//rows, err := f.GetRows(SIGN_TABLE_SHEET_STRIPPED)
	//if err != nil {
	//	fmt.Println(err)
	//	log.Fatal()
	//}
	//for _, row := range rows {
	//	for _, colCell := range row {
	//		fmt.Print(colCell, "\t")
	//	}
	//	fmt.Println()
	//}

	for i := 0; i < NUMBER_OF_SIGNS; i++ {
		for j := 0; j < NUMBER_OF_SIGNS; j++ {
			cellCoordinate, coErr := excelize.CoordinatesToCellName(j+1, i+1) //+1 => index to cell number; //Inverting J and I because of Cell coordinates.
			if coErr != nil {
				PrintError(logger, coErr, "Error converting coordinates to cell value.")
			}

			cellValueString, err := f.GetCellValue(SIGN_TABLE_SHEET_STRIPPED, cellCoordinate)
			if err != nil {
				PrintError(logger, err, "Error fetching cell value.")
			}

			fmt.Printf("Cell coordinate for : [%d, %d] => %v\n", i, j, cellCoordinate)

			cellValue := StringToFloat64(cellValueString)
			SIGN_TABLE[i][j] = cellValue

			logger.Info("SIGN_TABLE[%v][%v] = %v", SIGNS.GetSign(i).ToString(), SIGNS.GetSign(j).ToString(), cellValue)
		}
	}

	logger.Info("Sign table initialized successfully!!")

	fmt.Print("\t")
	for j := 0; j < NUMBER_OF_SIGNS; j++ {
		fmt.Print(SIGNS.GetSign(j).name, "\t")
	}
	fmt.Print("\n")

	for i := 0; i < NUMBER_OF_SIGNS; i++ {
		fmt.Print(SIGNS.GetSign(i).name, "\t")
		for j := 0; j < NUMBER_OF_SIGNS; j++ {
			fmt.Print(SIGN_TABLE[i][j], "\t")
		}
		fmt.Print("\n")
	}
}
