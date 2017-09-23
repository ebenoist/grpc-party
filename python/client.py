import party_pb2_grpc
import grpc
import party_pb2

channel = grpc.insecure_channel('localhost:3001')
stub = party_pb2_grpc.PartyStub(channel)

print stub.PartyHard(party_pb2.User(name="Python"))
