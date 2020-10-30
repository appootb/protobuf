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
  static const Component COMPONENT_CAPTCHA = Component._(1000, 'COMPONENT_CAPTCHA');
  static const Component COMPONENT_ACCOUNT = Component._(2000, 'COMPONENT_ACCOUNT');
  static const Component COMPONENT_STRAINER = Component._(3000, 'COMPONENT_STRAINER');
  static const Component COMPONENT_COUNTER = Component._(4000, 'COMPONENT_COUNTER');
  static const Component COMPONENT_RELATION = Component._(5000, 'COMPONENT_RELATION');
  static const Component COMPONENT_NOTIFICATION = Component._(6000, 'COMPONENT_NOTIFICATION');
  static const Component COMPONENT_MESSAGE = Component._(7000, 'COMPONENT_MESSAGE');
  static const Component COMPONENT_WALLET = Component._(8000, 'COMPONENT_WALLET');

  static const $core.List<Component> values = <Component> [
    COMPONENT_UNSPECIFIED,
    COMPONENT_CAPTCHA,
    COMPONENT_ACCOUNT,
    COMPONENT_STRAINER,
    COMPONENT_COUNTER,
    COMPONENT_RELATION,
    COMPONENT_NOTIFICATION,
    COMPONENT_MESSAGE,
    COMPONENT_WALLET,
  ];

  static final $core.Map<$core.int, Component> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Component valueOf($core.int value) => _byValue[value];

  const Component._($core.int v, $core.String n) : super(v, n);
}

class InnerService extends $pb.ProtobufEnum {
  static const InnerService INNER_SERVICE_UNSPECIFIED = InnerService._(0, 'INNER_SERVICE_UNSPECIFIED');
  static const InnerService INNER_SERVICE_CAPTCHA_CODE = InnerService._(1001, 'INNER_SERVICE_CAPTCHA_CODE');
  static const InnerService INNER_SERVICE_ACCOUNT_SECRET = InnerService._(2001, 'INNER_SERVICE_ACCOUNT_SECRET');
  static const InnerService INNER_SERVICE_MESSAGE_ROUTER = InnerService._(7001, 'INNER_SERVICE_MESSAGE_ROUTER');
  static const InnerService INNER_SERVICE_MESSAGE_SESSION = InnerService._(7002, 'INNER_SERVICE_MESSAGE_SESSION');
  static const InnerService INNER_SERVICE_MESSAGE_CHAT = InnerService._(7003, 'INNER_SERVICE_MESSAGE_CHAT');
  static const InnerService INNER_SERVICE_MESSAGE_ROOM = InnerService._(7004, 'INNER_SERVICE_MESSAGE_ROOM');

  static const $core.List<InnerService> values = <InnerService> [
    INNER_SERVICE_UNSPECIFIED,
    INNER_SERVICE_CAPTCHA_CODE,
    INNER_SERVICE_ACCOUNT_SECRET,
    INNER_SERVICE_MESSAGE_ROUTER,
    INNER_SERVICE_MESSAGE_SESSION,
    INNER_SERVICE_MESSAGE_CHAT,
    INNER_SERVICE_MESSAGE_ROOM,
  ];

  static final $core.Map<$core.int, InnerService> _byValue = $pb.ProtobufEnum.initByValue(values);
  static InnerService valueOf($core.int value) => _byValue[value];

  const InnerService._($core.int v, $core.String n) : super(v, n);
}

