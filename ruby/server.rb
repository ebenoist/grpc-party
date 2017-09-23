$:.unshift File.dirname(__FILE__)

require "party_pb"
require "party_services_pb"

class Server < Party::Service
  def party_hard(user, _call)
    PartyResp.new(msg: "ruby wants #{user.name} to party hard 🤘")
  end
end

addr = "0.0.0.0:3001"
s = GRPC::RpcServer.new
s.add_http2_port(addr, :this_port_is_insecure)
s.handle(Server.new())
s.run_till_terminated
