#php_grpc_client_generator

A tool can generate php-grpc client from .proto(grpc)

I don't kown why protoc can not generate php-client,no one use grpc with php?


Installation
------------
go get -u github.com/jichre/php_grpc_client_generator


Run
------------
func_prefix: Will add prefix to function name if true

php_grpc_client_generator --input_dir=./proto --output_dir=./ --func_prefix=true


