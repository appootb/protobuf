///
//  Generated code. Do not modify.
//  source: chat.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'include.pb.dart' as $0;
export 'chat.pb.dart';

class ChatClient extends $grpc.Client {
  static final _$interact = $grpc.ClientMethod<$0.Stream, $0.Stream>(
      '/appootb.message.Chat/Interact',
      ($0.Stream value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Stream.fromBuffer(value));
  static final _$send = $grpc.ClientMethod<$0.Post, $0.Postmark>(
      '/appootb.message.Chat/Send',
      ($0.Post value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Postmark.fromBuffer(value));

  ChatClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

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

abstract class ChatServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.message.Chat';

  ChatServiceBase() {
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

  $async.Future<$0.Postmark> send_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Post> request) async {
    return send(call, await request);
  }

  $async.Stream<$0.Stream> interact(
      $grpc.ServiceCall call, $async.Stream<$0.Stream> request);
  $async.Future<$0.Postmark> send($grpc.ServiceCall call, $0.Post request);
}
