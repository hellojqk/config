import { PageContainer } from '@ant-design/pro-layout';
import React, { useState, useEffect, ReactNode } from 'react';
import { Button, message, Spin } from 'antd';
import styles from './index.less';
import JSONSchemaForm from '@/components/JSONSchemaForm';
import type { JSONSchema7, JSONSchema7Object } from 'json-schema';
import { findOneStruct } from '@/services/config/struct';
import { findData, findOneData, insertData, updateData } from '@/services/config/data';
import ProTable from '@ant-design/pro-table';

const columnKey = {
    title: "唯一标识",
    dataIndex: "key",
    key: "key",
    description: "将作为配置数据存储和获取的唯一标识，供程序识别",
    tooltip: "将作为配置数据存储和获取的唯一标识，供程序识别"
}

type DataEditorProps = {
    structKey?: string
}

const DataEditor: React.FC<DataEditorProps> = (props) => {
    const { structKey } = props;
    const [structInfo, setStructInfo] = useState<config.ConfigStruct>()
    const [structDataInfo, setStructDataInfo] = useState<config.ConfigData>()

    const [structDataColumns, setStructDataColumns] = useState<any[]>([])

    const [addFormStatus, setAddFormStatus] = useState("")

    useEffect(() => {
        console.log("useEffect", structKey)
        if (!structKey) {
            return
        }
        setStructInfo({})
        setStructDataInfo({})
        findOneStruct(structKey).then(result => {
            setStructInfo(result)
            if (result.array && result.schema) {
                const dataSchema = JSON.parse(result.schema) as JSONSchema7;
                let columns: any[] = [columnKey]
                if (!dataSchema.properties) {
                    return
                }
                Object.keys(dataSchema.properties).forEach(fieldKey => {
                    const fieldSchema = dataSchema.properties && dataSchema.properties[fieldKey] as JSONSchema7;
                    if (!fieldSchema) {
                        return;
                    }
                    if (fieldSchema.type == "object" || fieldSchema.type == "array" || !fieldSchema.type) {
                        return;
                    }
                    const column = {
                        title: fieldSchema.title,
                        dataIndex: `${fieldKey}`,
                        key: `${fieldKey}`,
                        description: fieldSchema.description,
                        tooltip: fieldSchema.description,
                    }
                    columns.push(column)
                })
                console.log("setStructDataColumns", columns)
                setStructDataColumns(columns)
            } else {
                findOneData(structKey, structKey).then(result => {
                    setStructDataInfo(result)
                })
            }
        })


    }, [structKey]);

    const onSave = (values: any) => {
        console.log("onSave", values)
        if (!structKey) {
            return
        }
        let key = values.key;
        if (!key) {
            key = structDataInfo?.key;
        }
        if (!structDataInfo?.key) {
            insertData(structKey, values).then(result => {
                message.info("新增成功");
            })
        } else {
            updateData(structKey, structDataInfo.key, values).then(result => {
                message.info("更新成功");
            })
        }
    }
    return (
        <>
            <div>{structInfo?.title}数据管理</div>
            {structInfo && structInfo.schema && <>
                {!structInfo.array ? <JSONSchemaForm editStatus={"update"} values={{ ...structDataInfo?.data, key: structInfo.key }} onSave={onSave} schema={JSON.parse(structInfo?.schema)}></JSONSchemaForm> :
                    <ProTable size='small' columns={structDataColumns}
                        toolBarRender={
                            () => [<Button type="primary" size='small' onClick={() => {
                                setAddFormStatus("true")
                            }}>新增</Button>]
                        }
                        request={(params) => (findData(structKey, { page_num: params.current, page_size: params.pageSize }).then(result => {
                            console.log("findData", result)
                            result.data.map(item => {
                                const keys = Object.keys(item)
                                for (let i = 0; i < keys.length; i++) {
                                    const key = keys[i];
                                    if (item[key] instanceof Object || item[key] instanceof Array) {
                                        item[key] = null
                                    }
                                }
                                return item
                            })
                            return result
                        }))}
                    ></ProTable>}
                {addFormStatus && <JSONSchemaForm editStatus={"insert"} onSave={onSave} schema={JSON.parse(structInfo?.schema)} />}
            </>}
        </>
    );
};

export default DataEditor;
