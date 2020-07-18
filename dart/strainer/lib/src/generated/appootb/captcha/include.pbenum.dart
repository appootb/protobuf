///
//  Generated code. Do not modify.
//  source: appootb/captcha/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Channel extends $pb.ProtobufEnum {
  static const Channel CHANNEL_UNSPECIFIED = Channel._(0, 'CHANNEL_UNSPECIFIED');
  static const Channel CHANNEL_EMAIL = Channel._(1, 'CHANNEL_EMAIL');
  static const Channel CHANNEL_SMS = Channel._(2, 'CHANNEL_SMS');
  static const Channel CHANNEL_PHONE = Channel._(3, 'CHANNEL_PHONE');
  static const Channel CHANNEL_OTP = Channel._(4, 'CHANNEL_OTP');

  static const $core.List<Channel> values = <Channel> [
    CHANNEL_UNSPECIFIED,
    CHANNEL_EMAIL,
    CHANNEL_SMS,
    CHANNEL_PHONE,
    CHANNEL_OTP,
  ];

  static final $core.Map<$core.int, Channel> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Channel valueOf($core.int value) => _byValue[value];

  const Channel._($core.int v, $core.String n) : super(v, n);
}

class Category extends $pb.ProtobufEnum {
  static const Category CATEGORY_UNSPECIFIED = Category._(0, 'CATEGORY_UNSPECIFIED');
  static const Category CATEGORY_REGISTER = Category._(1, 'CATEGORY_REGISTER');
  static const Category CATEGORY_LOGIN = Category._(2, 'CATEGORY_LOGIN');
  static const Category CATEGORY_RESET_PWD = Category._(3, 'CATEGORY_RESET_PWD');
  static const Category CATEGORY_BIND = Category._(4, 'CATEGORY_BIND');
  static const Category CATEGORY_UNBIND = Category._(5, 'CATEGORY_UNBIND');

  static const $core.List<Category> values = <Category> [
    CATEGORY_UNSPECIFIED,
    CATEGORY_REGISTER,
    CATEGORY_LOGIN,
    CATEGORY_RESET_PWD,
    CATEGORY_BIND,
    CATEGORY_UNBIND,
  ];

  static final $core.Map<$core.int, Category> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Category valueOf($core.int value) => _byValue[value];

  const Category._($core.int v, $core.String n) : super(v, n);
}

