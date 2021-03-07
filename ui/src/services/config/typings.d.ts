// @ts-ignore
/* eslint-disable */

declare namespace config {
    type Base = {
        create_time?: number;
        create_user_id?: number;
        update_time?: number;
        update_user_id?: number;
    }
    type ConfigStruct = Base & {
        key?: string;
        title?: string;
        description?: string;
        secret?: boolean;
        array?: boolean;
        schema?: string;
    }
}
