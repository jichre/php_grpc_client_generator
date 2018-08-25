<?php
namespace Helloworld;
/**
 * service Helloworld{}
 * 生成 (gprc 定义 Helloworld 服务)的客户端
 */
class HelloworldClient extends \Grpc\BaseStub {

	public function __construct($hostname, $opts, $channel = null) {
		parent::__construct($hostname, $opts, $channel);
	}

	// Sends a greetin
	public function SayHello(\Helloworld\HelloRequest $argument, $metadata = [], $options = []) {
		return $this->_simpleRequest('/helloworld.Greeter/SayHello',
			$argument,
			['\helloworld\HelloReply', 'decode'],
			$metadata, $options);
	}
	
	// 放心发斯蒂芬斯蒂芬似懂非懂
	//sdgsdfsdfsdf
	//sdgsdfsdfd
	//sgsd
	public function SayMyHello(\Helloworld\HelloRequest $argument, $metadata = [], $options = []) {
		return $this->_simpleRequest('/helloworld.Greeter/SayMyHello',
			$argument,
			['\helloworld\HelloReply', 'decode'],
			$metadata, $options);
	}


}
?>

?>
