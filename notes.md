Protobuf Notes:
- Numbers 1-15 for field numbers use 1 byte, while 16-2047 use 2 bytes. So 1-5 should be used for frequently occurring elements.
- Unlike proto2, where declaring a field to be optional was explicit, in proto3 all the fields are optional by default.