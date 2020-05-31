///
//  Generated code. Do not modify.
//  source: inner_filter.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'common.pb.dart' as $0;
import 'google/protobuf/empty.pb.dart' as $1;
export 'inner_filter.pb.dart';

class InnerFilterClient extends $grpc.Client {
  static final _$filter = $grpc.ClientMethod<$0.FilterRequest, $1.Empty>(
      '/appootb.strainer.InnerFilter/Filter',
      ($0.FilterRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));

  InnerFilterClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.Empty> filter($0.FilterRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$filter, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class InnerFilterServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.strainer.InnerFilter';

  InnerFilterServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.FilterRequest, $1.Empty>(
        'Filter',
        filter_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.FilterRequest.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$1.Empty> filter_Pre(
      $grpc.ServiceCall call, $async.Future<$0.FilterRequest> request) async {
    return filter(call, await request);
  }

  $async.Future<$1.Empty> filter(
      $grpc.ServiceCall call, $0.FilterRequest request);
}
