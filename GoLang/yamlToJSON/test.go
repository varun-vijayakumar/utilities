package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	//"path/filepath"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
// DeviceJSON represents unsupported device list data
type DeviceJSON struct {
	CnC map[string]Feature
}

// Feature represet list of features and supported devices
type Feature struct {
	Host string
	Urls []string
}
*/

func random() {
	/*
		cncInfo := make(map[string]string)
		uuid := "JL1369A"
		featureList := []string{"dot1x_enabled", "radius_enabled", "mac_auth_enabled", "invalid"}
		//fmt.Println(featureList)
		for _, feature := range featureList {
			cncInfo[feature] = "false"
		}
		fmt.Println(cncInfo)

		fmt.Println("validating cnc info")
		checkIfFeatureDisabledOnTheDevice(featureList, uuid, cncInfo)
		fmt.Println(cncInfo)
	*/
	/*
		var filePath = "/Users/varunvijayakumar/Desktop/FILES/NESTLE/GoLang/requiredCnCs.json"
		fmt.Printf("cncFile : filePath %s\n", filepath.Clean(filePath))
		rawSettings, readErr := ioutil.ReadFile(filepath.Clean(filePath))
		if readErr != nil {
			return
		}
		fmt.Printf("cncs file: %s\n", rawSettings)

		var result map[string]interface{}
		json.Unmarshal(rawSettings, &result)

		cncs := result["cncs"]
		fmt.Printf("cncs --> : %s\n", cncs)
		fmt.Println(reflect.TypeOf(cncs))

		var CnCInfo map[string]string = make(map[string]string)

		for _, value := range cncs.([]interface{}) {
			CnCInfo[value.(string)] = "100"
		}

		fmt.Println(CnCInfo)
	*/
	//
	//}

	/*
		var DeviceInfo map[string]string = make(map[string]string)

		kconfigcncs := result["kconfigcncs"]
		fmt.Printf("kconfig --> : %s\n", kconfigcncs)
		for _, value := range kconfigcncs.([]interface{}) {
			DeviceInfo[value.(string)] = ""
		}
		fmt.Println(DeviceInfo)

		for item := range CnCInfo {
			CnCInfo[item] = ""
		}
		fmt.Println(CnCInfo)
	*/
}
func checkIfFeatureDisabledOnTheDevice(featureList []string, uuid string, cncInfo map[string]string) {
	var filePath = "/Users/varunvijayakumar/Desktop/FILES/NESTLE/GoLang/unsupported.json"
	fmt.Printf("cncFile : filePath %s\n", filepath.Clean(filePath))
	rawSettings, readErr := ioutil.ReadFile(filepath.Clean(filePath))
	if readErr != nil {
		//fmt.Errorf("cnc : Fail to read settings file: %s", filePath)
	}
	fmt.Printf("file : %s\n", rawSettings)

	var result map[string]interface{}
	json.Unmarshal(rawSettings, &result)

	feat := result["features"]

	//fmt.Printf("features %s\n", feat)

	//fmt.Printf("type  %s\n", reflect.TypeOf(feat))
	fmt.Printf("searching uuid : %s\n", uuid)
	found := false
	for k, v := range feat.(map[string]interface{}) {
		fmt.Println("k:", k, "v:", v)
		feature, ok := v.(map[string]interface{})
		if !ok {
			fmt.Printf("feature %s not found\n\n", feature)
			continue
		}
		fmt.Println("Value : ", feature)
		unsupporteDevices := feature["unsupportedDevices"]
		if unsupporteDevices == nil {
			fmt.Printf("unsupportedDevices  not found\n")
		}
		fmt.Println("unsupportedDevices : ", unsupporteDevices)
		for _, device := range unsupporteDevices.([]interface{}) {
			if device == uuid {
				fmt.Printf("device found\n\n")
				cncInfo[k] = "true"
				found = true
			}
		}
		if !found {
			fmt.Printf("device not found\n\n")
		}
	}

	/*
			for feature, v := range feat.(map[string]interface{}) {
				for _, cncItem := range feature.([]interface{}) {

					item, okT := cncItem.(map[string]interface{})
					if okT {
						fmt.Printf("debug -- %s found\n", item)
					}
					unsupporteDevices := item["unsupportedDevices"]
					if unsupporteDevices == nil {
						fmt.Printf("unsupportedDevices  not found\n")
						continue
					}

					for _, device := range unsupporteDevices.([]interface{}) {
						if device == uuid {
							fmt.Printf("device found\n\n")
						}
					}
				}
			}


		md, ok := feat.(map[string]interface{})
		if !ok {
			return
		}
		found := false
		for _, item := range featureList {
			uuid := "JL1365A"
			fmt.Printf("searching feature : %s uuid : %s\n", item, uuid)
			feature, ok := md[item].(map[string]interface{})
			if !ok {
				fmt.Printf("feature %s not found\n\n", item)
				continue
			}
			unsupporteDevices := feature["unsupportedDevices"]
			if unsupporteDevices == nil {
				fmt.Printf("unsupportedDevices  not found for %s\n\n", item)
				continue
			}

			for _, device := range unsupporteDevices.([]interface{}) {
				if device == uuid {
					fmt.Printf("device found\n\n")
					cncInfo[item] = "true"
					found = true
				}
			}
			if !found {
				fmt.Printf("device not found\n\n")
			}
		}
	*/
}

