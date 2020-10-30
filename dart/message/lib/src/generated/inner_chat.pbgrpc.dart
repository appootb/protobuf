///
//  Generated code. Do not modify.
//  source: inner_chat.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'include.pb.dart' as $0;
import 'google/protobuf/empty.pb.dart' as $1;
export 'inner_chat.pb.dart';

class InnerChatClient extends $grpc.Client {
  static final _$deliver = $grpc.ClientMethod<$0.Package, $1.Empty>(
      '/appootb.message.InnerChat/Deliver',
      ($0.Package value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$kick = $grpc.ClientMethod<$0.Connections, $1.Empty>(
      '/appootb.message.InnerChat/Kick',
      ($0.Connections value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));

  InnerChatClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.Empty> deliver($0.Package request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$deliver, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Empty> kick($0.Connections request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$kick, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class InnerChatServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.message.InnerChat';

  InnerChatServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Package, $1.Empty>(
        'Deliver',
        deliver_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Package.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Connections, $1.Empty>(
        'Kick',
        kick_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Connections.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
  }

  $async.Future<$1.Empty> deliver_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Package> request) async {
    return deliver(call, await request);
  }

  $async.Future<$1.Empty> kick_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Connections> request) async {
    return kick(call, await request);
  }

  $async.Future<$1.Empty> deliver($grpc.ServiceCall call, $0.Package request);
  $async.Future<$1.Empty> kick($grpc.ServiceCall call, $0.Connections request);
}
