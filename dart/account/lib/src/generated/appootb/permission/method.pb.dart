///
//  Generated code. Do not modify.
//  source: method.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'method.pbenum.dart';

export 'method.pbenum.dart';

class Token extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Token', package: const $pb.PackageName('appootb.permission.method'), createEmptyInstance: create)
    ..e<TokenLevel>(1, 'required', $pb.PbFieldType.OE, defaultOrMaker: TokenLevel.NONE_TOKEN, valueOf: TokenLevel.valueOf, enumValues: TokenLevel.values)
    ..hasRequiredFields = false
  ;

  Token._() : super();
  factory Token() => create();
  factory Token.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Token.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Token clone() => Token()..mergeFromMessage(this);
  Token copyWith(void Function(Token) updates) => super.copyWith((message) => updates(message as Token));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Token create() => Token._();
  Token createEmptyInstance() => create();
  static $pb.PbList<Token> createRepeated() => $pb.PbList<Token>();
  @$core.pragma('dart2js:noInline')
  static Token getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Token>(create);
  static Token _defaultInstance;

  @$pb.TagNumber(1)
  TokenLevel get required => $_getN(0);
  @$pb.TagNumber(1)
  set required(TokenLevel v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasRequired() => $_has(0);
  @$pb.TagNumber(1)
  void clearRequired() => clearField(1);
}

class Method {
  static final $pb.Extension token = $pb.Extension<Token>('google.protobuf.MethodOptions', 'token', 2507, $pb.PbFieldType.OM, defaultOrMaker: Token.getDefault, subBuilder: Token.create);
  static void registerAllExtensions($pb.ExtensionRegistry registry) {
    registry.add(token);
  }
}

