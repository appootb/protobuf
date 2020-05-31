///
//  Generated code. Do not modify.
//  source: network.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Network extends $pb.ProtobufEnum {
  static const Network NETWORK_UNSPECIFIED = Network._(0, 'NETWORK_UNSPECIFIED');
  static const Network NETWORK_ETHERNET = Network._(1, 'NETWORK_ETHERNET');
  static const Network NETWORK_WIFI = Network._(2, 'NETWORK_WIFI');
  static const Network NETWORK_CELLULAR = Network._(3, 'NETWORK_CELLULAR');

  static const $core.List<Network> values = <Network> [
    NETWORK_UNSPECIFIED,
    NETWORK_ETHERNET,
    NETWORK_WIFI,
    NETWORK_CELLULAR,
  ];

  static final $core.Map<$core.int, Network> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Network valueOf($core.int value) => _byValue[value];

  const Network._($core.int v, $core.String n) : super(v, n);
}

