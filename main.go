package main

import (
	"flag"
	"fmt"
	"github.com/jichre/php_grpc_client_generator/analyze"
	"github.com/jichre/php_grpc_client_generator/template"
	"os"
	"path/filepath"
)

var (
	inputDir    = flag.String("input_dir", "", "General inputDir")
	outputDir   = flag.String("output_dir", "", "General outputDir")
	func_prefix = flag.String("func_prefix", "false", "General funcPrefix")
)

func main() {
	flag.Parse()

	if *inputDir == "" || *outputDir == "" {
		fmt.Println("input_dir or output_dir is empty!!!")
	}

	tl := &template.GrpcTemplate{}

	filepath.Walk(*inputDir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		result := analyze.AnalysisProtoFile(path)
		upName := analyze.GetStringFirstUp(result.PackageName)

		tl.AddStart(upName)
		for _, service := range result.Service {
			prefix := ""
			if *func_prefix == "true" {
				prefix = analyze.GetStringFirstUp(service.ServiceName)
			}
			for _, method := range service.Methods {
				tl.SetServiceFunc()
				tl.Replace(upName, template.TagSpace)
				tl.Replace(result.PackageName, template.TagPackage)
				tl.Replace(service.ServiceName, template.TagServiceName)
				tl.Replace(method.Note, template.TagRpcNode)
				tl.Replace(prefix, template.TageFuncPreifx)
				tl.Replace(method.FunName, template.TagServiceFunc)
				tl.Replace(method.ResponseName, template.TagResponse)
				tl.Replace(method.RequestName, template.TagRequest)
				tl.WriteServiceFunc()
			}
		}
		outDir := *outputDir
		if '/' == outDir[len(outDir)-1] {
			tl.WriteToFile(outDir + upName + "Client.php")
		} else {
			tl.WriteToFile(outDir + "/" + upName + "Client.php")
		}

		fmt.Println(outDir + upName + "Client.php")
		return nil
	})
}
