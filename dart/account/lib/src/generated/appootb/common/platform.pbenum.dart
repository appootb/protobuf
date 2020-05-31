///
//  Generated code. Do not modify.
//  source: platform.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Platform extends $pb.ProtobufEnum {
  static const Platform PLATFORM_UNSPECIFIED = Platform._(0, 'PLATFORM_UNSPECIFIED');
  static const Platform PLATFORM_H5 = Platform._(1, 'PLATFORM_H5');
  static const Platform PLATFORM_WEB = Platform._(2, 'PLATFORM_WEB');
  static const Platform PLATFORM_CHROME = Platform._(3, 'PLATFORM_CHROME');
  static const Platform PLATFORM_LINUX = Platform._(4, 'PLATFORM_LINUX');
  static const Platform PLATFORM_WINDOWS = Platform._(5, 'PLATFORM_WINDOWS');
  static const Platform PLATFORM_DARWIN = Platform._(6, 'PLATFORM_DARWIN');
  static const Platform PLATFORM_ANDROID = Platform._(7, 'PLATFORM_ANDROID');
  static const Platform PLATFORM_IOS = Platform._(8, 'PLATFORM_IOS');
  static const Platform PLATFORM_SERVER = Platform._(9, 'PLATFORM_SERVER');

  static const $core.List<Platform> values = <Platform> [
    PLATFORM_UNSPECIFIED,
    PLATFORM_H5,
    PLATFORM_WEB,
    PLATFORM_CHROME,
    PLATFORM_LINUX,
    PLATFORM_WINDOWS,
    PLATFORM_DARWIN,
    PLATFORM_ANDROID,
    PLATFORM_IOS,
    PLATFORM_SERVER,
  ];

  static final $core.Map<$core.int, Platform> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Platform valueOf($core.int value) => _byValue[value];

  const Platform._($core.int v, $core.String n) : super(v, n);
}

