package file

import "io/ioutil"

func Read(name string) ([]byte, error) {
	bts, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return bts, nil
}
