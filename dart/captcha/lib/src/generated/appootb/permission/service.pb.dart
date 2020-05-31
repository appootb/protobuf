///
//  Generated code. Do not modify.
//  source: service.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'service.pbenum.dart';

export 'service.pbenum.dart';

class ServiceVisible extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ServiceVisible', package: const $pb.PackageName('appootb.permission.service'), createEmptyInstance: create)
    ..e<VisibleScope>(1, 'scope', $pb.PbFieldType.OE, defaultOrMaker: VisibleScope.DEFAULT_SCOPE, valueOf: VisibleScope.valueOf, enumValues: VisibleScope.values)
    ..hasRequiredFields = false
  ;

  ServiceVisible._() : super();
  factory ServiceVisible() => create();
  factory ServiceVisible.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServiceVisible.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ServiceVisible clone() => ServiceVisible()..mergeFromMessage(this);
  ServiceVisible copyWith(void Function(ServiceVisible) updates) => super.copyWith((message) => updates(message as ServiceVisible));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ServiceVisible create() => ServiceVisible._();
  ServiceVisible createEmptyInstance() => create();
  static $pb.PbList<ServiceVisible> createRepeated() => $pb.PbList<ServiceVisible>();
  @$core.pragma('dart2js:noInline')
  static ServiceVisible getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServiceVisible>(create);
  static ServiceVisible _defaultInstance;

  @$pb.TagNumber(1)
  VisibleScope get scope => $_getN(0);
  @$pb.TagNumber(1)
  set scope(VisibleScope v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasScope() => $_has(0);
  @$pb.TagNumber(1)
  void clearScope() => clearField(1);
}

class Service {
  static final $pb.Extension visible = $pb.Extension<ServiceVisible>('google.protobuf.ServiceOptions', 'visible', 1507, $pb.PbFieldType.OM, defaultOrMaker: ServiceVisible.getDefault, subBuilder: ServiceVisible.create);
  static void registerAllExtensions($pb.ExtensionRegistry registry) {
    registry.add(visible);
  }
}

