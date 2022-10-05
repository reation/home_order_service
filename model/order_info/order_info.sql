CREATE TABLE `order_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_num` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `u_id` int(10) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `order_price` float NOT NULL DEFAULT 0 COMMENT '订单金额',
  `discount_price` float NOT NULL COMMENT '优惠金额',
  `real_price` float NOT NULL DEFAULT 0 COMMENT '实际支付金额',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '订单状态 1:待支付，2：已支付，3：取消，4：退款中，5：已退款，6：已完成',
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `order_num`(`order_num`) USING BTREE,
  INDEX `u_id`(`u_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci;

