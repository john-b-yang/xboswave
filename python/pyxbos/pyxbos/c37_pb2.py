# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: c37.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='c37.proto',
  package='xbospb',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\tc37.proto\x12\x06xbospb\"\x91\x01\n\x0c\x43\x33\x37\x44\x61taFrame\x12\x13\n\x0bstationName\x18\x01 \x01(\t\x12\x0e\n\x06idCode\x18\x02 \x01(\r\x12-\n\x0ephasorChannels\x18\x03 \x03(\x0b\x32\x15.xbospb.PhasorChannel\x12-\n\x0escalarChannels\x18\x04 \x03(\x0b\x32\x15.xbospb.ScalarChannel\"P\n\rPhasorChannel\x12\x13\n\x0b\x63hannelName\x18\x01 \x01(\t\x12\x0c\n\x04unit\x18\x02 \x01(\t\x12\x1c\n\x04\x64\x61ta\x18\x03 \x03(\x0b\x32\x0e.xbospb.Phasor\"N\n\x06Phasor\x12\x0c\n\x04time\x18\x01 \x01(\x03\x12\r\n\x05\x61ngle\x18\x02 \x01(\x01\x12\x11\n\tmagnitude\x18\x03 \x01(\x01\x12\t\n\x01P\x18\x04 \x01(\x01\x12\t\n\x01Q\x18\x05 \x01(\x01\"P\n\rScalarChannel\x12\x13\n\x0b\x63hannelName\x18\x01 \x01(\t\x12\x0c\n\x04unit\x18\x02 \x01(\t\x12\x1c\n\x04\x64\x61ta\x18\x03 \x03(\x0b\x32\x0e.xbospb.Scalar\"%\n\x06Scalar\x12\x0c\n\x04time\x18\x01 \x01(\x03\x12\r\n\x05value\x18\x02 \x01(\x01\x62\x06proto3')
)




_C37DATAFRAME = _descriptor.Descriptor(
  name='C37DataFrame',
  full_name='xbospb.C37DataFrame',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='stationName', full_name='xbospb.C37DataFrame.stationName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='idCode', full_name='xbospb.C37DataFrame.idCode', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phasorChannels', full_name='xbospb.C37DataFrame.phasorChannels', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='scalarChannels', full_name='xbospb.C37DataFrame.scalarChannels', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=22,
  serialized_end=167,
)


_PHASORCHANNEL = _descriptor.Descriptor(
  name='PhasorChannel',
  full_name='xbospb.PhasorChannel',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='channelName', full_name='xbospb.PhasorChannel.channelName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='unit', full_name='xbospb.PhasorChannel.unit', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data', full_name='xbospb.PhasorChannel.data', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=169,
  serialized_end=249,
)


_PHASOR = _descriptor.Descriptor(
  name='Phasor',
  full_name='xbospb.Phasor',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='time', full_name='xbospb.Phasor.time', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='angle', full_name='xbospb.Phasor.angle', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='magnitude', full_name='xbospb.Phasor.magnitude', index=2,
      number=3, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='P', full_name='xbospb.Phasor.P', index=3,
      number=4, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Q', full_name='xbospb.Phasor.Q', index=4,
      number=5, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=251,
  serialized_end=329,
)


_SCALARCHANNEL = _descriptor.Descriptor(
  name='ScalarChannel',
  full_name='xbospb.ScalarChannel',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='channelName', full_name='xbospb.ScalarChannel.channelName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='unit', full_name='xbospb.ScalarChannel.unit', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data', full_name='xbospb.ScalarChannel.data', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=331,
  serialized_end=411,
)


_SCALAR = _descriptor.Descriptor(
  name='Scalar',
  full_name='xbospb.Scalar',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='time', full_name='xbospb.Scalar.time', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='xbospb.Scalar.value', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=413,
  serialized_end=450,
)

_C37DATAFRAME.fields_by_name['phasorChannels'].message_type = _PHASORCHANNEL
_C37DATAFRAME.fields_by_name['scalarChannels'].message_type = _SCALARCHANNEL
_PHASORCHANNEL.fields_by_name['data'].message_type = _PHASOR
_SCALARCHANNEL.fields_by_name['data'].message_type = _SCALAR
DESCRIPTOR.message_types_by_name['C37DataFrame'] = _C37DATAFRAME
DESCRIPTOR.message_types_by_name['PhasorChannel'] = _PHASORCHANNEL
DESCRIPTOR.message_types_by_name['Phasor'] = _PHASOR
DESCRIPTOR.message_types_by_name['ScalarChannel'] = _SCALARCHANNEL
DESCRIPTOR.message_types_by_name['Scalar'] = _SCALAR
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

C37DataFrame = _reflection.GeneratedProtocolMessageType('C37DataFrame', (_message.Message,), dict(
  DESCRIPTOR = _C37DATAFRAME,
  __module__ = 'c37_pb2'
  # @@protoc_insertion_point(class_scope:xbospb.C37DataFrame)
  ))
_sym_db.RegisterMessage(C37DataFrame)

PhasorChannel = _reflection.GeneratedProtocolMessageType('PhasorChannel', (_message.Message,), dict(
  DESCRIPTOR = _PHASORCHANNEL,
  __module__ = 'c37_pb2'
  # @@protoc_insertion_point(class_scope:xbospb.PhasorChannel)
  ))
_sym_db.RegisterMessage(PhasorChannel)

Phasor = _reflection.GeneratedProtocolMessageType('Phasor', (_message.Message,), dict(
  DESCRIPTOR = _PHASOR,
  __module__ = 'c37_pb2'
  # @@protoc_insertion_point(class_scope:xbospb.Phasor)
  ))
_sym_db.RegisterMessage(Phasor)

ScalarChannel = _reflection.GeneratedProtocolMessageType('ScalarChannel', (_message.Message,), dict(
  DESCRIPTOR = _SCALARCHANNEL,
  __module__ = 'c37_pb2'
  # @@protoc_insertion_point(class_scope:xbospb.ScalarChannel)
  ))
_sym_db.RegisterMessage(ScalarChannel)

Scalar = _reflection.GeneratedProtocolMessageType('Scalar', (_message.Message,), dict(
  DESCRIPTOR = _SCALAR,
  __module__ = 'c37_pb2'
  # @@protoc_insertion_point(class_scope:xbospb.Scalar)
  ))
_sym_db.RegisterMessage(Scalar)


# @@protoc_insertion_point(module_scope)
