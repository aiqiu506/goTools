package main

import "fmt"

func main() {

	/*params := os.Args
	if len(params) != 3 {
		fmt.Println("params is error")
		return
	}
	var srcFileName = params[1]
	var dstFileName = params[2]
	srcFile, error := os.Open(srcFileName)
	if error != nil {
		fmt.Println("open srcFile error:", error.Error())
		return
	}
	dstFile, error := os.Create(dstFileName)
	if error != nil {
		fmt.Println("crate dstFile error:", error.Error())
		return
	}
	defer srcFile.Close()
	defer dstFile.Close()

	var buf = make([]byte, 4*1024)
	n, err := srcFile.Read(buf)
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("readFile error:", err)
	}
	dstFile.Write(buf[:n])*/

	/*	//读写csv文件
		type dataTemp struct {
			id      string `title:"标id"`
			id_card string `title:"身份证"`
			amount  string `title:"金额"`
		}
		data := make([]dataTemp, 0)
		csvHanler := &grqExport.Csv{
			FileName: "/Users/intro/Downloads/list.csv",
			DealFunc: func(record []string) error {
				data = append(data, dataTemp{record[0], record[1], record[2]})
				return nil
			},
		}
		file, _ := csvHanler.Open()

		csvHanler.Read(file)
		defer csvHanler.Close()

		//再生成一个文件
		headT := dataTemp{}
		oo := grqExport.OutCsv{Head: headT, Data: data}
		grqExport.OutPutCsv("/Users/intro/Downloads/list2.csv", &oo)
		grqStruct.StructToMap()*/

	type A struct {
		id int
	}
	var a map[string]struct{ id int }
	a = make(map[string]struct{ id int })
	a["a"] = struct {
					id int
				}{12}
	//a["a"].id = 2
	b := a["a"].id
	fmt.Println(b)

}
