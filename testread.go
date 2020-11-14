package main

/*
import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) ([]string) {
	file, _ := os.Open(path)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main(){
	mapp := readLines("txt.txt")

	var myMap [2][3]string
	fmt.Println(mapp)
	fmt.Println(mapp[2][0])
	fmt.Println(mapp[1][1])

	var code uint8
	code = 206
	fmt.Print(string(code))

	for i:=0; i<400; i++{
		var code uint8
		code = uint8(i)
		stringcode := string(code)
		fmt.Printf("%s - %s", i, stringcode)
		fmt.Print("\n")
	}/*
	fmt.Println(myMap)
	/*for i:=0; i<len(mapp); i++{
		j := 0
		for{
			if mapp[i] == " "{
				break
			}else{
				myMap[j][i] = mapp[i]
				j++
			}

		}

	}



}
*/