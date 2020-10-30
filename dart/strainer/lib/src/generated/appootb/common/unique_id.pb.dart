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

class UUID extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UUID', package: const $pb.PackageName('appootb.common'), createEmptyInstance: create)
    ..aQS(1, 'uuid')
  ;

  UUID._() : super();
  factory UUID() => create();
  factory UUID.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UUID.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UUID clone() => UUID()..mergeFromMessage(this);
  UUID copyWith(void Function(UUID) updates) => super.copyWith((message) => updates(message as UUID));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UUID create() => UUID._();
  UUID createEmptyInstance() => create();
  static $pb.PbList<UUID> createRepeated() => $pb.PbList<UUID>();
  @$core.pragma('dart2js:noInline')
  static UUID getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UUID>(create);
  static UUID _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get uuid => $_getSZ(0);
  @$pb.TagNumber(1)
  set uuid($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUuid() => $_has(0);
  @$pb.TagNumber(1)
  void clearUuid() => clearField(1);
}

class UUIDs extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('UUIDs', package: const $pb.PackageName('appootb.common'), createEmptyInstance: create)
    ..pPS(1, 'uuids')
    ..hasRequiredFields = false
  ;

  UUIDs._() : super();
  factory UUIDs() => create();
  factory UUIDs.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UUIDs.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  UUIDs clone() => UUIDs()..mergeFromMessage(this);
  UUIDs copyWith(void Function(UUIDs) updates) => super.copyWith((message) => updates(message as UUIDs));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UUIDs create() => UUIDs._();
  UUIDs createEmptyInstance() => create();
  static $pb.PbList<UUIDs> createRepeated() => $pb.PbList<UUIDs>();
  @$core.pragma('dart2js:noInline')
  static UUIDs getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UUIDs>(create);
  static UUIDs _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.String> get uuids => $_getList(0);
}

class PaginationIdRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('PaginationIdRequest', package: const $pb.PackageName('appootb.common'), createEmptyInstance: create)
    ..aOM<UniqueId>(1, 'target', subBuilder: UniqueId.create)
    ..a<$fixnum.Int64>(2, 'offset', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$core.int>(3, 'count', $pb.PbFieldType.Q3)
  ;

  PaginationIdRequest._() : super();
  factory PaginationIdRequest() => create();
  factory PaginationIdRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PaginationIdRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  PaginationIdRequest clone() => PaginationIdRequest()..mergeFromMessage(this);
  PaginationIdRequest copyWith(void Function(PaginationIdRequest) updates) => super.copyWith((message) => updates(message as PaginationIdRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PaginationIdRequest create() => PaginationIdRequest._();
  PaginationIdRequest createEmptyInstance() => create();
  static $pb.PbList<PaginationIdRequest> createRepeated() => $pb.PbList<PaginationIdRequest>();
  @$core.pragma('dart2js:noInline')
  static PaginationIdRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PaginationIdRequest>(create);
  static PaginationIdRequest _defaultInstance;

  @$pb.TagNumber(1)
  UniqueId get target => $_getN(0);
  @$pb.TagNumber(1)
  set target(UniqueId v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTarget() => $_has(0);
  @$pb.TagNumber(1)
  void clearTarget() => clearField(1);
  @$pb.TagNumber(1)
  UniqueId ensureTarget() => $_ensure(0);

  @$pb.TagNumber(2)
  $fixnum.Int64 get offset => $_getI64(1);
  @$pb.TagNumber(2)
  set offset($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasOffset() => $_has(1);
  @$pb.TagNumber(2)
  void clearOffset() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get count => $_getIZ(2);
  @$pb.TagNumber(3)
  set count($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasCount() => $_has(2);
  @$pb.TagNumber(3)
  void clearCount() => clearField(3);
}

class PaginationIdResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('PaginationIdResponse', package: const $pb.PackageName('appootb.common'), createEmptyInstance: create)
    ..a<$core.bool>(1, 'more', $pb.PbFieldType.QB)
    ..a<$fixnum.Int64>(2, 'next', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aQM<UniqueIds>(3, 'data', subBuilder: UniqueIds.create)
  ;

  PaginationIdResponse._() : super();
  factory PaginationIdResponse() => create();
  factory PaginationIdResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PaginationIdResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  PaginationIdResponse clone() => PaginationIdResponse()..mergeFromMessage(this);
  PaginationIdResponse copyWith(void Function(PaginationIdResponse) updates) => super.copyWith((message) => updates(message as PaginationIdResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PaginationIdResponse create() => PaginationIdResponse._();
  PaginationIdResponse createEmptyInstance() => create();
  static $pb.PbList<PaginationIdResponse> createRepeated() => $pb.PbList<PaginationIdResponse>();
  @$core.pragma('dart2js:noInline')
  static PaginationIdResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PaginationIdResponse>(create);
  static PaginationIdResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get more => $_getBF(0);
  @$pb.TagNumber(1)
  set more($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasMore() => $_has(0);
  @$pb.TagNumber(1)
  void clearMore() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get next => $_getI64(1);
  @$pb.TagNumber(2)
  set next($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasNext() => $_has(1);
  @$pb.TagNumber(2)
  void clearNext() => clearField(2);

  @$pb.TagNumber(3)
  UniqueIds get data => $_getN(2);
  @$pb.TagNumber(3)
  set data(UniqueIds v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasData() => $_has(2);
  @$pb.TagNumber(3)
  void clearData() => clearField(3);
  @$pb.TagNumber(3)
  UniqueIds ensureData() => $_ensure(2);
}

