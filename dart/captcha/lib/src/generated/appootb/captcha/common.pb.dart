///
//  Generated code. Do not modify.
//  source: appootb/captcha/common.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'common.pbenum.dart';

export 'common.pbenum.dart';

class CodeRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('CodeRequest', package: const $pb.PackageName('appootb.captcha'), createEmptyInstance: create)
    ..e<Channel>(1, 'channel', $pb.PbFieldType.QE, defaultOrMaker: Channel.CHANNEL_UNSPECIFIED, valueOf: Channel.valueOf, enumValues: Channel.values)
    ..e<Category>(2, 'category', $pb.PbFieldType.QE, defaultOrMaker: Category.CATEGORY_UNSPECIFIED, valueOf: Category.valueOf, enumValues: Category.values)
    ..aQS(3, 'target')
    ..aOS(4, 'value')
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
  Channel get channel => $_getN(0);
  @$pb.TagNumber(1)
  set channel(Channel v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasChannel() => $_has(0);
  @$pb.TagNumber(1)
  void clearChannel() => clearField(1);

  @$pb.TagNumber(2)
  Category get category => $_getN(1);
  @$pb.TagNumber(2)
  set category(Category v) { setField(2, v); }
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

  @$pb.TagNumber(4)
  $core.String get value => $_getSZ(3);
  @$pb.TagNumber(4)
  set value($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasValue() => $_has(3);
  @$pb.TagNumber(4)
  void clearValue() => clearField(4);
}

