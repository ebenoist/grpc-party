import party_pb2_grpc
import party_pb2
import grpc
import time

from concurrent import futures

class PartyServicer(party_pb2_grpc.PartyServicer):
    def PartyHard(self, request, context):
        msg = "python wants {0} to party hard!".format(request.name)
        return party_pb2.PartyResp(msg=msg)

server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
party_pb2_grpc.add_PartyServicer_to_server(
      PartyServicer(), server)
server.add_insecure_port('[::]:3001')
server.start()

while True:
    time.sleep(5)
