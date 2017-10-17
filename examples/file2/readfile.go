package file

import "io/ioutil"

func Read(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}
