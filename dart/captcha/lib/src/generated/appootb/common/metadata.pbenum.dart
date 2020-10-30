///
//  Generated code. Do not modify.
//  source: metadata.proto
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

class Platform extends $pb.ProtobufEnum {
  static const Platform PLATFORM_UNSPECIFIED = Platform._(0, 'PLATFORM_UNSPECIFIED');
  static const Platform PLATFORM_H5 = Platform._(1, 'PLATFORM_H5');
  static const Platform PLATFORM_BROWSER = Platform._(2, 'PLATFORM_BROWSER');
  static const Platform PLATFORM_CHROME = Platform._(4, 'PLATFORM_CHROME');
  static const Platform PLATFORM_WEB = Platform._(7, 'PLATFORM_WEB');
  static const Platform PLATFORM_LINUX = Platform._(16, 'PLATFORM_LINUX');
  static const Platform PLATFORM_WINDOWS = Platform._(32, 'PLATFORM_WINDOWS');
  static const Platform PLATFORM_DARWIN = Platform._(64, 'PLATFORM_DARWIN');
  static const Platform PLATFORM_PC = Platform._(112, 'PLATFORM_PC');
  static const Platform PLATFORM_ANDROID = Platform._(256, 'PLATFORM_ANDROID');
  static const Platform PLATFORM_IOS = Platform._(512, 'PLATFORM_IOS');
  static const Platform PLATFORM_MOBILE = Platform._(768, 'PLATFORM_MOBILE');
  static const Platform PLATFORM_SERVER = Platform._(4096, 'PLATFORM_SERVER');

  static const $core.List<Platform> values = <Platform> [
    PLATFORM_UNSPECIFIED,
    PLATFORM_H5,
    PLATFORM_BROWSER,
    PLATFORM_CHROME,
    PLATFORM_WEB,
    PLATFORM_LINUX,
    PLATFORM_WINDOWS,
    PLATFORM_DARWIN,
    PLATFORM_PC,
    PLATFORM_ANDROID,
    PLATFORM_IOS,
    PLATFORM_MOBILE,
    PLATFORM_SERVER,
  ];

  static final $core.Map<$core.int, Platform> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Platform valueOf($core.int value) => _byValue[value];

  const Platform._($core.int v, $core.String n) : super(v, n);
}

