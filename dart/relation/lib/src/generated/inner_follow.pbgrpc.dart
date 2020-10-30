///
//  Generated code. Do not modify.
//  source: inner_follow.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'appootb/common/unique_id.pb.dart' as $0;
import 'include.pb.dart' as $2;
export 'inner_follow.pb.dart';

class InnerFollowClient extends $grpc.Client {
  static final _$isFollowing =
      $grpc.ClientMethod<$0.UniqueIds, $2.StatusResponse>(
          '/appootb.relation.InnerFollow/IsFollowing',
          ($0.UniqueIds value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.StatusResponse.fromBuffer(value));
  static final _$isFriends =
      $grpc.ClientMethod<$0.UniqueIds, $2.StatusResponse>(
          '/appootb.relation.InnerFollow/IsFriends',
          ($0.UniqueIds value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.StatusResponse.fromBuffer(value));

  InnerFollowClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$2.StatusResponse> isFollowing($0.UniqueIds request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$isFollowing, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.StatusResponse> isFriends($0.UniqueIds request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$isFriends, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class InnerFollowServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.relation.InnerFollow';

  InnerFollowServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.UniqueIds, $2.StatusResponse>(
        'IsFollowing',
        isFollowing_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UniqueIds.fromBuffer(value),
        ($2.StatusResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UniqueIds, $2.StatusResponse>(
        'IsFriends',
        isFriends_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UniqueIds.fromBuffer(value),
        ($2.StatusResponse value) => value.writeToBuffer()));
  }

  $async.Future<$2.StatusResponse> isFollowing_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UniqueIds> request) async {
    return isFollowing(call, await request);
  }

  $async.Future<$2.StatusResponse> isFriends_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UniqueIds> request) async {
    return isFriends(call, await request);
  }

  $async.Future<$2.StatusResponse> isFollowing(
      $grpc.ServiceCall call, $0.UniqueIds request);
  $async.Future<$2.StatusResponse> isFriends(
      $grpc.ServiceCall call, $0.UniqueIds request);
}
