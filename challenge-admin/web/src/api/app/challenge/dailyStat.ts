import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ChallengeDailyStatModel {
	statDate?: string;
	joinUserCnt?: number;
	successUserCnt?: number;
	failUserCnt?: number;
	joinAmount?: string;
	successAmount?: string;
	failAmount?: string;
	platformBonus?: string;
	poolAmount?: string;
	createdAt?: string;
}

export const getDailyStatPageApi = (params: ReqPage) =>
	request.get<ResPage<ChallengeDailyStatModel>>("/admin-api/v1/challenge/daily_stat", {
		...params,
		pageIndex: params?.current
	});

export const exportDailyStatApi = (query: object) => request.download("/admin-api/v1/challenge/daily_stat/export", query);

export const addDailyStatApi = (data: object) => request.post<object>("/admin-api/v1/challenge/daily_stat", data);

export const updateDailyStatApi = (statDate: string, data: object) =>
	request.put<object>(`/admin-api/v1/challenge/daily_stat/${statDate}`, data);

export const delDailyStatApi = (statDate: string) => request.delete<object>(`/admin-api/v1/challenge/daily_stat/${statDate}`);
