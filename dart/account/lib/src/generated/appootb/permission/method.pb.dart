///
//  Generated code. Do not modify.
//  source: method.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'method.pbenum.dart';

export 'method.pbenum.dart';

class Method {
  static final $pb.Extension token = $pb.Extension<Audience>.repeated('google.protobuf.MethodOptions', 'token', 2507, $pb.PbFieldType.PE, check: $pb.getCheckFunction($pb.PbFieldType.PE), valueOf: Audience.valueOf, enumValues: Audience.values);
  static void registerAllExtensions($pb.ExtensionRegistry registry) {
    registry.add(token);
  }
}

