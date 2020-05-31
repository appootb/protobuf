///
//  Generated code. Do not modify.
//  source: unique_id.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class UniqueId extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UniqueId', package: const $pb.PackageName('appootb.common'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, 'id', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
  ;

  UniqueId._() : super();
  factory UniqueId() => create();
  factory UniqueId.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UniqueId.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UniqueId clone() => UniqueId()..mergeFromMessage(this);
  UniqueId copyWith(void Function(UniqueId) updates) => super.copyWith((message) => updates(message as UniqueId));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UniqueId create() => UniqueId._();
  UniqueId createEmptyInstance() => create();
  static $pb.PbList<UniqueId> createRepeated() => $pb.PbList<UniqueId>();
  @$core.pragma('dart2js:noInline')
  static UniqueId getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UniqueId>(create);
  static UniqueId _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class UniqueIds extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UniqueIds', package: const $pb.PackageName('appootb.common'), createEmptyInstance: create)
    ..p<$fixnum.Int64>(1, 'ids', $pb.PbFieldType.PU6)
    ..hasRequiredFields = false
  ;

  UniqueIds._() : super();
  factory UniqueIds() => create();
  factory UniqueIds.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UniqueIds.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UniqueIds clone() => UniqueIds()..mergeFromMessage(this);
  UniqueIds copyWith(void Function(UniqueIds) updates) => super.copyWith((message) => updates(message as UniqueIds));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UniqueIds create() => UniqueIds._();
  UniqueIds createEmptyInstance() => create();
  static $pb.PbList<UniqueIds> createRepeated() => $pb.PbList<UniqueIds>();
  @$core.pragma('dart2js:noInline')
  static UniqueIds getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UniqueIds>(create);
  static UniqueIds _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$fixnum.Int64> get ids => $_getList(0);
}

