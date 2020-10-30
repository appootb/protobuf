///
//  Generated code. Do not modify.
//  source: metadata.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'metadata.pbenum.dart';

export 'metadata.pbenum.dart';

class Metadata extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Metadata', package: const $pb.PackageName('appootb.common'), createEmptyInstance: create)
    ..aOS(1, 'token')
    ..e<Platform>(2, 'platform', $pb.PbFieldType.QE, defaultOrMaker: Platform.PLATFORM_UNSPECIFIED, valueOf: Platform.valueOf, enumValues: Platform.values)
    ..e<Network>(3, 'network', $pb.PbFieldType.OE, defaultOrMaker: Network.NETWORK_UNSPECIFIED, valueOf: Network.valueOf, enumValues: Network.values)
    ..aOS(4, 'package')
    ..aQS(5, 'version')
    ..aOS(6, 'osVersion')
    ..aOS(7, 'brand')
    ..aOS(8, 'model')
    ..aQS(9, 'deviceId')
    ..a<$fixnum.Int64>(10, 'timestamp', $pb.PbFieldType.Q6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$core.bool>(11, 'isEmulator', $pb.PbFieldType.OB, defaultOrMaker: true)
    ..aOB(12, 'isDebug')
    ..aOS(13, 'latitude')
    ..aOS(14, 'longitude')
    ..aOS(15, 'locale')
    ..aOS(16, 'clientIp')
    ..aOS(17, 'channel')
    ..aOS(18, 'product')
    ..aOS(19, 'traceId')
  ;

  Metadata._() : super();
  factory Metadata() => create();
  factory Metadata.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Metadata.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Metadata clone() => Metadata()..mergeFromMessage(this);
  Metadata copyWith(void Function(Metadata) updates) => super.copyWith((message) => updates(message as Metadata));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Metadata create() => Metadata._();
  Metadata createEmptyInstance() => create();
  static $pb.PbList<Metadata> createRepeated() => $pb.PbList<Metadata>();
  @$core.pragma('dart2js:noInline')
  static Metadata getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Metadata>(create);
  static Metadata _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get token => $_getSZ(0);
  @$pb.TagNumber(1)
  set token($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearToken() => clearField(1);

  @$pb.TagNumber(2)
  Platform get platform => $_getN(1);
  @$pb.TagNumber(2)
  set platform(Platform v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasPlatform() => $_has(1);
  @$pb.TagNumber(2)
  void clearPlatform() => clearField(2);

  @$pb.TagNumber(3)
  Network get network => $_getN(2);
  @$pb.TagNumber(3)
  set network(Network v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasNetwork() => $_has(2);
  @$pb.TagNumber(3)
  void clearNetwork() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get package => $_getSZ(3);
  @$pb.TagNumber(4)
  set package($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasPackage() => $_has(3);
  @$pb.TagNumber(4)
  void clearPackage() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get version => $_getSZ(4);
  @$pb.TagNumber(5)
  set version($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasVersion() => $_has(4);
  @$pb.TagNumber(5)
  void clearVersion() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get osVersion => $_getSZ(5);
  @$pb.TagNumber(6)
  set osVersion($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasOsVersion() => $_has(5);
  @$pb.TagNumber(6)
  void clearOsVersion() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get brand => $_getSZ(6);
  @$pb.TagNumber(7)
  set brand($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasBrand() => $_has(6);
  @$pb.TagNumber(7)
  void clearBrand() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get model => $_getSZ(7);
  @$pb.TagNumber(8)
  set model($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasModel() => $_has(7);
  @$pb.TagNumber(8)
  void clearModel() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get deviceId => $_getSZ(8);
  @$pb.TagNumber(9)
  set deviceId($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasDeviceId() => $_has(8);
  @$pb.TagNumber(9)
  void clearDeviceId() => clearField(9);

  @$pb.TagNumber(10)
  $fixnum.Int64 get timestamp => $_getI64(9);
  @$pb.TagNumber(10)
  set timestamp($fixnum.Int64 v) { $_setInt64(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasTimestamp() => $_has(9);
  @$pb.TagNumber(10)
  void clearTimestamp() => clearField(10);

  @$pb.TagNumber(11)
  $core.bool get isEmulator => $_getB(10, true);
  @$pb.TagNumber(11)
  set isEmulator($core.bool v) { $_setBool(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasIsEmulator() => $_has(10);
  @$pb.TagNumber(11)
  void clearIsEmulator() => clearField(11);

  @$pb.TagNumber(12)
  $core.bool get isDebug => $_getBF(11);
  @$pb.TagNumber(12)
  set isDebug($core.bool v) { $_setBool(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasIsDebug() => $_has(11);
  @$pb.TagNumber(12)
  void clearIsDebug() => clearField(12);

  @$pb.TagNumber(13)
  $core.String get latitude => $_getSZ(12);
  @$pb.TagNumber(13)
  set latitude($core.String v) { $_setString(12, v); }
  @$pb.TagNumber(13)
  $core.bool hasLatitude() => $_has(12);
  @$pb.TagNumber(13)
  void clearLatitude() => clearField(13);

  @$pb.TagNumber(14)
  $core.String get longitude => $_getSZ(13);
  @$pb.TagNumber(14)
  set longitude($core.String v) { $_setString(13, v); }
  @$pb.TagNumber(14)
  $core.bool hasLongitude() => $_has(13);
  @$pb.TagNumber(14)
  void clearLongitude() => clearField(14);

  @$pb.TagNumber(15)
  $core.String get locale => $_getSZ(14);
  @$pb.TagNumber(15)
  set locale($core.String v) { $_setString(14, v); }
  @$pb.TagNumber(15)
  $core.bool hasLocale() => $_has(14);
  @$pb.TagNumber(15)
  void clearLocale() => clearField(15);

  @$pb.TagNumber(16)
  $core.String get clientIp => $_getSZ(15);
  @$pb.TagNumber(16)
  set clientIp($core.String v) { $_setString(15, v); }
  @$pb.TagNumber(16)
  $core.bool hasClientIp() => $_has(15);
  @$pb.TagNumber(16)
  void clearClientIp() => clearField(16);

  @$pb.TagNumber(17)
  $core.String get channel => $_getSZ(16);
  @$pb.TagNumber(17)
  set channel($core.String v) { $_setString(16, v); }
  @$pb.TagNumber(17)
  $core.bool hasChannel() => $_has(16);
  @$pb.TagNumber(17)
  void clearChannel() => clearField(17);

  @$pb.TagNumber(18)
  $core.String get product => $_getSZ(17);
  @$pb.TagNumber(18)
  set product($core.String v) { $_setString(17, v); }
  @$pb.TagNumber(18)
  $core.bool hasProduct() => $_has(17);
  @$pb.TagNumber(18)
  void clearProduct() => clearField(18);

  @$pb.TagNumber(19)
  $core.String get traceId => $_getSZ(18);
  @$pb.TagNumber(19)
  set traceId($core.String v) { $_setString(18, v); }
  @$pb.TagNumber(19)
  $core.bool hasTraceId() => $_has(18);
  @$pb.TagNumber(19)
  void clearTraceId() => clearField(19);
}

