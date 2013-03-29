package comptbfile

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Compare(fname1, fname2 string, beginLine, endLine, cellNum int) {
	f1, err := os.Open(fname1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f1.Close()

	f2, ef2 := os.Open(fname2)
	if ef2 != nil {
		panic(ef2)
	}
	defer f2.Close()

	f3, ef3 := os.Create("result.txt")
	if ef3 != nil {
		panic(ef3)
	}
	defer f3.Close()
	buf2 := bufio.NewReader(f2)
	buf1 := bufio.NewReader(f1)
	buf3 := bufio.NewWriter(f3)

	lineNum := 1
	// begin := 1826
	// end := 2745
	for {
		line1, err1 := buf1.ReadString('\n')
		if err1 != nil && err1 == io.EOF {
			break
		}
		line2, err2 := buf2.ReadString('\n')
		if err2 != nil && err2 == io.EOF {
			break
		}

		if lineNum >= beginLine && lineNum <= endLine {
			l2 := strings.Split(line2, "\t")
			l1 := strings.Split(line1, "\t")
			if lineNum == beginLine {
				_, err3 := buf3.WriteString(l1[0] + "\t" + l2[0] + "\r\n")
				if err3 != nil {
					panic(err3)
				}
			}
			//fmt.Println(l1[35], "=", l2[35])
			_, err3 := buf3.WriteString(l1[cellNum] + "====" + l2[cellNum] + "\r\n")
			if err3 != nil {
				panic(err3)
			}

			if lineNum == endLine {
				_, err3 := buf3.WriteString(l1[0] + "\t" + l2[0] + "\r\n")
				if err3 != nil {
					panic(err3)
				}
				buf3.Flush()
			}
		}
		lineNum++
	}
}
