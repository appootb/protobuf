///
//  Generated code. Do not modify.
//  source: method.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Audience extends $pb.ProtobufEnum {
  static const Audience NONE = Audience._(0, 'NONE');
  static const Audience GUEST = Audience._(1, 'GUEST');
  static const Audience WEB = Audience._(8, 'WEB');
  static const Audience PC = Audience._(64, 'PC');
  static const Audience MOBILE = Audience._(512, 'MOBILE');
  static const Audience LOGGED_IN = Audience._(584, 'LOGGED_IN');
  static const Audience CLIENT = Audience._(585, 'CLIENT');
  static const Audience SERVER = Audience._(4096, 'SERVER');
  static const Audience ANY = Audience._(65535, 'ANY');

  static const $core.List<Audience> values = <Audience> [
    NONE,
    GUEST,
    WEB,
    PC,
    MOBILE,
    LOGGED_IN,
    CLIENT,
    SERVER,
    ANY,
  ];

  static final $core.Map<$core.int, Audience> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Audience valueOf($core.int value) => _byValue[value];

  const Audience._($core.int v, $core.String n) : super(v, n);
}

