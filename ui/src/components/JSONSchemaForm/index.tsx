import React, { useEffect } from 'react';

import type { JSONSchema7 } from 'json-schema';

// import styles from './index.less';
import { Button, Form, } from 'antd';
import { CheckOutlined, CloseOutlined } from '@ant-design/icons';
import ProForm, { ProFormDigit, ProFormGroup, ProFormList, ProFormSwitch, ProFormText } from '@ant-design/pro-form';

const formItemLayout = { labelCol: { span: 4 }, wrapperCol: { span: 20 } }

type JSONSchemaFormProps = {
    schema: JSONSchema7;
    onSave?: (values: any) => void,
    values?: any
}
type JSONSchemaFormObjectProps = JSONSchemaFormProps & {
    fieldKey: any[];
}

const JSONSchemaFormObject: React.FC<JSONSchemaFormObjectProps> = (props) => {
    const { schema, fieldKey } = props;
    return <>
        {Object.keys(schema.properties || {}).map(childPropKey => {
            const childPropSchema = schema.properties && (schema.properties[childPropKey] as JSONSchema7)
            if (!childPropSchema) {
                return <></>
            }
            // console.log(JSON.stringify(childPropSchema))
            switch (childPropSchema.type) {
                case "string":
                    return <ProFormText fieldProps={{ min: childPropSchema?.minLength, max: childPropSchema?.maxLength }} {...formItemLayout} name={[...fieldKey, childPropKey]} label={childPropSchema?.title} tooltip={childPropSchema?.description} initialValue={childPropSchema?.default} required={schema.required?.includes(childPropKey)} />
                case "integer":
                    return <ProFormDigit fieldProps={{ precision: 0 }} min={childPropSchema?.minimum} max={childPropSchema?.maximum} {...formItemLayout} name={[...fieldKey, childPropKey]} label={childPropSchema?.title} tooltip={childPropSchema?.description} initialValue={childPropSchema?.default} required={schema.required?.includes(childPropKey)} />
                case "number":
                    return <ProFormDigit min={childPropSchema?.minimum} max={childPropSchema?.maximum} {...formItemLayout} name={[...fieldKey, childPropKey]} label={childPropSchema?.title} tooltip={childPropSchema?.description} initialValue={childPropSchema?.default} required={schema.required?.includes(childPropKey)} />
                case "boolean":
                    return <ProFormSwitch checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} {...formItemLayout} name={[...fieldKey, childPropKey]} label={childPropSchema?.title} tooltip={childPropSchema?.description} initialValue={childPropSchema?.default} required={schema.required?.includes(childPropKey)} />
                case "object":
                    return <Form.Item {...formItemLayout} label={childPropSchema?.title} tooltip={childPropSchema?.description} initialValue={childPropSchema?.default} required={schema.required?.includes(childPropKey)} >
                        {/* {JSON.stringify([...fieldKey, childPropKey])} */}
                        <JSONSchemaFormObject fieldKey={[...fieldKey, childPropKey]} schema={childPropSchema} />
                    </Form.Item>
                case "array":
                    const arrayFieldSchema = childPropSchema.items as JSONSchema7
                    let childComps;
                    switch (arrayFieldSchema.type) {
                        case "string":
                            childComps = <ProFormText fieldProps={{ min: arrayFieldSchema?.minLength, max: arrayFieldSchema?.maxLength }} {...formItemLayout} name={[]} label={arrayFieldSchema?.title} tooltip={arrayFieldSchema?.description} initialValue={arrayFieldSchema?.default} required={schema.required?.includes(childPropKey)} />
                            break
                        case "integer":
                            childComps = <ProFormDigit fieldProps={{ precision: 0 }} min={arrayFieldSchema?.minimum} max={arrayFieldSchema?.maximum} {...formItemLayout} name={[]} label={arrayFieldSchema?.title} tooltip={arrayFieldSchema?.description} initialValue={arrayFieldSchema?.default} required={schema.required?.includes(childPropKey)} />
                            break
                        case "number":
                            childComps = <ProFormDigit min={arrayFieldSchema?.minimum} max={arrayFieldSchema?.maximum} {...formItemLayout} name={[]} label={arrayFieldSchema?.title} tooltip={arrayFieldSchema?.description} initialValue={arrayFieldSchema?.default} required={schema.required?.includes(childPropKey)} />
                            break
                        case "boolean":
                            childComps = <ProFormSwitch checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} {...formItemLayout} name={[]} label={arrayFieldSchema?.title} tooltip={arrayFieldSchema?.description} initialValue={arrayFieldSchema?.default} required={schema.required?.includes(childPropKey)} />
                            break
                        case "object":
                            // console.log(JSON.stringify(arrayFieldSchema))
                            childComps = <ProFormGroup style={{ width: "200%" }}><JSONSchemaFormObject fieldKey={[...fieldKey]} schema={arrayFieldSchema} /></ProFormGroup>
                            break
                        default:
                            return <></>
                    }
                    return <ProFormList name={[...fieldKey, childPropKey]} label={childPropSchema?.title} tooltip={childPropSchema?.description}>
                        {childComps}
                    </ProFormList>
                default:
                    return <></>
            }
        })}
    </>
}

const JSONSchemaForm: React.FC<JSONSchemaFormProps> = (props) => {
    const [form] = Form.useForm();

    const { schema, onSave, values } = props;

    useEffect(() => {
        form.setFieldsValue(values);
    }, [values])

    const onFinish = (values: any) => {
        console.log('表单数据: ', values);
        onSave && onSave(values);
    };

    const onValuesChange = (changedValues: any, allValues: any) => {
        console.log(changedValues, allValues)
    }


    return (
        <>
            <ProForm {...formItemLayout}
                layout='horizontal'
                form={form}
                onFinish={onFinish}
                onValuesChange={onValuesChange}
            >
                <JSONSchemaFormObject fieldKey={[]} schema={schema} ></JSONSchemaFormObject>
                <Form.Item wrapperCol={{ offset: 4 }}>
                    <Button type="primary" htmlType="submit">保存</Button>
                </Form.Item>
            </ProForm>
        </>
    );
};

export default JSONSchemaForm;
