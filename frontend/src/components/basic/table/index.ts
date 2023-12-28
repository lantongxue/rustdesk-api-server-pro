export interface ITableColumn {
    // 表格字段
    field: string,
    // 表格标题
    title: string,
    type?: ColumnType,
    render?: Function,
    labelWidth?: number,
    height?: number,
    width?: number,
    fixed?: boolean|string,
    typeOptions?: {
        switchTag?: {
            trueText?: string,
            falseText?: string,
        },
        multiple?: boolean,
        size?: 'small' | 'medium' | 'large' | number | any,
        options?: Array<any>,
        type?:any,
    },
    hide?: boolean,
    // 结尾
    suffix?: string
}


type ColumnType = 'image' | 'switch' | 'select' | 'date' | 'datetime' | 'avatar' | 'string' | 'tag'


