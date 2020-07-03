///
//  Generated code. Do not modify.
//  source: account.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'include.pbenum.dart' as $1;

class VariableInfo extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('VariableInfo', package: const $pb.PackageName('appootb.account'), createEmptyInstance: create)
    ..aOS(1, 'nickname')
    ..aOS(2, 'avatar')
    ..aOS(3, 'signature')
    ..e<$1.Gender>(4, 'gender', $pb.PbFieldType.OE, defaultOrMaker: $1.Gender.GENDER_UNSPECIFIED, valueOf: $1.Gender.valueOf, enumValues: $1.Gender.values)
    ..hasRequiredFields = false
  ;

  VariableInfo._() : super();
  factory VariableInfo() => create();
  factory VariableInfo.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory VariableInfo.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  VariableInfo clone() => VariableInfo()..mergeFromMessage(this);
  VariableInfo copyWith(void Function(VariableInfo) updates) => super.copyWith((message) => updates(message as VariableInfo));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static VariableInfo create() => VariableInfo._();
  VariableInfo createEmptyInstance() => create();
  static $pb.PbList<VariableInfo> createRepeated() => $pb.PbList<VariableInfo>();
  @$core.pragma('dart2js:noInline')
  static VariableInfo getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<VariableInfo>(create);
  static VariableInfo _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get nickname => $_getSZ(0);
  @$pb.TagNumber(1)
  set nickname($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasNickname() => $_has(0);
  @$pb.TagNumber(1)
  void clearNickname() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get avatar => $_getSZ(1);
  @$pb.TagNumber(2)
  set avatar($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasAvatar() => $_has(1);
  @$pb.TagNumber(2)
  void clearAvatar() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get signature => $_getSZ(2);
  @$pb.TagNumber(3)
  set signature($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasSignature() => $_has(2);
  @$pb.TagNumber(3)
  void clearSignature() => clearField(3);

  @$pb.TagNumber(4)
  $1.Gender get gender => $_getN(3);
  @$pb.TagNumber(4)
  set gender($1.Gender v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasGender() => $_has(3);
  @$pb.TagNumber(4)
  void clearGender() => clearField(4);
}

