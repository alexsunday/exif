package exif

type EntryFormat int

const (
	FormatUnsignedByte     EntryFormat = 1
	FormatAscii            EntryFormat = 2
	FormatUnsignedShort    EntryFormat = 3
	FormatUnsignedLong     EntryFormat = 4
	FormatUnsignedRational EntryFormat = 5
	FormatSignedByte       EntryFormat = 6
	FormatUndefined        EntryFormat = 7
	FormatSignedShort      EntryFormat = 8
	FormatSignedLong       EntryFormat = 9
	FormatSignedRational   EntryFormat = 10
	FormatFloat            EntryFormat = 11
	FormatDouble           EntryFormat = 12
)

type Ifd uint16

const (
	Ifd0 Ifd = 0 + iota
	Ifd1
	IfdExif
	IfdGps
	IfdInterOperability
	IfdMaxCount
)

type Tag uint16

const (
	EXIF_TAG_INTEROPERABILITY_INDEX                   Tag = 0x0001
	EXIF_TAG_INTEROPERABILITY_VERSION                 Tag = 0x0002
	EXIF_TAG_NEW_SUBFILE_TYPE                         Tag = 0x00fe
	EXIF_TAG_IMAGE_WIDTH                              Tag = 0x0100
	EXIF_TAG_IMAGE_LENGTH                             Tag = 0x0101
	EXIF_TAG_BITS_PER_SAMPLE                          Tag = 0x0102
	EXIF_TAG_COMPRESSION                              Tag = 0x0103
	EXIF_TAG_PHOTOMETRIC_INTERPRETATION               Tag = 0x0106
	EXIF_TAG_FILL_ORDER                               Tag = 0x010a
	EXIF_TAG_DOCUMENT_NAME                            Tag = 0x010d
	EXIF_TAG_IMAGE_DESCRIPTION                        Tag = 0x010e
	EXIF_TAG_MAKE                                     Tag = 0x010f
	EXIF_TAG_MODEL                                    Tag = 0x0110
	EXIF_TAG_STRIP_OFFSETS                            Tag = 0x0111
	EXIF_TAG_ORIENTATION                              Tag = 0x0112
	EXIF_TAG_SAMPLES_PER_PIXEL                        Tag = 0x0115
	EXIF_TAG_ROWS_PER_STRIP                           Tag = 0x0116
	EXIF_TAG_STRIP_BYTE_COUNTS                        Tag = 0x0117
	EXIF_TAG_X_RESOLUTION                             Tag = 0x011a
	EXIF_TAG_Y_RESOLUTION                             Tag = 0x011b
	EXIF_TAG_PLANAR_CONFIGURATION                     Tag = 0x011c
	EXIF_TAG_RESOLUTION_UNIT                          Tag = 0x0128
	EXIF_TAG_TRANSFER_FUNCTION                        Tag = 0x012d
	EXIF_TAG_SOFTWARE                                 Tag = 0x0131
	EXIF_TAG_DATE_TIME                                Tag = 0x0132
	EXIF_TAG_ARTIST                                   Tag = 0x013b
	EXIF_TAG_WHITE_POINT                              Tag = 0x013e
	EXIF_TAG_PRIMARY_CHROMATICITIES                   Tag = 0x013f
	EXIF_TAG_SUB_IFDS                                 Tag = 0x014a
	EXIF_TAG_TRANSFER_RANGE                           Tag = 0x0156
	EXIF_TAG_JPEG_PROC                                Tag = 0x0200
	EXIF_TAG_JPEG_INTERCHANGE_FORMAT                  Tag = 0x0201
	EXIF_TAG_JPEG_INTERCHANGE_FORMAT_LENGTH           Tag = 0x0202
	EXIF_TAG_YCBCR_COEFFICIENTS                       Tag = 0x0211
	EXIF_TAG_YCBCR_SUB_SAMPLING                       Tag = 0x0212
	EXIF_TAG_YCBCR_POSITIONING                        Tag = 0x0213
	EXIF_TAG_REFERENCE_BLACK_WHITE                    Tag = 0x0214
	EXIF_TAG_XML_PACKET                               Tag = 0x02bc
	EXIF_TAG_RELATED_IMAGE_FILE_FORMAT                Tag = 0x1000
	EXIF_TAG_RELATED_IMAGE_WIDTH                      Tag = 0x1001
	EXIF_TAG_RELATED_IMAGE_LENGTH                     Tag = 0x1002
	EXIF_TAG_CFA_REPEAT_PATTERN_DIM                   Tag = 0x828d
	EXIF_TAG_CFA_PATTERN                              Tag = 0x828e
	EXIF_TAG_BATTERY_LEVEL                            Tag = 0x828f
	EXIF_TAG_COPYRIGHT                                Tag = 0x8298
	EXIF_TAG_EXPOSURE_TIME                            Tag = 0x829a
	EXIF_TAG_FNUMBER                                  Tag = 0x829d
	EXIF_TAG_IPTC_NAA                                 Tag = 0x83bb
	EXIF_TAG_IMAGE_RESOURCES                          Tag = 0x8649
	EXIF_TAG_EXIF_IFD_POINTER                         Tag = 0x8769
	EXIF_TAG_INTER_COLOR_PROFILE                      Tag = 0x8773
	EXIF_TAG_EXPOSURE_PROGRAM                         Tag = 0x8822
	EXIF_TAG_SPECTRAL_SENSITIVITY                     Tag = 0x8824
	EXIF_TAG_GPS_INFO_IFD_POINTER                     Tag = 0x8825
	EXIF_TAG_ISO_SPEED_RATINGS                        Tag = 0x8827
	EXIF_TAG_OECF                                     Tag = 0x8828
	EXIF_TAG_TIME_ZONE_OFFSET                         Tag = 0x882a
	EXIF_TAG_EXIF_VERSION                             Tag = 0x9000
	EXIF_TAG_DATE_TIME_ORIGINAL                       Tag = 0x9003
	EXIF_TAG_DATE_TIME_DIGITIZED                      Tag = 0x9004
	EXIF_TAG_COMPONENTS_CONFIGURATION                 Tag = 0x9101
	EXIF_TAG_COMPRESSED_BITS_PER_PIXEL                Tag = 0x9102
	EXIF_TAG_SHUTTER_SPEED_VALUE                      Tag = 0x9201
	EXIF_TAG_APERTURE_VALUE                           Tag = 0x9202
	EXIF_TAG_BRIGHTNESS_VALUE                         Tag = 0x9203
	EXIF_TAG_EXPOSURE_BIAS_VALUE                      Tag = 0x9204
	EXIF_TAG_MAX_APERTURE_VALUE                       Tag = 0x9205
	EXIF_TAG_SUBJECT_DISTANCE                         Tag = 0x9206
	EXIF_TAG_METERING_MODE                            Tag = 0x9207
	EXIF_TAG_LIGHT_SOURCE                             Tag = 0x9208
	EXIF_TAG_FLASH                                    Tag = 0x9209
	EXIF_TAG_FOCAL_LENGTH                             Tag = 0x920a
	EXIF_TAG_SUBJECT_AREA                             Tag = 0x9214
	EXIF_TAG_TIFF_EP_STANDARD_ID                      Tag = 0x9216
	EXIF_TAG_MAKER_NOTE                               Tag = 0x927c
	EXIF_TAG_USER_COMMENT                             Tag = 0x9286
	EXIF_TAG_SUB_SEC_TIME                             Tag = 0x9290
	EXIF_TAG_SUB_SEC_TIME_ORIGINAL                    Tag = 0x9291
	EXIF_TAG_SUB_SEC_TIME_DIGITIZED                   Tag = 0x9292
	EXIF_TAG_XP_TITLE                                 Tag = 0x9c9b
	EXIF_TAG_XP_COMMENT                               Tag = 0x9c9c
	EXIF_TAG_XP_AUTHOR                                Tag = 0x9c9d
	EXIF_TAG_XP_KEYWORDS                              Tag = 0x9c9e
	EXIF_TAG_XP_SUBJECT                               Tag = 0x9c9f
	EXIF_TAG_FLASH_PIX_VERSION                        Tag = 0xa000
	EXIF_TAG_COLOR_SPACE                              Tag = 0xa001
	EXIF_TAG_PIXEL_X_DIMENSION                        Tag = 0xa002
	EXIF_TAG_PIXEL_Y_DIMENSION                        Tag = 0xa003
	EXIF_TAG_RELATED_SOUND_FILE                       Tag = 0xa004
	EXIF_TAG_INTEROPERABILITY_IFD_POINTER             Tag = 0xa005
	EXIF_TAG_FLASH_ENERGY                             Tag = 0xa20b
	EXIF_TAG_SPATIAL_FREQUENCY_RESPONSE               Tag = 0xa20c
	EXIF_TAG_FOCAL_PLANE_X_RESOLUTION                 Tag = 0xa20e
	EXIF_TAG_FOCAL_PLANE_Y_RESOLUTION                 Tag = 0xa20f
	EXIF_TAG_FOCAL_PLANE_RESOLUTION_UNIT              Tag = 0xa210
	EXIF_TAG_SUBJECT_LOCATION                         Tag = 0xa214
	EXIF_TAG_EXPOSURE_INDEX                           Tag = 0xa215
	EXIF_TAG_SENSING_METHOD                           Tag = 0xa217
	EXIF_TAG_FILE_SOURCE                              Tag = 0xa300
	EXIF_TAG_SCENE_TYPE                               Tag = 0xa301
	EXIF_TAG_NEW_CFA_PATTERN                          Tag = 0xa302
	EXIF_TAG_CUSTOM_RENDERED                          Tag = 0xa401
	EXIF_TAG_EXPOSURE_MODE                            Tag = 0xa402
	EXIF_TAG_WHITE_BALANCE                            Tag = 0xa403
	EXIF_TAG_DIGITAL_ZOOM_RATIO                       Tag = 0xa404
	EXIF_TAG_FOCAL_LENGTH_IN_35MM_FILM                Tag = 0xa405
	EXIF_TAG_SCENE_CAPTURE_TYPE                       Tag = 0xa406
	EXIF_TAG_GAIN_CONTROL                             Tag = 0xa407
	EXIF_TAG_CONTRAST                                 Tag = 0xa408
	EXIF_TAG_SATURATION                               Tag = 0xa409
	EXIF_TAG_SHARPNESS                                Tag = 0xa40a
	EXIF_TAG_DEVICE_SETTING_DESCRIPTION               Tag = 0xa40b
	EXIF_TAG_SUBJECT_DISTANCE_RANGE                   Tag = 0xa40c
	EXIF_TAG_IMAGE_UNIQUE_ID                          Tag = 0xa420
	EXIF_TAG_CAMERA_OWNER_NAME                        Tag = 0xa430
	EXIF_TAG_BODY_SERIAL_NUMBER                       Tag = 0xa431
	EXIF_TAG_LENS_SPECIFICATION                       Tag = 0xa432
	EXIF_TAG_LENS_MAKE                                Tag = 0xa433
	EXIF_TAG_LENS_MODEL                               Tag = 0xa434
	EXIF_TAG_LENS_SERIAL_NUMBER                       Tag = 0xa435
	EXIF_TAG_COMPOSITE_IMAGE                          Tag = 0xa460
	EXIF_TAG_SOURCE_IMAGE_NUMBER_OF_COMPOSITE_IMAGE   Tag = 0xa461
	EXIF_TAG_SOURCE_EXPOSURE_TIMES_OF_COMPOSITE_IMAGE Tag = 0xa462
	EXIF_TAG_GAMMA                                    Tag = 0xa500
	EXIF_TAG_PRINT_IMAGE_MATCHING                     Tag = 0xc4a5
	EXIF_TAG_PADDING                                  Tag = 0xea1c

	EXIF_TAG_GPS_VERSION_ID          Tag = 0x0000
	EXIF_TAG_GPS_LATITUDE_REF        Tag = 0x0001
	EXIF_TAG_GPS_LATITUDE            Tag = 0x0002
	EXIF_TAG_GPS_LONGITUDE_REF       Tag = 0x0003
	EXIF_TAG_GPS_LONGITUDE           Tag = 0x0004
	EXIF_TAG_GPS_ALTITUDE_REF        Tag = 0x0005
	EXIF_TAG_GPS_ALTITUDE            Tag = 0x0006
	EXIF_TAG_GPS_TIME_STAMP          Tag = 0x0007
	EXIF_TAG_GPS_SATELLITES          Tag = 0x0008
	EXIF_TAG_GPS_STATUS              Tag = 0x0009
	EXIF_TAG_GPS_MEASURE_MODE        Tag = 0x000a
	EXIF_TAG_GPS_DOP                 Tag = 0x000b
	EXIF_TAG_GPS_SPEED_REF           Tag = 0x000c
	EXIF_TAG_GPS_SPEED               Tag = 0x000d
	EXIF_TAG_GPS_TRACK_REF           Tag = 0x000e
	EXIF_TAG_GPS_TRACK               Tag = 0x000f
	EXIF_TAG_GPS_IMG_DIRECTION_REF   Tag = 0x0010
	EXIF_TAG_GPS_IMG_DIRECTION       Tag = 0x0011
	EXIF_TAG_GPS_MAP_DATUM           Tag = 0x0012
	EXIF_TAG_GPS_DEST_LATITUDE_REF   Tag = 0x0013
	EXIF_TAG_GPS_DEST_LATITUDE       Tag = 0x0014
	EXIF_TAG_GPS_DEST_LONGITUDE_REF  Tag = 0x0015
	EXIF_TAG_GPS_DEST_LONGITUDE      Tag = 0x0016
	EXIF_TAG_GPS_DEST_BEARING_REF    Tag = 0x0017
	EXIF_TAG_GPS_DEST_BEARING        Tag = 0x0018
	EXIF_TAG_GPS_DEST_DISTANCE_REF   Tag = 0x0019
	EXIF_TAG_GPS_DEST_DISTANCE       Tag = 0x001a
	EXIF_TAG_GPS_PROCESSING_METHOD   Tag = 0x001b
	EXIF_TAG_GPS_AREA_INFORMATION    Tag = 0x001c
	EXIF_TAG_GPS_DATE_STAMP          Tag = 0x001d
	EXIF_TAG_GPS_DIFFERENTIAL        Tag = 0x001e
	EXIF_TAG_GPS_H_POSITIONING_ERROR Tag = 0x001f
)
