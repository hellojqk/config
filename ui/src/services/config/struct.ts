// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 此处后端没有提供注释 GET /api/notices */
export async function addStruct(struct: config.ConfigStruct) {
    return request<config.ConfigStruct>('/api/struct', {
        method: 'POST',
        data: struct
    });
}

export async function updateStruct(structKey: string, struct: config.ConfigStruct) {
    return request<config.ConfigStruct>(`/api/struct/${structKey}`, {
        method: 'PUT',
        data: struct
    });
}


export async function findStruct(params: any) {
    return request<config.ConfigStruct>('/api/struct', {
        method: 'GET',
        params: params
    });
}