import { PageContainer } from '@ant-design/pro-layout';
import type { ReactNode } from 'react';
import { useState, useEffect } from 'react';
import {
    FormInstance, Table, Tabs
} from 'antd';
import {
    Row, Col, Form,
    Input,
    Button,
    Select,
    Switch,
    Space
} from 'antd';
import styles from './index.less';

import type { JSONSchema7 } from 'json-schema';
import { CheckOutlined, CloseOutlined, MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';
import { addStruct, findStruct } from '@/services/config/struct';
import type config from 'config/config';
import JsonSchemaForm from '@/components/JsonSchemaFrom';
import type { ProColumns } from '@ant-design/pro-table';
import ProTable from '@ant-design/pro-table';
import moment from 'moment';

// Make modifications to the theme with your own fields and widgets

const formItemLayout = {
    labelCol: { span: 2 },
    wrapperCol: { span: 22 },
};

const { TabPane } = Tabs;

export default () => {

    const [form] = Form.useForm();

    const columnKey = {
        title: "唯一标识",
        dataIndex: "key",
        key: "key",
        description: "将作为配置数据存储和获取的唯一标识，供程序识别",
        tooltip: "将作为配置数据存储和获取的唯一标识，供程序识别"
    }
    const columnTitle = {
        title: "结构名称",
        dataIndex: "title",
        key: "title",
        description: "配置名称，供配置人员识别",
        tooltip: "配置名称，供配置人员识别"
    }
    const columnDescription = {
        title: "结构说明",
        dataIndex: "description",
        key: "description",
        description: "配置说明，供配置人员了解配置项作用",
        tooltip: "配置说明，供配置人员了解配置项作用",
    }
    const columnSecret = {
        title: "全局保密",
        dataIndex: "secret",
        key: "secret",
        description: "加密的数据无法通过http接口访问，如果仅部分字段加密，此配置不要开启",
        tooltip: "加密的数据无法通过http接口访问，如果仅部分字段加密，此配置不要开启",
        render: (text: ReactNode, record: config.ConfigStruct) => <Switch disabled checked={record.secret} checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} />
    }
    const columnArray = {
        title: "数组类型",
        dataIndex: "array",
        key: "array",
        description: "默认情况下所有配置存储在同一个表中，不支持复杂查询。如有大量重复结构的数据，开启此配置后将独立存储，支持独立列表查询。",
        tooltip: "默认情况下所有配置存储在同一个表中，不支持复杂查询。如有大量重复结构的数据，开启此配置后将独立存储，支持独立列表查询。",
        render: (text: ReactNode, record: config.ConfigStruct) => <Switch disabled checked={record.secret} checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} />
    }
    const columnSchema = {
        title: "结构",
        dataIndex: "schema",
        key: "schema",
        description: "json-schema格式数据",
        hideInTable: true,
        hideInSearch: true,
    }

    const columns: ProColumns<config.ConfigStruct>[] = [
        columnKey,
        columnTitle,
        columnDescription,
        columnSecret,
        columnArray,
        columnSchema,
        {
            title: '创建时间',
            dataIndex: "create_time",
            key: 'create_time',
            hideInSearch: true,
            renderText: (text) => {
                return moment(text).format('YYYY-MM-DD HH:mm:ss')
            }
        },
        {
            title: '修改时间',
            dataIndex: "update_time",
            key: 'update_time',
            hideInSearch: true,
            renderText: (text) => {
                return moment(text).format('YYYY-MM-DD HH:mm:ss')
            }
        },
        {
            title: 'Action',
            key: 'action',
            valueType: 'option',
            render: () => [
                <a key="edit">编辑</a>,
                <a key="edit">历史</a>,
                <a key="edit">查看</a>,
                <a key="edit">查看数据</a>,
            ],
        },
    ];


    const onFinish = (values: config.ConfigStruct) => {
        console.log('表单数据: ', values);
        addStruct(values).then(result => {
            console.log(result)
        })
    };

    const onValuesChange = (changedValues: any, allValues: any) => {
        console.log(changedValues, allValues, form)
    }

    const [schema, setSchema] = useState<any>({})
    useEffect(() => {
        setSchema({
            "type": "object",
            "properties": {
                "name": { "type": "string", title: "名称", description: "这是名称介绍", default: "哈哈哈哈" },
                "email": { "type": "string", title: "邮箱", description: "这是名称介绍" },
                "address": { "type": "string", title: "地址", description: "这是名称介绍" },
                "telephone": { "type": "string", title: "手机号", description: "这是名称介绍" },
                "secret": { "type": "boolean", title: "保密", description: "这是名称介绍" },
                "qty": { "type": "integer", title: "数量", description: "这是名称介绍", minimum: 0, maximum: 10 },
                "price": { "type": "number", title: "价格", description: "这是名称介绍", minimum: 1, maximum: 99 },
                "children": {
                    "type": "object",
                    title: "子对象", description: "子对象", default: "子对象",
                    "properties": {
                        "name1": { "type": "string", title: "名称c", description: "这是名称介绍", default: "哈哈哈哈" },
                        "email": { "type": "string", title: "邮箱c", description: "这是名称介绍" },
                        "children": {
                            "type": "object",
                            title: "子对象", description: "子对象", default: "子对象",
                            "properties": {
                                "name1": { "type": "string", title: "名称c", description: "这是名称介绍", default: "哈哈哈哈" },
                                "email": { "type": "string", title: "邮箱c", description: "这是名称介绍" },
                            },
                            "required": ["name1", "email"]
                        },
                    },
                    "required": ["name1", "email"]
                },
                "item_sku_codes": {
                    "type": "array",
                    title: "商品列表",
                    description: "这是名称介绍",
                    default: "哈哈哈哈",
                    "items": {
                        "type": "string",
                        title: "商品编码",
                        description: "这是名称介绍",
                    },
                },
                "coupon": {
                    "type": "array",
                    title: "优惠券列表",
                    description: "这是名称介绍",
                    default: "哈哈哈哈",
                    "items": {
                        "type": "object",
                        "properties": {
                            "name": { "type": "string", title: "名称", description: "这是名称介绍" },
                            "email": { "type": "string", title: "邮箱", description: "这是名称介绍" },
                            "secret": { "type": "boolean", title: "保密", description: "这是名称介绍" },
                        }
                    }
                },
                "required": ["name", "email"]
            }
        })
    }, []);
    return (
        <PageContainer title={false} className={styles.main}>
            <Tabs defaultActiveKey="1">
                <TabPane tab="查看" key="1">
                    <ProTable columns={columns} request={(params) => (findStruct({ page_num: params.current, page_size: params.pageSize }).then((result) => { return { data: result } }))}></ProTable>
                </TabPane>
                <TabPane tab="编辑" key="2">
                    <Form
                        form={form}
                        name="validate_other"
                        {...formItemLayout}
                        onFinish={onFinish}
                        onValuesChange={onValuesChange}
                        initialValues={{
                            "title": "title",
                            "description": "description",
                            "secret": false,
                            "array": false,
                            "schema": ""
                        }}
                    >
                        <Form.Item name="key" wrapperCol={{ span: 4 }} required label={columnKey.title} tooltip={columnKey.description}
                            rules={[{ required: true, message: "请输入结构标识" }, { pattern: /^\w+$/, message: '仅支持数字、字母及下划线' }]}
                        >
                            <Input />
                        </Form.Item>
                        <Form.Item name="title" wrapperCol={{ span: 8 }} required label={columnTitle.title} tooltip={columnTitle.description}
                            rules={[{ required: true, message: "请输入结构名称" }]}
                        >
                            <Input />
                        </Form.Item>
                        <Form.Item name="description" wrapperCol={{ span: 8 }} label={columnDescription.title} tooltip={columnDescription.description}
                        >
                            <Input.TextArea />
                        </Form.Item>
                        <Form.Item name="secret" valuePropName="checked" label={columnSecret.title} tooltip={columnSecret.description}>
                            <Switch checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} />
                        </Form.Item>
                        <Form.Item name="array" valuePropName="checked" label={columnArray.title} tooltip={columnArray.description}>
                            <Switch checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} />
                        </Form.Item>
                        <Form.Item wrapperCol={{ span: 8 }} label="结构" tooltip='json-schema格式数据' required>
                            <Form.Item name='schema' noStyle>
                                <Input.TextArea />
                            </Form.Item>
                            {/* <JsonSchemaForm propKey={["schema"]} schema={schema}></JsonSchemaForm> */}
                        </Form.Item>

                        {/* <Form.Item name="fields" label="结构字段" valuePropName="array" tooltip='数据字段' required>
                    <Form.List name="fields">
                        {(fields, { add, move, remove }) => (
                            <>
                                {fields.map(field => (
                                    <Space key={field.key} align="baseline">
                                        <Form.Item
                                            {...field}
                                            label="标识"
                                            name={[field.name, 'key']}
                                            fieldKey={[field.fieldKey, 'key']}
                                            rules={[{ required: true, message: "请输入唯一标识" }, { pattern: /^\w+$/, message: '仅支持数字、字母及下划线' }]}
                                        >
                                            <Input />
                                        </Form.Item>
                                        <Form.Item
                                            {...field}
                                            label="名称"
                                            name={[field.name, 'title']}
                                            fieldKey={[field.fieldKey, 'title']}
                                            rules={[{ required: true, message: "请输入名称" }]}
                                        >
                                            <Input />
                                        </Form.Item>
                                        <Form.Item
                                            {...field}
                                            label="类型"
                                            name={[field.name, 'type']}
                                            fieldKey={[field.fieldKey, 'type']}
                                            rules={[{ required: true, message: '请选择类型' }]}
                                            initialValue="txt"
                                        >
                                            <Select style={{ width: 100 }}>
                                                <Select.Option value='string'>文本</Select.Option>
                                                <Select.Option value='number'>数字</Select.Option>
                                                <Select.Option value='integer'>手机号</Select.Option>
                                                <Select.Option value='email'>邮箱</Select.Option>
                                                <Select.Option value='image'>图片</Select.Option>
                                            </Select>
                                        </Form.Item>
                                        <Form.Item
                                            {...field}
                                            label="数组"
                                            tooltip='数据是否是数据类型'
                                            name={[field.name, 'array']}
                                            fieldKey={[field.fieldKey, 'array']}
                                            rules={[{ required: false, message: 'Missing required' }]}
                                            initialValue={false}
                                        >
                                            <Switch checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} />
                                        </Form.Item>
                                        <Form.Item
                                            {...field}
                                            label="保密"
                                            tooltip='保密字段无法通过http接口获取，如果父级别数据为保密数据'
                                            name={[field.name, 'secret']}
                                            fieldKey={[field.fieldKey, 'secret']}
                                            rules={[{ required: false, message: 'Missing required' }]}
                                        >
                                            <Switch checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} />
                                        </Form.Item>
                                        <Form.Item
                                            {...field}
                                            label="必填"
                                            name={[field.name, 'required']}
                                            fieldKey={[field.fieldKey, 'required']}
                                            rules={[{ required: false, message: 'Missing required' }]}
                                        >
                                            <Switch checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} />
                                        </Form.Item>
                                        <Form.Item
                                            {...field}
                                            label="说明"
                                            name={[field.name, 'description']}
                                            fieldKey={[field.fieldKey, 'description']}
                                            rules={[{ required: false, message: 'Missing required' }]}
                                        >
                                            <Input />
                                        </Form.Item>
                                        <MinusCircleOutlined onClick={() => remove(field.name)} />
                                    </Space>
                                ))}

                                <Form.Item wrapperCol={{ span: 14 }}>
                                    <Button type="dashed" onClick={() => add()} block icon={<PlusOutlined />}>添加结构字段</Button>
                                </Form.Item>
                            </>
                        )}
                    </Form.List>
                </Form.Item> */}
                        <Form.Item wrapperCol={{ offset: 4 }}>
                            <Button type="primary" htmlType="submit">保存</Button>
                        </Form.Item>
                    </Form>
                </TabPane>
            </Tabs>

            {/* <JsonSchemaForm propKey='' schema={schema}></JsonSchemaForm> */}
        </PageContainer >
    );
};
