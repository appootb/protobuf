///
//  Generated code. Do not modify.
//  source: group.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'group.pb.dart' as $4;
import 'appootb/common/unique_id.pb.dart' as $0;
export 'group.pb.dart';

class GroupClient extends $grpc.Client {
  static final _$create = $grpc.ClientMethod<$4.GroupInfo, $4.GroupInfo>(
      '/appootb.relation.Group/Create',
      ($4.GroupInfo value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $4.GroupInfo.fromBuffer(value));
  static final _$get = $grpc.ClientMethod<$0.UniqueId, $4.GroupInfo>(
      '/appootb.relation.Group/Get',
      ($0.UniqueId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $4.GroupInfo.fromBuffer(value));
  static final _$update = $grpc.ClientMethod<$4.GroupInfo, $4.GroupInfo>(
      '/appootb.relation.Group/Update',
      ($4.GroupInfo value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $4.GroupInfo.fromBuffer(value));
  static final _$add = $grpc.ClientMethod<$0.UniqueIds, $4.GroupInfo>(
      '/appootb.relation.Group/Add',
      ($0.UniqueIds value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $4.GroupInfo.fromBuffer(value));
  static final _$remove = $grpc.ClientMethod<$0.UniqueId, $4.GroupInfo>(
      '/appootb.relation.Group/Remove',
      ($0.UniqueId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $4.GroupInfo.fromBuffer(value));

  GroupClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$4.GroupInfo> create($4.GroupInfo request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$create, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$4.GroupInfo> get($0.UniqueId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$get, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$4.GroupInfo> update($4.GroupInfo request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$update, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$4.GroupInfo> add($0.UniqueIds request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$add, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$4.GroupInfo> remove($0.UniqueId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$remove, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class GroupServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.relation.Group';

  GroupServiceBase() {
    $addMethod($grpc.ServiceMethod<$4.GroupInfo, $4.GroupInfo>(
        'Create',
        create_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.GroupInfo.fromBuffer(value),
        ($4.GroupInfo value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UniqueId, $4.GroupInfo>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UniqueId.fromBuffer(value),
        ($4.GroupInfo value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.GroupInfo, $4.GroupInfo>(
        'Update',
        update_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.GroupInfo.fromBuffer(value),
        ($4.GroupInfo value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UniqueIds, $4.GroupInfo>(
        'Add',
        add_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UniqueIds.fromBuffer(value),
        ($4.GroupInfo value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UniqueId, $4.GroupInfo>(
        'Remove',
        remove_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UniqueId.fromBuffer(value),
        ($4.GroupInfo value) => value.writeToBuffer()));
  }

  $async.Future<$4.GroupInfo> create_Pre(
      $grpc.ServiceCall call, $async.Future<$4.GroupInfo> request) async {
    return create(call, await request);
  }

  $async.Future<$4.GroupInfo> get_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UniqueId> request) async {
    return get(call, await request);
  }

  $async.Future<$4.GroupInfo> update_Pre(
      $grpc.ServiceCall call, $async.Future<$4.GroupInfo> request) async {
    return update(call, await request);
  }

  $async.Future<$4.GroupInfo> add_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UniqueIds> request) async {
    return add(call, await request);
  }

  $async.Future<$4.GroupInfo> remove_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UniqueId> request) async {
    return remove(call, await request);
  }

  $async.Future<$4.GroupInfo> create(
      $grpc.ServiceCall call, $4.GroupInfo request);
  $async.Future<$4.GroupInfo> get($grpc.ServiceCall call, $0.UniqueId request);
  $async.Future<$4.GroupInfo> update(
      $grpc.ServiceCall call, $4.GroupInfo request);
  $async.Future<$4.GroupInfo> add($grpc.ServiceCall call, $0.UniqueIds request);
  $async.Future<$4.GroupInfo> remove(
      $grpc.ServiceCall call, $0.UniqueId request);
}
