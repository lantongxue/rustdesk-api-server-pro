import {request} from "@/service/http";

export const login = async (data) => {
    return await request.post<any>('/login/admin1', {}, {params: data})
}