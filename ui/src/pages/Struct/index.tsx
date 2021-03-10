import { PageContainer } from '@ant-design/pro-layout';
import type { ReactNode } from 'react';
import { useState, useEffect } from 'react';
import {
    FormInstance, message, Table, Tabs
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
import { insertStruct, findStruct, updateStruct } from '@/services/config/struct';
import type config from 'config/config';
import type { ProColumns } from '@ant-design/pro-table';
import ProTable from '@ant-design/pro-table';
import moment from 'moment';
import DataEditor from './components/Data';
import JSONSchemaForm from '@/components/JSONSchemaForm';
import ProForm, { ProFormSwitch, ProFormText, ProFormTextArea } from '@ant-design/pro-form';
import FormItemInput from 'antd/lib/form/FormItemInput';

// Make modifications to the theme with your own fields and widgets

const formItemLayout = {
    labelCol: { span: 4 },
    wrapperCol: { span: 20 },
};

const { TabPane } = Tabs;

export default () => {

    const [form] = Form.useForm();
    const [structFormData, setStructFormData] = useState<config.ConfigStruct>()
    const [activeKey, setActiveKey] = useState<string>("list")
    const [schema, setSchema] = useState<any>({})
    const [editStatus, setEditStatus] = useState<string>("add")

    const [structKey, setStructKey] = useState<string>("")

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
        render: (text: ReactNode, record: config.ConfigStruct) => <Switch disabled checked={record.array} checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} />
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
            render: (_, record: config.ConfigStruct) => [
                <a key="edit" onClick={() => {
                    try {
                        setSchema(JSON.parse(record.schema || "{}"));
                    } catch (error) {

                    }
                    form.setFieldsValue(record);
                    setActiveKey("edit");
                    setEditStatus("update")
                }}>编辑</a>,
                <a key="edit">历史</a>,
                <a key="edit">查看</a>,
                <a key="edit" onClick={() => {
                    setActiveKey("structDataManager");
                    if (record.key) {
                        setStructKey(record.key);
                    }
                }}>数据维护</a>,
            ],
        },
    ];



    const onFinish = (values: config.ConfigStruct) => {
        console.log('表单数据: ', values);
        if (editStatus == "update") {
            updateStruct(values.key as string, values).then(result => {
                // console.log(result)
                message.info(result);
            })
        } else {
            insertStruct(values).then(result => {
                message.info(result);
            })
        }
    };

    const onValuesChange = (changedValues: any, allValues: any) => {
        console.log(changedValues, allValues, form)

        try {
            const json = JSON.parse(allValues.schema);
            setSchema(json);
        } catch (error) {
            console.log(error);
        }
    }


    useEffect(() => {
    }, []);
    return (
        <PageContainer title={false} className={styles.main}>
            <Tabs defaultActiveKey="1" activeKey={activeKey} onTabClick={(key) => { setActiveKey(key) }}>
                <TabPane tab="结构列表" key="list">
                    <ProTable size='small' columns={columns}
                        toolBarRender={
                            () => [<Button type="primary" size='small' onClick={() => {
                                setActiveKey("edit");
                                setEditStatus("add")
                            }}>新增</Button>]
                        }
                        request={(params) => (findStruct({ page_num: params.current, page_size: params.pageSize }))}></ProTable>
                </TabPane>
                <TabPane tab="结构新增" key="edit">
                    <Row>
                        <Col span={10}><ProForm
                            layout='horizontal'
                            form={form}
                            name="validate_other"
                            {...formItemLayout}
                            onFinish={onFinish}
                            onReset={() => { setStructFormData({}); setSchema({}); setEditStatus("add"); form.resetFields() }}
                            onValuesChange={onValuesChange}
                        >
                            <ProFormText disabled={editStatus === "update"} name={"key"} required label={columnKey.title} tooltip={columnKey.description}
                                rules={[{ required: true, message: "请输入结构标识" }, { pattern: /^\w+$/, message: '仅支持数字、字母及下划线' }]} />
                            <ProFormText name="title" required label={columnTitle.title} tooltip={columnTitle.description}
                                rules={[{ required: true, message: "请输入结构名称" }]} />
                            <ProFormTextArea name="description" label={columnDescription.title} tooltip={columnDescription.description} />
                            <ProFormSwitch name="secret" valuePropName="checked" label={columnSecret.title} tooltip={columnSecret.description}
                                checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} />

                            <ProFormSwitch name="array" valuePropName="checked" label={columnArray.title} tooltip={columnArray.description}
                                checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} />
                            <ProFormTextArea name="schema" label="结构" tooltip='json-schema格式数据' fieldProps={{ autoSize: true }} required />
                        </ProForm></Col>
                        <Col span={14}><JSONSchemaForm editStatus={"insert"} schema={schema} /></Col>
                    </Row>
                </TabPane>
                <TabPane tab="结构数据维护" key="structDataManager">
                    <DataEditor structKey={structKey} />
                </TabPane>
            </Tabs>
        </PageContainer >
    );
};
