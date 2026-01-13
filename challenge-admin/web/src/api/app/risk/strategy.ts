import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface RiskStrategyModel {
	id?: number;
	scene?: string;
	ruleCode?: string;
	identityType?: string;
	windowSeconds?: number;
	threshold?: number;
	action?: string;
	actionValue?: number;
	status?: number;
	priority?: number;
	remark?: string;
	createdAt?: string;
	updatedAt?: string;
}

export const getRiskStrategyPageApi = (params: ReqPage) =>
	request.get<ResPage<RiskStrategyModel>>("/admin-api/v1/risk/strategy", { ...params, pageIndex: params?.current });

export const exportRiskStrategyApi = (query: object) => request.download("/admin-api/v1/risk/strategy/export", query);

export const addRiskStrategyApi = (data: object) => request.post<object>("/admin-api/v1/risk/strategy", data);

export const updateRiskStrategyApi = (id: number, data: object) => request.put<object>(`/admin-api/v1/risk/strategy/${id}`, data);

export const delRiskStrategyApi = (id: number) => request.delete<object>(`/admin-api/v1/risk/strategy/${id}`);
