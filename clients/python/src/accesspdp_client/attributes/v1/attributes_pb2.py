# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: attributes/v1/attributes.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1e\x61ttributes/v1/attributes.proto\x12\rattributes.v1\"[\n\x11\x41ttributeInstance\x12\x1c\n\tauthority\x18\x01 \x01(\tR\tauthority\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n\x05value\x18\x03 \x01(\tR\x05value\"\xe5\x01\n\x13\x41ttributeDefinition\x12\x1c\n\tauthority\x18\x01 \x01(\tR\tauthority\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x12\n\x04rule\x18\x03 \x01(\tR\x04rule\x12\x19\n\x05state\x18\x04 \x01(\tH\x00R\x05state\x88\x01\x01\x12@\n\x08group_by\x18\x05 \x01(\x0b\x32 .attributes.v1.AttributeInstanceH\x01R\x07groupBy\x88\x01\x01\x12\x14\n\x05order\x18\x06 \x03(\tR\x05orderB\x08\n\x06_stateB\x0b\n\t_group_byB2Z0github.com/virtru/access-pdp/proto/attributes/v1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'attributes.v1.attributes_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z0github.com/virtru/access-pdp/proto/attributes/v1'
  _ATTRIBUTEINSTANCE._serialized_start=49
  _ATTRIBUTEINSTANCE._serialized_end=140
  _ATTRIBUTEDEFINITION._serialized_start=143
  _ATTRIBUTEDEFINITION._serialized_end=372
# @@protoc_insertion_point(module_scope)
