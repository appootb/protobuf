///
//  Generated code. Do not modify.
//  source: appootb/account/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/timestamp.pb.dart' as $0;
import '../../google/protobuf/any.pb.dart' as $1;

import 'include.pbenum.dart';

export 'include.pbenum.dart';

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

class Info extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Info', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, 'uniqueId', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aQS(2, 'nickname')
    ..aQS(3, 'avatar')
    ..e<Status>(4, 'status', $pb.PbFieldType.QE, defaultOrMaker: Status.STATUS_UNSPECIFIED, valueOf: Status.valueOf, enumValues: Status.values)
    ..aOS(5, 'signature')
    ..e<Gender>(6, 'gender', $pb.PbFieldType.OE, defaultOrMaker: Gender.GENDER_UNSPECIFIED, valueOf: Gender.valueOf, enumValues: Gender.values)
    ..aOS(7, 'signs')
    ..aOS(8, 'location')
    ..aOM<Secret>(9, 'secret', subBuilder: Secret.create)
    ..aOM<$0.Timestamp>(98, 'createAt', subBuilder: $0.Timestamp.create)
    ..aOM<$1.Any>(99, 'extend', subBuilder: $1.Any.create)
  ;

  Info._() : super();
  factory Info() => create();
  factory Info.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Info.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Info clone() => Info()..mergeFromMessage(this);
  Info copyWith(void Function(Info) updates) => super.copyWith((message) => updates(message as Info));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Info create() => Info._();
  Info createEmptyInstance() => create();
  static $pb.PbList<Info> createRepeated() => $pb.PbList<Info>();
  @$core.pragma('dart2js:noInline')
  static Info getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Info>(create);
  static Info _defaultInstance;

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
  Status get status => $_getN(3);
  @$pb.TagNumber(4)
  set status(Status v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasStatus() => $_has(3);
  @$pb.TagNumber(4)
  void clearStatus() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get signature => $_getSZ(4);
  @$pb.TagNumber(5)
  set signature($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasSignature() => $_has(4);
  @$pb.TagNumber(5)
  void clearSignature() => clearField(5);

  @$pb.TagNumber(6)
  Gender get gender => $_getN(5);
  @$pb.TagNumber(6)
  set gender(Gender v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasGender() => $_has(5);
  @$pb.TagNumber(6)
  void clearGender() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get signs => $_getSZ(6);
  @$pb.TagNumber(7)
  set signs($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasSigns() => $_has(6);
  @$pb.TagNumber(7)
  void clearSigns() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get location => $_getSZ(7);
  @$pb.TagNumber(8)
  set location($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasLocation() => $_has(7);
  @$pb.TagNumber(8)
  void clearLocation() => clearField(8);

  @$pb.TagNumber(9)
  Secret get secret => $_getN(8);
  @$pb.TagNumber(9)
  set secret(Secret v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasSecret() => $_has(8);
  @$pb.TagNumber(9)
  void clearSecret() => clearField(9);
  @$pb.TagNumber(9)
  Secret ensureSecret() => $_ensure(8);

  @$pb.TagNumber(98)
  $0.Timestamp get createAt => $_getN(9);
  @$pb.TagNumber(98)
  set createAt($0.Timestamp v) { setField(98, v); }
  @$pb.TagNumber(98)
  $core.bool hasCreateAt() => $_has(9);
  @$pb.TagNumber(98)
  void clearCreateAt() => clearField(98);
  @$pb.TagNumber(98)
  $0.Timestamp ensureCreateAt() => $_ensure(9);

  @$pb.TagNumber(99)
  $1.Any get extend => $_getN(10);
  @$pb.TagNumber(99)
  set extend($1.Any v) { setField(99, v); }
  @$pb.TagNumber(99)
  $core.bool hasExtend() => $_has(10);
  @$pb.TagNumber(99)
  void clearExtend() => clearField(99);
  @$pb.TagNumber(99)
  $1.Any ensureExtend() => $_ensure(10);
}

class Infos extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Infos', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..m<$fixnum.Int64, Info>(1, 'accounts', entryClassName: 'Infos.AccountsEntry', keyFieldType: $pb.PbFieldType.OU6, valueFieldType: $pb.PbFieldType.OM, valueCreator: Info.create, packageName: const $pb.PackageName('appootb.account'))
  ;

  Infos._() : super();
  factory Infos() => create();
  factory Infos.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Infos.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Infos clone() => Infos()..mergeFromMessage(this);
  Infos copyWith(void Function(Infos) updates) => super.copyWith((message) => updates(message as Infos));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Infos create() => Infos._();
  Infos createEmptyInstance() => create();
  static $pb.PbList<Infos> createRepeated() => $pb.PbList<Infos>();
  @$core.pragma('dart2js:noInline')
  static Infos getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Infos>(create);
  static Infos _defaultInstance;

  @$pb.TagNumber(1)
  $core.Map<$fixnum.Int64, Info> get accounts => $_getMap(0);
}

