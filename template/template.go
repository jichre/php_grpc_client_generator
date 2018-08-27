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
 * Generate (gprc {namespace}) client
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
	public function {prefix}{serviceFunc}(\{namespace}\{request} $argument, $metadata = [], $options = []) {
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
	TageFuncPreifx = "{prefix}"
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

func (g *GrpcTemplate) WriteServiceFunc() {
	g.buffer.WriteString(g.temp)
	g.temp = ""
}

func (g *GrpcTemplate) WriteToFile(fileName string) {
	g.buffer.WriteString(tEnd)
	// fmt.Println(g.buffer.String())

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	file.Write(g.buffer.Bytes())

	g.buffer.Reset()
}
