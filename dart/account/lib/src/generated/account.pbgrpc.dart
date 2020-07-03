///
//  Generated code. Do not modify.
//  source: account.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'appootb/common/unique_id.pb.dart' as $0;
import 'include.pb.dart' as $1;
import 'account.pb.dart' as $2;
export 'account.pb.dart';

class AccountClient extends $grpc.Client {
  static final _$getInfo = $grpc.ClientMethod<$0.UniqueId, $1.Info>(
      '/appootb.account.Account/GetInfo',
      ($0.UniqueId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Info.fromBuffer(value));
  static final _$getInfos = $grpc.ClientMethod<$0.UniqueIds, $1.Infos>(
      '/appootb.account.Account/GetInfos',
      ($0.UniqueIds value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Infos.fromBuffer(value));
  static final _$updateInfo = $grpc.ClientMethod<$2.VariableInfo, $1.Info>(
      '/appootb.account.Account/UpdateInfo',
      ($2.VariableInfo value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.Info.fromBuffer(value));

  AccountClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.Info> getInfo($0.UniqueId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$getInfo, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Infos> getInfos($0.UniqueIds request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$getInfos, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.Info> updateInfo($2.VariableInfo request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$updateInfo, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class AccountServiceBase extends $grpc.Service {
  $core.String get $name => 'appootb.account.Account';

  AccountServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.UniqueId, $1.Info>(
        'GetInfo',
        getInfo_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UniqueId.fromBuffer(value),
        ($1.Info value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UniqueIds, $1.Infos>(
        'GetInfos',
        getInfos_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UniqueIds.fromBuffer(value),
        ($1.Infos value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.VariableInfo, $1.Info>(
        'UpdateInfo',
        updateInfo_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.VariableInfo.fromBuffer(value),
        ($1.Info value) => value.writeToBuffer()));
  }

  $async.Future<$1.Info> getInfo_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UniqueId> request) async {
    return getInfo(call, await request);
  }

  $async.Future<$1.Infos> getInfos_Pre(
      $grpc.ServiceCall call, $async.Future<$0.UniqueIds> request) async {
    return getInfos(call, await request);
  }

  $async.Future<$1.Info> updateInfo_Pre(
      $grpc.ServiceCall call, $async.Future<$2.VariableInfo> request) async {
    return updateInfo(call, await request);
  }

  $async.Future<$1.Info> getInfo($grpc.ServiceCall call, $0.UniqueId request);
  $async.Future<$1.Infos> getInfos(
      $grpc.ServiceCall call, $0.UniqueIds request);
  $async.Future<$1.Info> updateInfo(
      $grpc.ServiceCall call, $2.VariableInfo request);
}
