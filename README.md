# 配置中心

## schema示例

```json
{
  "type": "object",
  "properties": {
    "name": {
      "type": "string",
      "title": "名称",
      "description": "这是名称介绍",
      "default": "哈哈哈哈"
    },
    "email": {
      "type": "string",
      "title": "邮箱",
      "description": "这是名称介绍"
    },
    "address": {
      "type": "string",
      "title": "地址",
      "description": "这是名称介绍"
    },
    "telephone": {
      "type": "string",
      "title": "手机号",
      "description": "这是名称介绍"
    },
    "secret": {
      "type": "boolean",
      "title": "保密",
      "description": "这是名称介绍"
    },
    "qty": {
      "type": "integer",
      "title": "数量",
      "description": "这是名称介绍",
      "minimum": 0,
      "maximum": 10
    },
    "price": {
      "type": "number",
      "title": "价格",
      "description": "这是名称介绍",
      "minimum": 1,
      "maximum": 99
    },
    "children": {
      "type": "object",
      "title": "子对象",
      "description": "子对象",
      "default": "子对象",
      "properties": {
        "name1": {
          "type": "string",
          "title": "名称c",
          "description": "这是名称介绍",
          "default": "哈哈哈哈"
        },
        "email": {
          "type": "string",
          "title": "邮箱c",
          "description": "这是名称介绍"
        },
        "children": {
          "type": "object",
          "title": "子对象",
          "description": "子对象",
          "default": "子对象",
          "properties": {
            "name1": {
              "type": "string",
              "title": "名称c",
              "description": "这是名称介绍",
              "default": "哈哈哈哈"
            },
            "email": {
              "type": "string",
              "title": "邮箱c",
              "description": "这是名称介绍"
            }
          },
          "required": [
            "name1",
            "email"
          ]
        }
      },
      "required": [
        "name1",
        "email"
      ]
    },
    "item_sku_codes": {
      "type": "array",
      "title": "商品列表",
      "description": "这是名称介绍",
      "default": "哈哈哈哈",
      "items": {
        "type": "string",
        "title": "商品编码",
        "description": "这是名称介绍"
      }
    },
    "coupon": {
      "type": "array",
      "title": "优惠券列表",
      "description": "这是名称介绍",
      "default": "哈哈哈哈",
      "items": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string",
            "title": "名称",
            "description": "这是名称介绍"
          },
          "email": {
            "type": "string",
            "title": "邮箱",
            "description": "这是名称介绍"
          },
          "secret": {
            "type": "boolean",
            "title": "保密",
            "description": "这是名称介绍"
          }
        }
      }
    },
    "required": [
      "name",
      "email"
    ]
  }
}
```

## TODO

- 第一优先级
  - 日期类型支持
    - 2004-05-03T17:30:08+08:00
    - 2004-05-03
    - 时间戳
  - KeyValue数据支持，例如：选择类型名称，保存类型ID
  - 选择框支持远程接口数据 接口返回数据列表 支持搜索等
  - 完善schema表单参数校验 正则
  - 支持表单字段依赖性配置

- 第二优先级
  - 支持redis缓存
  - 支持HTTP方式接口通过KEY访问非敏感数据。
  - 支持集群化和应用内存级别缓存
    - 数据更新时添加redis发布订阅，各节点更新并更新本地缓存
  - 支持文件缓存
    - 数据更新后直接缓存成文件key.json，提供HTTP直接访问缓存文件，后期接入对象存储中心支持直接上传。
  - 提供GO语言SDK（GRPC?）或 redis直连

- 第三优先级
  - 支持级联
  - 对象存储管理，并支持在表单中选择对应对象
    - 接入各厂商对象存储
  - 支持权限管理