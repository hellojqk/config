// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 此处后端没有提供注释 GET /api/notices */
export async function insertData(structKey: string, data: config.ConfigData) {
    return request<config.ConfigData>(`/api/struct/${structKey}/data`, {
        method: 'POST',
        data: data
    });
}

export async function updateData(structKey: string, dataKey: string, data: config.ConfigData) {
    return request<config.ConfigData>(`/api/struct/${structKey}/data/${dataKey}`, {
        method: 'PUT',
        data: data
    });
}


export async function findData(structKey: string, params: any) {
    return request<config.ConfigData>(`/api/struct/${structKey}/data`, {
        method: 'GET',
        params: params
    });
}

export async function findOneData(structKey: string, dataKey: string) {
    return request<config.ConfigData>(`/api/struct/${structKey}/data/${dataKey}`, {
        method: 'GET'
    });
}