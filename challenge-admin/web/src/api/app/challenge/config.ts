import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ChallengeConfigModel {
	id?: number;
	dayCount?: number;
	amount?: string;
	checkinStart?: string;
	checkinEnd?: string;
	platformBonus?: string;
	status?: number;
	sort?: number;
	createdAt?: string;
}

export const getConfigPageApi = (params: ReqPage) =>
	request.get<ResPage<ChallengeConfigModel>>("/admin-api/v1/challenge/config", {
		...params,
		pageIndex: params?.current
	});

export const exportConfigApi = (query: object) => request.download("/admin-api/v1/challenge/config/export", query);

export const addConfigApi = (data: object) => request.post<object>("/admin-api/v1/challenge/config", data);

export const updateConfigApi = (id: number, data: object) => request.put<object>(`/admin-api/v1/challenge/config/${id}`, data);

export const delConfigApi = (id: number) => request.delete<object>(`/admin-api/v1/challenge/config/${id}`);
