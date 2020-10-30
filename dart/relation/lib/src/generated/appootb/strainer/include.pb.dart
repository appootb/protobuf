///
//  Generated code. Do not modify.
//  source: appootb/strainer/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'include.pbenum.dart';

export 'include.pbenum.dart';

class FilterRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FilterRequest', package: const $pb.PackageName('appootb.strainer'), createEmptyInstance: create)
    ..e<Channel>(1, 'channel', $pb.PbFieldType.OE, defaultOrMaker: Channel.CHANNEL_UNSPECIFIED, valueOf: Channel.valueOf, enumValues: Channel.values)
    ..e<Target>(2, 'type', $pb.PbFieldType.QE, defaultOrMaker: Target.TARGET_UNSPECIFIED, valueOf: Target.valueOf, enumValues: Target.values)
    ..aQS(3, 'data')
  ;

  FilterRequest._() : super();
  factory FilterRequest() => create();
  factory FilterRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FilterRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  FilterRequest clone() => FilterRequest()..mergeFromMessage(this);
  FilterRequest copyWith(void Function(FilterRequest) updates) => super.copyWith((message) => updates(message as FilterRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FilterRequest create() => FilterRequest._();
  FilterRequest createEmptyInstance() => create();
  static $pb.PbList<FilterRequest> createRepeated() => $pb.PbList<FilterRequest>();
  @$core.pragma('dart2js:noInline')
  static FilterRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FilterRequest>(create);
  static FilterRequest _defaultInstance;

  @$pb.TagNumber(1)
  Channel get channel => $_getN(0);
  @$pb.TagNumber(1)
  set channel(Channel v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasChannel() => $_has(0);
  @$pb.TagNumber(1)
  void clearChannel() => clearField(1);

  @$pb.TagNumber(2)
  Target get type => $_getN(1);
  @$pb.TagNumber(2)
  set type(Target v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasType() => $_has(1);
  @$pb.TagNumber(2)
  void clearType() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get data => $_getSZ(2);
  @$pb.TagNumber(3)
  set data($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasData() => $_has(2);
  @$pb.TagNumber(3)
  void clearData() => clearField(3);
}

