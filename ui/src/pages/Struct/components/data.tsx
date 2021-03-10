import { PageContainer } from '@ant-design/pro-layout';
import React, { useState, useEffect, ReactNode } from 'react';
import { message, Spin } from 'antd';
import styles from './index.less';
import JSONSchemaForm from '@/components/JSONSchemaForm';
import type { JSONSchema7, JSONSchema7Object } from 'json-schema';
import { findOneStruct } from '@/services/config/struct';
import { findData, findOneData, insertData, updateData } from '@/services/config/data';
import ProTable from '@ant-design/pro-table';


type DataEditorProps = {
    structKey?: string
}

const DataEditor: React.FC<DataEditorProps> = (props) => {
    const { structKey } = props;
    const [structInfo, setStructInfo] = useState<config.ConfigStruct>()
    const [structDataInfo, setStructDataInfo] = useState<config.ConfigData>()

    const [structDataColumns, setStructDataColumns] = useState<any[]>([])

    useEffect(() => {
        console.log("useEffect", structKey)
        if (!structKey) {
            return
        }
        setStructInfo({})
        setStructDataInfo({})
        findOneStruct(structKey).then(result => {
            setStructInfo(result)

            if (result.array) {
                const dataSchema = JSON.stringify(result.schema) as JSONSchema7;
                let columns: any[] = []
                if (!dataSchema.properties) {
                    return
                }
                Object.keys(dataSchema.properties).forEach(fieldKey => {
                    const fieldSchema = dataSchema.properties && dataSchema.properties[fieldKey] as JSONSchema7;
                    if (!fieldSchema) {
                        return;
                    }
                    const column = {
                        title: fieldSchema.title,
                        dataIndex: fieldKey,
                        key: fieldKey,
                        description: fieldSchema.description,
                        tooltip: fieldSchema.description,
                    }
                    columns.push(column)
                })
                setStructDataColumns(columns)
            }
        })

        findOneData(structKey, structKey).then(result => {
            setStructDataInfo(result)
        })
    }, [structKey]);

    const onSave = (values: any) => {
        console.log("onSave", values)
        if (!structKey) {
            return
        }
        if (!structDataInfo?.key) {
            insertData(structKey, { key: structKey, data: values }).then(result => {
                message.info(result);
            })
        } else {
            updateData(structKey, structDataInfo.key, { data: values }).then(result => {
                message.info(result);
            })
        }
    }
    return (
        <>
            <div>{structInfo?.title}数据管理</div>
            {structInfo && structInfo.schema && <>
                {!structInfo.array ? <JSONSchemaForm values={structDataInfo?.data} onSave={onSave} schema={JSON.parse(structInfo?.schema)}></JSONSchemaForm> :
                    <ProTable size='small' columns={structDataColumns}
                        request={(params) => (findData(structKey, { page_num: params.current, page_size: params.pageSize }))}
                    ></ProTable>}
            </>}
        </>
    );
};

export default DataEditor;
