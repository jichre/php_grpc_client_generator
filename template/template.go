package template

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const (
	tStart = `<?php
namespace {namespace};
/**
 * service {namespace}{}
 * 生成 (gprc 定义 {namespace} 服务)的客户端
 */
class {namespace}Client extends \Grpc\BaseStub {

	public function __construct($hostname, $opts, $channel = null) {
		parent::__construct($hostname, $opts, $channel);
	}

`
	tEnd = `

}
?>
`
	tServericeFuc = `{rpcFuncNote}
	public function {serviceFunc}(\{namespace}\{request} $argument, $metadata = [], $options = []) {
		return $this->_simpleRequest('/{packageName}.{serviceName}/{serviceFunc}',
			$argument,
			['\{packageName}\{response}', 'decode'],
			$metadata, $options);
	}
`

	TagSpace       = "{namespace}"
	TagServiceFunc = "{serviceFunc}"
	TagPackage     = "{packageName}"
	TagServiceName = "{serviceName}"
	TagResponse    = "{response}"
	TagRequest     = "{request}"
	TagRpcNode     = "{rpcFuncNote}"
)

type GrpcTemplate struct {
	buffer bytes.Buffer
	temp   string
}

func (g *GrpcTemplate) AddStart(nameSpace string) {
	s := strings.Replace(tStart, TagSpace, nameSpace, -1)
	g.buffer.WriteString(s)
}

func (g *GrpcTemplate) SetServiceFunc() {
	g.temp = tServericeFuc
}

func (g *GrpcTemplate) Replace(value, tag string) {
	g.temp = strings.Replace(g.temp, tag, value, -1)
}

func (g *GrpcTemplate) WriteTemp() {
	g.buffer.WriteString(g.temp)
	g.temp = ""
}

func (g *GrpcTemplate) WriteToFile(fileName string) {
	g.buffer.WriteString(tEnd)

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		fmt.Println(err)
		return
	}

	file.Write(g.buffer.Bytes())

	g.buffer.Reset()
}
