package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	//"path/filepath"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strconv"
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

func main() {

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
	var c conf
	c.getConf()
	fmt.Print(c)
}

type conf struct {
	Hits int64 `yaml:"hits"`
	Time int64 `yaml:"time"`
}

// InterfaceInfo interface map
type InterfaceInfo struct {
	PortInfo map[string]interface{}
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("ports.yaml")
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}
	//fmt.Printf("%s\n", yamlFile)

	var intfYamlMap map[string]interface{} = make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, &intfYamlMap)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}

	fmt.Printf("InterfaceMap2 type : %s\n", reflect.TypeOf(intfYamlMap))
	//fmt.Printf("%s\n", InterfaceMap2)
	//var ports []interface{}
	var portInfo map[string]interface{} = make(map[string]interface{})
	portInfo["parent_port_count"] = 0
	portInfo["total_port_count"] = 0
	jNumber := ""
	for k, v := range intfYamlMap {
		//fmt.Printf("%s : type : %s \n", k, reflect.TypeOf(v))
		if k == "jnumber" {

			jNumber = v.(string)
		}
		if k == "ports" {
			ports, ok := v.([]interface{})
			if !ok {
				fmt.Printf("extacting ports failed")
			}
			for _, port := range ports {
				//fmt.Printf("type : %s\n", reflect.TypeOf(port))
				if !process(port, portInfo) {
					fmt.Println("Port Info Fetch Failed")
				}
			}
			//fmt.Printf("Ports : %s\n", ports)
		}
	}
	fmt.Printf("JNUMBER : %s\n", jNumber)
	fmt.Printf("ParentPortCount : %d\n", portInfo["parent_port_count"].(int))
	fmt.Printf("TotalPortCount : %d\n", portInfo["total_port_count"].(int))
	/*fmt.Println("Dumping Port Info")
	for k, v := range portInfo {
		fmt.Printf("%s : %s\n", k, v)
	}
	*/
	cncInfo := &InterfaceInfo{portInfo}

	fmt.Printf("CnCInfo type : %s\n", reflect.TypeOf(cncInfo))

	//fmt.Printf("%s\n", cncInfo)
	jsonFile, err3 := json.Marshal(cncInfo)
	if err3 != nil {
		fmt.Print(err3)
	}
	//fmt.Print(jsonFile)

	fmt.Printf("jsonFile type : %s\n", reflect.TypeOf(jsonFile))
	err = ioutil.WriteFile("interface.json", jsonFile, 0644)
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}

	return c
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
		//fmt.Println(portCount)
		//fmt.Println(portInfo)
		return true
	default:
		fmt.Printf("%v is unknown \n ", value)
	}
	return false
}
