///
//  Generated code. Do not modify.
//  source: profile.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'google/protobuf/struct.pb.dart' as $12;

class Property extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Property', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..aQS(1, 'name')
    ..aOM<$12.Struct>(2, 'value', subBuilder: $12.Struct.create)
  ;

  Property._() : super();
  factory Property() => create();
  factory Property.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Property.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Property clone() => Property()..mergeFromMessage(this);
  Property copyWith(void Function(Property) updates) => super.copyWith((message) => updates(message as Property));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Property create() => Property._();
  Property createEmptyInstance() => create();
  static $pb.PbList<Property> createRepeated() => $pb.PbList<Property>();
  @$core.pragma('dart2js:noInline')
  static Property getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Property>(create);
  static Property _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $12.Struct get value => $_getN(1);
  @$pb.TagNumber(2)
  set value($12.Struct v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasValue() => $_has(1);
  @$pb.TagNumber(2)
  void clearValue() => clearField(2);
  @$pb.TagNumber(2)
  $12.Struct ensureValue() => $_ensure(1);
}

class Properties extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Properties', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..m<$core.String, $12.Struct>(1, 'kvs', entryClassName: 'Properties.KvsEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OM, valueCreator: $12.Struct.create, packageName: const $pb.PackageName('appootb.account'))
    ..hasRequiredFields = false
  ;

  Properties._() : super();
  factory Properties() => create();
  factory Properties.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Properties.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Properties clone() => Properties()..mergeFromMessage(this);
  Properties copyWith(void Function(Properties) updates) => super.copyWith((message) => updates(message as Properties));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Properties create() => Properties._();
  Properties createEmptyInstance() => create();
  static $pb.PbList<Properties> createRepeated() => $pb.PbList<Properties>();
  @$core.pragma('dart2js:noInline')
  static Properties getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Properties>(create);
  static Properties _defaultInstance;

  @$pb.TagNumber(1)
  $core.Map<$core.String, $12.Struct> get kvs => $_getMap(0);
}

