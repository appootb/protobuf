///
//  Generated code. Do not modify.
//  source: bind.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'common.pbenum.dart' as $1;

class BindInfo extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('BindInfo', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..e<$1.AuthType>(1, 'sourceType', $pb.PbFieldType.QE, defaultOrMaker: $1.AuthType.AUTH_TYPE_UNSPECIFIED, valueOf: $1.AuthType.valueOf, enumValues: $1.AuthType.values)
    ..aQS(2, 'sourceId')
  ;

  BindInfo._() : super();
  factory BindInfo() => create();
  factory BindInfo.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BindInfo.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  BindInfo clone() => BindInfo()..mergeFromMessage(this);
  BindInfo copyWith(void Function(BindInfo) updates) => super.copyWith((message) => updates(message as BindInfo));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BindInfo create() => BindInfo._();
  BindInfo createEmptyInstance() => create();
  static $pb.PbList<BindInfo> createRepeated() => $pb.PbList<BindInfo>();
  @$core.pragma('dart2js:noInline')
  static BindInfo getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BindInfo>(create);
  static BindInfo _defaultInstance;

  @$pb.TagNumber(1)
  $1.AuthType get sourceType => $_getN(0);
  @$pb.TagNumber(1)
  set sourceType($1.AuthType v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasSourceType() => $_has(0);
  @$pb.TagNumber(1)
  void clearSourceType() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get sourceId => $_getSZ(1);
  @$pb.TagNumber(2)
  set sourceId($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSourceId() => $_has(1);
  @$pb.TagNumber(2)
  void clearSourceId() => clearField(2);
}

class BindRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('BindRequest', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..aQM<BindInfo>(1, 'bind', subBuilder: BindInfo.create)
    ..aOS(2, 'code')
  ;

  BindRequest._() : super();
  factory BindRequest() => create();
  factory BindRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BindRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  BindRequest clone() => BindRequest()..mergeFromMessage(this);
  BindRequest copyWith(void Function(BindRequest) updates) => super.copyWith((message) => updates(message as BindRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BindRequest create() => BindRequest._();
  BindRequest createEmptyInstance() => create();
  static $pb.PbList<BindRequest> createRepeated() => $pb.PbList<BindRequest>();
  @$core.pragma('dart2js:noInline')
  static BindRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BindRequest>(create);
  static BindRequest _defaultInstance;

  @$pb.TagNumber(1)
  BindInfo get bind => $_getN(0);
  @$pb.TagNumber(1)
  set bind(BindInfo v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasBind() => $_has(0);
  @$pb.TagNumber(1)
  void clearBind() => clearField(1);
  @$pb.TagNumber(1)
  BindInfo ensureBind() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);
}

class AccountBinds extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AccountBinds', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..pc<BindInfo>(1, 'binds', $pb.PbFieldType.PM, subBuilder: BindInfo.create)
  ;

  AccountBinds._() : super();
  factory AccountBinds() => create();
  factory AccountBinds.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AccountBinds.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  AccountBinds clone() => AccountBinds()..mergeFromMessage(this);
  AccountBinds copyWith(void Function(AccountBinds) updates) => super.copyWith((message) => updates(message as AccountBinds));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AccountBinds create() => AccountBinds._();
  AccountBinds createEmptyInstance() => create();
  static $pb.PbList<AccountBinds> createRepeated() => $pb.PbList<AccountBinds>();
  @$core.pragma('dart2js:noInline')
  static AccountBinds getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AccountBinds>(create);
  static AccountBinds _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<BindInfo> get binds => $_getList(0);
}

