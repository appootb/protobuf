///
//  Generated code. Do not modify.
//  source: appootb/message/include.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const Category$json = const {
  '1': 'Category',
  '2': const [
    const {'1': 'CATEGORY_UNSPECIFIED', '2': 0},
    const {'1': 'CATEGORY_PRIVATE', '2': 1},
    const {'1': 'CATEGORY_GROUP', '2': 2},
    const {'1': 'CATEGORY_CHANNEL', '2': 3},
    const {'1': 'CATEGORY_ROOM', '2': 4},
  ],
};

const Type$json = const {
  '1': 'Type',
  '2': const [
    const {'1': 'TYPE_UNSPECIFIED', '2': 0},
    const {'1': 'TYPE_TEXT', '2': 1},
    const {'1': 'TYPE_AUDIO', '2': 2},
    const {'1': 'TYPE_STICKER', '2': 3},
    const {'1': 'TYPE_PHOTO', '2': 4},
    const {'1': 'TYPE_VIDEO', '2': 5},
    const {'1': 'TYPE_GIF', '2': 6},
    const {'1': 'TYPE_FILE', '2': 7},
    const {'1': 'TYPE_LOCATION', '2': 8},
    const {'1': 'TYPE_CONTACT', '2': 9},
    const {'1': 'TYPE_RECALL', '2': 10},
    const {'1': 'TYPE_READ', '2': 11},
  ],
};

const PayloadType$json = const {
  '1': 'PayloadType',
  '2': const [
    const {'1': 'PAYLOAD_TYPE_UNSPECIFIED', '2': 0},
    const {'1': 'PAYLOAD_TYPE_SYNC', '2': 1},
    const {'1': 'PAYLOAD_TYPE_MESSAGES', '2': 2},
    const {'1': 'PAYLOAD_TYPE_RECEIPTS', '2': 3},
  ],
};

const Postmark$json = const {
  '1': 'Postmark',
  '2': const [
    const {'1': 'category', '3': 1, '4': 2, '5': 14, '6': '.appootb.message.Category', '10': 'category'},
    const {'1': 'sender', '3': 2, '4': 2, '5': 4, '10': 'sender'},
    const {'1': 'receiver', '3': 3, '4': 2, '5': 4, '10': 'receiver'},
    const {'1': 'unique_id', '3': 4, '4': 1, '5': 4, '10': 'uniqueId'},
    const {'1': 'expire', '3': 8, '4': 1, '5': 11, '6': '.google.protobuf.Duration', '10': 'expire'},
    const {'1': 'timestamp', '3': 9, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'timestamp'},
  ],
};

const Text$json = const {
  '1': 'Text',
  '2': const [
    const {'1': 'content', '3': 1, '4': 2, '5': 9, '10': 'content'},
  ],
};

const Media$json = const {
  '1': 'Media',
  '2': const [
    const {'1': 'url', '3': 1, '4': 2, '5': 9, '10': 'url'},
    const {'1': 'thumbnail', '3': 2, '4': 1, '5': 9, '10': 'thumbnail'},
  ],
};

const Location$json = const {
  '1': 'Location',
  '2': const [
    const {'1': 'latitude', '3': 1, '4': 2, '5': 2, '10': 'latitude'},
    const {'1': 'longitude', '3': 2, '4': 2, '5': 2, '10': 'longitude'},
  ],
};

const Post$json = const {
  '1': 'Post',
  '2': const [
    const {'1': 'postmark', '3': 1, '4': 2, '5': 11, '6': '.appootb.message.Postmark', '10': 'postmark'},
    const {'1': 'type', '3': 2, '4': 2, '5': 14, '6': '.appootb.message.Type', '10': 'type'},
    const {'1': 'text', '3': 10, '4': 1, '5': 11, '6': '.appootb.message.Text', '9': 0, '10': 'text'},
    const {'1': 'resource', '3': 11, '4': 1, '5': 11, '6': '.appootb.message.Media', '9': 0, '10': 'resource'},
    const {'1': 'coordinate', '3': 12, '4': 1, '5': 11, '6': '.appootb.message.Location', '9': 0, '10': 'coordinate'},
  ],
  '8': const [
    const {'1': 'envelope'},
  ],
};

const Postbox$json = const {
  '1': 'Postbox',
  '2': const [
    const {'1': 'posts', '3': 1, '4': 3, '5': 11, '6': '.appootb.message.Post', '10': 'posts'},
  ],
};

const Receipts$json = const {
  '1': 'Receipts',
  '2': const [
    const {'1': 'local_ids', '3': 1, '4': 3, '5': 11, '6': '.appootb.message.ClientId', '10': 'localIds'},
    const {'1': 'postmarks', '3': 2, '4': 3, '5': 11, '6': '.appootb.message.Postmark', '10': 'postmarks'},
  ],
};

const ClientId$json = const {
  '1': 'ClientId',
  '2': const [
    const {'1': 'category', '3': 1, '4': 2, '5': 14, '6': '.appootb.message.Category', '10': 'category'},
    const {'1': 'last_id', '3': 2, '4': 1, '5': 4, '10': 'lastId'},
  ],
};

const ClientIds$json = const {
  '1': 'ClientIds',
  '2': const [
    const {'1': 'ids', '3': 1, '4': 3, '5': 11, '6': '.appootb.message.ClientId', '10': 'ids'},
  ],
};

const Stream$json = const {
  '1': 'Stream',
  '2': const [
    const {'1': 'sn', '3': 1, '4': 2, '5': 9, '10': 'sn'},
    const {'1': 'target', '3': 2, '4': 1, '5': 4, '10': 'target'},
    const {'1': 'type', '3': 3, '4': 2, '5': 14, '6': '.appootb.message.PayloadType', '10': 'type'},
    const {'1': 'local', '3': 10, '4': 1, '5': 11, '6': '.appootb.message.ClientIds', '9': 0, '10': 'local'},
    const {'1': 'postbox', '3': 11, '4': 1, '5': 11, '6': '.appootb.message.Postbox', '9': 0, '10': 'postbox'},
    const {'1': 'receipts', '3': 12, '4': 1, '5': 11, '6': '.appootb.message.Receipts', '9': 0, '10': 'receipts'},
  ],
  '8': const [
    const {'1': 'payload'},
  ],
};

const Connections$json = const {
  '1': 'Connections',
  '2': const [
    const {'1': 'product', '3': 1, '4': 2, '5': 9, '10': 'product'},
    const {'1': 'unique_ids', '3': 2, '4': 3, '5': 9, '10': 'uniqueIds'},
  ],
};

const Tag$json = const {
  '1': 'Tag',
  '2': const [
    const {'1': 'product', '3': 1, '4': 2, '5': 9, '10': 'product'},
    const {'1': 'name', '3': 2, '4': 2, '5': 9, '10': 'name'},
    const {'1': 'prefixed', '3': 3, '4': 1, '5': 8, '10': 'prefixed'},
  ],
};

const Package$json = const {
  '1': 'Package',
  '2': const [
    const {'1': 'tag', '3': 1, '4': 1, '5': 11, '6': '.appootb.message.Tag', '9': 0, '10': 'tag'},
    const {'1': 'conn', '3': 2, '4': 1, '5': 11, '6': '.appootb.message.Connections', '9': 0, '10': 'conn'},
    const {'1': 'sn', '3': 8, '4': 2, '5': 9, '10': 'sn'},
    const {'1': 'posts', '3': 9, '4': 3, '5': 11, '6': '.appootb.message.Post', '10': 'posts'},
  ],
  '8': const [
    const {'1': 'target'},
  ],
};

