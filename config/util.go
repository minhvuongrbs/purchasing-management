package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

type Configurable interface {
	Verify()
	SetDefault()
}

func verifyConfig(base interface{}) {
	baseType := reflect.TypeOf(base)
	modelType := reflect.TypeOf((*Configurable)(nil)).Elem()

	baseValue := reflect.ValueOf(base)
	for i := 0; i < baseType.NumField(); i++ {
		fieldType := baseType.Field(i).Type
		implemented := fieldType.Implements(modelType)
		if !implemented {
			continue
		}
		fieldValue := baseValue.Field(i)
		if fieldValue.IsNil() {
			log.Panicf("Config is nil, type=%v", fieldValue.Type())
		}
		fieldValue.MethodByName("Verify").Call([]reflect.Value{})
		fieldValue.MethodByName("SetDefault").Call([]reflect.Value{})
	}
}

// loadConfig load config from path then unmarshal to out interface
func loadConfig(path string, out interface{}) (err error) {
	viper.SetConfigFile(path)
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	return viper.Unmarshal(out)
}

// getConfigPath parse program args and get config path, pass -config=<path> to the program arguments
// Default get from environment variable (CONFIG_PATH) if -config absent
func getConfigPath() (configPath string, err error) {
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.StringVar(&configPath, "config", os.Getenv("CONFIG_PATH"), "Path to config file")
	err = fs.Parse(os.Args[1:])
	return
}

// printConfig print provided config to console
func printConfig(conf interface{}) {
	prettyJSON, _ := json.MarshalIndent(conf, "", "\t")
	fmt.Printf("Config: %v", string(prettyJSON))
	fmt.Println()
}
