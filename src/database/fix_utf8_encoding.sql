-- 修复 UTF-8 双重编码导致的中文乱码
-- 适用于 MySQL 客户端以 latin1 导入 utf8mb4 数据后产生的 mojibake
SET NAMES utf8mb4;
SET CHARACTER SET utf8mb4;

USE tennis_booking;

UPDATE courts SET
  name = CONVERT(BINARY CONVERT(name USING latin1) USING utf8mb4),
  location = CONVERT(BINARY CONVERT(location USING latin1) USING utf8mb4),
  description = CONVERT(BINARY CONVERT(description USING latin1) USING utf8mb4),
  facilities = CONVERT(BINARY CONVERT(facilities USING latin1) USING utf8mb4)
WHERE name REGEXP '[\\x80-\\xFF]';

UPDATE admins SET
  real_name = CONVERT(BINARY CONVERT(real_name USING latin1) USING utf8mb4)
WHERE real_name REGEXP '[\\x80-\\xFF]';

UPDATE system_configs SET
  description = CONVERT(BINARY CONVERT(description USING latin1) USING utf8mb4)
WHERE description REGEXP '[\\x80-\\xFF]';

UPDATE shops SET
  shop_name = CONVERT(BINARY CONVERT(shop_name USING latin1) USING utf8mb4),
  description = CONVERT(BINARY CONVERT(description USING latin1) USING utf8mb4),
  address = CONVERT(BINARY CONVERT(address USING latin1) USING utf8mb4),
  manager_name = CONVERT(BINARY CONVERT(manager_name USING latin1) USING utf8mb4),
  business_hours = CONVERT(BINARY CONVERT(business_hours USING latin1) USING utf8mb4)
WHERE shop_name REGEXP '[\\x80-\\xFF]';
