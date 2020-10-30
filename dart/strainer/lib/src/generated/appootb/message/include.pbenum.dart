///
//  Generated code. Do not modify.
//  source: appootb/message/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Category extends $pb.ProtobufEnum {
  static const Category CATEGORY_UNSPECIFIED = Category._(0, 'CATEGORY_UNSPECIFIED');
  static const Category CATEGORY_PRIVATE = Category._(1, 'CATEGORY_PRIVATE');
  static const Category CATEGORY_GROUP = Category._(2, 'CATEGORY_GROUP');
  static const Category CATEGORY_CHANNEL = Category._(3, 'CATEGORY_CHANNEL');
  static const Category CATEGORY_ROOM = Category._(4, 'CATEGORY_ROOM');

  static const $core.List<Category> values = <Category> [
    CATEGORY_UNSPECIFIED,
    CATEGORY_PRIVATE,
    CATEGORY_GROUP,
    CATEGORY_CHANNEL,
    CATEGORY_ROOM,
  ];

  static final $core.Map<$core.int, Category> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Category valueOf($core.int value) => _byValue[value];

  const Category._($core.int v, $core.String n) : super(v, n);
}

class Type extends $pb.ProtobufEnum {
  static const Type TYPE_UNSPECIFIED = Type._(0, 'TYPE_UNSPECIFIED');
  static const Type TYPE_TEXT = Type._(1, 'TYPE_TEXT');
  static const Type TYPE_AUDIO = Type._(2, 'TYPE_AUDIO');
  static const Type TYPE_STICKER = Type._(3, 'TYPE_STICKER');
  static const Type TYPE_PHOTO = Type._(4, 'TYPE_PHOTO');
  static const Type TYPE_VIDEO = Type._(5, 'TYPE_VIDEO');
  static const Type TYPE_GIF = Type._(6, 'TYPE_GIF');
  static const Type TYPE_FILE = Type._(7, 'TYPE_FILE');
  static const Type TYPE_LOCATION = Type._(8, 'TYPE_LOCATION');
  static const Type TYPE_CONTACT = Type._(9, 'TYPE_CONTACT');
  static const Type TYPE_RECALL = Type._(10, 'TYPE_RECALL');
  static const Type TYPE_READ = Type._(11, 'TYPE_READ');

  static const $core.List<Type> values = <Type> [
    TYPE_UNSPECIFIED,
    TYPE_TEXT,
    TYPE_AUDIO,
    TYPE_STICKER,
    TYPE_PHOTO,
    TYPE_VIDEO,
    TYPE_GIF,
    TYPE_FILE,
    TYPE_LOCATION,
    TYPE_CONTACT,
    TYPE_RECALL,
    TYPE_READ,
  ];

  static final $core.Map<$core.int, Type> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Type valueOf($core.int value) => _byValue[value];

  const Type._($core.int v, $core.String n) : super(v, n);
}

class PayloadType extends $pb.ProtobufEnum {
  static const PayloadType PAYLOAD_TYPE_UNSPECIFIED = PayloadType._(0, 'PAYLOAD_TYPE_UNSPECIFIED');
  static const PayloadType PAYLOAD_TYPE_SYNC = PayloadType._(1, 'PAYLOAD_TYPE_SYNC');
  static const PayloadType PAYLOAD_TYPE_MESSAGES = PayloadType._(2, 'PAYLOAD_TYPE_MESSAGES');
  static const PayloadType PAYLOAD_TYPE_RECEIPTS = PayloadType._(3, 'PAYLOAD_TYPE_RECEIPTS');

  static const $core.List<PayloadType> values = <PayloadType> [
    PAYLOAD_TYPE_UNSPECIFIED,
    PAYLOAD_TYPE_SYNC,
    PAYLOAD_TYPE_MESSAGES,
    PAYLOAD_TYPE_RECEIPTS,
  ];

  static final $core.Map<$core.int, PayloadType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static PayloadType valueOf($core.int value) => _byValue[value];

  const PayloadType._($core.int v, $core.String n) : super(v, n);
}

