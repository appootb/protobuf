///
//  Generated code. Do not modify.
//  source: component.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Component extends $pb.ProtobufEnum {
  static const Component COMPONENT_UNSPECIFIED = Component._(0, 'COMPONENT_UNSPECIFIED');
  static const Component COMPONENT_CAPTCHA = Component._(1, 'COMPONENT_CAPTCHA');
  static const Component COMPONENT_ACCOUNT = Component._(2, 'COMPONENT_ACCOUNT');
  static const Component COMPONENT_STRAINER = Component._(3, 'COMPONENT_STRAINER');

  static const $core.List<Component> values = <Component> [
    COMPONENT_UNSPECIFIED,
    COMPONENT_CAPTCHA,
    COMPONENT_ACCOUNT,
    COMPONENT_STRAINER,
  ];

  static final $core.Map<$core.int, Component> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Component valueOf($core.int value) => _byValue[value];

  const Component._($core.int v, $core.String n) : super(v, n);
}

