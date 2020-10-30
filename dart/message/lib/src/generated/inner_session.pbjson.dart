///
//  Generated code. Do not modify.
//  source: inner_session.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const SessionType$json = const {
  '1': 'SessionType',
  '2': const [
    const {'1': 'SESSION_TYPE_UNSPECIFIED', '2': 0},
    const {'1': 'SESSION_TYPE_CHAT', '2': 1},
    const {'1': 'SESSION_TYPE_ROOM', '2': 2},
  ],
};

const SessionId$json = const {
  '1': 'SessionId',
  '2': const [
    const {'1': 'type', '3': 1, '4': 2, '5': 14, '6': '.appootb.message.SessionType', '10': 'type'},
    const {'1': 'account', '3': 2, '4': 2, '5': 4, '10': 'account'},
    const {'1': 'device_id', '3': 3, '4': 1, '5': 9, '10': 'deviceId'},
  ],
};

const SessionIds$json = const {
  '1': 'SessionIds',
  '2': const [
    const {'1': 'ids', '3': 1, '4': 3, '5': 11, '6': '.appootb.message.SessionId', '10': 'ids'},
  ],
};

const Session$json = const {
  '1': 'Session',
  '2': const [
    const {'1': 'type', '3': 1, '4': 2, '5': 14, '6': '.appootb.message.SessionType', '10': 'type'},
    const {'1': 'entry', '3': 2, '4': 2, '5': 9, '10': 'entry'},
    const {'1': 'metadata', '3': 3, '4': 2, '5': 11, '6': '.appootb.common.Metadata', '10': 'metadata'},
    const {'1': 'created_at', '3': 4, '4': 2, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'createdAt'},
  ],
};

const Sessions$json = const {
  '1': 'Sessions',
  '2': const [
    const {'1': 'sessions', '3': 1, '4': 3, '5': 11, '6': '.appootb.message.Session', '10': 'sessions'},
  ],
};