// InterfaceInfo interface map
type InterfaceInfo struct {
	PortInfo map[string]interface{}
}

func process(value interface{}, portInfo map[string]interface{}) bool {
	var speedMap map[string]interface{} = make(map[string]interface{})

	switch value.(type) {
	case string:
		fmt.Printf("%v is an string \n ", value.(string))
	case bool:
		fmt.Printf("%v is bool \n ", value.(bool))
	case float64:
		fmt.Printf("%v is float64 \n ", value.(float64))
	case []interface{}:
		fmt.Printf("%v is a slice of interface \n ", value)
		for _, v := range value.([]interface{}) {
			process(v, portInfo)
		}
	case map[string]interface{}:
		fmt.Printf("%v is a map \n ", value)
		for _, v := range value.(map[string]interface{}) {
			process(v, portInfo)
		}
	case map[interface{}]interface{}:
		var curName = "1/1/"
		var speedList interface{}
		parentPortCount := portInfo["parent_port_count"].(int)
		totalPortCount := portInfo["total_port_count"].(int)
		portInfo["ports"] = make(map[string][]interface{})

		for k, v := range value.(map[interface{}]interface{}) {
			if k == "name" {
				str := fmt.Sprintf("%v", v)
				curName += str
				totalPortCount++
			}
			if k == "speeds" {
				speedList = v
			}
			if k == "subports" {
				parentPortCount++
			}
		}
		var portSpeeds []interface{}
		for _, s := range speedList.([]interface{}) {
			portSpeeds = append(portSpeeds, strconv.Itoa(s.(int)))
		}
		speedMap["speeds"] = portSpeeds
		portInfo[curName] = speedMap
		portInfo["parent_port_count"] = parentPortCount
		portInfo["total_port_count"] = totalPortCount
		// fmt.Println(parentPortCount)
		// fmt.Println(totalPortCount)
		return true
	default:
		fmt.Printf("%v is unknown \n ", value)
	}
	return false
}

