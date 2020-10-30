///
//  Generated code. Do not modify.
//  source: inner_group.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'appootb/common/unique_id.pb.dart' as $0;
import 'include.pb.dart' as $2;
export 'inner_group.pb.dart';

class InnerGroupClient extends $grpc.Client {
  static final _$getMembers = $grpc.ClientMethod<$0.UniqueId, $0.UniqueIds>(
      '/appootb.relation.InnerGroup/GetMembers',
      ($0.UniqueId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.UniqueIds.fromBuffer(value));
  static final _$isMembers =
      $grpc.ClientMethod<$2.StatusRequest, $2.StatusResponse>(
          '/appootb.relation.InnerGroup/IsMembers',
          ($2.StatusRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.StatusResponse.fromBuffer(value));

  InnerGroupClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$0.UniqueIds> getMembers($0.UniqueId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$getMembers, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.StatusResponse> isMembers($2.StatusRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$isMembers, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class InnerGroupServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.relation.InnerGroup';

  InnerGroupServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.UniqueId, $0.UniqueIds>(
        'GetMembers',
        getMembers_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UniqueId.fromBuffer(value),
        ($0.UniqueIds value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.StatusRequest, $2.StatusResponse>(
        'IsMembers',
        isMembers_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.StatusRequest.fromBuffer(value),
        ($2.StatusResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.UniqueIds> getMembers_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UniqueId> request) async {
    return getMembers(call, await request);
  }

  $async.Future<$2.StatusResponse> isMembers_Pre(
      $grpc.ServiceCall call, $async.Future<$2.StatusRequest> request) async {
    return isMembers(call, await request);
  }

  $async.Future<$0.UniqueIds> getMembers(
      $grpc.ServiceCall call, $0.UniqueId request);
  $async.Future<$2.StatusResponse> isMembers(
      $grpc.ServiceCall call, $2.StatusRequest request);
}
