///
//  Generated code. Do not modify.
//  source: appootb/secret/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Type extends $pb.ProtobufEnum {
  static const Type CLIENT = Type._(0, 'CLIENT');
  static const Type SERVER = Type._(1, 'SERVER');

  static const $core.List<Type> values = <Type> [
    CLIENT,
    SERVER,
  ];

  static final $core.Map<$core.int, Type> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Type valueOf($core.int value) => _byValue[value];

  const Type._($core.int v, $core.String n) : super(v, n);
}

class Algorithm extends $pb.ProtobufEnum {
  static const Algorithm None = Algorithm._(0, 'None');
  static const Algorithm HMAC = Algorithm._(1, 'HMAC');
  static const Algorithm RSA = Algorithm._(2, 'RSA');
  static const Algorithm PSS = Algorithm._(3, 'PSS');
  static const Algorithm ECDSA = Algorithm._(4, 'ECDSA');
  static const Algorithm EdDSA = Algorithm._(5, 'EdDSA');

  static const $core.List<Algorithm> values = <Algorithm> [
    None,
    HMAC,
    RSA,
    PSS,
    ECDSA,
    EdDSA,
  ];

  static final $core.Map<$core.int, Algorithm> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Algorithm valueOf($core.int value) => _byValue[value];

  const Algorithm._($core.int v, $core.String n) : super(v, n);
}

