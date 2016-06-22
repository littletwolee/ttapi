package tools

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	Conf *ConfigHelper
	conflist []map[string]map[string]string
)
type ConfigHelper struct {
	Filepath string                         
}

func init() {
	Conf = new(ConfigHelper)
	Conf.Filepath = "./conf/app.conf"
	Conf.ReadList()
}

func (c *ConfigHelper) GetValue(section, name string) string {
	// if conflist == nil{
	// 	c.ReadList()	
	// }
	for _, v := range conflist {
		for key, value := range v {
			if key == section {
				return value[name]
			}
		}
	}
	return "no value"
}

func (c *ConfigHelper) ReadList() []map[string]map[string]string {
	file, err := os.Open(c.Filepath)
	if err != nil {
		CheckErr(err)
	}
	defer file.Close()
	var data map[string]map[string]string
	var section string
	buf := bufio.NewReader(file)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				CheckErr(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		case line[0] == '[' && line[len(line)-1] == ']':
			section = strings.TrimSpace(line[1 : len(line)-1])
			data = make(map[string]map[string]string)
			data[section] = make(map[string]string)
		default:
			i := strings.IndexAny(line, "=")
			value := strings.TrimSpace(line[i+1 : len(line)])
			data[section][strings.TrimSpace(line[0:i])] = value
			if c.uniquappend(section) == true {
				conflist = append(conflist, data)
			}
		}

	}

	return conflist
}

func CheckErr(err error) string {
	if err != nil {
		return fmt.Sprintf("Error is :'%s'", err.Error())
	}
	return "Notfound this error"
}


func (c *ConfigHelper) uniquappend(conf string) bool {
	for _, v := range conflist {
		for k, _ := range v {
			if k == conf {
				return false
			}
		}
	}
	return true
}

func AppConfig() *ConfigHelper {
    return Conf
}
