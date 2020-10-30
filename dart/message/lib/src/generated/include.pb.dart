///
//  Generated code. Do not modify.
//  source: include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'google/protobuf/duration.pb.dart' as $5;
import 'google/protobuf/timestamp.pb.dart' as $6;

import 'include.pbenum.dart';

export 'include.pbenum.dart';

class Postmark extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Postmark', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..e<Category>(1, 'category', $pb.PbFieldType.QE, defaultOrMaker: Category.CATEGORY_UNSPECIFIED, valueOf: Category.valueOf, enumValues: Category.values)
    ..a<$fixnum.Int64>(2, 'sender', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(3, 'receiver', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(4, 'uniqueId', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOM<$5.Duration>(8, 'expire', subBuilder: $5.Duration.create)
    ..aOM<$6.Timestamp>(9, 'timestamp', subBuilder: $6.Timestamp.create)
  ;

  Postmark._() : super();
  factory Postmark() => create();
  factory Postmark.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Postmark.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Postmark clone() => Postmark()..mergeFromMessage(this);
  Postmark copyWith(void Function(Postmark) updates) => super.copyWith((message) => updates(message as Postmark));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Postmark create() => Postmark._();
  Postmark createEmptyInstance() => create();
  static $pb.PbList<Postmark> createRepeated() => $pb.PbList<Postmark>();
  @$core.pragma('dart2js:noInline')
  static Postmark getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Postmark>(create);
  static Postmark _defaultInstance;

  @$pb.TagNumber(1)
  Category get category => $_getN(0);
  @$pb.TagNumber(1)
  set category(Category v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasCategory() => $_has(0);
  @$pb.TagNumber(1)
  void clearCategory() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get sender => $_getI64(1);
  @$pb.TagNumber(2)
  set sender($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSender() => $_has(1);
  @$pb.TagNumber(2)
  void clearSender() => clearField(2);

  @$pb.TagNumber(3)
  $fixnum.Int64 get receiver => $_getI64(2);
  @$pb.TagNumber(3)
  set receiver($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasReceiver() => $_has(2);
  @$pb.TagNumber(3)
  void clearReceiver() => clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get uniqueId => $_getI64(3);
  @$pb.TagNumber(4)
  set uniqueId($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasUniqueId() => $_has(3);
  @$pb.TagNumber(4)
  void clearUniqueId() => clearField(4);

  @$pb.TagNumber(8)
  $5.Duration get expire => $_getN(4);
  @$pb.TagNumber(8)
  set expire($5.Duration v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasExpire() => $_has(4);
  @$pb.TagNumber(8)
  void clearExpire() => clearField(8);
  @$pb.TagNumber(8)
  $5.Duration ensureExpire() => $_ensure(4);

  @$pb.TagNumber(9)
  $6.Timestamp get timestamp => $_getN(5);
  @$pb.TagNumber(9)
  set timestamp($6.Timestamp v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasTimestamp() => $_has(5);
  @$pb.TagNumber(9)
  void clearTimestamp() => clearField(9);
  @$pb.TagNumber(9)
  $6.Timestamp ensureTimestamp() => $_ensure(5);
}

class Text extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Text', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..aQS(1, 'content')
  ;

  Text._() : super();
  factory Text() => create();
  factory Text.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Text.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Text clone() => Text()..mergeFromMessage(this);
  Text copyWith(void Function(Text) updates) => super.copyWith((message) => updates(message as Text));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Text create() => Text._();
  Text createEmptyInstance() => create();
  static $pb.PbList<Text> createRepeated() => $pb.PbList<Text>();
  @$core.pragma('dart2js:noInline')
  static Text getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Text>(create);
  static Text _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get content => $_getSZ(0);
  @$pb.TagNumber(1)
  set content($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasContent() => $_has(0);
  @$pb.TagNumber(1)
  void clearContent() => clearField(1);
}

class Media extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Media', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..aQS(1, 'url')
    ..aOS(2, 'thumbnail')
  ;

  Media._() : super();
  factory Media() => create();
  factory Media.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Media.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Media clone() => Media()..mergeFromMessage(this);
  Media copyWith(void Function(Media) updates) => super.copyWith((message) => updates(message as Media));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Media create() => Media._();
  Media createEmptyInstance() => create();
  static $pb.PbList<Media> createRepeated() => $pb.PbList<Media>();
  @$core.pragma('dart2js:noInline')
  static Media getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Media>(create);
  static Media _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get url => $_getSZ(0);
  @$pb.TagNumber(1)
  set url($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUrl() => $_has(0);
  @$pb.TagNumber(1)
  void clearUrl() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get thumbnail => $_getSZ(1);
  @$pb.TagNumber(2)
  set thumbnail($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasThumbnail() => $_has(1);
  @$pb.TagNumber(2)
  void clearThumbnail() => clearField(2);
}

class Location extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Location', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..a<$core.double>(1, 'latitude', $pb.PbFieldType.QF)
    ..a<$core.double>(2, 'longitude', $pb.PbFieldType.QF)
  ;

  Location._() : super();
  factory Location() => create();
  factory Location.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Location.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Location clone() => Location()..mergeFromMessage(this);
  Location copyWith(void Function(Location) updates) => super.copyWith((message) => updates(message as Location));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Location create() => Location._();
  Location createEmptyInstance() => create();
  static $pb.PbList<Location> createRepeated() => $pb.PbList<Location>();
  @$core.pragma('dart2js:noInline')
  static Location getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Location>(create);
  static Location _defaultInstance;

  @$pb.TagNumber(1)
  $core.double get latitude => $_getN(0);
  @$pb.TagNumber(1)
  set latitude($core.double v) { $_setFloat(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasLatitude() => $_has(0);
  @$pb.TagNumber(1)
  void clearLatitude() => clearField(1);

  @$pb.TagNumber(2)
  $core.double get longitude => $_getN(1);
  @$pb.TagNumber(2)
  set longitude($core.double v) { $_setFloat(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLongitude() => $_has(1);
  @$pb.TagNumber(2)
  void clearLongitude() => clearField(2);
}

enum Post_Envelope {
  text, 
  resource, 
  coordinate, 
  notSet
}

class Post extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, Post_Envelope> _Post_EnvelopeByTag = {
    10 : Post_Envelope.text,
    11 : Post_Envelope.resource,
    12 : Post_Envelope.coordinate,
    0 : Post_Envelope.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Post', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..oo(0, [10, 11, 12])
    ..aQM<Postmark>(1, 'postmark', subBuilder: Postmark.create)
    ..e<Type>(2, 'type', $pb.PbFieldType.QE, defaultOrMaker: Type.TYPE_UNSPECIFIED, valueOf: Type.valueOf, enumValues: Type.values)
    ..aOM<Text>(10, 'text', subBuilder: Text.create)
    ..aOM<Media>(11, 'resource', subBuilder: Media.create)
    ..aOM<Location>(12, 'coordinate', subBuilder: Location.create)
  ;

  Post._() : super();
  factory Post() => create();
  factory Post.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Post.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Post clone() => Post()..mergeFromMessage(this);
  Post copyWith(void Function(Post) updates) => super.copyWith((message) => updates(message as Post));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Post create() => Post._();
  Post createEmptyInstance() => create();
  static $pb.PbList<Post> createRepeated() => $pb.PbList<Post>();
  @$core.pragma('dart2js:noInline')
  static Post getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Post>(create);
  static Post _defaultInstance;

  Post_Envelope whichEnvelope() => _Post_EnvelopeByTag[$_whichOneof(0)];
  void clearEnvelope() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  Postmark get postmark => $_getN(0);
  @$pb.TagNumber(1)
  set postmark(Postmark v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasPostmark() => $_has(0);
  @$pb.TagNumber(1)
  void clearPostmark() => clearField(1);
  @$pb.TagNumber(1)
  Postmark ensurePostmark() => $_ensure(0);

  @$pb.TagNumber(2)
  Type get type => $_getN(1);
  @$pb.TagNumber(2)
  set type(Type v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasType() => $_has(1);
  @$pb.TagNumber(2)
  void clearType() => clearField(2);

  @$pb.TagNumber(10)
  Text get text => $_getN(2);
  @$pb.TagNumber(10)
  set text(Text v) { setField(10, v); }
  @$pb.TagNumber(10)
  $core.bool hasText() => $_has(2);
  @$pb.TagNumber(10)
  void clearText() => clearField(10);
  @$pb.TagNumber(10)
  Text ensureText() => $_ensure(2);

  @$pb.TagNumber(11)
  Media get resource => $_getN(3);
  @$pb.TagNumber(11)
  set resource(Media v) { setField(11, v); }
  @$pb.TagNumber(11)
  $core.bool hasResource() => $_has(3);
  @$pb.TagNumber(11)
  void clearResource() => clearField(11);
  @$pb.TagNumber(11)
  Media ensureResource() => $_ensure(3);

  @$pb.TagNumber(12)
  Location get coordinate => $_getN(4);
  @$pb.TagNumber(12)
  set coordinate(Location v) { setField(12, v); }
  @$pb.TagNumber(12)
  $core.bool hasCoordinate() => $_has(4);
  @$pb.TagNumber(12)
  void clearCoordinate() => clearField(12);
  @$pb.TagNumber(12)
  Location ensureCoordinate() => $_ensure(4);
}

class Postbox extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Postbox', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..pc<Post>(1, 'posts', $pb.PbFieldType.PM, subBuilder: Post.create)
  ;

  Postbox._() : super();
  factory Postbox() => create();
  factory Postbox.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Postbox.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Postbox clone() => Postbox()..mergeFromMessage(this);
  Postbox copyWith(void Function(Postbox) updates) => super.copyWith((message) => updates(message as Postbox));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Postbox create() => Postbox._();
  Postbox createEmptyInstance() => create();
  static $pb.PbList<Postbox> createRepeated() => $pb.PbList<Postbox>();
  @$core.pragma('dart2js:noInline')
  static Postbox getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Postbox>(create);
  static Postbox _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Post> get posts => $_getList(0);
}

class Receipts extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Receipts', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..pc<ClientId>(1, 'localIds', $pb.PbFieldType.PM, subBuilder: ClientId.create)
    ..pc<Postmark>(2, 'postmarks', $pb.PbFieldType.PM, subBuilder: Postmark.create)
  ;

  Receipts._() : super();
  factory Receipts() => create();
  factory Receipts.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Receipts.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Receipts clone() => Receipts()..mergeFromMessage(this);
  Receipts copyWith(void Function(Receipts) updates) => super.copyWith((message) => updates(message as Receipts));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Receipts create() => Receipts._();
  Receipts createEmptyInstance() => create();
  static $pb.PbList<Receipts> createRepeated() => $pb.PbList<Receipts>();
  @$core.pragma('dart2js:noInline')
  static Receipts getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Receipts>(create);
  static Receipts _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<ClientId> get localIds => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<Postmark> get postmarks => $_getList(1);
}

class ClientId extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ClientId', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..e<Category>(1, 'category', $pb.PbFieldType.QE, defaultOrMaker: Category.CATEGORY_UNSPECIFIED, valueOf: Category.valueOf, enumValues: Category.values)
    ..a<$fixnum.Int64>(2, 'lastId', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
  ;

  ClientId._() : super();
  factory ClientId() => create();
  factory ClientId.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ClientId.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ClientId clone() => ClientId()..mergeFromMessage(this);
  ClientId copyWith(void Function(ClientId) updates) => super.copyWith((message) => updates(message as ClientId));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ClientId create() => ClientId._();
  ClientId createEmptyInstance() => create();
  static $pb.PbList<ClientId> createRepeated() => $pb.PbList<ClientId>();
  @$core.pragma('dart2js:noInline')
  static ClientId getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ClientId>(create);
  static ClientId _defaultInstance;

  @$pb.TagNumber(1)
  Category get category => $_getN(0);
  @$pb.TagNumber(1)
  set category(Category v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasCategory() => $_has(0);
  @$pb.TagNumber(1)
  void clearCategory() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get lastId => $_getI64(1);
  @$pb.TagNumber(2)
  set lastId($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLastId() => $_has(1);
  @$pb.TagNumber(2)
  void clearLastId() => clearField(2);
}

class ClientIds extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ClientIds', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..pc<ClientId>(1, 'ids', $pb.PbFieldType.PM, subBuilder: ClientId.create)
  ;

  ClientIds._() : super();
  factory ClientIds() => create();
  factory ClientIds.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ClientIds.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ClientIds clone() => ClientIds()..mergeFromMessage(this);
  ClientIds copyWith(void Function(ClientIds) updates) => super.copyWith((message) => updates(message as ClientIds));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ClientIds create() => ClientIds._();
  ClientIds createEmptyInstance() => create();
  static $pb.PbList<ClientIds> createRepeated() => $pb.PbList<ClientIds>();
  @$core.pragma('dart2js:noInline')
  static ClientIds getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ClientIds>(create);
  static ClientIds _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<ClientId> get ids => $_getList(0);
}

enum Stream_Payload {
  local, 
  postbox, 
  receipts, 
  notSet
}

class Stream extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, Stream_Payload> _Stream_PayloadByTag = {
    10 : Stream_Payload.local,
    11 : Stream_Payload.postbox,
    12 : Stream_Payload.receipts,
    0 : Stream_Payload.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Stream', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..oo(0, [10, 11, 12])
    ..aQS(1, 'sn')
    ..a<$fixnum.Int64>(2, 'target', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..e<PayloadType>(3, 'type', $pb.PbFieldType.QE, defaultOrMaker: PayloadType.PAYLOAD_TYPE_UNSPECIFIED, valueOf: PayloadType.valueOf, enumValues: PayloadType.values)
    ..aOM<ClientIds>(10, 'local', subBuilder: ClientIds.create)
    ..aOM<Postbox>(11, 'postbox', subBuilder: Postbox.create)
    ..aOM<Receipts>(12, 'receipts', subBuilder: Receipts.create)
  ;

  Stream._() : super();
  factory Stream() => create();
  factory Stream.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Stream.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Stream clone() => Stream()..mergeFromMessage(this);
  Stream copyWith(void Function(Stream) updates) => super.copyWith((message) => updates(message as Stream));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Stream create() => Stream._();
  Stream createEmptyInstance() => create();
  static $pb.PbList<Stream> createRepeated() => $pb.PbList<Stream>();
  @$core.pragma('dart2js:noInline')
  static Stream getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Stream>(create);
  static Stream _defaultInstance;

  Stream_Payload whichPayload() => _Stream_PayloadByTag[$_whichOneof(0)];
  void clearPayload() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $core.String get sn => $_getSZ(0);
  @$pb.TagNumber(1)
  set sn($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSn() => $_has(0);
  @$pb.TagNumber(1)
  void clearSn() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get target => $_getI64(1);
  @$pb.TagNumber(2)
  set target($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTarget() => $_has(1);
  @$pb.TagNumber(2)
  void clearTarget() => clearField(2);

  @$pb.TagNumber(3)
  PayloadType get type => $_getN(2);
  @$pb.TagNumber(3)
  set type(PayloadType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasType() => $_has(2);
  @$pb.TagNumber(3)
  void clearType() => clearField(3);

  @$pb.TagNumber(10)
  ClientIds get local => $_getN(3);
  @$pb.TagNumber(10)
  set local(ClientIds v) { setField(10, v); }
  @$pb.TagNumber(10)
  $core.bool hasLocal() => $_has(3);
  @$pb.TagNumber(10)
  void clearLocal() => clearField(10);
  @$pb.TagNumber(10)
  ClientIds ensureLocal() => $_ensure(3);

  @$pb.TagNumber(11)
  Postbox get postbox => $_getN(4);
  @$pb.TagNumber(11)
  set postbox(Postbox v) { setField(11, v); }
  @$pb.TagNumber(11)
  $core.bool hasPostbox() => $_has(4);
  @$pb.TagNumber(11)
  void clearPostbox() => clearField(11);
  @$pb.TagNumber(11)
  Postbox ensurePostbox() => $_ensure(4);

  @$pb.TagNumber(12)
  Receipts get receipts => $_getN(5);
  @$pb.TagNumber(12)
  set receipts(Receipts v) { setField(12, v); }
  @$pb.TagNumber(12)
  $core.bool hasReceipts() => $_has(5);
  @$pb.TagNumber(12)
  void clearReceipts() => clearField(12);
  @$pb.TagNumber(12)
  Receipts ensureReceipts() => $_ensure(5);
}

class Connections extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Connections', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..aQS(1, 'product')
    ..pPS(2, 'uniqueIds')
  ;

  Connections._() : super();
  factory Connections() => create();
  factory Connections.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Connections.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Connections clone() => Connections()..mergeFromMessage(this);
  Connections copyWith(void Function(Connections) updates) => super.copyWith((message) => updates(message as Connections));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Connections create() => Connections._();
  Connections createEmptyInstance() => create();
  static $pb.PbList<Connections> createRepeated() => $pb.PbList<Connections>();
  @$core.pragma('dart2js:noInline')
  static Connections getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Connections>(create);
  static Connections _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get product => $_getSZ(0);
  @$pb.TagNumber(1)
  set product($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasProduct() => $_has(0);
  @$pb.TagNumber(1)
  void clearProduct() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.String> get uniqueIds => $_getList(1);
}

class Tag extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Tag', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..aQS(1, 'product')
    ..aQS(2, 'name')
    ..aOB(3, 'prefixed')
  ;

  Tag._() : super();
  factory Tag() => create();
  factory Tag.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Tag.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Tag clone() => Tag()..mergeFromMessage(this);
  Tag copyWith(void Function(Tag) updates) => super.copyWith((message) => updates(message as Tag));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Tag create() => Tag._();
  Tag createEmptyInstance() => create();
  static $pb.PbList<Tag> createRepeated() => $pb.PbList<Tag>();
  @$core.pragma('dart2js:noInline')
  static Tag getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Tag>(create);
  static Tag _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get product => $_getSZ(0);
  @$pb.TagNumber(1)
  set product($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasProduct() => $_has(0);
  @$pb.TagNumber(1)
  void clearProduct() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get prefixed => $_getBF(2);
  @$pb.TagNumber(3)
  set prefixed($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasPrefixed() => $_has(2);
  @$pb.TagNumber(3)
  void clearPrefixed() => clearField(3);
}

enum Package_Target {
  tag, 
  conn, 
  notSet
}

class Package extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, Package_Target> _Package_TargetByTag = {
    1 : Package_Target.tag,
    2 : Package_Target.conn,
    0 : Package_Target.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Package', package: const $pb.PackageName('appootb.message'), createEmptyInstance: create)
    ..oo(0, [1, 2])
    ..aOM<Tag>(1, 'tag', subBuilder: Tag.create)
    ..aOM<Connections>(2, 'conn', subBuilder: Connections.create)
    ..aQS(8, 'sn')
    ..pc<Post>(9, 'posts', $pb.PbFieldType.PM, subBuilder: Post.create)
  ;

  Package._() : super();
  factory Package() => create();
  factory Package.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Package.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Package clone() => Package()..mergeFromMessage(this);
  Package copyWith(void Function(Package) updates) => super.copyWith((message) => updates(message as Package));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Package create() => Package._();
  Package createEmptyInstance() => create();
  static $pb.PbList<Package> createRepeated() => $pb.PbList<Package>();
  @$core.pragma('dart2js:noInline')
  static Package getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Package>(create);
  static Package _defaultInstance;

  Package_Target whichTarget() => _Package_TargetByTag[$_whichOneof(0)];
  void clearTarget() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  Tag get tag => $_getN(0);
  @$pb.TagNumber(1)
  set tag(Tag v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTag() => $_has(0);
  @$pb.TagNumber(1)
  void clearTag() => clearField(1);
  @$pb.TagNumber(1)
  Tag ensureTag() => $_ensure(0);

  @$pb.TagNumber(2)
  Connections get conn => $_getN(1);
  @$pb.TagNumber(2)
  set conn(Connections v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasConn() => $_has(1);
  @$pb.TagNumber(2)
  void clearConn() => clearField(2);
  @$pb.TagNumber(2)
  Connections ensureConn() => $_ensure(1);

  @$pb.TagNumber(8)
  $core.String get sn => $_getSZ(2);
  @$pb.TagNumber(8)
  set sn($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(8)
  $core.bool hasSn() => $_has(2);
  @$pb.TagNumber(8)
  void clearSn() => clearField(8);

  @$pb.TagNumber(9)
  $core.List<Post> get posts => $_getList(3);
}

