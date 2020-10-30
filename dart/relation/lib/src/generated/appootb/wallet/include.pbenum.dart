///
//  Generated code. Do not modify.
//  source: appootb/wallet/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class AccountType extends $pb.ProtobufEnum {
  static const AccountType ACCOUNT_TYPE_UNSPECIFIED = AccountType._(0, 'ACCOUNT_TYPE_UNSPECIFIED');
  static const AccountType ACCOUNT_TYPE_NORMAL = AccountType._(1, 'ACCOUNT_TYPE_NORMAL');
  static const AccountType ACCOUNT_TYPE_SYSTEM = AccountType._(2, 'ACCOUNT_TYPE_SYSTEM');

  static const $core.List<AccountType> values = <AccountType> [
    ACCOUNT_TYPE_UNSPECIFIED,
    ACCOUNT_TYPE_NORMAL,
    ACCOUNT_TYPE_SYSTEM,
  ];

  static final $core.Map<$core.int, AccountType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static AccountType valueOf($core.int value) => _byValue[value];

  const AccountType._($core.int v, $core.String n) : super(v, n);
}

class ActiveStatus extends $pb.ProtobufEnum {
  static const ActiveStatus ACTIVE_STATUS_UNSPECIFIED = ActiveStatus._(0, 'ACTIVE_STATUS_UNSPECIFIED');

  static const $core.List<ActiveStatus> values = <ActiveStatus> [
    ACTIVE_STATUS_UNSPECIFIED,
  ];

  static final $core.Map<$core.int, ActiveStatus> _byValue = $pb.ProtobufEnum.initByValue(values);
  static ActiveStatus valueOf($core.int value) => _byValue[value];

  const ActiveStatus._($core.int v, $core.String n) : super(v, n);
}

class Channel extends $pb.ProtobufEnum {
  static const Channel CHANNEL_UNSPECIFIED = Channel._(0, 'CHANNEL_UNSPECIFIED');
  static const Channel CHANNEL_WECHAT = Channel._(1, 'CHANNEL_WECHAT');
  static const Channel CHANNEL_ALIPAY = Channel._(2, 'CHANNEL_ALIPAY');
  static const Channel CHANNEL_IN_APP_APPLE = Channel._(3, 'CHANNEL_IN_APP_APPLE');
  static const Channel CHANNEL_IN_APP_GOOGLE = Channel._(4, 'CHANNEL_IN_APP_GOOGLE');

  static const $core.List<Channel> values = <Channel> [
    CHANNEL_UNSPECIFIED,
    CHANNEL_WECHAT,
    CHANNEL_ALIPAY,
    CHANNEL_IN_APP_APPLE,
    CHANNEL_IN_APP_GOOGLE,
  ];

  static final $core.Map<$core.int, Channel> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Channel valueOf($core.int value) => _byValue[value];

  const Channel._($core.int v, $core.String n) : super(v, n);
}

class OrderStatus extends $pb.ProtobufEnum {
  static const OrderStatus ORDER_STATUS_UNSPECIFIED = OrderStatus._(0, 'ORDER_STATUS_UNSPECIFIED');

  static const $core.List<OrderStatus> values = <OrderStatus> [
    ORDER_STATUS_UNSPECIFIED,
  ];

  static final $core.Map<$core.int, OrderStatus> _byValue = $pb.ProtobufEnum.initByValue(values);
  static OrderStatus valueOf($core.int value) => _byValue[value];

  const OrderStatus._($core.int v, $core.String n) : super(v, n);
}

class SubjectType extends $pb.ProtobufEnum {
  static const SubjectType SUBJECT_TYPE_UNSPECIFIED = SubjectType._(0, 'SUBJECT_TYPE_UNSPECIFIED');

  static const $core.List<SubjectType> values = <SubjectType> [
    SUBJECT_TYPE_UNSPECIFIED,
  ];

  static final $core.Map<$core.int, SubjectType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static SubjectType valueOf($core.int value) => _byValue[value];

  const SubjectType._($core.int v, $core.String n) : super(v, n);
}

class ActionType extends $pb.ProtobufEnum {
  static const ActionType ACTION_TYPE_UNSPECIFIED = ActionType._(0, 'ACTION_TYPE_UNSPECIFIED');

  static const $core.List<ActionType> values = <ActionType> [
    ACTION_TYPE_UNSPECIFIED,
  ];

  static final $core.Map<$core.int, ActionType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static ActionType valueOf($core.int value) => _byValue[value];

  const ActionType._($core.int v, $core.String n) : super(v, n);
}

class JournalType extends $pb.ProtobufEnum {
  static const JournalType JOURNAL_TYPE_UNSPECIFIED = JournalType._(0, 'JOURNAL_TYPE_UNSPECIFIED');

  static const $core.List<JournalType> values = <JournalType> [
    JOURNAL_TYPE_UNSPECIFIED,
  ];

  static final $core.Map<$core.int, JournalType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static JournalType valueOf($core.int value) => _byValue[value];

  const JournalType._($core.int v, $core.String n) : super(v, n);
}

