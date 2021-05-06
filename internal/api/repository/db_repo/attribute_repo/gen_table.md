#### saas_mis_split_4.attribute 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint(20) unsigned | PRI | NO | auto_increment |  |
| 2 | spu_id | 主题编号 | varchar(255) |  | NO |  |  |
| 3 | opt_name | 标题 | varchar(255) |  | NO |  |  |
| 4 | img | 详情图 | varchar(255) |  | NO |  |  |
| 5 | value | 值 | varchar(255) |  | NO |  |  |
| 6 | status | 是否禁用   1启用   非1禁用 | tinyint(4) |  | NO |  | 1 |
| 7 | created_at |  | timestamp |  | YES |  |  |
| 8 | updated_at |  | timestamp |  | YES |  |  |
