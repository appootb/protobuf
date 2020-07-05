///
//  Generated code. Do not modify.
//  source: method.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Subject extends $pb.ProtobufEnum {
  static const Subject NONE = Subject._(0, 'NONE');
  static const Subject GUEST = Subject._(1, 'GUEST');
  static const Subject WEB = Subject._(8, 'WEB');
  static const Subject PC = Subject._(64, 'PC');
  static const Subject MOBILE = Subject._(512, 'MOBILE');
  static const Subject LOGGED_IN = Subject._(584, 'LOGGED_IN');
  static const Subject CLIENT = Subject._(585, 'CLIENT');
  static const Subject SERVER = Subject._(4096, 'SERVER');
  static const Subject ANY = Subject._(65535, 'ANY');

  static const $core.List<Subject> values = <Subject> [
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

  static final $core.Map<$core.int, Subject> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Subject valueOf($core.int value) => _byValue[value];

  const Subject._($core.int v, $core.String n) : super(v, n);
}

