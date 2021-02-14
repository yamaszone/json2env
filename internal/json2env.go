package json2env

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type JsonData map[string]string

func ToEnv(filename string, exportable bool) error {
	var jd JsonData

	f, err := os.Open(filename)
	if err != nil {
		log.Printf("Error: failed to open %s. %s", filename, err)
		return err
	}
	defer f.Close()

	res, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("Error: failed to read %s. %s", filename, err)
		return err
	}

	err = json.Unmarshal(res, &jd)
	if err != nil {
		log.Printf("Error: failed to parse input file %s. %s", filename, err)
		return err
	}

	//fmt.Println(jd)
	for k, v := range jd {
		if exportable {
			fmt.Println(fmt.Sprintf("export %s=%s", k, v))
		} else {
			fmt.Println(fmt.Sprintf("%s=%s", k, v))
		}
	}

	return nil
}
