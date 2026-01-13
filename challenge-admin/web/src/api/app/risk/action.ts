import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface RiskActionModel {
	code?: string;
	type?: string;
	defaultValue?: number;
	remark?: string;
	createdAt?: string;
}

export const getRiskActionPageApi = (params: ReqPage) =>
	request.get<ResPage<RiskActionModel>>("/admin-api/v1/risk/action", { ...params, pageIndex: params?.current });

export const exportRiskActionApi = (query: object) => request.download("/admin-api/v1/risk/action/export", query);

export const addRiskActionApi = (data: object) => request.post<object>("/admin-api/v1/risk/action", data);

export const updateRiskActionApi = (code: string, data: object) => request.put<object>(`/admin-api/v1/risk/action/${code}`, data);

export const delRiskActionApi = (code: string) => request.delete<object>(`/admin-api/v1/risk/action/${code}`);
