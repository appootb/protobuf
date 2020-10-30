///
//  Generated code. Do not modify.
//  source: group.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'google/protobuf/timestamp.pb.dart' as $5;

class GroupInfo extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('GroupInfo', package: const $pb.PackageName('appootb.relation'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(2, 'name')
    ..aOS(3, 'notice')
    ..p<$fixnum.Int64>(4, 'members', $pb.PbFieldType.PU6)
    ..aOM<$5.Timestamp>(10, 'createdAt', subBuilder: $5.Timestamp.create)
    ..aOM<$5.Timestamp>(11, 'updatedAt', subBuilder: $5.Timestamp.create)
    ..hasRequiredFields = false
  ;

  GroupInfo._() : super();
  factory GroupInfo() => create();
  factory GroupInfo.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GroupInfo.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  GroupInfo clone() => GroupInfo()..mergeFromMessage(this);
  GroupInfo copyWith(void Function(GroupInfo) updates) => super.copyWith((message) => updates(message as GroupInfo));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GroupInfo create() => GroupInfo._();
  GroupInfo createEmptyInstance() => create();
  static $pb.PbList<GroupInfo> createRepeated() => $pb.PbList<GroupInfo>();
  @$core.pragma('dart2js:noInline')
  static GroupInfo getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GroupInfo>(create);
  static GroupInfo _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get notice => $_getSZ(2);
  @$pb.TagNumber(3)
  set notice($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasNotice() => $_has(2);
  @$pb.TagNumber(3)
  void clearNotice() => clearField(3);

  @$pb.TagNumber(4)
  $core.List<$fixnum.Int64> get members => $_getList(3);

  @$pb.TagNumber(10)
  $5.Timestamp get createdAt => $_getN(4);
  @$pb.TagNumber(10)
  set createdAt($5.Timestamp v) { setField(10, v); }
  @$pb.TagNumber(10)
  $core.bool hasCreatedAt() => $_has(4);
  @$pb.TagNumber(10)
  void clearCreatedAt() => clearField(10);
  @$pb.TagNumber(10)
  $5.Timestamp ensureCreatedAt() => $_ensure(4);

  @$pb.TagNumber(11)
  $5.Timestamp get updatedAt => $_getN(5);
  @$pb.TagNumber(11)
  set updatedAt($5.Timestamp v) { setField(11, v); }
  @$pb.TagNumber(11)
  $core.bool hasUpdatedAt() => $_has(5);
  @$pb.TagNumber(11)
  void clearUpdatedAt() => clearField(11);
  @$pb.TagNumber(11)
  $5.Timestamp ensureUpdatedAt() => $_ensure(5);
}

