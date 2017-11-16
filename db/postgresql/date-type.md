
### pg数据类型

数据类型名称                           | 别名           |  简单描述
----------------------------------  | -------------  | ------------
bigint	                            | int8           |	有符号8字节整数
bigserial	                        | serial8	     |  自增8字节整数
bit [(n)]	                        |                |  固定长度位串
bit varying [(n)]	                | varbit	     |  可变长度位串
boolean	                            | bool	         |  布尔值 (true/false)
box	 	                            |                |  rectangular box on a plane
bytea	 	                        |                |  binary data (“byte array”)
character [(n)]	                    | char [(n)]	 |  定长字符串
character varying [(n)]	            | varchar [(n)]	 |  变长字符串
cidr	 	                        |                |  IPv4 或 IPv6 网络地址
circle	 	                        |                |  circle on a plane
date	 	                        |                |  日历日期 (year, month, day)
double precision	                |float8	         |  双精度浮点数(8字节)
inet	 	                        |                |  IPv4 或 IPv6 主机地址
integer	                            | int, int4	     |  有符号4字节整数
interval [fields] [(p)]	 	        |                |  时间间隔
json	 	                        |                |  文本JSON数据 写入快 读取慢
jsonb	 	                        |                |  分解的JSON二进制数据 写入慢 读取快
line	 	                        |                |  infinite line on a plane
lseg	 	                        |                |  line segment on a plane
macaddr	 	                        |                |  MAC (Media Access Control) address
macaddr8	 	                    |                |  MAC (Media Access Control) address (EUI-64 format)
money	 	                        |                |  货币金额
numeric [(p, s)]	                | decimal[(p, s)]|	可选精度的准确数值数据类型
path	 	                        |                |  geometric path on a plane
pg_lsn	 	                        |                |  PostgreSQL Log Sequence Number
point	 	                        |                |  geometric point on a plane
polygon	 	                        |                |  closed geometric path on a plane
real	                            | float4	     |  single precision floating-point number (4 bytes)
smallint	                        | int2	         |  有符号2字节整数
smallserial	                        | serial2	     |  自增2字节整数
serial	                            | serial4	     |  自增4字节整数
text	 	                        |                |  可变长字符串
time [(p)] [without time zone]	    |                |  time of day (no time zone)
time [(p)] with time zone	        | timetz	     |  time of day, including time zone
timestamp [(p)] [without time zone]	|                |  date and time (no time zone)
timestamp [(p)] with time zone	    | timestamptz	 |  date and time, including time zone
tsquery	 	                        |                |  text search query
tsvector	 	                    |                |  text search document
txid_snapshot	 	                |                |  user-level transaction ID snapshot
uuid	 	                        |                |  唯一ID
xml	 	                            | XML            |  XML数据