package src

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func changeEmptyFieldToDefault(data string) string {
	if data == "" {
		return "Default"
	}
	return data
}

func FindBuSize(terminalId string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	csvfile, err := os.Open(os.Getenv("MERCHANT"))
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if newTerminalId := changeEmptyFieldToDefault(terminalId); record[2] == newTerminalId {
			return record[4]
		}
	}
	return ""
}

func FindCapabilityNonOil(buSize string) int {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	csvfile, err := os.Open(os.Getenv("CAPABILITY_NON_OIL"))
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if newBuSize := changeEmptyFieldToDefault(buSize); record[0] == newBuSize {
			maxCapabilityNonOil, err := strconv.Atoi(record[1])
			if err != nil {
				fmt.Println(err)
				return -1
			}
			return maxCapabilityNonOil
		}
	}
	return -1
}

func FindCapabilityOil(memberClass, productGroup string) int {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	csvfile, err := os.Open(os.Getenv("CAPABILITY_OIL"))
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if newMemberClass := changeEmptyFieldToDefault(memberClass); record[0] == newMemberClass && record[1] == productGroup {
			maxCapabilityOil, err := strconv.Atoi(record[2])
			if err != nil {
				fmt.Println(err)
				return -1
			}
			return maxCapabilityOil
		}
	}
	return -1
}

func FindProduct(productCode string) map[string]string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	csvfile, err := os.Open(os.Getenv("PRODUCT"))
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	currentProduct := make(map[string]string)
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] == productCode {
			currentProduct["productCode"] = record[0]
			currentProduct["productName"] = record[1]
			currentProduct["productGroup"] = record[2]
			currentProduct["productType"] = record[3]
			currentProduct["samePriseDelta"] = record[4]
			currentProduct["priceRiseStartDate"] = record[5]
			currentProduct["priceRiseEndDate"] = record[6]
			return currentProduct
		}
	}
	return currentProduct
}

func FindMember(memberId string) map[string]string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	csvfile, err := os.Open(os.Getenv("MEMBER"))
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	currentMember := make(map[string]string)
	if len(memberId) == 16 {
		// Iterate through the records
		for {
			// Read each record from csv
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			if record[0] == memberId {
				currentMember["cardNumber"] = record[0]
				currentMember["mobileNumber"] = record[1]
				currentMember["cardType"] = record[2]
				currentMember["memberClass"] = record[3]
				return currentMember
			}
		}
	} else if len(memberId) == 10 {
		for {
			// Read each record from csv
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			if record[1] == memberId {
				currentMember["cardNumber"] = record[0]
				currentMember["mobileNumber"] = record[1]
				currentMember["cardType"] = record[2]
				currentMember["memberClass"] = record[3]
				return currentMember
			}
		}
	}
	return currentMember
}
