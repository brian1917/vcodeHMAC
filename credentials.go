package vcodeHMAC

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func getCredentials(fileString string) [2]string {
	var credentials [2]string

	file, err := os.Open(fileString)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//We remove spaces to account for discrepancies in user configuration of creds file
		if strings.Contains(scanner.Text(), "veracode_api_key_id") {
			removeSpaces := strings.Replace(scanner.Text(), " ", "", -1)
			credentials[0] = strings.Replace(removeSpaces, "veracode_api_key_id=", "", -1)
		} else if strings.Contains(scanner.Text(), "veracode_api_key_secret") {
			removeSpaces := strings.Replace(scanner.Text(), " ", "", -1)
			credentials[1] = strings.Replace(removeSpaces, "veracode_api_key_secret=", "", -1)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return credentials
}
