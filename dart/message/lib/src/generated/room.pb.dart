///
//  Generated code. Do not modify.
//  source: room.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class RoomServerOption extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RoomServerOption', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..aQS(1, 'host')
    ..a<$core.int>(2, 'port', $pb.PbFieldType.OU3, defaultOrMaker: 443)
    ..a<$core.bool>(3, 'secure', $pb.PbFieldType.OB, defaultOrMaker: true)
  ;

  RoomServerOption._() : super();
  factory RoomServerOption() => create();
  factory RoomServerOption.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoomServerOption.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RoomServerOption clone() => RoomServerOption()..mergeFromMessage(this);
  RoomServerOption copyWith(void Function(RoomServerOption) updates) => super.copyWith((message) => updates(message as RoomServerOption));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RoomServerOption create() => RoomServerOption._();
  RoomServerOption createEmptyInstance() => create();
  static $pb.PbList<RoomServerOption> createRepeated() => $pb.PbList<RoomServerOption>();
  @$core.pragma('dart2js:noInline')
  static RoomServerOption getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoomServerOption>(create);
  static RoomServerOption _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get host => $_getSZ(0);
  @$pb.TagNumber(1)
  set host($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasHost() => $_has(0);
  @$pb.TagNumber(1)
  void clearHost() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get port => $_getI(1, 443);
  @$pb.TagNumber(2)
  set port($core.int v) { $_setUnsignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPort() => $_has(1);
  @$pb.TagNumber(2)
  void clearPort() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get secure => $_getB(2, true);
  @$pb.TagNumber(3)
  set secure($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasSecure() => $_has(2);
  @$pb.TagNumber(3)
  void clearSecure() => clearField(3);
}

