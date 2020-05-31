///
//  Generated code. Do not modify.
//  source: service.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class VisibleScope extends $pb.ProtobufEnum {
  static const VisibleScope DEFAULT_SCOPE = VisibleScope._(0, 'DEFAULT_SCOPE');
  static const VisibleScope INNER_SCOPE = VisibleScope._(100, 'INNER_SCOPE');
  static const VisibleScope ALL_SCOPES = VisibleScope._(999, 'ALL_SCOPES');

  static const $core.List<VisibleScope> values = <VisibleScope> [
    DEFAULT_SCOPE,
    INNER_SCOPE,
    ALL_SCOPES,
  ];

  static final $core.Map<$core.int, VisibleScope> _byValue = $pb.ProtobufEnum.initByValue(values);
  static VisibleScope valueOf($core.int value) => _byValue[value];

  const VisibleScope._($core.int v, $core.String n) : super(v, n);
}