func createAndWriteFile(directory string, data []byte) {
	if !createDirectory(directory) {
		return
	}

	fileName := directory + "/interface.json"
	err := ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func createDirectory(directory string) bool {
	err := os.MkdirAll(directory, 0755)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func createEmptyFile(fileName string) {
	d := []byte("")
	err := ioutil.WriteFile(fileName, d, 0644)
	if err != nil {
		log.Fatalf("failed to create file  #%v ", err)
	}
}

func extractFiles(directory string) bool {
	log.Println(directory)
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
		return false
	}
	for _, f := range files {
		log.Printf("file format %s", reflect.TypeOf(files))

		fileName := f.Name()
		if strings.Contains(fileName, ".json") {
			fileName = directory + fileName
			log.Printf("fileName : %s\n", fileName)
			file, err2 := ioutil.ReadFile(fileName)
			if err2 != nil {
				log.Fatal("read failed")
			}

			var intfYamlMap map[string]interface{} = make(map[string]interface{})
			err = json.Unmarshal(file, &intfYamlMap)
			if err != nil {
				fmt.Printf("Unmarshal: %v", err)
			}
			log.Println(intfYamlMap)
		}
	}
	return true
}

func extract(parentDirectory string, jNumber string) bool {
	log.Printf("Extracting File parentDirectory : %s jNumber : %s\n", parentDirectory, jNumber)
	intfYamlFile := parentDirectory + "ports.yaml"

	log.Printf("Extracting file %s\n", intfYamlFile)
	yamlFile, err := ioutil.ReadFile(intfYamlFile)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
		return false
	}
	// fmt.Printf("%s\n", yamlFile)

	var intfYamlMap map[string]interface{} = make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, &intfYamlMap)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
		return false
	}

	var portInfo map[string]interface{} = make(map[string]interface{})
	portInfo["parent_port_count"] = 0
	portInfo["total_port_count"] = 0
	portInfoCount := 0
	ok2 := false
	for k, v := range intfYamlMap {
		if k == "ports" {
			ports, ok := v.([]interface{})
			if !ok {
				fmt.Printf("extacting ports failed")
				return false
			}
			for _, port := range ports {
				//fmt.Printf("type : %s\n", reflect.TypeOf(port))
				if !process(port, portInfo) {
					fmt.Println("Port Info Fetch Failed")
					return false
				}
			}
		}
		if k == "port_info" {
			portInfo := v.(map[interface{}]interface{})
			portInfoCount, ok2 = portInfo["number_ports"].(int)
			if !ok2 {
				log.Fatal("port info count not found")
			}
			log.Printf("Port info : %d\n", portInfoCount)
		}
	}

	if portInfoCount == 0 {
		log.Fatalf("number_ports zero in %s", intfYamlFile)
		return false
	}

	totalPortCount := portInfo["total_port_count"].(int)
	if portInfoCount != totalPortCount {
		log.Printf("number_ports %d and total_port_count %d mismatch in %s",
			portInfoCount, totalPortCount, intfYamlFile)
		portInfoCount = totalPortCount
	}

	log.Printf("number_ports %d and total_port_count %d", portInfoCount, totalPortCount)
	if portInfo["parent_port_count"].(int) == 0 {
		portInfo["parent_port_count"] = portInfoCount
		log.Printf("ParentPortCount is 0 , set to total_port_count %d\n", portInfo["parent_port_count"].(int))
	}

	cncInfo := &InterfaceInfo{portInfo}
	jsonFile, err3 := json.Marshal(cncInfo)
	if err3 != nil {
		log.Fatalf("json marshall failed for %s err: %s", intfYamlFile, err3)
		return false
	}

	createAndWriteFile(parentDirectory, jsonFile)

	return true
}

func listDirectories(parentDirectory string) bool {
	log.Println(parentDirectory)
	directories, err := ioutil.ReadDir(parentDirectory)
	if err != nil {
		log.Fatal(err)
		return false
	}

	log.Printf("Listing subdir/parent/child")
	for _, item := range directories {
		//log.Println("% ", item.Name(), item.IsDir())
		if item.IsDir() {
			directoryName := item.Name()
			srcDirectory := parentDirectory + item.Name() + "/"
			log.Printf("%s %s\n", srcDirectory, directoryName)
			if !extract(srcDirectory, directoryName) {
				return false
			}
		}
	}
	return true
}

func main() {
	parentDirectory := "/Users/varunvijayakumar/Desktop/FILES/NESTLE/UTILITIES/GoLang/yamlToJSON/devices/"
	listDirectories(parentDirectory)
}
