import party_pb2_grpc
import grpc
import party_pb2
import random

port = 3000 + random.randint(0,2)
print 'python connecting to {0}'.format(port)
channel = grpc.insecure_channel('localhost:{0}'.format(port))
stub = party_pb2_grpc.PartyStub(channel)

print stub.PartyHard(party_pb2.User(name="Python")).msg
