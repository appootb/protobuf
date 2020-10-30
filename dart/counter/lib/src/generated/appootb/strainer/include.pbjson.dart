///
//  Generated code. Do not modify.
//  source: appootb/strainer/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const Channel$json = const {
  '1': 'Channel',
  '2': const [
    const {'1': 'CHANNEL_UNSPECIFIED', '2': 0},
  ],
};

const Target$json = const {
  '1': 'Target',
  '2': const [
    const {'1': 'TARGET_UNSPECIFIED', '2': 0},
    const {'1': 'TARGET_NICKNAME', '2': 1},
    const {'1': 'TARGET_AVATAR', '2': 2},
    const {'1': 'TARGET_SIGNATURE', '2': 3},
    const {'1': 'TARGET_TEXT', '2': 4},
    const {'1': 'TARGET_IMAGE', '2': 5},
    const {'1': 'TARGET_VIDEO', '2': 6},
    const {'1': 'TARGET_AUDIO', '2': 7},
  ],
};

const FilterRequest$json = const {
  '1': 'FilterRequest',
  '2': const [
    const {'1': 'channel', '3': 1, '4': 1, '5': 14, '6': '.appootb.strainer.Channel', '10': 'channel'},
    const {'1': 'type', '3': 2, '4': 2, '5': 14, '6': '.appootb.strainer.Target', '10': 'type'},
    const {'1': 'data', '3': 3, '4': 2, '5': 9, '10': 'data'},
  ],
};

