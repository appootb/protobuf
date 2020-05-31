///
//  Generated code. Do not modify.
//  source: inner_secret.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'appootb/permission/method.pbenum.dart' as $12;

class SecretInfo extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('SecretInfo', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..aQS(1, 'accountId')
    ..aQS(2, 'deviceId')
    ..e<$12.TokenLevel>(9, 'level', $pb.PbFieldType.QE, defaultOrMaker: $12.TokenLevel.NONE_TOKEN, valueOf: $12.TokenLevel.valueOf, enumValues: $12.TokenLevel.values)
  ;

  SecretInfo._() : super();
  factory SecretInfo() => create();
  factory SecretInfo.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SecretInfo.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  SecretInfo clone() => SecretInfo()..mergeFromMessage(this);
  SecretInfo copyWith(void Function(SecretInfo) updates) => super.copyWith((message) => updates(message as SecretInfo));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SecretInfo create() => SecretInfo._();
  SecretInfo createEmptyInstance() => create();
  static $pb.PbList<SecretInfo> createRepeated() => $pb.PbList<SecretInfo>();
  @$core.pragma('dart2js:noInline')
  static SecretInfo getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SecretInfo>(create);
  static SecretInfo _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get accountId => $_getSZ(0);
  @$pb.TagNumber(1)
  set accountId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccountId() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccountId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get deviceId => $_getSZ(1);
  @$pb.TagNumber(2)
  set deviceId($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDeviceId() => $_has(1);
  @$pb.TagNumber(2)
  void clearDeviceId() => clearField(2);

  @$pb.TagNumber(9)
  $12.TokenLevel get level => $_getN(2);
  @$pb.TagNumber(9)
  set level($12.TokenLevel v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasLevel() => $_has(2);
  @$pb.TagNumber(9)
  void clearLevel() => clearField(9);
}

