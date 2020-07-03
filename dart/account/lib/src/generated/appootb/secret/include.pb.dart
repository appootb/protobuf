///
//  Generated code. Do not modify.
//  source: appootb/secret/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/timestamp.pb.dart' as $0;

import 'include.pbenum.dart';
import '../permission/method.pbenum.dart' as $2;

export 'include.pbenum.dart';

class Info extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Info', package: const $pb.PackageName('appootb.secret'), createEmptyInstance: create)
    ..e<Type>(1, 'type', $pb.PbFieldType.QE, defaultOrMaker: Type.CLIENT, valueOf: Type.valueOf, enumValues: Type.values)
    ..aQS(2, 'issuer')
    ..aQS(3, 'subject')
    ..a<$fixnum.Int64>(4, 'account', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(5, 'keyId', $pb.PbFieldType.Q6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..pPS(6, 'roles')
    ..e<$2.Audience>(11, 'audience', $pb.PbFieldType.QE, defaultOrMaker: $2.Audience.NONE, valueOf: $2.Audience.valueOf, enumValues: $2.Audience.values)
    ..aQM<$0.Timestamp>(21, 'issuedAt', subBuilder: $0.Timestamp.create)
    ..aQM<$0.Timestamp>(22, 'expiredAt', subBuilder: $0.Timestamp.create)
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
  Type get type => $_getN(0);
  @$pb.TagNumber(1)
  set type(Type v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasType() => $_has(0);
  @$pb.TagNumber(1)
  void clearType() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get issuer => $_getSZ(1);
  @$pb.TagNumber(2)
  set issuer($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasIssuer() => $_has(1);
  @$pb.TagNumber(2)
  void clearIssuer() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get subject => $_getSZ(2);
  @$pb.TagNumber(3)
  set subject($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasSubject() => $_has(2);
  @$pb.TagNumber(3)
  void clearSubject() => clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get account => $_getI64(3);
  @$pb.TagNumber(4)
  set account($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasAccount() => $_has(3);
  @$pb.TagNumber(4)
  void clearAccount() => clearField(4);

  @$pb.TagNumber(5)
  $fixnum.Int64 get keyId => $_getI64(4);
  @$pb.TagNumber(5)
  set keyId($fixnum.Int64 v) { $_setInt64(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasKeyId() => $_has(4);
  @$pb.TagNumber(5)
  void clearKeyId() => clearField(5);

  @$pb.TagNumber(6)
  $core.List<$core.String> get roles => $_getList(5);

  @$pb.TagNumber(11)
  $2.Audience get audience => $_getN(6);
  @$pb.TagNumber(11)
  set audience($2.Audience v) { setField(11, v); }
  @$pb.TagNumber(11)
  $core.bool hasAudience() => $_has(6);
  @$pb.TagNumber(11)
  void clearAudience() => clearField(11);

  @$pb.TagNumber(21)
  $0.Timestamp get issuedAt => $_getN(7);
  @$pb.TagNumber(21)
  set issuedAt($0.Timestamp v) { setField(21, v); }
  @$pb.TagNumber(21)
  $core.bool hasIssuedAt() => $_has(7);
  @$pb.TagNumber(21)
  void clearIssuedAt() => clearField(21);
  @$pb.TagNumber(21)
  $0.Timestamp ensureIssuedAt() => $_ensure(7);

  @$pb.TagNumber(22)
  $0.Timestamp get expiredAt => $_getN(8);
  @$pb.TagNumber(22)
  set expiredAt($0.Timestamp v) { setField(22, v); }
  @$pb.TagNumber(22)
  $core.bool hasExpiredAt() => $_has(8);
  @$pb.TagNumber(22)
  void clearExpiredAt() => clearField(22);
  @$pb.TagNumber(22)
  $0.Timestamp ensureExpiredAt() => $_ensure(8);
}

