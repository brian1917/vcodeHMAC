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
		if strings.Contains(scanner.Text(), "veracode_api_key_id") {
			credentials[0] = strings.Replace(scanner.Text(),"veracode_api_key_id = ","", -1)
		} else if strings.Contains(scanner.Text(), "veracode_api_key_secret"){
			credentials[1] = strings.Replace(scanner.Text(),"veracode_api_key_secret = ","", -1)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return credentials
}
