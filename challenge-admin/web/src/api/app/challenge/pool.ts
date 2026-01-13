import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ChallengePoolModel {
	id?: number;
	configId?: number;
	startDate?: string;
	endDate?: string;
	totalAmount?: string;
	settled?: number;
	createdAt?: string;
}

export const getPoolPageApi = (params: ReqPage) =>
	request.get<ResPage<ChallengePoolModel>>("/admin-api/v1/challenge/pool", {
		...params,
		pageIndex: params?.current
	});

export const exportPoolApi = (query: object) => request.download("/admin-api/v1/challenge/pool/export", query);

export const addPoolApi = (data: object) => request.post<object>("/admin-api/v1/challenge/pool", data);

export const updatePoolApi = (id: number, data: object) => request.put<object>(`/admin-api/v1/challenge/pool/${id}`, data);

export const delPoolApi = (id: number) => request.delete<object>(`/admin-api/v1/challenge/pool/${id}`);
