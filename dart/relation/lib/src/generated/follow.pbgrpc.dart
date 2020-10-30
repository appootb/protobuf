///
//  Generated code. Do not modify.
//  source: follow.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'include.pb.dart' as $2;
import 'google/protobuf/empty.pb.dart' as $1;
import 'appootb/common/unique_id.pb.dart' as $0;
export 'follow.pb.dart';

class FollowClient extends $grpc.Client {
  static final _$add = $grpc.ClientMethod<$2.ApplyRequest, $1.Empty>(
      '/appootb.relation.Follow/Add',
      ($2.ApplyRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$cancel = $grpc.ClientMethod<$0.UniqueId, $1.Empty>(
      '/appootb.relation.Follow/Cancel',
      ($0.UniqueId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$getFollowings = $grpc.ClientMethod<$0.UniqueId, $0.UniqueIds>(
      '/appootb.relation.Follow/GetFollowings',
      ($0.UniqueId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.UniqueIds.fromBuffer(value));
  static final _$getFollowers =
      $grpc.ClientMethod<$0.PaginationIdRequest, $0.PaginationIdResponse>(
          '/appootb.relation.Follow/GetFollowers',
          ($0.PaginationIdRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.PaginationIdResponse.fromBuffer(value));
  static final _$getFriends = $grpc.ClientMethod<$1.Empty, $0.UniqueIds>(
      '/appootb.relation.Follow/GetFriends',
      ($1.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.UniqueIds.fromBuffer(value));

  FollowClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.Empty> add($2.ApplyRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$add, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Empty> cancel($0.UniqueId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$cancel, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.UniqueIds> getFollowings($0.UniqueId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$getFollowings, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.PaginationIdResponse> getFollowers(
      $0.PaginationIdRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$getFollowers, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.UniqueIds> getFriends($1.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$getFriends, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class FollowServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.relation.Follow';

  FollowServiceBase() {
    $addMethod($grpc.ServiceMethod<$2.ApplyRequest, $1.Empty>(
        'Add',
        add_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.ApplyRequest.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UniqueId, $1.Empty>(
        'Cancel',
        cancel_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UniqueId.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UniqueId, $0.UniqueIds>(
        'GetFollowings',
        getFollowings_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UniqueId.fromBuffer(value),
        ($0.UniqueIds value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$0.PaginationIdRequest, $0.PaginationIdResponse>(
            'GetFollowers',
            getFollowers_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.PaginationIdRequest.fromBuffer(value),
            ($0.PaginationIdResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Empty, $0.UniqueIds>(
        'GetFriends',
        getFriends_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($0.UniqueIds value) => value.writeToBuffer()));
  }

  $async.Future<$1.Empty> add_Pre(
      $grpc.ServiceCall call, $async.Future<$2.ApplyRequest> request) async {
    return add(call, await request);
  }

  $async.Future<$1.Empty> cancel_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UniqueId> request) async {
    return cancel(call, await request);
  }

  $async.Future<$0.UniqueIds> getFollowings_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UniqueId> request) async {
    return getFollowings(call, await request);
  }

  $async.Future<$0.PaginationIdResponse> getFollowers_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.PaginationIdRequest> request) async {
    return getFollowers(call, await request);
  }

  $async.Future<$0.UniqueIds> getFriends_Pre(
      $grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return getFriends(call, await request);
  }

  $async.Future<$1.Empty> add($grpc.ServiceCall call, $2.ApplyRequest request);
  $async.Future<$1.Empty> cancel($grpc.ServiceCall call, $0.UniqueId request);
  $async.Future<$0.UniqueIds> getFollowings(
      $grpc.ServiceCall call, $0.UniqueId request);
  $async.Future<$0.PaginationIdResponse> getFollowers(
      $grpc.ServiceCall call, $0.PaginationIdRequest request);
  $async.Future<$0.UniqueIds> getFriends(
      $grpc.ServiceCall call, $1.Empty request);
}
