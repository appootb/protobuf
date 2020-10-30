///
//  Generated code. Do not modify.
//  source: appootb/relation/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class ApplyRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ApplyRequest', package: const $pb.PackageName('appootb.relation'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, 'target', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(2, 'reason')
  ;

  ApplyRequest._() : super();
  factory ApplyRequest() => create();
  factory ApplyRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ApplyRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ApplyRequest clone() => ApplyRequest()..mergeFromMessage(this);
  ApplyRequest copyWith(void Function(ApplyRequest) updates) => super.copyWith((message) => updates(message as ApplyRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ApplyRequest create() => ApplyRequest._();
  ApplyRequest createEmptyInstance() => create();
  static $pb.PbList<ApplyRequest> createRepeated() => $pb.PbList<ApplyRequest>();
  @$core.pragma('dart2js:noInline')
  static ApplyRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ApplyRequest>(create);
  static ApplyRequest _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get target => $_getI64(0);
  @$pb.TagNumber(1)
  set target($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTarget() => $_has(0);
  @$pb.TagNumber(1)
  void clearTarget() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get reason => $_getSZ(1);
  @$pb.TagNumber(2)
  set reason($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasReason() => $_has(1);
  @$pb.TagNumber(2)
  void clearReason() => clearField(2);
}

class StatusRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('StatusRequest', package: const $pb.PackageName('appootb.relation'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, 'target', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..p<$fixnum.Int64>(2, 'ids', $pb.PbFieldType.PU6)
  ;

  StatusRequest._() : super();
  factory StatusRequest() => create();
  factory StatusRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StatusRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  StatusRequest clone() => StatusRequest()..mergeFromMessage(this);
  StatusRequest copyWith(void Function(StatusRequest) updates) => super.copyWith((message) => updates(message as StatusRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StatusRequest create() => StatusRequest._();
  StatusRequest createEmptyInstance() => create();
  static $pb.PbList<StatusRequest> createRepeated() => $pb.PbList<StatusRequest>();
  @$core.pragma('dart2js:noInline')
  static StatusRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StatusRequest>(create);
  static StatusRequest _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get target => $_getI64(0);
  @$pb.TagNumber(1)
  set target($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTarget() => $_has(0);
  @$pb.TagNumber(1)
  void clearTarget() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$fixnum.Int64> get ids => $_getList(1);
}

class StatusResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('StatusResponse', package: const $pb.PackageName('appootb.relation'), createEmptyInstance: create)
    ..m<$fixnum.Int64, $core.bool>(1, 'state', entryClassName: 'StatusResponse.StateEntry', keyFieldType: $pb.PbFieldType.OU6, valueFieldType: $pb.PbFieldType.OB, packageName: const $pb.PackageName('appootb.relation'))
    ..hasRequiredFields = false
  ;

  StatusResponse._() : super();
  factory StatusResponse() => create();
  factory StatusResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory StatusResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  StatusResponse clone() => StatusResponse()..mergeFromMessage(this);
  StatusResponse copyWith(void Function(StatusResponse) updates) => super.copyWith((message) => updates(message as StatusResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static StatusResponse create() => StatusResponse._();
  StatusResponse createEmptyInstance() => create();
  static $pb.PbList<StatusResponse> createRepeated() => $pb.PbList<StatusResponse>();
  @$core.pragma('dart2js:noInline')
  static StatusResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<StatusResponse>(create);
  static StatusResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.Map<$fixnum.Int64, $core.bool> get state => $_getMap(0);
}

