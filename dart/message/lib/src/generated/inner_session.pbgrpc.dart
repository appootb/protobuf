///
//  Generated code. Do not modify.
//  source: inner_session.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'inner_session.pb.dart' as $2;
import 'google/protobuf/empty.pb.dart' as $1;
export 'inner_session.pb.dart';

class InnerSessionClient extends $grpc.Client {
  static final _$open = $grpc.ClientMethod<$2.Session, $2.Sessions>(
      '/appootb.message.InnerSession/Open',
      ($2.Session value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Sessions.fromBuffer(value));
  static final _$close = $grpc.ClientMethod<$2.SessionId, $1.Empty>(
      '/appootb.message.InnerSession/Close',
      ($2.SessionId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Empty.fromBuffer(value));
  static final _$gets = $grpc.ClientMethod<$2.SessionIds, $2.Sessions>(
      '/appootb.message.InnerSession/Gets',
      ($2.SessionIds value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.Sessions.fromBuffer(value));

  InnerSessionClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$2.Sessions> open($2.Session request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$open, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Empty> close($2.SessionId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$close, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.Sessions> gets($2.SessionIds request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$gets, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class InnerSessionServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.message.InnerSession';

  InnerSessionServiceBase() {
    $addMethod($grpc.ServiceMethod<$2.Session, $2.Sessions>(
        'Open',
        open_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Session.fromBuffer(value),
        ($2.Sessions value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.SessionId, $1.Empty>(
        'Close',
        close_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.SessionId.fromBuffer(value),
        ($1.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.SessionIds, $2.Sessions>(
        'Gets',
        gets_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.SessionIds.fromBuffer(value),
        ($2.Sessions value) => value.writeToBuffer()));
  }

  $async.Future<$2.Sessions> open_Pre(
      $grpc.ServiceCall call, $async.Future<$2.Session> request) async {
    return open(call, await request);
  }

  $async.Future<$1.Empty> close_Pre(
      $grpc.ServiceCall call, $async.Future<$2.SessionId> request) async {
    return close(call, await request);
  }

  $async.Future<$2.Sessions> gets_Pre(
      $grpc.ServiceCall call, $async.Future<$2.SessionIds> request) async {
    return gets(call, await request);
  }

  $async.Future<$2.Sessions> open($grpc.ServiceCall call, $2.Session request);
  $async.Future<$1.Empty> close($grpc.ServiceCall call, $2.SessionId request);
  $async.Future<$2.Sessions> gets(
      $grpc.ServiceCall call, $2.SessionIds request);
}
