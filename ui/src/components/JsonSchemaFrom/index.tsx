import React from 'react';

import type { JSONSchema7 } from 'json-schema';

import styles from './index.less';
import { Button, Form, Input, InputNumber, Switch } from 'antd';
import { CheckOutlined, CloseOutlined, MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';

const formItemLayout = { labelCol: { span: 4 }, wrapperCol: { span: 18 } }

const formItemLayoutWithOutLabel = {
    wrapperCol: {
        xs: { span: 12, offset: 0 },
        sm: { span: 12, offset: 2 },
    },
};
type JsonSchemaFormProps = {
    schema: JSONSchema7;
    fieldKey: any[];
}

// const JsonSchemaFormItemArray: React.FC<JsonSchemaFormProps> = (props) => {
//     const { propKey, schema } = props;
//     return <>
//         <Form.Item  {...formItemLayout} name={propKey} fieldKey={propKey} label={schema?.title} tooltip={schema?.description} initialValue={schema?.default} required={schema.required?.includes(propKey)} key={propKey}>
//             {!(schema.items instanceof Array) &&
//                 <Form.List name={propKey}>
//                     {(fields, { add, remove }, { errors }) => (
//                         <>
//                             {fields.map((field, index) => (
//                                 <Form.Item {...formItemLayoutWithOutLabel} required={false}>
//                                     {schema.items && (schema.items as JSONSchema7).type === 'string' && <JsonSchemaFormItemString propKey={propKey} schema={schema.items as JSONSchema7} />}
//                                     {schema.items && (schema.items as JSONSchema7).type === 'boolean' && <JsonSchemaFormItemBoolean propKey={propKey} schema={schema.items as JSONSchema7} />}
//                                     {schema.items && (schema.items as JSONSchema7).type === 'integer' && <JsonSchemaFormItemInteger propKey={propKey} schema={schema.items as JSONSchema7} />}
//                                     {schema.items && (schema.items as JSONSchema7).type === 'number' && <JsonSchemaFormItemNumber propKey={propKey} schema={schema.items as JSONSchema7} />}
//                                     {/* {schema.items && (schema.items as JSONSchema7).type === 'object' && <JsonSchemaFormItemObject propKey={propKey} schema={schema.items as JSONSchema7} />} */}
//                                     {fields.length > 1 ? (<MinusCircleOutlined className="dynamic-delete-button" onClick={() => remove(field.name)} />
//                                     ) : null}
//                                 </Form.Item>
//                             ))}
//                             <Form.Item {...formItemLayoutWithOutLabel}>
//                                 <Button type="dashed" onClick={() => add()} style={{ width: '60%' }} icon={<PlusOutlined />}>添加</Button>
//                                 <Form.ErrorList errors={errors} />
//                             </Form.Item>
//                         </>
//                     )}
//                 </Form.List>
//             }
//         </Form.Item>
//     </>
// }

const JsonSchemaFormItemObject: React.FC<JsonSchemaFormProps> = (props) => {
    const { schema, fieldKey } = props;
    return <>
        {Object.keys(schema.properties || {}).map(childPropKey => {
            const childPropSchema = schema.properties && (schema.properties[childPropKey] as JSONSchema7)
            if (!childPropSchema) {
                return <></>
            }
            switch (childPropSchema.type) {
                case "string":
                    return <Form.Item {...formItemLayout} name={[...fieldKey, childPropKey]} label={childPropSchema?.title} tooltip={childPropSchema?.description} initialValue={childPropSchema?.default} required={schema.required?.includes(childPropKey)} >
                        <Input />
                    </Form.Item>
                case "integer":
                    return <Form.Item {...formItemLayout} name={[...fieldKey, childPropKey]} label={childPropSchema?.title} tooltip={childPropSchema?.description} initialValue={childPropSchema?.default} required={schema.required?.includes(childPropKey)} >
                        <InputNumber min={childPropSchema?.minimum} max={childPropSchema?.maximum} />
                    </Form.Item>
                case "number":
                    return <Form.Item {...formItemLayout} name={[...fieldKey, childPropKey]} label={childPropSchema?.title} tooltip={childPropSchema?.description} initialValue={childPropSchema?.default} required={schema.required?.includes(childPropKey)} >
                        <InputNumber min={childPropSchema?.minimum} max={childPropSchema?.maximum} />
                    </Form.Item>
                case "boolean":
                    return <Form.Item {...formItemLayout} name={[...fieldKey, childPropKey]} label={childPropSchema?.title} tooltip={childPropSchema?.description} initialValue={childPropSchema?.default} required={schema.required?.includes(childPropKey)} >
                        <Switch checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} />
                    </Form.Item>
                case "object":
                    return <Form.Item {...formItemLayout} label={childPropSchema?.title} tooltip={childPropSchema?.description} initialValue={childPropSchema?.default} required={schema.required?.includes(childPropKey)} >
                        {/* {JSON.stringify([...fieldKey, childPropKey])} */}
                        <JsonSchemaFormItemObject fieldKey={[...fieldKey, childPropKey]} schema={childPropSchema} />
                    </Form.Item>
                case "array":
                    return <Form.Item {...formItemLayout} label={childPropSchema?.title} tooltip={childPropSchema?.description} initialValue={childPropSchema?.default} required={schema.required?.includes(childPropKey)} >
                        {/* {JSON.stringify(childPropSchema)} */}
                        {childPropSchema.items &&
                            <Form.List name={[...fieldKey, childPropKey]}>
                                {(fields, { add, remove }, { errors }) => (
                                    <>
                                        {fields.map((field) => {
                                            const fieldSchema = childPropSchema.items as JSONSchema7
                                            switch (fieldSchema.type) {
                                                case "string":
                                                    return <Form.Item required={false} key={field.key}>
                                                        <Form.Item {...field} style={{ width: "50%" }}>
                                                            <Input />
                                                        </Form.Item>
                                                        {fields.length > 1 ? (<MinusCircleOutlined className={styles['dynamic-delete-button']} onClick={() => remove(field.name)} />) : null}
                                                    </Form.Item>
                                                case "integer":
                                                    return <Form.Item required={false} key={field.key}>
                                                        <Form.Item {...field} style={{ width: "50%" }}>
                                                            <InputNumber min={fieldSchema?.minimum} max={fieldSchema?.maximum} />
                                                        </Form.Item>
                                                        {fields.length > 1 ? (<MinusCircleOutlined className={styles['dynamic-delete-button']} onClick={() => remove(field.name)} />) : null}
                                                    </Form.Item>
                                                case "number":
                                                    return <Form.Item required={false} key={field.key}>
                                                        <Form.Item {...field} style={{ width: "50%" }}>
                                                            <InputNumber min={fieldSchema?.minimum} max={fieldSchema?.maximum} />
                                                        </Form.Item>
                                                        {fields.length > 1 ? (<MinusCircleOutlined className={styles['dynamic-delete-button']} onClick={() => remove(field.name)} />) : null}
                                                    </Form.Item>
                                                case "boolean":
                                                    return <Form.Item required={false} key={field.key}>
                                                        <Form.Item {...field} style={{ width: "50%" }}>
                                                            <Switch checkedChildren={<CheckOutlined />} unCheckedChildren={<CloseOutlined />} />
                                                        </Form.Item>
                                                        {fields.length > 1 ? (<MinusCircleOutlined className={styles['dynamic-delete-button']} onClick={() => remove(field.name)} />) : null}
                                                    </Form.Item>
                                                case "object":
                                                    return <Form.Item required={false} key={field.key}>
                                                        <JsonSchemaFormItemObject fieldKey={[field.fieldKey]} schema={fieldSchema} />
                                                        {fields.length > 1 ? (<MinusCircleOutlined className={styles['dynamic-delete-button']} onClick={() => remove(field.name)} />) : null}
                                                    </Form.Item>
                                                default:
                                                    return <></>
                                            }
                                            // return <Form.Item required={false} key={field.key}>
                                            //     {JSON.stringify(childPropSchema.items)}
                                            //     <JsonSchemaFormItemObject fieldKey={[field.fieldKey]} schema={childPropSchema.items as JSONSchema7} />
                                            //     {/* <Form.Item {...field} style={{ width: "50%" }}>
                                            //     <Input />
                                            // </Form.Item> */}
                                            //     {fields.length > 1 ? (
                                            //         <MinusCircleOutlined
                                            //             className={styles['dynamic-delete-button']}
                                            //             onClick={() => remove(field.name)}
                                            //         />
                                            //     ) : null}
                                            // </Form.Item>
                                        })}
                                        <Form.Item>
                                            <Button type="dashed" onClick={() => add()} style={{ width: '60%' }} icon={<PlusOutlined />}>添加</Button>
                                            <Form.ErrorList errors={errors} />
                                        </Form.Item>
                                    </>
                                )}
                            </Form.List>
                        }
                    </Form.Item>
                default:
                    return <Form.Item {...formItemLayout} name={childPropKey} label={childPropSchema?.title} tooltip={childPropSchema?.description} initialValue={childPropSchema?.default} required={schema.required?.includes(childPropKey)} >
                        <Input />
                    </Form.Item>
            }
        })}
    </>
}

const JsonSchemaForm: React.FC<JsonSchemaFormProps> = (props) => {
    const [form] = Form.useForm();

    const { schema } = props;

    const onFinish = (values: any) => {
        console.log('表单数据: ', values);
    };

    const onValuesChange = (changedValues: any, allValues: any) => {
        console.log(changedValues, allValues)
    }


    return (
        <>
            <Form {...formItemLayout}
                form={form}
                onFinish={onFinish}
                onValuesChange={onValuesChange}
            >
                <JsonSchemaFormItemObject fieldKey={[]} schema={schema} ></JsonSchemaFormItemObject>
                <Form.Item wrapperCol={{ offset: 4 }}>
                    <Button type="primary" htmlType="submit">保存</Button>
                </Form.Item>
            </Form>
        </>
    );
};

export default JsonSchemaForm;
