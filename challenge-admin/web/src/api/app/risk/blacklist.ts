import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface RiskBlacklistModel {
	id?: number;
	type?: string;
	value?: string;
	riskLevel?: number;
	reason?: string;
	status?: string;
	createdAt?: string;
}

export const getRiskBlacklistPageApi = (params: ReqPage) =>
	request.get<ResPage<RiskBlacklistModel>>("/admin-api/v1/risk/blacklist", { ...params, pageIndex: params?.current });

export const exportRiskBlacklistApi = (query: object) => request.download("/admin-api/v1/risk/blacklist/export", query);

export const addRiskBlacklistApi = (data: object) => request.post<object>("/admin-api/v1/risk/blacklist", data);

export const updateRiskBlacklistApi = (id: number, data: object) => request.put<object>(`/admin-api/v1/risk/blacklist/${id}`, data);

// 后端暂未提供删除接口，前端不做删除
