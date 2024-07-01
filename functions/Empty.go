package skills

import (
	"os"
	"io"
	"bufio"
)


func IsEmpty(filename string) (bool,error) {
	file , err := os.Open(filename) 
	if err != nil{
		return false,err
	}

	reader := bufio.NewReader(file)
	_, err = reader.Peek(1)

	if err == io.EOF  {
		return true, nil
	}else {
		return false, err
	}

}