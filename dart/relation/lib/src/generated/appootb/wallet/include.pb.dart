///
//  Generated code. Do not modify.
//  source: appootb/wallet/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

export 'include.pbenum.dart';

class WalletId extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('WalletId', package: const $pb.PackageName('appootb.wallet'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, 'account', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(2, 'currency', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
  ;

  WalletId._() : super();
  factory WalletId() => create();
  factory WalletId.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory WalletId.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  WalletId clone() => WalletId()..mergeFromMessage(this);
  WalletId copyWith(void Function(WalletId) updates) => super.copyWith((message) => updates(message as WalletId));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static WalletId create() => WalletId._();
  WalletId createEmptyInstance() => create();
  static $pb.PbList<WalletId> createRepeated() => $pb.PbList<WalletId>();
  @$core.pragma('dart2js:noInline')
  static WalletId getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<WalletId>(create);
  static WalletId _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get account => $_getI64(0);
  @$pb.TagNumber(1)
  set account($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAccount() => $_has(0);
  @$pb.TagNumber(1)
  void clearAccount() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get currency => $_getI64(1);
  @$pb.TagNumber(2)
  set currency($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCurrency() => $_has(1);
  @$pb.TagNumber(2)
  void clearCurrency() => clearField(2);
}

