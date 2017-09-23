$:.unshift File.dirname(__FILE__)

require "party_pb"
require "party_services_pb"

port = 3000 + rand(3)

puts "Ruby connecting to #{port}"
stub = Party::Stub.new("localhost:#{port}", :this_channel_is_insecure)
resp = stub.party_hard(User.new(name: "ruby"))

puts resp.msg
