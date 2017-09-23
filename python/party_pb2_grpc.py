# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import party_pb2 as party__pb2


class PartyStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.PartyHard = channel.unary_unary(
        '/Party/PartyHard',
        request_serializer=party__pb2.User.SerializeToString,
        response_deserializer=party__pb2.PartyResp.FromString,
        )


class PartyServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def PartyHard(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_PartyServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'PartyHard': grpc.unary_unary_rpc_method_handler(
          servicer.PartyHard,
          request_deserializer=party__pb2.User.FromString,
          response_serializer=party__pb2.PartyResp.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'Party', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))