///
//  Generated code. Do not modify.
//  source: secret.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'google/protobuf/timestamp.pb.dart' as $0;

class Secret extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Secret', package: const $pb.PackageName('appootb.permission.account'), createEmptyInstance: create)
    ..aQS(1, 'issuer')
    ..aQS(2, 'subject')
    ..a<$fixnum.Int64>(3, 'accountId', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aQS(4, 'keyId')
    ..pPS(5, 'roles')
    ..m<$core.String, $core.String>(6, 'metadata', entryClassName: 'Secret.MetadataEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OS, packageName: const $pb.PackageName('appootb.permission.account'))
    ..aQM<$0.Timestamp>(7, 'issuedAt', subBuilder: $0.Timestamp.create)
    ..aQM<$0.Timestamp>(8, 'expiredAt', subBuilder: $0.Timestamp.create)
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
  $core.String get issuer => $_getSZ(0);
  @$pb.TagNumber(1)
  set issuer($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIssuer() => $_has(0);
  @$pb.TagNumber(1)
  void clearIssuer() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get subject => $_getSZ(1);
  @$pb.TagNumber(2)
  set subject($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSubject() => $_has(1);
  @$pb.TagNumber(2)
  void clearSubject() => clearField(2);

  @$pb.TagNumber(3)
  $fixnum.Int64 get accountId => $_getI64(2);
  @$pb.TagNumber(3)
  set accountId($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasAccountId() => $_has(2);
  @$pb.TagNumber(3)
  void clearAccountId() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get keyId => $_getSZ(3);
  @$pb.TagNumber(4)
  set keyId($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasKeyId() => $_has(3);
  @$pb.TagNumber(4)
  void clearKeyId() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$core.String> get roles => $_getList(4);

  @$pb.TagNumber(6)
  $core.Map<$core.String, $core.String> get metadata => $_getMap(5);

  @$pb.TagNumber(7)
  $0.Timestamp get issuedAt => $_getN(6);
  @$pb.TagNumber(7)
  set issuedAt($0.Timestamp v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasIssuedAt() => $_has(6);
  @$pb.TagNumber(7)
  void clearIssuedAt() => clearField(7);
  @$pb.TagNumber(7)
  $0.Timestamp ensureIssuedAt() => $_ensure(6);

  @$pb.TagNumber(8)
  $0.Timestamp get expiredAt => $_getN(7);
  @$pb.TagNumber(8)
  set expiredAt($0.Timestamp v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasExpiredAt() => $_has(7);
  @$pb.TagNumber(8)
  void clearExpiredAt() => clearField(8);
  @$pb.TagNumber(8)
  $0.Timestamp ensureExpiredAt() => $_ensure(7);
}

