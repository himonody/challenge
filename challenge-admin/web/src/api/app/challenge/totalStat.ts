import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ChallengeTotalStatModel {
	id?: number;
	totalUserCnt?: number;
	totalJoinCnt?: number;
	totalSuccessCnt?: number;
	totalFailCnt?: number;
	totalJoinAmount?: string;
	totalSuccessAmount?: string;
	totalFailAmount?: string;
	totalPlatformBonus?: string;
	totalPoolAmount?: string;
	updatedAt?: string;
}

export const getTotalStatPageApi = (params: ReqPage) =>
	request.get<ResPage<ChallengeTotalStatModel>>("/admin-api/v1/challenge/total_stat", {
		...params,
		pageIndex: params?.current
	});

export const exportTotalStatApi = (query: object) => request.download("/admin-api/v1/challenge/total_stat/export", query);

export const addTotalStatApi = (data: object) => request.post<object>("/admin-api/v1/challenge/total_stat", data);

export const updateTotalStatApi = (id: number, data: object) =>
	request.put<object>(`/admin-api/v1/challenge/total_stat/${id}`, data);

export const delTotalStatApi = (id: number) => request.delete<object>(`/admin-api/v1/challenge/total_stat/${id}`);
