///
//  Generated code. Do not modify.
//  source: inner_hash.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class HashField extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HashField', package: const $pb.PackageName('appootb.counter'), createEmptyInstance: create)
    ..aQS(1, 'product')
    ..aQS(2, 'type')
    ..aQS(3, 'relateId')
    ..aQS(4, 'field')
    ..aInt64(5, 'value')
  ;

  HashField._() : super();
  factory HashField() => create();
  factory HashField.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HashField.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  HashField clone() => HashField()..mergeFromMessage(this);
  HashField copyWith(void Function(HashField) updates) => super.copyWith((message) => updates(message as HashField));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HashField create() => HashField._();
  HashField createEmptyInstance() => create();
  static $pb.PbList<HashField> createRepeated() => $pb.PbList<HashField>();
  @$core.pragma('dart2js:noInline')
  static HashField getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HashField>(create);
  static HashField _defaultInstance;

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
  $core.String get field_4 => $_getSZ(3);
  @$pb.TagNumber(4)
  set field_4($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasField_4() => $_has(3);
  @$pb.TagNumber(4)
  void clearField_4() => clearField(4);

  @$pb.TagNumber(5)
  $fixnum.Int64 get value => $_getI64(4);
  @$pb.TagNumber(5)
  set value($fixnum.Int64 v) { $_setInt64(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasValue() => $_has(4);
  @$pb.TagNumber(5)
  void clearValue() => clearField(5);
}

class HashFields extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HashFields', package: const $pb.PackageName('appootb.counter'), createEmptyInstance: create)
    ..aQS(1, 'product')
    ..aQS(2, 'type')
    ..aQS(3, 'relateId')
    ..pPS(4, 'fields')
    ..m<$core.String, $fixnum.Int64>(5, 'values', entryClassName: 'HashFields.ValuesEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.O6, packageName: const $pb.PackageName('appootb.counter'))
  ;

  HashFields._() : super();
  factory HashFields() => create();
  factory HashFields.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HashFields.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  HashFields clone() => HashFields()..mergeFromMessage(this);
  HashFields copyWith(void Function(HashFields) updates) => super.copyWith((message) => updates(message as HashFields));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HashFields create() => HashFields._();
  HashFields createEmptyInstance() => create();
  static $pb.PbList<HashFields> createRepeated() => $pb.PbList<HashFields>();
  @$core.pragma('dart2js:noInline')
  static HashFields getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HashFields>(create);
  static HashFields _defaultInstance;

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
  $core.List<$core.String> get fields => $_getList(3);

  @$pb.TagNumber(5)
  $core.Map<$core.String, $fixnum.Int64> get values => $_getMap(4);
}

class HashKeys extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HashKeys', package: const $pb.PackageName('appootb.counter'), createEmptyInstance: create)
    ..aQS(1, 'product')
    ..aQS(2, 'type')
    ..pPS(3, 'relateIds')
  ;

  HashKeys._() : super();
  factory HashKeys() => create();
  factory HashKeys.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HashKeys.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  HashKeys clone() => HashKeys()..mergeFromMessage(this);
  HashKeys copyWith(void Function(HashKeys) updates) => super.copyWith((message) => updates(message as HashKeys));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HashKeys create() => HashKeys._();
  HashKeys createEmptyInstance() => create();
  static $pb.PbList<HashKeys> createRepeated() => $pb.PbList<HashKeys>();
  @$core.pragma('dart2js:noInline')
  static HashKeys getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HashKeys>(create);
  static HashKeys _defaultInstance;

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
}

class HashValue extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HashValue', package: const $pb.PackageName('appootb.counter'), createEmptyInstance: create)
    ..m<$core.String, $fixnum.Int64>(1, 'fieldValues', entryClassName: 'HashValue.FieldValuesEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.O6, packageName: const $pb.PackageName('appootb.counter'))
    ..hasRequiredFields = false
  ;

  HashValue._() : super();
  factory HashValue() => create();
  factory HashValue.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HashValue.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  HashValue clone() => HashValue()..mergeFromMessage(this);
  HashValue copyWith(void Function(HashValue) updates) => super.copyWith((message) => updates(message as HashValue));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HashValue create() => HashValue._();
  HashValue createEmptyInstance() => create();
  static $pb.PbList<HashValue> createRepeated() => $pb.PbList<HashValue>();
  @$core.pragma('dart2js:noInline')
  static HashValue getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HashValue>(create);
  static HashValue _defaultInstance;

  @$pb.TagNumber(1)
  $core.Map<$core.String, $fixnum.Int64> get fieldValues => $_getMap(0);
}

class HashValues extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HashValues', package: const $pb.PackageName('appootb.counter'), createEmptyInstance: create)
    ..m<$core.String, HashValue>(1, 'values', entryClassName: 'HashValues.ValuesEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OM, valueCreator: HashValue.create, packageName: const $pb.PackageName('appootb.counter'))
    ..hasRequiredFields = false
  ;

  HashValues._() : super();
  factory HashValues() => create();
  factory HashValues.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HashValues.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  HashValues clone() => HashValues()..mergeFromMessage(this);
  HashValues copyWith(void Function(HashValues) updates) => super.copyWith((message) => updates(message as HashValues));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HashValues create() => HashValues._();
  HashValues createEmptyInstance() => create();
  static $pb.PbList<HashValues> createRepeated() => $pb.PbList<HashValues>();
  @$core.pragma('dart2js:noInline')
  static HashValues getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HashValues>(create);
  static HashValues _defaultInstance;

  @$pb.TagNumber(1)
  $core.Map<$core.String, HashValue> get values => $_getMap(0);
}

