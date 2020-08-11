CREATE TABLE `record_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `track_id` varchar(128) NOT NULL DEFAULT ' ' COMMENT '轨迹id',
  `start_time` bigint(20) unsigned DEFAULT NULL,
  `end_time` bigint(20) unsigned DEFAULT NULL,
  `videoname` varchar(255) DEFAULT NULL  comment '视频名称',
  PRIMARY KEY (`id`),
  KEY `idx_annotation_info_deleted_at` (`deleted_at`),
  KEY `idx_track` (`track_id`),
  KEY `id_video_label_track` (`videoname`,`track_id`),
  KEY `id_start_end_tm` (`start_time`, `end_time`)
) ENGINE=InnoDB AUTO_INCREMENT=2601 DEFAULT CHARSET=latin1