///
//  Generated code. Do not modify.
//  source: appootb/strainer/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Channel extends $pb.ProtobufEnum {
  static const Channel CHANNEL_UNSPECIFIED = Channel._(0, 'CHANNEL_UNSPECIFIED');

  static const $core.List<Channel> values = <Channel> [
    CHANNEL_UNSPECIFIED,
  ];

  static final $core.Map<$core.int, Channel> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Channel valueOf($core.int value) => _byValue[value];

  const Channel._($core.int v, $core.String n) : super(v, n);
}

class Target extends $pb.ProtobufEnum {
  static const Target TARGET_UNSPECIFIED = Target._(0, 'TARGET_UNSPECIFIED');
  static const Target TARGET_NICKNAME = Target._(1, 'TARGET_NICKNAME');
  static const Target TARGET_AVATAR = Target._(2, 'TARGET_AVATAR');
  static const Target TARGET_SIGNATURE = Target._(3, 'TARGET_SIGNATURE');
  static const Target TARGET_TEXT = Target._(4, 'TARGET_TEXT');
  static const Target TARGET_IMAGE = Target._(5, 'TARGET_IMAGE');
  static const Target TARGET_VIDEO = Target._(6, 'TARGET_VIDEO');
  static const Target TARGET_AUDIO = Target._(7, 'TARGET_AUDIO');

  static const $core.List<Target> values = <Target> [
    TARGET_UNSPECIFIED,
    TARGET_NICKNAME,
    TARGET_AVATAR,
    TARGET_SIGNATURE,
    TARGET_TEXT,
    TARGET_IMAGE,
    TARGET_VIDEO,
    TARGET_AUDIO,
  ];

  static final $core.Map<$core.int, Target> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Target valueOf($core.int value) => _byValue[value];

  const Target._($core.int v, $core.String n) : super(v, n);
}

