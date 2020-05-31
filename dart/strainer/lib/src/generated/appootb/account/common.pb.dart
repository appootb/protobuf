///
//  Generated code. Do not modify.
//  source: appootb/account/common.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/timestamp.pb.dart' as $0;
import '../../google/protobuf/any.pb.dart' as $1;

import 'common.pbenum.dart';

export 'common.pbenum.dart';

class Secret extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Secret', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..aQS(1, 'token')
  ;

  Secret._() : super();
  factory Secret() => create();
  factory Secret.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Secret.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Secret clone() => Secret()..mergeFromMessage(this);
  Secret copyWith(void Function(Secret) updates) => super.copyWith((message) => updates(message as Secret));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Secret create() => Secret._();
  Secret createEmptyInstance() => create();
  static $pb.PbList<Secret> createRepeated() => $pb.PbList<Secret>();
  @$core.pragma('dart2js:noInline')
  static Secret getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Secret>(create);
  static Secret _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get token => $_getSZ(0);
  @$pb.TagNumber(1)
  set token($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearToken() => clearField(1);
}

class AccountInfo extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AccountInfo', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, 'uniqueId', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aQS(2, 'nickname')
    ..aQS(3, 'avatar')
    ..aOS(4, 'signature')
    ..e<Gender>(5, 'gender', $pb.PbFieldType.OE, defaultOrMaker: Gender.GENDER_UNSPECIFIED, valueOf: Gender.valueOf, enumValues: Gender.values)
    ..aOS(6, 'signs')
    ..aOS(7, 'location')
    ..aOM<Secret>(8, 'secret', subBuilder: Secret.create)
    ..aOM<$0.Timestamp>(98, 'createAt', subBuilder: $0.Timestamp.create)
    ..aOM<$1.Any>(99, 'extend', subBuilder: $1.Any.create)
  ;

  AccountInfo._() : super();
  factory AccountInfo() => create();
  factory AccountInfo.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AccountInfo.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  AccountInfo clone() => AccountInfo()..mergeFromMessage(this);
  AccountInfo copyWith(void Function(AccountInfo) updates) => super.copyWith((message) => updates(message as AccountInfo));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AccountInfo create() => AccountInfo._();
  AccountInfo createEmptyInstance() => create();
  static $pb.PbList<AccountInfo> createRepeated() => $pb.PbList<AccountInfo>();
  @$core.pragma('dart2js:noInline')
  static AccountInfo getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AccountInfo>(create);
  static AccountInfo _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get uniqueId => $_getI64(0);
  @$pb.TagNumber(1)
  set uniqueId($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUniqueId() => $_has(0);
  @$pb.TagNumber(1)
  void clearUniqueId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get nickname => $_getSZ(1);
  @$pb.TagNumber(2)
  set nickname($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasNickname() => $_has(1);
  @$pb.TagNumber(2)
  void clearNickname() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get avatar => $_getSZ(2);
  @$pb.TagNumber(3)
  set avatar($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasAvatar() => $_has(2);
  @$pb.TagNumber(3)
  void clearAvatar() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get signature => $_getSZ(3);
  @$pb.TagNumber(4)
  set signature($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasSignature() => $_has(3);
  @$pb.TagNumber(4)
  void clearSignature() => clearField(4);

  @$pb.TagNumber(5)
  Gender get gender => $_getN(4);
  @$pb.TagNumber(5)
  set gender(Gender v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasGender() => $_has(4);
  @$pb.TagNumber(5)
  void clearGender() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get signs => $_getSZ(5);
  @$pb.TagNumber(6)
  set signs($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasSigns() => $_has(5);
  @$pb.TagNumber(6)
  void clearSigns() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get location => $_getSZ(6);
  @$pb.TagNumber(7)
  set location($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasLocation() => $_has(6);
  @$pb.TagNumber(7)
  void clearLocation() => clearField(7);

  @$pb.TagNumber(8)
  Secret get secret => $_getN(7);
  @$pb.TagNumber(8)
  set secret(Secret v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasSecret() => $_has(7);
  @$pb.TagNumber(8)
  void clearSecret() => clearField(8);
  @$pb.TagNumber(8)
  Secret ensureSecret() => $_ensure(7);

  @$pb.TagNumber(98)
  $0.Timestamp get createAt => $_getN(8);
  @$pb.TagNumber(98)
  set createAt($0.Timestamp v) { setField(98, v); }
  @$pb.TagNumber(98)
  $core.bool hasCreateAt() => $_has(8);
  @$pb.TagNumber(98)
  void clearCreateAt() => clearField(98);
  @$pb.TagNumber(98)
  $0.Timestamp ensureCreateAt() => $_ensure(8);

  @$pb.TagNumber(99)
  $1.Any get extend => $_getN(9);
  @$pb.TagNumber(99)
  set extend($1.Any v) { setField(99, v); }
  @$pb.TagNumber(99)
  $core.bool hasExtend() => $_has(9);
  @$pb.TagNumber(99)
  void clearExtend() => clearField(99);
  @$pb.TagNumber(99)
  $1.Any ensureExtend() => $_ensure(9);
}

class AccountInfos extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AccountInfos', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..pc<AccountInfo>(1, 'infos', $pb.PbFieldType.PM, subBuilder: AccountInfo.create)
  ;

  AccountInfos._() : super();
  factory AccountInfos() => create();
  factory AccountInfos.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AccountInfos.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  AccountInfos clone() => AccountInfos()..mergeFromMessage(this);
  AccountInfos copyWith(void Function(AccountInfos) updates) => super.copyWith((message) => updates(message as AccountInfos));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AccountInfos create() => AccountInfos._();
  AccountInfos createEmptyInstance() => create();
  static $pb.PbList<AccountInfos> createRepeated() => $pb.PbList<AccountInfos>();
  @$core.pragma('dart2js:noInline')
  static AccountInfos getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AccountInfos>(create);
  static AccountInfos _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<AccountInfo> get infos => $_getList(0);
}

