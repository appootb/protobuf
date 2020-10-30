///
//  Generated code. Do not modify.
//  source: room.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'appootb/common/unique_id.pb.dart' as $3;
import 'room.pb.dart' as $4;
import 'include.pb.dart' as $0;
export 'room.pb.dart';

class RoomClient extends $grpc.Client {
  static final _$dispatch =
      $grpc.ClientMethod<$3.UniqueId, $4.RoomServerOption>(
          '/appootb.message.Room/Dispatch',
          ($3.UniqueId value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $4.RoomServerOption.fromBuffer(value));
  static final _$interact = $grpc.ClientMethod<$0.Stream, $0.Stream>(
      '/appootb.message.Room/Interact',
      ($0.Stream value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Stream.fromBuffer(value));
  static final _$send = $grpc.ClientMethod<$0.Post, $0.Postmark>(
      '/appootb.message.Room/Send',
      ($0.Post value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Postmark.fromBuffer(value));

  RoomClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$4.RoomServerOption> dispatch($3.UniqueId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$dispatch, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseStream<$0.Stream> interact($async.Stream<$0.Stream> request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$interact, request, options: options);
    return $grpc.ResponseStream(call);
  }

  $grpc.ResponseFuture<$0.Postmark> send($0.Post request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$send, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class RoomServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.message.Room';

  RoomServiceBase() {
    $addMethod($grpc.ServiceMethod<$3.UniqueId, $4.RoomServerOption>(
        'Dispatch',
        dispatch_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.UniqueId.fromBuffer(value),
        ($4.RoomServerOption value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Stream, $0.Stream>(
        'Interact',
        interact,
        true,
        true,
        ($core.List<$core.int> value) => $0.Stream.fromBuffer(value),
        ($0.Stream value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Post, $0.Postmark>(
        'Send',
        send_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Post.fromBuffer(value),
        ($0.Postmark value) => value.writeToBuffer()));
  }

  $async.Future<$4.RoomServerOption> dispatch_Pre(
      $grpc.ServiceCall call, $async.Future<$3.UniqueId> request) async {
    return dispatch(call, await request);
  }

  $async.Future<$0.Postmark> send_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Post> request) async {
    return send(call, await request);
  }

  $async.Future<$4.RoomServerOption> dispatch(
      $grpc.ServiceCall call, $3.UniqueId request);
  $async.Stream<$0.Stream> interact(
      $grpc.ServiceCall call, $async.Stream<$0.Stream> request);
  $async.Future<$0.Postmark> send($grpc.ServiceCall call, $0.Post request);
}
