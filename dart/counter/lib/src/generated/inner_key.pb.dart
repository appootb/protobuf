///
//  Generated code. Do not modify.
//  source: inner_key.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'google/protobuf/duration.pb.dart' as $3;

class Key extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Key', package: const $pb.PackageName('appootb.counter'), createEmptyInstance: create)
    ..aQS(1, 'product')
    ..aQS(2, 'type')
    ..aQS(3, 'relateId')
    ..aInt64(4, 'value')
    ..aOM<$3.Duration>(5, 'expire', subBuilder: $3.Duration.create)
  ;

  Key._() : super();
  factory Key() => create();
  factory Key.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Key.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Key clone() => Key()..mergeFromMessage(this);
  Key copyWith(void Function(Key) updates) => super.copyWith((message) => updates(message as Key));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Key create() => Key._();
  Key createEmptyInstance() => create();
  static $pb.PbList<Key> createRepeated() => $pb.PbList<Key>();
  @$core.pragma('dart2js:noInline')
  static Key getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Key>(create);
  static Key _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get product => $_getSZ(0);
  @$pb.TagNumber(1)
  set product($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasProduct() => $_has(0);
  @$pb.TagNumber(1)
  void clearProduct() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get type => $_getSZ(1);
  @$pb.TagNumber(2)
  set type($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasType() => $_has(1);
  @$pb.TagNumber(2)
  void clearType() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get relateId => $_getSZ(2);
  @$pb.TagNumber(3)
  set relateId($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasRelateId() => $_has(2);
  @$pb.TagNumber(3)
  void clearRelateId() => clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get value => $_getI64(3);
  @$pb.TagNumber(4)
  set value($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasValue() => $_has(3);
  @$pb.TagNumber(4)
  void clearValue() => clearField(4);

  @$pb.TagNumber(5)
  $3.Duration get expire => $_getN(4);
  @$pb.TagNumber(5)
  set expire($3.Duration v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasExpire() => $_has(4);
  @$pb.TagNumber(5)
  void clearExpire() => clearField(5);
  @$pb.TagNumber(5)
  $3.Duration ensureExpire() => $_ensure(4);
}

class Keys extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Keys', package: const $pb.PackageName('appootb.counter'), createEmptyInstance: create)
    ..aQS(1, 'product')
    ..aQS(2, 'type')
    ..pPS(3, 'relateIds')
    ..m<$core.String, $fixnum.Int64>(4, 'values', entryClassName: 'Keys.ValuesEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.O6, packageName: const $pb.PackageName('appootb.counter'))
  ;

  Keys._() : super();
  factory Keys() => create();
  factory Keys.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Keys.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Keys clone() => Keys()..mergeFromMessage(this);
  Keys copyWith(void Function(Keys) updates) => super.copyWith((message) => updates(message as Keys));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Keys create() => Keys._();
  Keys createEmptyInstance() => create();
  static $pb.PbList<Keys> createRepeated() => $pb.PbList<Keys>();
  @$core.pragma('dart2js:noInline')
  static Keys getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Keys>(create);
  static Keys _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get product => $_getSZ(0);
  @$pb.TagNumber(1)
  set product($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasProduct() => $_has(0);
  @$pb.TagNumber(1)
  void clearProduct() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get type => $_getSZ(1);
  @$pb.TagNumber(2)
  set type($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasType() => $_has(1);
  @$pb.TagNumber(2)
  void clearType() => clearField(2);

  @$pb.TagNumber(3)
  $core.List<$core.String> get relateIds => $_getList(2);

  @$pb.TagNumber(4)
  $core.Map<$core.String, $fixnum.Int64> get values => $_getMap(3);
}

