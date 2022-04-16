require "rubygems"
require "bundler/setup"

require 'grpc'
require "protosum"


s = GRPC::RpcServer.new
s.add_http2_port('0.0.0.0:50052', :this_port_is_insecure)
s.run_till_terminated_or_interrupted([1, 'int', 'SIGQUIT'])
