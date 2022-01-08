package defcore

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Stash interface {
	Init(sessionId string)
	ReadData(sessionId, item string) string
	WriteData(sessionId, item, data string)

	Add(sessionId, item string) error
}

type Disk struct {
	RootDir     string
	DataDirName string
	VarDirName  string
}

func LinuxDisk() Stash {
	return &Disk{
		RootDir:     "/tmp/deflang/",
		DataDirName: "data",
		VarDirName:  "vars",
	}
}

func (disk *Disk) dataFile(sessionId, item string) string {
	return disk.RootDir + "/" + sessionId + "/" + disk.DataDirName + "/" + item
}

func (disk *Disk) varFile(sessionId, item string) string {
	return disk.RootDir + "/" + sessionId + "/" + disk.VarDirName + "/" + item
}

func (disk *Disk) Init(sessionId string) {
	os.MkdirAll(disk.RootDir+"/"+sessionId+"/"+disk.DataDirName, 0700)
	os.MkdirAll(disk.RootDir+"/"+sessionId+"/"+disk.VarDirName, 0700)
}

func (disk *Disk) ReadData(sessionId, item string) string {
	content, err := os.ReadFile(disk.dataFile(sessionId, item))
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(content)
}

func (disk *Disk) WriteData(sessionId, item, data string) {
	err := ioutil.WriteFile(disk.dataFile(sessionId, item), []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (disk *Disk) Add(sessionId, item string) error {
	_, err := copyFile("./defvar", disk.varFile(sessionId, item))
	if err != nil {
		log.Fatal(err)
	}
	os.Chmod(disk.varFile(sessionId, item), 0764)
	dataFile, e := os.Create(disk.dataFile(sessionId, item))
	if e != nil {
		log.Fatal(e)
	}
	dataFile.Close()
	return nil
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
