///
//  Generated code. Do not modify.
//  source: include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Gender extends $pb.ProtobufEnum {
  static const Gender GENDER_UNSPECIFIED = Gender._(0, 'GENDER_UNSPECIFIED');
  static const Gender GENDER_MALE = Gender._(1, 'GENDER_MALE');
  static const Gender GENDER_FEMALE = Gender._(2, 'GENDER_FEMALE');
  static const Gender GENDER_OTHER = Gender._(3, 'GENDER_OTHER');

  static const $core.List<Gender> values = <Gender> [
    GENDER_UNSPECIFIED,
    GENDER_MALE,
    GENDER_FEMALE,
    GENDER_OTHER,
  ];

  static final $core.Map<$core.int, Gender> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Gender valueOf($core.int value) => _byValue[value];

  const Gender._($core.int v, $core.String n) : super(v, n);
}

class AuthType extends $pb.ProtobufEnum {
  static const AuthType AUTH_TYPE_UNSPECIFIED = AuthType._(0, 'AUTH_TYPE_UNSPECIFIED');
  static const AuthType AUTH_TYPE_APPLE_ID = AuthType._(1, 'AUTH_TYPE_APPLE_ID');
  static const AuthType AUTH_TYPE_ACCOUNT = AuthType._(2, 'AUTH_TYPE_ACCOUNT');
  static const AuthType AUTH_TYPE_WECHAT = AuthType._(3, 'AUTH_TYPE_WECHAT');
  static const AuthType AUTH_TYPE_QQ = AuthType._(4, 'AUTH_TYPE_QQ');
  static const AuthType AUTH_TYPE_WEIBO = AuthType._(5, 'AUTH_TYPE_WEIBO');

  static const $core.List<AuthType> values = <AuthType> [
    AUTH_TYPE_UNSPECIFIED,
    AUTH_TYPE_APPLE_ID,
    AUTH_TYPE_ACCOUNT,
    AUTH_TYPE_WECHAT,
    AUTH_TYPE_QQ,
    AUTH_TYPE_WEIBO,
  ];

  static final $core.Map<$core.int, AuthType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static AuthType valueOf($core.int value) => _byValue[value];

  const AuthType._($core.int v, $core.String n) : super(v, n);
}

class Status extends $pb.ProtobufEnum {
  static const Status STATUS_UNSPECIFIED = Status._(0, 'STATUS_UNSPECIFIED');
  static const Status STATUS_ACTIVE = Status._(1, 'STATUS_ACTIVE');
  static const Status STATUS_BLOCKED = Status._(2, 'STATUS_BLOCKED');

  static const $core.List<Status> values = <Status> [
    STATUS_UNSPECIFIED,
    STATUS_ACTIVE,
    STATUS_BLOCKED,
  ];

  static final $core.Map<$core.int, Status> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Status valueOf($core.int value) => _byValue[value];

  const Status._($core.int v, $core.String n) : super(v, n);
}

