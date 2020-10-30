///
//  Generated code. Do not modify.
//  source: inner_router.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
export 'inner_router.pb.dart';

class InnerRouterClient extends $grpc.Client {
  InnerRouterClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);
}

abstract class InnerRouterServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.message.InnerRouter';

  InnerRouterServiceBase() {}
}
