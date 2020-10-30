///
//  Generated code. Do not modify.
//  source: block.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'appootb/common/unique_id.pb.dart' as $0;
import 'google/protobuf/empty.pb.dart' as $1;
export 'block.pb.dart';

class BlockClient extends $grpc.Client {
  static final _$add = $grpc.ClientMethod<$0.UniqueId, $1.Empty>(
      '/appootb.relation.Block/Add',
      ($0.UniqueId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$cancel = $grpc.ClientMethod<$0.UniqueId, $1.Empty>(
      '/appootb.relation.Block/Cancel',
      ($0.UniqueId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));

  BlockClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.Empty> add($0.UniqueId request,
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
}

abstract class BlockServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.relation.Block';

  BlockServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.UniqueId, $1.Empty>(
        'Add',
        add_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UniqueId.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UniqueId, $1.Empty>(
        'Cancel',
        cancel_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UniqueId.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$1.Empty> add_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UniqueId> request) async {
    return add(call, await request);
  }

  $async.Future<$1.Empty> cancel_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UniqueId> request) async {
    return cancel(call, await request);
  }

  $async.Future<$1.Empty> add($grpc.ServiceCall call, $0.UniqueId request);
  $async.Future<$1.Empty> cancel($grpc.ServiceCall call, $0.UniqueId request);
}
