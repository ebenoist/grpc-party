$:.unshift File.dirname(__FILE__)

require "party_pb"
require "party_services_pb"

stub = Party::Stub.new('localhost:50051', :this_channel_is_insecure)
resp = stub.party_hard(User.new(name: "ruby"))

puts resp.msg
