///
//  Generated code. Do not modify.
//  source: channel.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'include.pb.dart' as $2;
import 'channel.pb.dart' as $3;
import 'appootb/common/unique_id.pb.dart' as $0;
import 'google/protobuf/empty.pb.dart' as $1;
export 'channel.pb.dart';

class ChannelClient extends $grpc.Client {
  static final _$add = $grpc.ClientMethod<$2.ApplyRequest, $3.ChannelInfo>(
      '/appootb.relation.Channel/Add',
      ($2.ApplyRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $3.ChannelInfo.fromBuffer(value));
  static final _$cancel = $grpc.ClientMethod<$0.UniqueId, $1.Empty>(
      '/appootb.relation.Channel/Cancel',
      ($0.UniqueId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));

  ChannelClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$3.ChannelInfo> add($2.ApplyRequest request,
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

abstract class ChannelServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.relation.Channel';

  ChannelServiceBase() {
    $addMethod($grpc.ServiceMethod<$2.ApplyRequest, $3.ChannelInfo>(
        'Add',
        add_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.ApplyRequest.fromBuffer(value),
        ($3.ChannelInfo value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UniqueId, $1.Empty>(
        'Cancel',
        cancel_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UniqueId.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$3.ChannelInfo> add_Pre(
      $grpc.ServiceCall call, $async.Future<$2.ApplyRequest> request) async {
    return add(call, await request);
  }

  $async.Future<$1.Empty> cancel_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UniqueId> request) async {
    return cancel(call, await request);
  }

  $async.Future<$3.ChannelInfo> add(
      $grpc.ServiceCall call, $2.ApplyRequest request);
  $async.Future<$1.Empty> cancel($grpc.ServiceCall call, $0.UniqueId request);
}
