package grqExport

import (
	"os"
	"encoding/csv"
	"io"
)
type  Csv struct {
	FileName string
	FileHandle *os.File
	DealFunc DelFunc
}
type putcsvI interface {
	WriteHead(w *csv.Writer)
	WriteBody(w *csv.Writer)
}
func OutPutCsv(fileName string,i putcsvI) error {

	f, err := os.Create(fileName)//创建文件
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF")
	w:=csv.NewWriter(f)
	i.WriteHead(w)
	i.WriteBody(w)
	w.Flush()
	return nil
}
type DelFunc func (record []string) error
/**
  打开文件
 */
func (c *Csv)Open() (*csv.Reader,error) {
	file,err:=os.Open(c.FileName)
	if err != nil {
		return nil,err
	}
	reader := csv.NewReader(file)
	c.FileHandle=file
	return reader,nil
}
/**
  关闭文件
 */
func (c *Csv)Close(){
	c.FileHandle.Close()
}



/**
  读取文件
 */
func (c *Csv)Read(reader *csv.Reader) error{
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		err=c.DealFunc(record)
		if err!=nil{
			return  err
		}
	}
   return nil
}