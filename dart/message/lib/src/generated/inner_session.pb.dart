///
//  Generated code. Do not modify.
//  source: inner_session.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'appootb/common/metadata.pb.dart' as $7;
import 'google/protobuf/timestamp.pb.dart' as $6;

import 'inner_session.pbenum.dart';

export 'inner_session.pbenum.dart';

class SessionId extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('SessionId', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..e<SessionType>(1, 'type', $pb.PbFieldType.QE, defaultOrMaker: SessionType.SESSION_TYPE_UNSPECIFIED, valueOf: SessionType.valueOf, enumValues: SessionType.values)
    ..a<$fixnum.Int64>(2, 'account', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(3, 'deviceId')
  ;

  SessionId._() : super();
  factory SessionId() => create();
  factory SessionId.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SessionId.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  SessionId clone() => SessionId()..mergeFromMessage(this);
  SessionId copyWith(void Function(SessionId) updates) => super.copyWith((message) => updates(message as SessionId));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SessionId create() => SessionId._();
  SessionId createEmptyInstance() => create();
  static $pb.PbList<SessionId> createRepeated() => $pb.PbList<SessionId>();
  @$core.pragma('dart2js:noInline')
  static SessionId getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SessionId>(create);
  static SessionId _defaultInstance;

  @$pb.TagNumber(1)
  SessionType get type => $_getN(0);
  @$pb.TagNumber(1)
  set type(SessionType v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasType() => $_has(0);
  @$pb.TagNumber(1)
  void clearType() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get account => $_getI64(1);
  @$pb.TagNumber(2)
  set account($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasAccount() => $_has(1);
  @$pb.TagNumber(2)
  void clearAccount() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get deviceId => $_getSZ(2);
  @$pb.TagNumber(3)
  set deviceId($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDeviceId() => $_has(2);
  @$pb.TagNumber(3)
  void clearDeviceId() => clearField(3);
}

class SessionIds extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('SessionIds', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..pc<SessionId>(1, 'ids', $pb.PbFieldType.PM, subBuilder: SessionId.create)
  ;

  SessionIds._() : super();
  factory SessionIds() => create();
  factory SessionIds.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SessionIds.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  SessionIds clone() => SessionIds()..mergeFromMessage(this);
  SessionIds copyWith(void Function(SessionIds) updates) => super.copyWith((message) => updates(message as SessionIds));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SessionIds create() => SessionIds._();
  SessionIds createEmptyInstance() => create();
  static $pb.PbList<SessionIds> createRepeated() => $pb.PbList<SessionIds>();
  @$core.pragma('dart2js:noInline')
  static SessionIds getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SessionIds>(create);
  static SessionIds _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<SessionId> get ids => $_getList(0);
}

class Session extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Session', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..e<SessionType>(1, 'type', $pb.PbFieldType.QE, defaultOrMaker: SessionType.SESSION_TYPE_UNSPECIFIED, valueOf: SessionType.valueOf, enumValues: SessionType.values)
    ..aQS(2, 'entry')
    ..aQM<$7.Metadata>(3, 'metadata', subBuilder: $7.Metadata.create)
    ..aQM<$6.Timestamp>(4, 'createdAt', subBuilder: $6.Timestamp.create)
  ;

  Session._() : super();
  factory Session() => create();
  factory Session.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Session.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Session clone() => Session()..mergeFromMessage(this);
  Session copyWith(void Function(Session) updates) => super.copyWith((message) => updates(message as Session));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Session create() => Session._();
  Session createEmptyInstance() => create();
  static $pb.PbList<Session> createRepeated() => $pb.PbList<Session>();
  @$core.pragma('dart2js:noInline')
  static Session getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Session>(create);
  static Session _defaultInstance;

  @$pb.TagNumber(1)
  SessionType get type => $_getN(0);
  @$pb.TagNumber(1)
  set type(SessionType v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasType() => $_has(0);
  @$pb.TagNumber(1)
  void clearType() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get entry => $_getSZ(1);
  @$pb.TagNumber(2)
  set entry($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasEntry() => $_has(1);
  @$pb.TagNumber(2)
  void clearEntry() => clearField(2);

  @$pb.TagNumber(3)
  $7.Metadata get metadata => $_getN(2);
  @$pb.TagNumber(3)
  set metadata($7.Metadata v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasMetadata() => $_has(2);
  @$pb.TagNumber(3)
  void clearMetadata() => clearField(3);
  @$pb.TagNumber(3)
  $7.Metadata ensureMetadata() => $_ensure(2);

  @$pb.TagNumber(4)
  $6.Timestamp get createdAt => $_getN(3);
  @$pb.TagNumber(4)
  set createdAt($6.Timestamp v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasCreatedAt() => $_has(3);
  @$pb.TagNumber(4)
  void clearCreatedAt() => clearField(4);
  @$pb.TagNumber(4)
  $6.Timestamp ensureCreatedAt() => $_ensure(3);
}

class Sessions extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Sessions', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..pc<Session>(1, 'sessions', $pb.PbFieldType.PM, subBuilder: Session.create)
  ;

  Sessions._() : super();
  factory Sessions() => create();
  factory Sessions.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Sessions.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Sessions clone() => Sessions()..mergeFromMessage(this);
  Sessions copyWith(void Function(Sessions) updates) => super.copyWith((message) => updates(message as Sessions));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Sessions create() => Sessions._();
  Sessions createEmptyInstance() => create();
  static $pb.PbList<Sessions> createRepeated() => $pb.PbList<Sessions>();
  @$core.pragma('dart2js:noInline')
  static Sessions getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Sessions>(create);
  static Sessions _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Session> get sessions => $_getList(0);
}

