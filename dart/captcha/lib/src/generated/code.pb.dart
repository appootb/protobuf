///
//  Generated code. Do not modify.
//  source: code.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'include.pbenum.dart' as $3;

class Region extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Region', package: const $pb.PackageName('appootb.captcha'), createEmptyInstance: create)
    ..aQS(1, 'name')
    ..aQS(2, 'local')
    ..aQS(3, 'abbr')
    ..aQS(4, 'code')
  ;

  Region._() : super();
  factory Region() => create();
  factory Region.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Region.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Region clone() => Region()..mergeFromMessage(this);
  Region copyWith(void Function(Region) updates) => super.copyWith((message) => updates(message as Region));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Region create() => Region._();
  Region createEmptyInstance() => create();
  static $pb.PbList<Region> createRepeated() => $pb.PbList<Region>();
  @$core.pragma('dart2js:noInline')
  static Region getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Region>(create);
  static Region _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get local => $_getSZ(1);
  @$pb.TagNumber(2)
  set local($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLocal() => $_has(1);
  @$pb.TagNumber(2)
  void clearLocal() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get abbr => $_getSZ(2);
  @$pb.TagNumber(3)
  set abbr($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasAbbr() => $_has(2);
  @$pb.TagNumber(3)
  void clearAbbr() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get code => $_getSZ(3);
  @$pb.TagNumber(4)
  set code($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasCode() => $_has(3);
  @$pb.TagNumber(4)
  void clearCode() => clearField(4);
}

class Regions extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Regions', package: const $pb.PackageName('appootb.captcha'), createEmptyInstance: create)
    ..pc<Region>(1, 'regions', $pb.PbFieldType.PM, subBuilder: Region.create)
  ;

  Regions._() : super();
  factory Regions() => create();
  factory Regions.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Regions.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Regions clone() => Regions()..mergeFromMessage(this);
  Regions copyWith(void Function(Regions) updates) => super.copyWith((message) => updates(message as Regions));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Regions create() => Regions._();
  Regions createEmptyInstance() => create();
  static $pb.PbList<Regions> createRepeated() => $pb.PbList<Regions>();
  @$core.pragma('dart2js:noInline')
  static Regions getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Regions>(create);
  static Regions _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Region> get regions => $_getList(0);
}

class CodeRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('CodeRequest', package: const $pb.PackageName('appootb.captcha'), createEmptyInstance: create)
    ..e<$3.Channel>(1, 'channel', $pb.PbFieldType.QE, defaultOrMaker: $3.Channel.CHANNEL_UNSPECIFIED, valueOf: $3.Channel.valueOf, enumValues: $3.Channel.values)
    ..e<$3.Category>(2, 'category', $pb.PbFieldType.QE, defaultOrMaker: $3.Category.CATEGORY_UNSPECIFIED, valueOf: $3.Category.valueOf, enumValues: $3.Category.values)
    ..aQS(3, 'target')
  ;

  CodeRequest._() : super();
  factory CodeRequest() => create();
  factory CodeRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CodeRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  CodeRequest clone() => CodeRequest()..mergeFromMessage(this);
  CodeRequest copyWith(void Function(CodeRequest) updates) => super.copyWith((message) => updates(message as CodeRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CodeRequest create() => CodeRequest._();
  CodeRequest createEmptyInstance() => create();
  static $pb.PbList<CodeRequest> createRepeated() => $pb.PbList<CodeRequest>();
  @$core.pragma('dart2js:noInline')
  static CodeRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CodeRequest>(create);
  static CodeRequest _defaultInstance;

  @$pb.TagNumber(1)
  $3.Channel get channel => $_getN(0);
  @$pb.TagNumber(1)
  set channel($3.Channel v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasChannel() => $_has(0);
  @$pb.TagNumber(1)
  void clearChannel() => clearField(1);

  @$pb.TagNumber(2)
  $3.Category get category => $_getN(1);
  @$pb.TagNumber(2)
  set category($3.Category v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasCategory() => $_has(1);
  @$pb.TagNumber(2)
  void clearCategory() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get target => $_getSZ(2);
  @$pb.TagNumber(3)
  set target($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasTarget() => $_has(2);
  @$pb.TagNumber(3)
  void clearTarget() => clearField(3);
}

