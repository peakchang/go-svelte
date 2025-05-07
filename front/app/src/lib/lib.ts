import axios, { AxiosError, type AxiosResponse } from 'axios';

interface ErrorResponse {
    message?: string;
    [key: string]: any;
}

export async function axiosPost<T>(link: string, obj: object): Promise<AxiosResponse<T> | null> {
    try {
        const res = await axios.post<T>(link, obj);
        return res; // ✅ 타입 일치
    } catch (error) {
        const err = error as AxiosError<ErrorResponse>;
        if (err.response?.data?.message) {
            alert(err.response.data.message);
        } else {
            alert("서버와의 통신에 실패했습니다. 다시 시도해주세요.");
            console.error(error);
        }
        return null;
    }
}