package main

import (
	"os"
	"fmt"
	"encoding/csv"
	"io"
	"bytes"
)

func main() {
	file,err:=os.Open("/Users/intro/Downloads/list.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var sql bytes.Buffer
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		sql.WriteString("update tender_user_account set old_income="+record[2]+" where uid="+record[0]+";\n\r")

	}
	file1,err:=os.OpenFile("/Users/intro/Downloads/list.sql",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0644)
	if err != nil {
		fmt.Println("Failed to open the file",err.Error())
		os.Exit(2)
	}
	file1.WriteString(sql.String())
	defer file1.Close()
	fmt.Println(sql.String())
}