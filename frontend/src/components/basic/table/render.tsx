import {NTag, NSpace, NImage, NAvatar} from 'naive-ui'
import {unref} from "vue";
import {ITableColumn} from "./index";

export function generateTableColumnRender(field: ITableColumn) {
    let res = null
    switch (field.type) {
        case "switch":
            res = (data: any) => {
                return (
                    <NTag round type={data[field.field] ? 'success' : 'error'}
                          size='small'>{data[field.field] ? field?.typeOptions.switchTag?.trueText ?? '开启' :
                        field?.typeOptions.switchTag?.falseText ?? '关闭'}</NTag>
                )
            }
            break
        case "select":
            res = (data: any) => {
                const val = data[field.field]
                if (field.typeOptions.options) {
                    return (
                        <NSpace size='small' justify='center'>{
                            unref(field.typeOptions.options).map((item: any) => {
                                if (field.typeOptions.multiple && data[field.field] && data[field.field].length) {
                                    if (data[field.field].includes(item.value)) {
                                        return (
                                            <NTag round type={item.tagType ?? 'primary'} size='small'>{item.label}</NTag>
                                        )
                                    }
                                } else {
                                    if (item.value === data[field.field]) {
                                        return (
                                            <NTag round type={item.tagType ?? 'primary'} size='small'>{item.label}</NTag>
                                        )
                                    }
                                }
                            })
                        }</NSpace>
                    )
                }
                return (
                    <span>{val}</span>
                )
            }
            break
        case "date":
        case "datetime":
            res = (data: any) => {
                return (
                    <span>{data[field.field]}</span>
                )
            }
            break
        case "avatar":
            res = (data: any) => {
                return (
                    <NAvatar
                        round
                        size={field.typeOptions.size}
                        src={data[field.field]}
                    />
                )
            }
            break
        case "image":
            res = (data: any) => {
                return (
                    <NImage
                        src={data[field.field]}
                        width={field.width}
                        height={field.height}
                    />
                )
            }
            break
        case "tag":
            res = (data: any) => {
                return (
                    <NTag round type={field.typeOptions.type} size={field.typeOptions.size}>{data[field.field]}</NTag>
                )
            }
            break
        default:
            res = (data: any) => {
                return (
                    <div>
                        <span>{data[field.field]}{field.suffix}</span>
                    </div>
                )
            }
    }
    return res;
}