///
//  Generated code. Do not modify.
//  source: bind.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'include.pbenum.dart' as $1;

class Binding extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Binding', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..e<$1.AuthType>(1, 'sourceType', $pb.PbFieldType.QE, defaultOrMaker: $1.AuthType.AUTH_TYPE_UNSPECIFIED, valueOf: $1.AuthType.valueOf, enumValues: $1.AuthType.values)
    ..aQS(2, 'sourceId')
  ;

  Binding._() : super();
  factory Binding() => create();
  factory Binding.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Binding.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Binding clone() => Binding()..mergeFromMessage(this);
  Binding copyWith(void Function(Binding) updates) => super.copyWith((message) => updates(message as Binding));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Binding create() => Binding._();
  Binding createEmptyInstance() => create();
  static $pb.PbList<Binding> createRepeated() => $pb.PbList<Binding>();
  @$core.pragma('dart2js:noInline')
  static Binding getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Binding>(create);
  static Binding _defaultInstance;

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
    ..aQM<Binding>(1, 'bind', subBuilder: Binding.create)
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
  Binding get bind => $_getN(0);
  @$pb.TagNumber(1)
  set bind(Binding v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasBind() => $_has(0);
  @$pb.TagNumber(1)
  void clearBind() => clearField(1);
  @$pb.TagNumber(1)
  Binding ensureBind() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.String get code => $_getSZ(1);
  @$pb.TagNumber(2)
  set code($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);
}

class Bindings extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Bindings', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..pc<Binding>(1, 'binds', $pb.PbFieldType.PM, subBuilder: Binding.create)
  ;

  Bindings._() : super();
  factory Bindings() => create();
  factory Bindings.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Bindings.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Bindings clone() => Bindings()..mergeFromMessage(this);
  Bindings copyWith(void Function(Bindings) updates) => super.copyWith((message) => updates(message as Bindings));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Bindings create() => Bindings._();
  Bindings createEmptyInstance() => create();
  static $pb.PbList<Bindings> createRepeated() => $pb.PbList<Bindings>();
  @$core.pragma('dart2js:noInline')
  static Bindings getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Bindings>(create);
  static Bindings _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Binding> get binds => $_getList(0);
}

