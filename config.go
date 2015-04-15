package config

import (
    "os"
    "bufio"
    "strings"
    "fmt"
)

func Unmarshal(filePath string)(jsonStr string,err error){
    return unmarshal(filePath)
}

func unmarshal(filePath string) (jsonStr string,err error) {
    //configMap, err := readFile("/Users/heguangxiu/go/src/config/test.conf")
    jsonStr, err = readFile(filePath)
    if err != nil{
        return
    }
    return
}

func readFile(filePath string)(jsonStr string ,err error){
    f, err := os.Open(filePath)
    if err != nil{
        println("os open err", err)
        return
    }
    defer f.Close()
    tmpJson := []string{}
    scanner := bufio.NewScanner(f)
    for scanner.Scan(){
        str := scanner.Text()
        if strings.HasPrefix(str, "#") {
            //fmt.Printf("注释 %s \n", str)
            continue
        }
        s := strings.Split(scanner.Text(),":")
        tmp:=fmt.Sprintf("\"%s\":%s",s[0],strings.Join(s[1:],":"))
        tmpJson = append(tmpJson, tmp)
    }
    jsonStr = fmt.Sprintf("{%s}",strings.Join(tmpJson, ","))
    return
}