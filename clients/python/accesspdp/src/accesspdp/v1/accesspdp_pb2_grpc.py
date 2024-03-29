# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from accesspdp.v1 import accesspdp_pb2 as accesspdp_dot_v1_dot_accesspdp__pb2


class HealthStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Check = channel.unary_unary(
                '/accesspdp.v1.Health/Check',
                request_serializer=accesspdp_dot_v1_dot_accesspdp__pb2.HealthCheckRequest.SerializeToString,
                response_deserializer=accesspdp_dot_v1_dot_accesspdp__pb2.HealthCheckResponse.FromString,
                )


class HealthServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Check(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_HealthServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Check': grpc.unary_unary_rpc_method_handler(
                    servicer.Check,
                    request_deserializer=accesspdp_dot_v1_dot_accesspdp__pb2.HealthCheckRequest.FromString,
                    response_serializer=accesspdp_dot_v1_dot_accesspdp__pb2.HealthCheckResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'accesspdp.v1.Health', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Health(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Check(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/accesspdp.v1.Health/Check',
            accesspdp_dot_v1_dot_accesspdp__pb2.HealthCheckRequest.SerializeToString,
            accesspdp_dot_v1_dot_accesspdp__pb2.HealthCheckResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class AccessPDPEndpointStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.DetermineAccess = channel.unary_stream(
                '/accesspdp.v1.AccessPDPEndpoint/DetermineAccess',
                request_serializer=accesspdp_dot_v1_dot_accesspdp__pb2.DetermineAccessRequest.SerializeToString,
                response_deserializer=accesspdp_dot_v1_dot_accesspdp__pb2.DetermineAccessResponse.FromString,
                )


class AccessPDPEndpointServicer(object):
    """Missing associated documentation comment in .proto file."""

    def DetermineAccess(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AccessPDPEndpointServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'DetermineAccess': grpc.unary_stream_rpc_method_handler(
                    servicer.DetermineAccess,
                    request_deserializer=accesspdp_dot_v1_dot_accesspdp__pb2.DetermineAccessRequest.FromString,
                    response_serializer=accesspdp_dot_v1_dot_accesspdp__pb2.DetermineAccessResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'accesspdp.v1.AccessPDPEndpoint', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AccessPDPEndpoint(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def DetermineAccess(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/accesspdp.v1.AccessPDPEndpoint/DetermineAccess',
            accesspdp_dot_v1_dot_accesspdp__pb2.DetermineAccessRequest.SerializeToString,
            accesspdp_dot_v1_dot_accesspdp__pb2.DetermineAccessResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
