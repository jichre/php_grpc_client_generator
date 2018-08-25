package analyze

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type RpcMethod struct {
	Note         string
	FunName      string
	RequestName  string
	ResponseName string
}

type RpcService struct {
	Methods     []*RpcMethod
	ServiceName string
}

type RpcPackage struct {
	Service     []*RpcService
	PackageName string
}

const (
	tagPackage = "package"
	tagService = "service"
	tagRpc     = "rpc"
)

func AnalysisProtoFile(fileName string) *RpcPackage {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0777)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	rpcPackage := &RpcPackage{Service: []*RpcService{}}

	noteLineNumStart := 0
	noteLineNumEnd := 0
	lineNum := 0
	var noteBuffer bytes.Buffer

	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}

		lineNum++

		line = strings.TrimLeft(line, " ")

		//解析出package的名字
		if len(line) > 7 && line[:7] == tagPackage {
			end := strings.Index(line, ";")
			rpcPackage.PackageName = strings.Trim(line[8:end], " ")
		}

		//解析出service的名字
		if len(line) > 7 && line[:7] == tagService {
			noteLineNumStart = lineNum + 1
			noteBuffer.Reset()
			end := strings.Index(line, "{")
			name := strings.Trim(line[8:end], " ")
			rpcPackage.Service = append(rpcPackage.Service, &RpcService{ServiceName: name})
		}

		//解析出service中的方法
		if len(line) > 3 && line[:3] == tagRpc {
			noteLineNumEnd = lineNum - 1
			end := strings.Index(line, "(")
			//得到rpc方法名字
			fname := strings.Trim(line[4:end], " ")

			//获取方法注释，并将临时参数重置
			note := ""
			if noteLineNumEnd >= noteLineNumStart {
				noteBytes := noteBuffer.Bytes()
				note = string(noteBytes[:len(noteBytes)-2])
			}

			noteLineNumStart = lineNum + 1
			noteBuffer.Reset()

			method := &RpcMethod{FunName: fname, Note: note}

			//获取request类名
			rline := line[4:len(line)]
			lindex := strings.Index(rline, "(")
			rindex := strings.Index(rline, ")")

			if lindex < rindex {
				method.RequestName = rline[lindex+1 : rindex]
			}

			//获取response类名
			rline = line[rindex+4+1 : len(line)]
			lindex = strings.Index(rline, "(")
			rindex = strings.Index(rline, ")")
			if lindex < rindex {
				method.ResponseName = rline[lindex+1 : rindex]
			}

			//将方法放入service
			lastService := rpcPackage.Service[len(rpcPackage.Service)-1]
			lastService.Methods = append(lastService.Methods, method)
		} else {
			if noteLineNumStart != 0 && noteLineNumStart <= lineNum {
				noteBuffer.WriteString("	" + line)
			}
		}
	}

	return rpcPackage
}
